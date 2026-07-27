package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/KangBasrengg/MRI-Code/internal/graph"
)

var graphCmd = &cobra.Command{
	Use:   "graph [query]",
	Short: "Query the Neural Repository Graph from the terminal",
	Long:  "Inspect NRG nodes, edges, and topology directly from the CLI without launching the dashboard.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()

		cwd, _ := os.Getwd()
		dbPath := filepath.Join(cwd, ".codemri", "graph.db")

		sqliteStore, err := graph.NewSQLiteStorage(dbPath)
		if err != nil {
			fmt.Printf("❌ Cannot open NRG database: %v\n", err)
			fmt.Println("   Run 'codemri scan .' first to build the graph.")
			return
		}
		defer sqliteStore.Close()

		summary, err := sqliteStore.GetTopologySummary()
		if err != nil {
			fmt.Printf("❌ Failed to query topology: %v\n", err)
			return
		}

		fmt.Println(cyan("🧠 Neural Repository Graph (NRG) — Terminal Query Interface"))
		fmt.Println("─────────────────────────────────────────────────────────")

		totalNodes, _ := summary["total_nodes"].(int)
		totalEdges, _ := summary["total_edges"].(int)
		fmt.Printf("  📊 Total Nodes: %s\n", green(fmt.Sprintf("%d", totalNodes)))
		fmt.Printf("  🔗 Total Edges: %s\n", green(fmt.Sprintf("%d", totalEdges)))
		fmt.Printf("  💾 Engine:      %s\n", yellow(summary["storage_engine"]))
		fmt.Println()

		if nodeTypes, ok := summary["node_types"].(map[string]int); ok {
			fmt.Println(cyan("  🏗️  Symbol Topology:"))
			keys := make([]string, 0, len(nodeTypes))
			for k := range nodeTypes {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, k := range keys {
				fmt.Printf("      %-15s %s\n", k, green(fmt.Sprintf("%d", nodeTypes[k])))
			}
		}
		fmt.Println()

		if edgeTypes, ok := summary["edge_types"].(map[string]int); ok {
			fmt.Println(cyan("  ⚡ Relational Edges:"))
			keys := make([]string, 0, len(edgeTypes))
			for k := range edgeTypes {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, k := range keys {
				fmt.Printf("      %-15s %s\n", k, green(fmt.Sprintf("%d", edgeTypes[k])))
			}
		}

		// If a search query was provided, search nodes
		if len(args) > 0 {
			query := strings.ToLower(args[0])
			fmt.Println()
			fmt.Printf(cyan("  🔍 Searching NRG for: \"%s\"\n"), query)
			fmt.Println("  ─────────────────────────────────────────")

			nrg, loadErr := sqliteStore.LoadGraph(cwd)
			if loadErr != nil {
				fmt.Printf("  ⚠ Could not load full graph for search: %v\n", loadErr)
				return
			}

			matches := make([]*graph.Node, 0)
			for _, node := range nrg.Nodes {
				if strings.Contains(strings.ToLower(node.Name), query) ||
					strings.Contains(strings.ToLower(node.Path), query) {
					matches = append(matches, node)
				}
			}

			if len(matches) == 0 {
				fmt.Println("  (No matching nodes found)")
			} else {
				limit := 20
				if len(matches) < limit {
					limit = len(matches)
				}
				fmt.Printf("  Found %s matches (showing top %d):\n\n", green(fmt.Sprintf("%d", len(matches))), limit)
				for i := 0; i < limit; i++ {
					m := matches[i]
					fmt.Printf("    %s %-20s %s %s:%d-%d\n",
						yellow(fmt.Sprintf("[%s]", m.Type)),
						cyan(m.Name),
						"→",
						m.Path, m.StartLine, m.EndLine,
					)
				}
			}
		}

		// Also export summary as JSON if --json flag set
		jsonOutput, _ := cmd.Flags().GetBool("json")
		if jsonOutput {
			data, _ := json.MarshalIndent(summary, "", "  ")
			fmt.Println()
			fmt.Println(string(data))
		}
	},
}

func init() {
	graphCmd.Flags().Bool("json", false, "Output topology summary as JSON")
	rootCmd.AddCommand(graphCmd)
}
