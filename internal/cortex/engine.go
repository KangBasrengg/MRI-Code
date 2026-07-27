// Package cortex implements Phase 06 AI reasoning and conversational chatbot intelligence over the Neural Repository Graph.
// It features native integration with Freemodel.dev (OpenAI-compatible LLMs) enriched with local structural NRG context,
// and gracefully falls back to local offline structural reasoning when offline.
package cortex

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/KangBasrengg/MRI-Code/internal/graph"
)

// DefaultFreemodelAPIKey is the API key provided for freemodel.dev integrations.
const DefaultFreemodelAPIKey = "fe_oa_a86d249e95f7de55374c22071a87f4f62b71326084202490"

// OfflineMode enforces strict local structural reasoning and disables outbound cloud AI connections.
var OfflineMode bool

// ChatResponse represents the conversational output sent back to the interactive Web Analyzer UI or CLI.
type ChatResponse struct {
	Reply       string          `json:"reply"`
	Symbols     []SymbolSummary `json:"relevant_symbols"`
	HealthScore int             `json:"health_score,omitempty"`
	Engine      string          `json:"engine"`
}

// SymbolSummary summarizes a structural node in the NRG for display in the AI Chat interface.
type SymbolSummary struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Path         string `json:"path"`
	StartLine    int    `json:"start_line"`
	EndLine      int    `json:"end_line"`
	CallersCount int    `json:"callers_count"`
	DepsCount    int    `json:"dependencies_count"`
}

