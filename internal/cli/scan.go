package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/KangBasrengg/MRI-Code/internal/analyzer"
	"github.com/KangBasrengg/MRI-Code/internal/graph"
	"github.com/KangBasrengg/MRI-Code/internal/scanner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type repositoryMeta struct {
	Platform       string                             `json:"platform"`
	Version        string                             `json:"version"`
	Codename       string                             `json:"codename"`
	ScannedRoot    string                             `json:"scanned_root"`
	Timestamp      time.Time                          `json:"timestamp"`
	Status         string                             `json:"status"`
	TotalFiles     int                                `json:"total_files"`
	TotalBytes     int64                              `json:"total_bytes"`
	TotalLOC       int                                `json:"total_loc"`
	TotalComments  int                                `json:"total_comments"`
	TotalBlank     int                                `json:"total_blank"`
	GraphNodeCount int                                `json:"graph_node_count"`
	GraphEdgeCount int                                `json:"graph_edge_count"`
	LanguageStats  map[string]*scanner.LanguageSummary `json:"language_stats"`
	ScanDuration   float64                            `json:"scan_duration_sec"`
}

var scanCmd = &cobra.Command{
	Use:   "scan [repository_path]",
	Short: "Scan repository, detect languages, run AST parsers, and compile Neural Repository Graph (NRG)",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
		green := color.New(color.FgGreen, color.Bold).SprintFunc()
		yellow := color.New(color.FgYellow, color.Bold).SprintFunc()
		magenta := color.New(color.FgMagenta, color.Bold).SprintFunc()

		targetDir := "."
		if len(args) > 0 {
			targetDir = args[0]
		}

		absPath, err := filepath.Abs(targetDir)
		if err != nil {
			fmt.Printf("Error resolving path: %v\n", err)
			return
		}

		fmt.Printf("%s Starting CodeMRI v1.0.0 (MRI) Repository Intelligence Scan on: %s\n", cyan("📡 [CodeMRI]"), yellow(absPath))
		fmt.Println("⚡ Running syntax parsers & indexing SQLite relational graph...")
		
		startTime := time.Now()
		ctx := context.Background()

		// Step 1: Initialize high-speed concurrent Atlas Engine
		engine := scanner.NewAtlasEngine()
		cfg := scanner.ScanConfig{
			RootPath:      absPath,
			MaxWorkers:    8,
			IncludeHidden: false,
		}

		res, err := engine.ExecuteScan(ctx, cfg)
		if err != nil {
			fmt.Printf("❌ Scan execution failed: %v\n", err)
			return
		}

		// Step 2: Assemble Neural Repository Graph (NRG) - Single Source of Truth
		nrg := graph.NewNRG(absPath, Version)
		for _, ast := range res.ParsedASTs {
			for _, node := range ast.Nodes {
				nrg.Nodes[node.ID] = node
			}
			nrg.Edges = append(nrg.Edges, ast.Edges...)
		}

		// Step 3: Ensure .codemri repository workspace exists
		dotDir := filepath.Join(absPath, ".codemri")
		if err := os.MkdirAll(dotDir, 0755); err != nil {
			fmt.Printf("❌ Failed to initialize local workspace (.codemri): %v\n", err)
			return
		}

		// Step 4: Write compiled NRG Graph DB equivalent to JSON (human-readable backup)
		graphBytes, _ := json.MarshalIndent(nrg, "", "  ")
		graphPath := filepath.Join(dotDir, "graph.json")
		if err := os.WriteFile(graphPath, graphBytes, 0644); err != nil {
			fmt.Printf("❌ Failed to save compiled Neural Repository Graph JSON: %v\n", err)
			return
		}

		// Step 4b: Persist NRG directly into SQLite database for microsecond relational queries (Phase 03 Neuron)
		dbPath := filepath.Join(dotDir, "graph.db")
		sqliteStore, err := graph.NewSQLiteStorage(dbPath)
		if err != nil {
			fmt.Printf("❌ Failed to open SQLite storage engine: %v\n", err)
			return
		}
		if err := sqliteStore.SaveGraph(nrg); err != nil {
			fmt.Printf("❌ Failed to persist NRG in SQLite database: %v\n", err)
			sqliteStore.Close()
			return
		}
		sqliteStore.Close()

		// Step 5: Format Language Stats for representation
		langStats := make(map[string]*scanner.LanguageSummary)
		for lang, sum := range res.LanguageStats {
			langStats[string(lang)] = sum
		}

		// Step 6: Save rich repository analytical metadata
		meta := repositoryMeta{
			Platform:       "CodeMRI Neural Repository Intelligence",
			Version:        Version,
			Codename:       Codename,
			ScannedRoot:    absPath,
			Timestamp:      time.Now(),
			Status:         "V1_0_0_MRI_COMPLETE (SQLite NRG + Technical Debt & Intelligence Engine Active)",
			TotalFiles:     res.TotalFiles,
			TotalBytes:     res.TotalBytes,
			TotalLOC:       res.TotalLOC,
			TotalComments:  res.TotalComments,
			TotalBlank:     res.TotalBlank,
			GraphNodeCount: len(nrg.Nodes),
			GraphEdgeCount: len(nrg.Edges),
			LanguageStats:  langStats,
			ScanDuration:   res.DurationSec,
		}

		metaBytes, _ := json.MarshalIndent(meta, "", "  ")
		metaPath := filepath.Join(dotDir, "repository.json")
		if err := os.WriteFile(metaPath, metaBytes, 0644); err != nil {
			fmt.Printf("❌ Failed to save analytical metadata: %v\n", err)
			return
		}

		// Step 7: Perform Phase 4 Pulse Architectural Technical Debt & Health Eval
		pulseReport := analyzer.Analyze(nrg, res.TotalLOC, res.TotalComments)
		pulseBytes, _ := json.MarshalIndent(pulseReport, "", "  ")
		pulsePath := filepath.Join(dotDir, "pulse.json")
		_ = os.WriteFile(pulsePath, pulseBytes, 0644)

		elapsed := time.Since(startTime)

		// Print comprehensive analytical dashboard summary
		fmt.Println("\n---------------------------------------------------------")
		fmt.Printf("%s AST, SQLite Relational Engine & Pulse Analytics Completed in %v\n", green("✔ [SUCCESS]"), elapsed)
		fmt.Printf("📂 Files Scanned : %s | 📄 LOC: %s | 💬 Comments: %d\n", yellow(fmt.Sprintf("%d", res.TotalFiles)), yellow(fmt.Sprintf("%d", res.TotalLOC)), res.TotalComments)
		fmt.Printf("🧠 Compiled NRG  : %s Nodes | %s Relational Edges\n", magenta(fmt.Sprintf("%d", len(nrg.Nodes))), magenta(fmt.Sprintf("%d", len(nrg.Edges))))
		fmt.Printf("💓 Pulse Health  : %s/100 (Grade %s) | %s\n", green(fmt.Sprintf("%d", pulseReport.Health.OverallScore)), green(pulseReport.Health.Grade), color.New(color.FgCyan).SprintFunc()(pulseReport.Health.DebtStatus))
		
		if len(res.LanguageStats) > 0 {
			fmt.Println("\n📊 Language Distribution:")
			for lang, stat := range res.LanguageStats {
				fmt.Printf("   ▪ %-12s : %6d lines (%.1f%%)\n", strings.ToUpper(string(lang)), stat.Lines, stat.Percentage)
			}
		}

		fmt.Println("---------------------------------------------------------")
		fmt.Printf("✔ SQLite relational database stored at: %s\n", green(dbPath))
		fmt.Printf("✔ Human-readable graph JSON backup at:  %s\n", green(graphPath))
		fmt.Printf("✔ Repository analytics saved at:         %s\n", green(metaPath))
		fmt.Printf("🚀 Run %s or type one word %s to launch interactive graphical dashboard!\n", cyan("codemri serve"), cyan("codemri"))
	},
}
