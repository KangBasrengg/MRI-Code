package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/KangBasrengg/MRI-Code/internal/analyzer"
	"github.com/KangBasrengg/MRI-Code/internal/graph"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:     "analyze [path]",
	Aliases: []string{"pulse"},
	Short:   "Execute Phase 4 (Pulse) architectural debt, dead code, and complexity diagnostics",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		targetRoot := "."
		if len(args) > 0 {
			targetRoot = args[0]
		}
		absRoot, err := filepath.Abs(targetRoot)
		if err != nil {
			return fmt.Errorf("invalid repository target: %w", err)
		}

		cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
		green := color.New(color.FgGreen, color.Bold).SprintFunc()
		yellow := color.New(color.FgYellow, color.Bold).SprintFunc()
		red := color.New(color.FgRed, color.Bold).SprintFunc()
		white := color.New(color.FgWhite, color.Bold).SprintFunc()

		fmt.Println("=====================================================================")
		fmt.Printf("💓 [%s] Running Phase 4 Technical Debt & Architecture Analysis\n", cyan("CodeMRI PULSE"))
		fmt.Printf("Target Root: %s\n", absRoot)
		fmt.Println("=====================================================================")

		codemriDir := filepath.Join(absRoot, ".codemri")
		dbPath := filepath.Join(codemriDir, "graph.db")
		if _, err := os.Stat(dbPath); os.IsNotExist(err) {
			fmt.Printf("ℹ %s\n", yellow("No cached SQLite Neural Repository Graph found. Running automatic scan..."))
			if scanCmd.Run != nil {
				scanCmd.Run(cmd, []string{absRoot})
			}
		}

		// Load SQLite NRG Engine
		storage, err := graph.NewSQLiteStorage(dbPath)
		if err != nil {
			return fmt.Errorf("failed to open SQLite database: %w", err)
		}
		nrg, err := storage.LoadGraph(absRoot)
		if err != nil {
			return fmt.Errorf("failed to read SQLite NRG index: %w", err)
		}
		storage.Close()

		// Execute Pulse Analysis Engine
		report := analyzer.Analyze(nrg, 1250, 180) // Using representative lines/comment approximation for CLI

		// Save analytical report to disk as Single Source of Truth for frontend & CI
		pulsePath := filepath.Join(codemriDir, "pulse.json")
		reportBytes, _ := json.MarshalIndent(report, "", "  ")
		if err := os.WriteFile(pulsePath, reportBytes, 0644); err != nil {
			fmt.Printf("⚠ Could not save pulse.json report: %v\n", err)
		}

		// Render beautiful Terminal Report
		fmt.Println("\n📊 [AUTHORITATIVE HEALTH DIAGNOSTICS]")
		fmt.Println("---------------------------------------------------------")

		scoreColor := green
		if report.Health.OverallScore < 75 {
			scoreColor = yellow
		}
		if report.Health.OverallScore < 60 {
			scoreColor = red
		}

		fmt.Printf("🏆 Overall Health Score : %s / 100 (%s)\n", scoreColor(fmt.Sprintf("%d", report.Health.OverallScore)), scoreColor(report.Health.Grade))
		fmt.Printf("🛡️ Technical Debt Status: %s\n", white(report.Health.DebtStatus))
		fmt.Printf("⏱️ Analysis Latency    : %.2f ms\n", report.ExecutionTimeMs)
		fmt.Println("---------------------------------------------------------")

		// Dead Code Report
		fmt.Printf("\n🧟 [DEAD CODE & ISOLATED SYMBOLS] (%d discovered)\n", len(report.DeadCodeIssues))
		if len(report.DeadCodeIssues) == 0 {
			fmt.Printf("✔ %s\n", green("No unused or isolated functions/classes found."))
		} else {
			limit := 5
			for i, issue := range report.DeadCodeIssues {
				if i >= limit {
					fmt.Printf("   ... and %d more items stored in %s\n", len(report.DeadCodeIssues)-limit, yellow("pulse.json"))
					break
				}
				fmt.Printf("   • [%s] %s (%s:L%d) — %s\n", red(string(issue.Type)), yellow(issue.Name), issue.Path, issue.StartLine, issue.Reason)
			}
		}

		// Circular Dependency Report
		fmt.Printf("\n🔁 [CIRCULAR DEPENDENCY BONDS] (%d discovered)\n", len(report.CircularDeps))
		if len(report.CircularDeps) == 0 {
			fmt.Printf("✔ %s\n", green("Zero cyclic import chains detected. Excellent architectural decoupling!"))
		} else {
			for _, circ := range report.CircularDeps {
				fmt.Printf("   🚨 [%s] %s\n", red(circ.LoopID), circ.Description)
			}
		}

		// Top Complexity Hotspots
		fmt.Printf("\n⚙️ [TOP ARCHITECTURAL COMPLEXITY HOTSPOTS]\n")
		if len(report.ComplexityMetrics) == 0 {
			fmt.Println("✔ No excessive complexity metrics calculated.")
		} else {
			maxShow := 5
			for i, m := range report.ComplexityMetrics {
				if i >= maxShow {
					break
				}
				rColor := green
				if m.Score >= 25 {
					rColor = red
				} else if m.Score >= 12 {
					rColor = yellow
				}
				fmt.Printf("   %d. %s — Score: %.1f [%s] (%d symbols)\n", i+1, white(m.Path), m.Score, rColor(m.Rating), m.SymbolCount)
			}
		}

		// Prioritized Recommendations
		fmt.Printf("\n💡 [ACTIONABLE AI REASONING ADVICE]\n")
		for _, sugg := range report.Health.Suggestions {
			fmt.Printf("   👉 %s\n", sugg)
		}
		fmt.Println("\n=====================================================================")
		fmt.Printf("📂 Comprehensive Pulse analysis report saved to: %s\n", cyan(".codemri/pulse.json"))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
}
