package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/KangBasrengg/MRI-Code/internal/cortex"
	"github.com/KangBasrengg/MRI-Code/internal/graph"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// explainCmd implements Phase 6 "Cortex" AI reasoning over the Neural Repository Graph.
// It uses local NRG structural intelligence to answer questions about the codebase
// without requiring external LLM API calls (offline-first reasoning).
var explainCmd = &cobra.Command{
	Use:   "explain <question>",
	Short: "AI reasoning engine — ask questions about your codebase architecture",
	Long: `Phase 6 Cortex: Ask natural language questions about your repository.
CodeMRI traverses the Neural Repository Graph to find relevant architectural
nodes and generates structured explanations without reading raw source code.

Examples:
  codemri explain "How does authentication work?"
  codemri explain "What are the main packages?"
  codemri explain "Show me the database layer"
  codemri explain "Which functions have the most dependencies?"`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		question := args[0]

		cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()
		white := color.New(color.FgWhite, color.Bold).SprintFunc()

		cwd, _ := os.Getwd()
		codemriDir := filepath.Join(cwd, ".codemri")
		dbPath := filepath.Join(codemriDir, "graph.db")

		fmt.Println(cyan("🧠 [CodeMRI CORTEX] Offline AI Reasoning Engine"))
		fmt.Println("─────────────────────────────────────────────────")
		fmt.Printf("❓ Question: %s\n\n", white(question))

		// Load NRG
		sqliteStore, err := graph.NewSQLiteStorage(dbPath)
		if err != nil {
			return fmt.Errorf("cannot open NRG database. Run 'codemri scan .' first: %w", err)
		}
		defer sqliteStore.Close()

		nrg, err := sqliteStore.LoadGraph(cwd)
		if err != nil {
			return fmt.Errorf("failed to load NRG: %w", err)
		}

		fmt.Printf("📊 NRG Loaded: %s nodes, %s edges\n\n", green(fmt.Sprintf("%d", len(nrg.Nodes))), green(fmt.Sprintf("%d", len(nrg.Edges))))

		// Extract search keywords from the question using cortex NLP syntax rules
		keywords := cortex.ExtractKeywords(question)
		fmt.Printf("🔍 Search Keywords: %s\n\n", yellow(strings.Join(keywords, ", ")))

		// Invoke Cortex conversational AI engine (Freemodel.dev Online or Local Offline Structural)
		chatResp, chatErr := cortex.GenerateChatReply(cwd, question)
		if chatErr == nil && chatResp != nil && len(chatResp.Reply) > 0 {
			fmt.Println(cyan("💡 CORTEX CONVERSATIONAL INTELLIGENCE (" + chatResp.Engine + ")"))
			fmt.Println("═══════════════════════════════════════════════════")
			fmt.Println(white(chatResp.Reply))
			fmt.Println("═══════════════════════════════════════════════════\n")
		}

		// Search NRG for relevant nodes
		relevantNodes := searchNRG(nrg, keywords)

		if len(relevantNodes) == 0 {
			if chatResp == nil || len(chatResp.Reply) == 0 {
				fmt.Println(yellow("⚠ No directly matching nodes found in the NRG for your query."))
				fmt.Println("  Try more specific terms, or run 'codemri graph' to see available symbols.")
				return nil
			}
			// General architectural query answered conversationally
			fmt.Println("═══════════════════════════════════════════════════")
			if cortex.OfflineMode || os.Getenv("CODEMRI_OFFLINE") == "1" || os.Getenv("CODEMRI_OFFLINE") == "true" {
				fmt.Printf("💡 %s: This analysis is generated entirely offline using your local NRG database.\n", cyan("Cortex"))
				fmt.Println("   No source code was transmitted to any external server (ADR-0002).")
			} else {
				fmt.Printf("💡 %s: Hybrid intelligence powered by freemodel.dev cloud models + local SQLite NRG.\n", cyan("Cortex"))
			}
			return nil
		}

		// Generate structural explanation
		fmt.Println(cyan("🔍 CORTEX STRUCTURAL SYMBOL TOPOLOGY"))
		fmt.Println("───────────────────────────────────────────────────")

		// Group by type
		typeGroups := make(map[string][]*graph.Node)
		for _, n := range relevantNodes {
			typeGroups[string(n.Type)] = append(typeGroups[string(n.Type)], n)
		}

		for nodeType, nodes := range typeGroups {
			fmt.Printf("\n  📦 %s (%d found):\n", white(strings.ToUpper(nodeType)), len(nodes))
			limit := 8
			for i, n := range nodes {
				if i >= limit {
					fmt.Printf("     ... and %d more\n", len(nodes)-limit)
					break
				}
				// Find connections
				inCount, outCount := countConnections(nrg, n.ID)
				connInfo := ""
				if inCount > 0 || outCount > 0 {
					connInfo = fmt.Sprintf(" [↑%d callers, ↓%d deps]", inCount, outCount)
				}
				fmt.Printf("     • %s %s:%d-%d%s\n",
					cyan(n.Name),
					n.Path, n.StartLine, n.EndLine,
					green(connInfo),
				)
			}
		}

		// Generate architectural insight
		fmt.Println("\n" + cyan("🧬 ARCHITECTURAL INSIGHT"))
		fmt.Println("───────────────────────────────────────────────────")
		generateInsight(nrg, relevantNodes, keywords, question)

		// Load pulse data if available for health context
		pulsePath := filepath.Join(codemriDir, "pulse.json")
		if data, err := os.ReadFile(pulsePath); err == nil {
			var pulse map[string]interface{}
			if json.Unmarshal(data, &pulse) == nil {
				if health, ok := pulse["health"].(map[string]interface{}); ok {
					if score, ok := health["overall_score"].(float64); ok {
						fmt.Printf("\n  💓 Repository Health: %s/100\n", green(fmt.Sprintf("%.0f", score)))
					}
				}
			}
		}

		fmt.Println("\n═══════════════════════════════════════════════════")
		if cortex.OfflineMode || os.Getenv("CODEMRI_OFFLINE") == "1" || os.Getenv("CODEMRI_OFFLINE") == "true" {
			fmt.Printf("💡 %s: This analysis is generated entirely offline using your local NRG database.\n", cyan("Cortex"))
			fmt.Println("   No source code was transmitted to any external server (ADR-0002).")
		} else {
			fmt.Printf("💡 %s: Hybrid intelligence powered by freemodel.dev cloud models + local SQLite NRG.\n", cyan("Cortex"))
		}

		return nil
	},
}

