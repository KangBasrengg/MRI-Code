package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/KangBasrengg/MRI-Code/internal/cortex"
	"github.com/KangBasrengg/MRI-Code/internal/graph"
	"github.com/KangBasrengg/MRI-Code/internal/remote"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/cobra"
)

var port string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launch interactive visualization server and architecture dashboard",
	Run: func(cmd *cobra.Command, args []string) {
		cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()

		// 1. Resolve an available TCP port
		activePort := resolveAvailablePort(port)
		if activePort != port {
			fmt.Printf("%s Port %s is already in use. Automatically switched to port %s.\n", yellow("⚠ [PORT CONFLICT]"), yellow(port), green(activePort))
		}

		// 2. Locate the .codemri workspace dynamically
		getDotFile := func(filename string) (string, string) {
			curr, err := os.Getwd()
			if err != nil {
				curr = "."
			}
			return curr, filepath.Join(curr, ".codemri", filename)
		}

		app := fiber.New(fiber.Config{
			AppName:               "CodeMRI v1.0.0 (MRI)",
			DisableStartupMessage: true,
			ReadBufferSize:        65536,
		})

		app.Use(logger.New(logger.Config{
			Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
		}))
		app.Use(cors.New())

		// ─── API: System Status ───
		app.Get("/api/status", func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"platform":        "CodeMRI Repository Intelligence Platform",
				"version":         Version,
				"codename":        Codename,
				"nrg_engine":      "NEURON_SQLITE_ONLINE",
				"pulse_engine":    "PULSE_ANALYTICS_ONLINE",
				"cortex_engine":   "CORTEX_REASONING_ONLINE",
				"shield_engine":   "SHIELD_SECURITY_ONLINE",
				"velocity_engine": "VELOCITY_PERFORMANCE_ONLINE",
				"storage_engine":  "Embedded SQLite Relational Index (.codemri/)",
				"architecture":    "Single Source of Truth (ADR-0001 & ADR-0002)",
			})
		})

		// ─── API: Pulse Analytics Diagnostics ───
		app.Get("/api/pulse", func(c *fiber.Ctx) error {
			_, pulsePath := getDotFile("pulse.json")
			data, err := os.ReadFile(pulsePath)
			if err != nil {
				return c.Status(404).JSON(fiber.Map{"error": "No .codemri/pulse.json found. Run: codemri analyze ."})
			}
			c.Set("Content-Type", "application/json")
			return c.Send(data)
		})

		// ─── API: Shield Security Diagnostics ───
		app.Get("/api/security", func(c *fiber.Ctx) error {
			_, secPath := getDotFile("security.json")
			data, err := os.ReadFile(secPath)
			if err != nil {
				return c.Status(404).JSON(fiber.Map{"error": "No .codemri/security.json found. Run: codemri analyze ."})
			}
			c.Set("Content-Type", "application/json")
			return c.Send(data)
		})

		// ─── API: Velocity Performance Diagnostics ───
		app.Get("/api/performance", func(c *fiber.Ctx) error {
			_, perfPath := getDotFile("performance.json")
			data, err := os.ReadFile(perfPath)
			if err != nil {
				return c.Status(404).JSON(fiber.Map{"error": "No .codemri/performance.json found. Run: codemri analyze ."})
			}
			c.Set("Content-Type", "application/json")
			return c.Send(data)
		})

		// ─── API: Repository Metadata ───
		app.Get("/api/repository", func(c *fiber.Ctx) error {
			_, repoMetaPath := getDotFile("repository.json")
			data, err := os.ReadFile(repoMetaPath)
			if err != nil {
				return c.Status(404).JSON(fiber.Map{"error": "No .codemri/repository.json found. Run: codemri scan ."})
			}
			c.Set("Content-Type", "application/json")
			return c.Send(data)
		})

		// ─── API: NRG Graph Data (Full Graph) ───
		app.Get("/api/graph", func(c *fiber.Ctx) error {
			cwd, dbPath := getDotFile("graph.db")
			_, graphPath := getDotFile("graph.json")
			// In Phase 3 Neuron, attempt to load cleanly from SQLite first
			if sqliteStore, err := graph.NewSQLiteStorage(dbPath); err == nil {
				defer sqliteStore.Close()
				if nrg, loadErr := sqliteStore.LoadGraph(cwd); loadErr == nil {
					return c.JSON(nrg)
				}
			}
			// Fallback to JSON if SQLite DB not found or path mismatches
			data, err := os.ReadFile(graphPath)
			if err != nil {
				return c.Status(404).JSON(fiber.Map{"error": "No .codemri/graph.db or graph.json found. Run: codemri scan ."})
			}
			c.Set("Content-Type", "application/json")
			return c.Send(data)
		})

		// ─── API: NRG Graph Summary (Microsecond SQLite Aggregation) ───
		app.Get("/api/graph/summary", func(c *fiber.Ctx) error {
			_, dbPath := getDotFile("graph.db")
			_, graphPath := getDotFile("graph.json")
			if sqliteStore, err := graph.NewSQLiteStorage(dbPath); err == nil {
				defer sqliteStore.Close()
				if summary, err := sqliteStore.GetTopologySummary(); err == nil {
					return c.JSON(summary)
				}
			}

			// Fallback for JSON
			data, err := os.ReadFile(graphPath)
			if err != nil {
				return c.Status(404).JSON(fiber.Map{"error": "graph not found"})
			}
			var raw map[string]json.RawMessage
			if err := json.Unmarshal(data, &raw); err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "parse error"})
			}

			// Count nodes by type
			var nodes map[string]json.RawMessage
			json.Unmarshal(raw["nodes"], &nodes)
			var edges []json.RawMessage
			json.Unmarshal(raw["edges"], &edges)

			typeCounts := make(map[string]int)
			for _, v := range nodes {
				var n struct{ Type string `json:"type"` }
				json.Unmarshal(v, &n)
				typeCounts[n.Type]++
			}

			return c.JSON(fiber.Map{
				"total_nodes":  len(nodes),
				"total_edges":  len(edges),
				"node_types":   typeCounts,
				"storage_type": "JSON Fallback",
			})
		})

		// ─── API: Relational Graph Querying (Phase 3 Neuron Feature) ───
		app.Get("/api/graph/node/:id", func(c *fiber.Ctx) error {
			_, dbPath := getDotFile("graph.db")
			sqliteStore, err := graph.NewSQLiteStorage(dbPath)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "sqlite database unavailable"})
			}
			defer sqliteStore.Close()

			nodeID := c.Params("id")
			node, err := sqliteStore.FindNodeByID(nodeID)
			if err != nil {
				return c.Status(404).JSON(fiber.Map{"error": err.Error()})
			}
			return c.JSON(node)
		})

		app.Get("/api/graph/neighbors/:id/:edge", func(c *fiber.Ctx) error {
			_, dbPath := getDotFile("graph.db")
			sqliteStore, err := graph.NewSQLiteStorage(dbPath)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "sqlite database unavailable"})
			}
			defer sqliteStore.Close()

			nodeID := c.Params("id")
			edgeType := graph.EdgeType(c.Params("edge"))
			neighbors, err := sqliteStore.FindNeighbors(nodeID, edgeType)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": err.Error()})
			}
			return c.JSON(neighbors)
		})

		// ─── API: Phase 05 ("Vision") Instant Impact Analysis & Subgraph Topology ───
		handleImpact := func(c *fiber.Ctx) error {
			cwd, dbPath := getDotFile("graph.db")
			_, graphPath := getDotFile("graph.json")
			var nrg *graph.NeuralRepositoryGraph
			if sqliteStore, err := graph.NewSQLiteStorage(dbPath); err == nil {
				nrg, _ = sqliteStore.LoadGraph(cwd)
				sqliteStore.Close()
			}
			if nrg == nil {
				data, err := os.ReadFile(graphPath)
				if err != nil {
					return c.Status(404).JSON(fiber.Map{"error": "graph not available for impact analysis"})
				}
				nrg = &graph.NeuralRepositoryGraph{}
				json.Unmarshal(data, nrg)
			}

			targetID := c.Query("id", c.Params("id"))
			targetNode, exists := nrg.Nodes[targetID]
			if !exists {
				return c.Status(404).JSON(fiber.Map{"error": "node not found in Neural Repository Graph"})
			}

			upstreamDependents := make([]*graph.Node, 0)
			downstreamDependencies := make([]*graph.Node, 0)

			for _, edge := range nrg.Edges {
				if edge.TargetID == targetID {
					if uNode, ok := nrg.Nodes[edge.SourceID]; ok {
						upstreamDependents = append(upstreamDependents, uNode)
					}
				} else if edge.SourceID == targetID {
					if dNode, ok := nrg.Nodes[edge.TargetID]; ok {
						downstreamDependencies = append(downstreamDependencies, dNode)
					}
				}
			}

			impactScore := (len(upstreamDependents)*3 + len(downstreamDependencies)) * 10
			if impactScore > 100 {
				impactScore = 100
			}
			severity := "Low Impact"
			if impactScore >= 70 {
				severity = "Critical Architectural Bottleneck"
			} else if impactScore >= 40 {
				severity = "Moderate Ripple Effect"
			}

			return c.JSON(fiber.Map{
				"target_node":             targetNode,
				"upstream_dependents":     upstreamDependents,
				"downstream_dependencies": downstreamDependencies,
				"impact_score":            impactScore,
				"severity":                severity,
				"advice":                  fmt.Sprintf("Modifying %s risks cascading changes across %d dependent symbol layers.", targetNode.Name, len(upstreamDependents)),
			})
		}
		app.Get("/api/graph/impact", handleImpact)
		app.Get("/api/graph/impact/:id", handleImpact)

		// ─── API: Cortex & Freemodel.dev AI Interactive Chatbot ───
		app.Post("/api/chat", func(c *fiber.Ctx) error {
			var req struct {
				Message string `json:"message"`
			}
			if err := c.BodyParser(&req); err != nil || strings.TrimSpace(req.Message) == "" {
				return c.Status(400).JSON(fiber.Map{"error": "Message content required"})
			}
			currDir, _ := os.Getwd()
			resp, err := cortex.GenerateChatReply(currDir, req.Message)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": err.Error()})
			}
			return c.JSON(resp)
		})

		// ─── API: Online/Offline Dynamic Repository Analyzer Switcher ───
		app.Post("/api/repo/scan", func(c *fiber.Ctx) error {
			var req struct {
				Target string `json:"target"`
			}
			if err := c.BodyParser(&req); err != nil || strings.TrimSpace(req.Target) == "" {
				return c.Status(400).JSON(fiber.Map{"error": "Target directory or GitHub URL required"})
			}

			fmt.Printf("\n🌐 [WEB ANALYZER] Received request to dynamically analyze: %s\n", req.Target)
			newPath, err := remote.ResolveTarget(req.Target)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": fmt.Sprintf("Failed to clone or open repository: %v", err)})
			}

			_ = os.Chdir(newPath)
			if scanCmd.Run != nil {
				scanCmd.Run(cmd, []string{newPath})
			}
			if analyzeCmd.RunE != nil {
				_ = analyzeCmd.RunE(cmd, []string{newPath})
			}

			return c.JSON(fiber.Map{
				"status": "success",
				"path":   newPath,
				"msg":    fmt.Sprintf("Repository successfully scanned and indexed! (Active target: %s)", req.Target),
			})
		})

		// ─── Dashboard: Rich Interactive HTML ───
		app.Get("/", func(c *fiber.Ctx) error {
			c.Set("Content-Type", "text/html; charset=utf-8")
			return c.SendString(dashboardHTML)
		})

		url := "http://localhost:" + activePort
		underline := color.New(color.Underline).SprintFunc()
		fmt.Printf("%s Dashboard Server Starting on port: %s\n", cyan("🌐 [CodeMRI Serve]"), green(activePort))
		fmt.Printf("👉 Access local intelligence UI at: %s\n", underline(url))
		fmt.Println("Press Ctrl+C to shut down the server.")

		go func() {
			time.Sleep(400 * time.Millisecond)
			openBrowser(url)
		}()

		if err := app.Listen(":" + activePort); err != nil {
			log.Fatalf("Server terminated: %v", err)
		}
	},
}

func resolveAvailablePort(startPortStr string) string {
	startPort, err := strconv.Atoi(startPortStr)
	if err != nil {
		return startPortStr
	}
	for i := 0; i < 20; i++ {
		currPort := strconv.Itoa(startPort + i)
		ln, err := net.Listen("tcp", ":"+currPort)
		if err == nil {
			ln.Close()
			return currPort
		}
	}
	return startPortStr
}

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = exec.Command("xdg-open", url).Start()
	}
	if err == nil {
		fmt.Println("🚀 Opened interactive dashboard in system browser automatically!")
	}
}

func init() {
	serveCmd.Flags().StringVarP(&port, "port", "p", "4000", "Port to serve the interactive dashboard on")
}