// GenerateChatReply analyzes a natural language question against the repository's cached NRG database,
// constructs an intelligent context prompt, and attempts to synthesize a deep natural language response
// via Freemodel.dev AI, falling back to local structural synthesis if offline.
func GenerateChatReply(repoPath string, question string) (*ChatResponse, error) {
	if repoPath == "" {
		repoPath = "."
	}
	absPath, err := filepath.Abs(repoPath)
	if err != nil {
		return nil, fmt.Errorf("invalid repository directory: %w", err)
	}

	codemriDir := filepath.Join(absPath, ".codemri")
	dbPath := filepath.Join(codemriDir, "graph.db")
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		return &ChatResponse{
			Reply:  "⚠️ **No analysis index found for this repository.**\nPlease run a scan first by pasting a GitHub repository URL in the search analyzer above and clicking **'⚡ Analyze & Clone'**, or executing `codemri scan .` in your terminal.",
			Engine: "Cortex Local Architecture (Offline)",
		}, nil
	}

	sqliteStore, err := graph.NewSQLiteStorage(dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open NRG SQLite database: %w", err)
	}
	defer sqliteStore.Close()

	nrg, err := sqliteStore.LoadGraph(absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load repository topology: %w", err)
	}

	// Read repository summary & health if available
	var healthScore int
	var debtStatus string
	pulsePath := filepath.Join(codemriDir, "pulse.json")
	if pulseData, err := os.ReadFile(pulsePath); err == nil {
		var pulse map[string]interface{}
		if json.Unmarshal(pulseData, &pulse) == nil {
			if health, ok := pulse["health"].(map[string]interface{}); ok {
				if score, ok := health["overall_score"].(float64); ok {
					healthScore = int(score)
				}
				if ds, ok := health["debt_status"].(string); ok {
					debtStatus = ds
				}
			}
		}
	}

	// Read security grade if available
	var secGrade string
	secPath := filepath.Join(codemriDir, "security.json")
	if secData, err := os.ReadFile(secPath); err == nil {
		var sec map[string]interface{}
		if json.Unmarshal(secData, &sec) == nil {
			if sg, ok := sec["security_grade"].(string); ok {
				secGrade = sg
			}
		}
	}

	// Extract keywords and search NRG for symbols relevant to the question
	keywords := ExtractKeywords(question)
	relevantNodes := SearchNRG(nrg, keywords)

	// If no direct keyword matches (e.g. general questions like "apa fungsi repo ini?"), fallback to top structural hub nodes
	if len(relevantNodes) == 0 && len(nrg.Nodes) > 0 {
		topNodes := make([]*graph.Node, 0, len(nrg.Nodes))
		for _, n := range nrg.Nodes {
			topNodes = append(topNodes, n)
		}
		sort.Slice(topNodes, func(i, j int) bool {
			in1, out1 := CountConnections(nrg, topNodes[i].ID)
			in2, out2 := CountConnections(nrg, topNodes[j].ID)
			return (in1 + out1) > (in2 + out2)
		})
		if len(topNodes) > 12 {
			relevantNodes = topNodes[:12]
		} else {
			relevantNodes = topNodes
		}
	}

	symbols := make([]SymbolSummary, 0)
	var symbolContext strings.Builder
	limit := 12
	for i, n := range relevantNodes {
		in, out := CountConnections(nrg, n.ID)
		rel, _ := filepath.Rel(absPath, n.Path)
		if rel == "" {
			rel = n.Path
		}
		rel = filepath.ToSlash(rel)

		symbols = append(symbols, SymbolSummary{
			Name:         n.Name,
			Type:         string(n.Type),
			Path:         rel,
			StartLine:    n.StartLine,
			EndLine:      n.EndLine,
			CallersCount: in,
			DepsCount:    out,
		})

		if i < limit {
			symbolContext.WriteString(fmt.Sprintf("- [%s] %s in %s (lines %d-%d) | Callers: %d | Dependencies: %d\n", n.Type, n.Name, rel, n.StartLine, n.EndLine, in, out))
		}
	}

	// Read README snippet if available for general context enrichment
	var readmeSnippet string
	for _, fname := range []string{"README.md", "readme.md", "README.txt", "readme.txt"} {
		if content, err := os.ReadFile(filepath.Join(absPath, fname)); err == nil {
			snippet := string(content)
			if len(snippet) > 800 {
				snippet = snippet[:800] + "..."
			}
			readmeSnippet = snippet
			break
		}
	}

	// Construct comprehensive AI system prompt & context
	systemPrompt := "You are CodeMRI Cortex, an elite AI software architecture consultant and codebase intelligence assistant created by Muhammad Nuril (@KangBasrengg). You analyze repositories using Neural Repository Graph (NRG) analytical data and documentation summaries. Answer the user's questions clearly, accurately, and concisely in GitHub-style Markdown. Highlight key architecture patterns, security practices, and symbol relationships when relevant."

	var promptContext strings.Builder
	promptContext.WriteString(fmt.Sprintf("Repository Architecture Profile:\n"))
	promptContext.WriteString(fmt.Sprintf("- Total Structural Symbols: %d Nodes, %d Relational Edges\n", len(nrg.Nodes), len(nrg.Edges)))
	if healthScore > 0 {
		promptContext.WriteString(fmt.Sprintf("- Pulse Health Score: %d / 100 (%s)\n", healthScore, debtStatus))
	}
	if secGrade != "" {
		promptContext.WriteString(fmt.Sprintf("- Shield Security Grade: %s\n", secGrade))
	}
	if readmeSnippet != "" {
		promptContext.WriteString(fmt.Sprintf("\nRepository README Snapshot:\n```\n%s\n```\n", readmeSnippet))
	}
	if symbolContext.Len() > 0 {
		promptContext.WriteString(fmt.Sprintf("\nCore Architectural Symbols & Dependencies Discovered:\n%s", symbolContext.String()))
	}

	userPrompt := fmt.Sprintf("Repository Context:\n%s\n\nUser Question:\n%s", promptContext.String(), question)

	// Attempt integration with Freemodel.dev AI unless strict offline mode is enabled
	if !OfflineMode && os.Getenv("CODEMRI_OFFLINE") != "1" && os.Getenv("CODEMRI_OFFLINE") != "true" {
		apiKey := os.Getenv("CODEMRI_API_KEY")
		if apiKey == "" {
			apiKey = DefaultFreemodelAPIKey
		}

		aiReply, aiErr := CallFreemodelAI(apiKey, systemPrompt, userPrompt)
		if aiErr == nil && len(strings.TrimSpace(aiReply)) > 10 {
			return &ChatResponse{
				Reply:       aiReply,
				Symbols:     symbols,
				HealthScore: healthScore,
				Engine:      "Cortex + Freemodel.dev AI (Online Intelligence)",
			}, nil
		}
	}

	// Graceful Fallback: Local Structural Synthesis (Offline-First)
	qLower := strings.ToLower(question)
	isGeneral := strings.Contains(qLower, "what is this") || strings.Contains(qLower, "overview") || strings.Contains(qLower, "summary") || strings.Contains(qLower, "apa fungsi") || strings.Contains(qLower, "fungsi dari") || strings.Contains(qLower, "repository ini") || strings.Contains(qLower, "tentang apa") || strings.Contains(qLower, "jelaskan") || len(keywords) == 0

	var fallbackReply strings.Builder
	if isGeneral {
		fallbackReply.WriteString("### 🧠 Repository Architectural Overview (Local Cortex Engine)\n\n")
		fallbackReply.WriteString(fmt.Sprintf("This repository is indexed into an offline Neural Repository Graph (NRG) with **%d structural nodes** and **%d relational bonds**.\n\n", len(nrg.Nodes), len(nrg.Edges)))
		if healthScore > 0 {
			fallbackReply.WriteString(fmt.Sprintf("💓 **Repository Health Score:** **%d / 100** (%s)\n\n", healthScore, debtStatus))
		}
		if secGrade != "" {
			fallbackReply.WriteString(fmt.Sprintf("🛡️ **Security Intelligence Grade:** **%s**\n\n", secGrade))
		}
		fallbackReply.WriteString("#### 🏗️ Notable Hub Symbols in Codebase:\n")
		for i, s := range symbols {
			if i >= 6 {
				break
			}
			fallbackReply.WriteString(fmt.Sprintf("- **`%s`** *(%s)* — **%d callers**, **%d dependencies** at `%s:%d`\n", s.Name, s.Type, s.CallersCount, s.DepsCount, s.Path, s.StartLine))
		}
	} else {
		fallbackReply.WriteString(fmt.Sprintf("### 🔍 Cortex Local Structural Answer\n\n"))
		fallbackReply.WriteString(fmt.Sprintf("Regarding **\"%s\"**, CodeMRI discovered **%d relevant structural symbols** matching keywords `[%s]`:\n\n", question, len(symbols), strings.Join(keywords, ", ")))
		for i, s := range symbols {
			if i >= 8 {
				break
			}
			fallbackReply.WriteString(fmt.Sprintf("- **`%s`** *(%s)* at `%s:%d-%d` [↑%d callers | ↓%d deps]\n", s.Name, s.Type, s.Path, s.StartLine, s.EndLine, s.CallersCount, s.DepsCount))
		}
		if len(symbols) == 0 {
			fallbackReply.WriteString("No direct symbol matches found. Try searching for domain concepts like `database`, `api`, `server`, `parse`, or `config`.\n")
		}
	}
	fallbackReply.WriteString("\n\n*🛡️ Offline Fallback Mode Active (ADR-0002): Computed locally via SQLite NRG structural traversal.*")

	engineTitle := "Cortex v1.0.0 Structural AI (Local Offline Fallback)"
	if OfflineMode || os.Getenv("CODEMRI_OFFLINE") == "1" || os.Getenv("CODEMRI_OFFLINE") == "true" {
		engineTitle = "Cortex v1.0.0 Structural AI (Strict Offline Mode - ADR-0002)"
	}

	return &ChatResponse{
		Reply:       fallbackReply.String(),
		Symbols:     symbols,
		HealthScore: healthScore,
		Engine:      engineTitle,
	}, nil
}

