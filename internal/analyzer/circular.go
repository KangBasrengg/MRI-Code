package analyzer

import (
	"fmt"
	"sort"
	"strings"

	"github.com/KangBasrengg/MRI-Code/internal/graph"
)

// DetectCircularDependencies executes Depth-First Search (DFS) with recursion stack state tracking
// across Neural Repository Graph edges to detect circular package and file import dependencies.
func DetectCircularDependencies(nrg *graph.NeuralRepositoryGraph) []CircularDependency {
	var cycles []CircularDependency

	// 1. Build directed adjacency map for dependency edges
	adj := make(map[string][]string)
	edgeTypes := make(map[string]string)

	for _, edge := range nrg.Edges {
		if edge.Type == graph.EdgeImports || edge.Type == graph.EdgeDependsOn {
			adj[edge.SourceID] = append(adj[edge.SourceID], edge.TargetID)
			edgeKey := edge.SourceID + "->" + edge.TargetID
			edgeTypes[edgeKey] = string(edge.Type)
		}
	}

	// State tracking: 0 = Unvisited, 1 = Visiting (in recursion stack), 2 = Fully Visited
	state := make(map[string]int)
	var stack []string
	seenLoops := make(map[string]bool)

	var dfs func(u string)
	dfs = func(u string) {
		state[u] = 1
		stack = append(stack, u)

		for _, v := range adj[u] {
			if state[v] == 1 {
				// We encountered a back-edge pointing to an ancestor currently in the DFS stack -> Circular loop found!
				loopStartIndex := -1
				for i, nodeID := range stack {
					if nodeID == v {
						loopStartIndex = i
						break
					}
				}

				if loopStartIndex != -1 {
					rawLoop := append([]string{}, stack[loopStartIndex:]...)
					rawLoop = append(rawLoop, v) // Complete the cycle representation

					// Convert Node IDs to human-readable symbol names or file paths
					var readableChain []string
					var usedEdgeTypes []string
					for k := 0; k < len(rawLoop); k++ {
						id := rawLoop[k]
						if node, exists := nrg.Nodes[id]; exists {
							if node.Path != "" {
								readableChain = append(readableChain, fmt.Sprintf("%s (%s)", node.Name, node.Path))
							} else {
								readableChain = append(readableChain, node.Name)
							}
						} else {
							readableChain = append(readableChain, id)
						}

						if k > 0 {
							eKey := rawLoop[k-1] + "->" + rawLoop[k]
							if t, ok := edgeTypes[eKey]; ok {
								usedEdgeTypes = append(usedEdgeTypes, t)
							} else {
								usedEdgeTypes = append(usedEdgeTypes, "IMPORTS")
							}
						}
					}

					// Deduplicate equivalent cycle rotations using sorted signature
					normChain := append([]string{}, rawLoop[:len(rawLoop)-1]...)
					sort.Strings(normChain)
					signature := strings.Join(normChain, "|")

					if !seenLoops[signature] && len(rawLoop) > 1 {
						seenLoops[signature] = true
						loopID := fmt.Sprintf("CIRC_DEP_%d", len(cycles)+1)
						desc := fmt.Sprintf("Circular architectural binding discovered across %d symbols: %s", len(rawLoop)-1, strings.Join(readableChain, " ➔ "))

						cycles = append(cycles, CircularDependency{
							LoopID:      loopID,
							Chain:       readableChain,
							EdgeTypes:   usedEdgeTypes,
							Description: desc,
						})
					}
				}
			} else if state[v] == 0 {
				dfs(v)
			}
		}

		stack = stack[:len(stack)-1]
		state[u] = 2
	}

	// Initiate DFS across all discovered nodes in NRG
	for id := range nrg.Nodes {
		if state[id] == 0 {
			dfs(id)
		}
	}

	return cycles
}
