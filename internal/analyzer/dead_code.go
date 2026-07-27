package analyzer

import (
	"strings"

	"github.com/KangBasrengg/MRI-Code/internal/graph"
)

// DetectDeadCode evaluates the Neural Repository Graph for functions and classes that have zero
// incoming usage edges (never invoked, imported, or called by any other discovered module).
func DetectDeadCode(nrg *graph.NeuralRepositoryGraph) []DeadCodeIssue {
	var issues []DeadCodeIssue

	// Map to track usage count for every Node ID
	incomingCount := make(map[string]int)
	for _, edge := range nrg.Edges {
		if edge.Type == graph.EdgeCalls || edge.Type == graph.EdgeImports || edge.Type == graph.EdgeDependsOn || edge.Type == graph.EdgeQueries {
			incomingCount[edge.TargetID]++
		}
	}

	for _, node := range nrg.Nodes {
		// Only analyze executable code symbols like Functions and Classes/Structs
		if node.Type != graph.NodeFunction && node.Type != graph.NodeClass {
			continue
		}

		// Exclude standard programmatic entrypoints, exported interfaces, and framework lifecycle methods
		if isExemptSymbol(node.Name) {
			continue
		}

		// Check if symbol receives zero incoming invocations or references
		if incomingCount[node.ID] == 0 {
			reason := "Private/Internal symbol receives 0 incoming calls within the scanned workspace."
			if node.Type == graph.NodeClass {
				reason = "Internal Class/Struct definition is declared but never instantiated in analyzed paths."
			}

			issues = append(issues, DeadCodeIssue{
				NodeID:    node.ID,
				Name:      node.Name,
				Type:      node.Type,
				Path:      node.Path,
				StartLine: node.StartLine,
				EndLine:   node.EndLine,
				Reason:    reason,
			})
		}
	}

	return issues
}

func isExemptSymbol(name string) bool {
	if len(name) == 0 {
		return true
	}

	// 1. In Go, TypeScript, and Python libraries, uppercase symbols are exported public APIs
	// Check if symbol or method name starts with an uppercase letter
	parts := strings.Split(name, ".")
	lastName := parts[len(parts)-1]
	if len(lastName) > 0 && lastName[0] >= 'A' && lastName[0] <= 'Z' {
		return true
	}

	lowerName := strings.ToLower(name)

	// 2. Standard language entry points, initializers, and test suites
	exemptions := []string{
		"main", "init", "app", "render", "setup", "teardown",
		"index", "root", "handler", "server", "execute", "run",
	}

	for _, ex := range exemptions {
		if lowerName == ex || strings.HasPrefix(lowerName, "test") || strings.HasSuffix(lowerName, "test") || strings.HasSuffix(lowerName, "handler") {
			return true
		}
	}
	return false
}