// CallFreemodelAI executes a chat completion HTTP request against Freemodel.dev API endpoints,
// automatically trying high-performance model aliases (gpt-5.6-luna, sol, terra) with intelligent failover.
func CallFreemodelAI(apiKey, systemPrompt, userPrompt string) (string, error) {
	url := os.Getenv("CODEMRI_AI_URL")
	if url == "" {
		url = "https://api.freemodel.dev/v1/chat/completions"
	}

	modelEnv := os.Getenv("CODEMRI_AI_MODEL")
	modelsToTry := []string{"gpt-5.6-luna", "gpt-5.6-sol", "gpt-5.6-terra"}
	if modelEnv != "" {
		modelsToTry = append([]string{modelEnv}, modelsToTry...)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 16*time.Second)
	defer cancel()
	client := &http.Client{Timeout: 16 * time.Second}

	var lastErr error
	for _, model := range modelsToTry {
		payload := map[string]interface{}{
			"model": model,
			"messages": []map[string]string{
				{"role": "system", "content": systemPrompt},
				{"role": "user", "content": userPrompt},
			},
			"temperature": 0.3,
			"max_tokens":  1200,
		}

		jsonBytes, err := json.Marshal(payload)
		if err != nil {
			return "", err
		}

		req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonBytes))
		if err != nil {
			return "", err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+apiKey)

		resp, err := client.Do(req)
		if err != nil {
			// Fallback trial for api.freemodels.dev plural variation if DNS fails
			if strings.Contains(url, "freemodel.dev") && !strings.Contains(url, "freemodels.dev") {
				altURL := "https://api.freemodels.dev/v1/chat/completions"
				reqAlt, _ := http.NewRequestWithContext(ctx, "POST", altURL, bytes.NewBuffer(jsonBytes))
				reqAlt.Header.Set("Content-Type", "application/json")
				reqAlt.Header.Set("Authorization", "Bearer "+apiKey)
				if altResp, altErr := client.Do(reqAlt); altErr == nil {
					resp = altResp
					err = nil
				} else {
					lastErr = altErr
					continue
				}
			} else {
				lastErr = err
				continue
			}
		}

		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			bodyBytes, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			lastErr = fmt.Errorf("AI API returned status %d for model %s: %s", resp.StatusCode, model, string(bodyBytes))
			continue
		}

		var result struct {
			Choices []struct {
				Message struct {
					Content string `json:"content"`
				} `json:"message"`
			} `json:"choices"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			lastErr = err
			continue
		}
		resp.Body.Close()

		if len(result.Choices) > 0 && len(strings.TrimSpace(result.Choices[0].Message.Content)) > 0 {
			return strings.TrimSpace(result.Choices[0].Message.Content), nil
		}
		lastErr = fmt.Errorf("empty AI choices returned for model %s", model)
	}
	return "", fmt.Errorf("all cloud AI model attempts failed: %v", lastErr)
}

// ExtractKeywords removes common English words and punctuation to isolate code technical concepts.
func ExtractKeywords(question string) []string {
	stopWords := map[string]bool{
		"how": true, "does": true, "the": true, "what": true, "are": true,
		"is": true, "in": true, "a": true, "an": true, "of": true,
		"to": true, "and": true, "or": true, "for": true, "with": true,
		"my": true, "me": true, "show": true, "which": true, "where": true,
		"can": true, "you": true, "this": true, "that": true, "work": true,
		"works": true, "main": true, "most": true, "have": true, "has": true,
		"about": true, "please": true, "tell": true, "find": true, "function": true,
		"repo": true, "repository": true, "code": true, "project": true, "file": true,
		// Indonesian Stop Words & Query Terms
		"apa": true, "fungsi": true, "dari": true, "ini": true, "adalah": true,
		"bagaimana": true, "kenapa": true, "dimana": true, "siapa": true, "untuk": true,
		"di": true, "ke": true, "pada": true, "dan": true, "atau": true,
		"yang": true, "dengan": true, "sebutkan": true, "jelaskan": true, "buat": true,
		"buatkan": true, "tentang": true, "bagaimanakah": true, "mengenai": true,
	}

	words := strings.Fields(strings.ToLower(question))
	keywords := make([]string, 0)
	for _, w := range words {
		w = strings.Trim(w, "?.,!\"'()[]{}<>;:")
		if len(w) > 2 && !stopWords[w] {
			keywords = append(keywords, w)
		}
	}
	return keywords
}

// SearchNRG searches nodes in the graph based on relevancy scoring against keywords.
func SearchNRG(nrg *graph.NeuralRepositoryGraph, keywords []string) []*graph.Node {
	type scoredNode struct {
		Node  *graph.Node
		Score int
	}

	scored := make([]scoredNode, 0)
	for _, node := range nrg.Nodes {
		score := 0
		nameLower := strings.ToLower(node.Name)
		pathLower := strings.ToLower(node.Path)

		for _, kw := range keywords {
			if strings.Contains(nameLower, kw) {
				score += 12
			}
			if strings.Contains(pathLower, kw) {
				score += 6
			}
		}

		if score > 0 {
			scored = append(scored, scoredNode{Node: node, Score: score})
		}
	}

	sort.Slice(scored, func(i, j int) bool {
		return scored[i].Score > scored[j].Score
	})

	result := make([]*graph.Node, 0)
	for i, s := range scored {
		if i >= 30 {
			break
		}
		result = append(result, s.Node)
	}
	return result
}

// CountConnections counts incoming caller edges and outgoing dependency edges for a node ID.
func CountConnections(nrg *graph.NeuralRepositoryGraph, nodeID string) (int, int) {
	incoming := 0
	outgoing := 0
	for _, edge := range nrg.Edges {
		if edge.TargetID == nodeID {
			incoming++
		}
		if edge.SourceID == nodeID {
			outgoing++
		}
	}
	return incoming, outgoing
}
