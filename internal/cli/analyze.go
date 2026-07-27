package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/KangBasrengg/MRI-Code/internal/analyzer"
	"github.com/KangBasrengg/MRI-Code/internal/graph"
	"github.com/KangBasrengg/MRI-Code/internal/performance"
	"github.com/KangBasrengg/MRI-Code/internal/remote"
	"github.com/KangBasrengg/MRI-Code/internal/security"
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
		absRoot, err := remote.ResolveTarget(targetRoot)
		if err != nil {
			return fmt.Errorf("invalid repository target or failed remote clone: %w", err)
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

		// ── Phase 7 Shield: Security Intelligence ──
		fmt.Println("\n🛡️ [SHIELD — SECURITY INTELLIGENCE SCAN]")
		fmt.Println("---------------------------------------------------------")
		secReport, secErr := security.RunSecurityScan(absRoot)
		if secErr != nil {
			fmt.Printf("   ⚠ Security scan error: %v\n", secErr)
		} else {
			secPath := filepath.Join(codemriDir, "security.json")
			security.SaveReport(secReport, secPath)
			gradeColor := green
			if secReport.Critical > 0 || secReport.High > 3 {
				gradeColor = red
			} else if secReport.High > 0 || secReport.Medium > 3 {
				gradeColor = yellow
			}
			fmt.Printf("   📊 Security Grade: %s\n", gradeColor(secReport.SecurityGrade))
			fmt.Printf("   🔍 Files Scanned: %d | Findings: %d\n", secReport.TotalFiles, secReport.TotalFindings)
			if secReport.Critical > 0 {
				fmt.Printf("   🚨 Critical: %s\n", red(fmt.Sprintf("%d", secReport.Critical)))
			}
			if secReport.High > 0 {
				fmt.Printf("   ⚠️  High:     %s\n", yellow(fmt.Sprintf("%d", secReport.High)))
			}
			if secReport.Medium > 0 {
				fmt.Printf("   ℹ️  Medium:   %d\n", secReport.Medium)
			}
			if secReport.TotalFindings == 0 {
				fmt.Printf("   ✔ %s\n", green("No security vulnerabilities detected. Excellent!"))
			} else {
				limit := 3
				for i, f := range secReport.Findings {
					if i >= limit {
						fmt.Printf("   ... and %d more in %s\n", secReport.TotalFindings-limit, yellow("security.json"))
						break
					}
					fmt.Printf("   • [%s] %s — %s:%d\n", red(string(f.Severity)), f.Description, f.FilePath, f.LineNumber)
				}
			}
			fmt.Printf("   📂 Full report saved to: %s\n", cyan(".codemri/security.json"))
		}

		// ── Phase 8 Velocity: Performance Intelligence ──
		fmt.Println("\n🚀 [VELOCITY — PERFORMANCE INTELLIGENCE SCAN]")
		fmt.Println("---------------------------------------------------------")
		perfReport, perfErr := performance.RunPerformanceScan(absRoot)
		if perfErr != nil {
			fmt.Printf("   ⚠ Performance scan error: %v\n", perfErr)
		} else {
			perfPath := filepath.Join(codemriDir, "performance.json")
			performance.SaveReport(perfReport, perfPath)
			fmt.Printf("   📊 Performance Grade: %s\n", green(perfReport.PerformanceGrade))
			fmt.Printf("   📦 Total Source Size: %s across %d files\n", yellow(humanizeBytes(perfReport.TotalSizeBytes)), perfReport.TotalFiles)
			if len(perfReport.LargeFiles) > 0 {
				fmt.Printf("   📏 Large Files: %d files exceed 50KB\n", len(perfReport.LargeFiles))
			}
			if len(perfReport.HeavyImports) > 0 {
				fmt.Printf("   📥 Heavy Imports: %d files with >15 imports\n", len(perfReport.HeavyImports))
			}
			for _, s := range perfReport.Suggestions {
				fmt.Printf("   💡 %s\n", s)
			}
			fmt.Printf("   📂 Full report saved to: %s\n", cyan(".codemri/performance.json"))
		}

		fmt.Println("\n=====================================================================")
		fmt.Printf("📂 All analysis reports saved to: %s\n", cyan(".codemri/"))
		return nil
	},
}

func humanizeBytes(b int64) string {
	switch {
	case b >= 1024*1024:
		return fmt.Sprintf("%.1f MB", float64(b)/(1024*1024))
	case b >= 1024:
		return fmt.Sprintf("%.1f KB", float64(b)/1024)
	default:
		return fmt.Sprintf("%d B", b)
	}
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
}
