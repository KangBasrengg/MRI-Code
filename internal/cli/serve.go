package cli

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
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

		app := fiber.New(fiber.Config{
			AppName:      "CodeMRI v0.1.0 (Genesis)",
			DisableStartupMessage: true,
		})

		app.Use(logger.New(logger.Config{
			Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
		}))

		// API endpoint checking system health and NRG status
		app.Get("/api/status", func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"platform":     "CodeMRI Repository Intelligence Platform",
				"version":      Version,
				"codename":     Codename,
				"nrg_engine":   "ONLINE_INITIALIZED",
				"architecture": "Single Source of Truth (ADR-0001)",
			})
		})

		// Simple preview page for Genesis dashboard
		app.Get("/", func(c *fiber.Ctx) error {
			html := `
			<!DOCTYPE html>
			<html>
			<head>
				<title>CodeMRI Dashboard (Genesis)</title>
				<style>
					body { font-family: 'Inter', system-ui, sans-serif; background: #0b0f19; color: #e2e8f0; display: flex; align-items: center; justify-content: center; height: 90vh; margin: 0; }
					.card { background: #1e293b; padding: 3rem; border-radius: 1rem; box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.5), 0 10px 10px -5px rgba(0, 0, 0, 0.2); border: 1px solid #334155; max-width: 600px; text-align: center; }
					h1 { color: #38bdf8; font-size: 2.2rem; margin-bottom: 0.5rem; }
					p { color: #94a3b8; line-height: 1.6; }
					.badge { display: inline-block; background: #3b82f622; color: #60a5fa; border: 1px solid #3b82f6; padding: 0.3rem 0.8rem; border-radius: 9999px; font-weight: bold; font-size: 0.85rem; margin-top: 1.5rem; }
				</style>
			</head>
			<body>
				<div class="card">
					<h1>🧠 CodeMRI Dashboard</h1>
					<p>"GitHub shows your files. CodeMRI shows how your software actually works."</p>
					<div class="badge">Version v0.1.0 "Genesis" Server Active</div>
					<p style="margin-top: 2rem; font-size: 0.9rem;">Neural Repository Graph (NRG) visualization & AI Reasoning Chat engine queued for Phase 05 (Vision).</p>
				</div>
			</body>
			</html>
			`
			c.Set("Content-Type", "text/html")
			return c.SendString(html)
		})

		underline := color.New(color.Underline).SprintFunc()
		fmt.Printf("%s Dashboard Server Starting on port: %s\n", cyan("🌐 [CodeMRI Serve]"), green(port))
		fmt.Printf("👉 Access local intelligence UI at: %s\n", underline("http://localhost:"+port))
		fmt.Println("Press Ctrl+C to shut down the server.")
		
		if err := app.Listen(":" + port); err != nil {
			log.Fatalf("Server terminated: %v", err)
		}
	},
}

func init() {
	serveCmd.Flags().StringVarP(&port, "port", "p", "4000", "Port to serve the interactive dashboard on")
}