// searchNRG finds nodes matching any of the given keywords.
func searchNRG(nrg *graph.NeuralRepositoryGraph, keywords []string) []*graph.Node {
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
				score += 10
			}
			if strings.Contains(pathLower, kw) {
				score += 5
			}
		}

		if score > 0 {
			scored = append(scored, scoredNode{Node: node, Score: score})
		}
	}

	// Sort by relevance score
	sort.Slice(scored, func(i, j int) bool {
		return scored[i].Score > scored[j].Score
	})

	result := make([]*graph.Node, 0)
	limit := 30
	for i, s := range scored {
		if i >= limit {
			break
		}
		result = append(result, s.Node)
	}
	return result
}

// countConnections counts incoming and outgoing edges for a node.
func countConnections(nrg *graph.NeuralRepositoryGraph, nodeID string) (int, int) {
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

// generateInsight produces a text-based architectural summary of matching nodes.
func generateInsight(nrg *graph.NeuralRepositoryGraph, nodes []*graph.Node, keywords []string, question string) {
	// Identify unique packages/directories
	packages := make(map[string]int)
	for _, n := range nodes {
		dir := filepath.Dir(n.Path)
		packages[dir]++
	}

	fmt.Printf("  The query \"%s\" maps to %d structural symbols across %d packages:\n\n",
		question, len(nodes), len(packages))

	for pkg, count := range packages {
		fmt.Printf("    📁 %s — %d relevant symbols\n", pkg, count)
	}

	// Find the most connected node
	var mostConnected *graph.Node
	maxConn := 0
	for _, n := range nodes {
		in, out := countConnections(nrg, n.ID)
		total := in + out
		if total > maxConn {
			maxConn = total
			mostConnected = n
		}
	}

	if mostConnected != nil && maxConn > 0 {
		in, out := countConnections(nrg, mostConnected.ID)
		fmt.Printf("\n  🎯 Central Node: %s (%s)\n", mostConnected.Name, mostConnected.Type)
		fmt.Printf("     ↑ %d upstream callers | ↓ %d downstream dependencies\n", in, out)
		fmt.Printf("     📂 Located at: %s (lines %d-%d)\n", mostConnected.Path, mostConnected.StartLine, mostConnected.EndLine)

		if in > 5 {
			fmt.Printf("     ⚠ This is a critical architectural hub. Modifying it may cascade across %d dependents.\n", in)
		}
	}
}

func init() {
	rootCmd.AddCommand(explainCmd)
}
