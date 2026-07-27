package analyzer

import (
	"math"
	"sort"

	"github.com/KangBasrengg/MRI-Code/internal/graph"
)

// EvaluateComplexity assesses the architectural density and structural cyclomatic weight
// of individual source files and modules registered in the Neural Repository Graph.
func EvaluateComplexity(nrg *graph.NeuralRepositoryGraph) []ComplexityMetric {
	var metrics []ComplexityMetric

	// Map file paths to symbol counts and structural edge activity
	fileSymbols := make(map[string]int)
	fileLines := make(map[string]int)
	fileEdges := make(map[string]int)

	// Populate node statistics per file path
	for _, node := range nrg.Nodes {
		if node.Path == "" || node.Path == "." {
			continue
		}
		if node.Type != graph.NodeFile {
			fileSymbols[node.Path]++
		} else {
			lines := node.EndLine - node.StartLine
			if lines > fileLines[node.Path] {
				fileLines[node.Path] = lines
			}
		}
	}

	// Associate relational edge weight to corresponding source paths
	for _, edge := range nrg.Edges {
		if sourceNode, exists := nrg.Nodes[edge.SourceID]; exists && sourceNode.Path != "" {
			fileEdges[sourceNode.Path]++
		}
	}

	// Calculate comprehensive score for every tracked path
	for path, symCount := range fileSymbols {
		lines := fileLines[path]
		edges := fileEdges[path]

		// Structural complexity heuristic weight calculation
		// Formula combines function/class density, branching dependencies, and total volume
		score := (float64(symCount) * 1.5) + (float64(edges) * 1.2) + (float64(lines) / 40.0)
		score = math.Round(score*10) / 10

		rating := "Low (Clean & Maintainable)"
		if score >= 45.0 {
			rating = "Extreme (High Technical Debt & Maintenance Risk)"
		} else if score >= 25.0 {
			rating = "High (Consider Refactoring & Decoupling)"
		} else if score >= 12.0 {
			rating = "Moderate (Healthy Structure)"
		}

		metrics = append(metrics, ComplexityMetric{
			TargetID:    "PATH_" + path,
			Path:        path,
			SymbolCount: symCount,
			Score:       score,
			Rating:      rating,
		})
	}

	// Sort descending by complexity score so engineers see top hotspots first
	sort.Slice(metrics, func(i, j int) bool {
		return metrics[i].Score > metrics[j].Score
	})

	// Return up to top 50 most complex hotspots to prevent JSON bloat
	if len(metrics) > 50 {
		return metrics[:50]
	}
	return metrics
}
