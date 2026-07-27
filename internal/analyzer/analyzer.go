package analyzer

import (
	"time"

	"github.com/KangBasrengg/MRI-Code/internal/graph"
)

// DeadCodeIssue represents a structural symbol in the Neural Repository Graph (NRG)
// that receives zero incoming relational usage bonds (never called or imported).
type DeadCodeIssue struct {
	NodeID    string           `json:"node_id"`
	Name      string           `json:"name"`
	Type      graph.NodeType   `json:"type"`
	Path      string           `json:"path"`
	StartLine int              `json:"start_line"`
	EndLine   int              `json:"end_line"`
	Reason    string           `json:"reason"`
}

// CircularDependency represents a detected cyclic loop among symbols or files.
type CircularDependency struct {
	LoopID      string   `json:"loop_id"`
	Chain       []string `json:"chain"` // List of file paths or symbol IDs forming the cycle
	EdgeTypes   []string `json:"edge_types"`
	Description string   `json:"description"`
}

// ComplexityMetric measures the structural complexity score of specific files or components.
type ComplexityMetric struct {
	TargetID    string  `json:"target_id"`
	Path        string  `json:"path"`
	SymbolCount int     `json:"symbol_count"`
	Score       float64 `json:"score"`
	Rating      string  `json:"rating"` // "Low", "Moderate", "High", "Extreme"
}

// HealthScore represents the automated evaluation grade of the codebase.
type HealthScore struct {
	OverallScore  int      `json:"overall_score"`  // 0 to 100
	Grade         string   `json:"grade"`          // "A+", "A", "B", "C", "D", "F"
	DebtStatus    string   `json:"debt_status"`    // e.g., "Negligible Debt", "Moderate Debt Required Action"
	CommentRatio  float64  `json:"comment_ratio"`  // Percentage of comments vs LOC
	DeadCodeRate  float64  `json:"dead_code_rate"` // Percentage of isolated symbols
	CircularCount int      `json:"circular_count"` // Total cyclic loops found
	Suggestions   []string `json:"suggestions"`    // Prioritized architectural improvements
}

// PulseReport is the single source of truth report produced by the Pulse Analysis Engine (v0.4.0).
type PulseReport struct {
	Version             string               `json:"version"`
	Codename            string               `json:"codename"`
	AnalyzedAt          time.Time            `json:"analyzed_at"`
	ExecutionTimeMs     float64              `json:"execution_time_ms"`
	Health              HealthScore          `json:"health"`
	DeadCodeIssues      []DeadCodeIssue      `json:"dead_code_issues"`
	CircularDeps        []CircularDependency `json:"circular_dependencies"`
	ComplexityMetrics   []ComplexityMetric   `json:"complexity_metrics"`
	TotalSymbolsAnalyzed int                 `json:"total_symbols_analyzed"`
	TotalEdgesAnalyzed  int                  `json:"total_edges_analyzed"`
}

// Analyze executes the complete suite of structural analyses on an in-memory or database NRG.
func Analyze(nrg *graph.NeuralRepositoryGraph, totalLOC int, totalComments int) *PulseReport {
	start := time.Now()

	report := &PulseReport{
		Version:              "v0.4.0",
		Codename:             "Pulse",
		AnalyzedAt:           time.Now(),
		TotalSymbolsAnalyzed: len(nrg.Nodes),
		TotalEdgesAnalyzed:   len(nrg.Edges),
	}

	// 1. Detect Dead Code / Isolated Unused Symbols
	report.DeadCodeIssues = DetectDeadCode(nrg)

	// 2. Detect Circular Dependency Chains
	report.CircularDeps = DetectCircularDependencies(nrg)

	// 3. Evaluate Cyclomatic & Structural Complexity
	report.ComplexityMetrics = EvaluateComplexity(nrg)

	// 4. Synthesize Overall Repository Health Score & Actionable Advice
	report.Health = CalculateHealthScore(
		len(nrg.Nodes),
		len(report.DeadCodeIssues),
		len(report.CircularDeps),
		totalLOC,
		totalComments,
		report.ComplexityMetrics,
	)

	report.ExecutionTimeMs = float64(time.Since(start).Microseconds()) / 1000.0
	return report
}
