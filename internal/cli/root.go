// Package cli defines all user-facing subcommands and CLI execution logic for CodeMRI.
package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	// Verbose toggles detailed debugging output across CLI commands.
	Verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "codemri [repository_path]",
	Short: color.CyanString("CodeMRI") + " - Neural Repository Intelligence Platform",
	Long: color.CyanString("🧠 CodeMRI - Neural Repository Intelligence Platform\n") +
		"\"GitHub shows your files. CodeMRI shows how your software actually works.\"\n\n" +
		"CodeMRI acts like an MRI machine for modern software codebases, scanning project\n" +
		"architectures in under 60 seconds and structuring them into an offline-first\n" +
		"Neural Repository Graph (NRG) for interactive exploration and AI reasoning.",
	SilenceErrors: true,
	SilenceUsage:  true,
	Args:          cobra.MaximumNArgs(1),
	// Killer Feature: Running `codemri` directly without args (or with `.`) auto-scans and serves UI!
	Run: func(cmd *cobra.Command, args []string) {
		cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
		green := color.New(color.FgGreen, color.Bold).SprintFunc()
		yellow := color.New(color.FgYellow, color.Bold).SprintFunc()

		var err error
		cwd := "."
		if len(args) > 0 {
			cwd = args[0]
		}
		cwd, err = filepath.Abs(cwd)
		if err != nil {
			cwd, _ = os.Getwd()
		}

		fmt.Println("=====================================================================")
		fmt.Printf("%s Automations Online | Target Root: %s\n", cyan("⚡ [CodeMRI ONE-COMMAND WORKFLOW]"), yellow(cwd))
		fmt.Println("=====================================================================")

		// 1. Detect if local .codemri repository workspace already exists
		dotDir := filepath.Join(cwd, ".codemri")
		metaPath := filepath.Join(dotDir, "repository.json")

		if _, err := os.Stat(metaPath); os.IsNotExist(err) {
			fmt.Println(yellow("📥 No local cache (.codemri) detected. Running instant architecture scan..."))
			// Automatically invoke scan routine on current project
			if scanCmd.Run != nil {
				scanCmd.Run(cmd, []string{cwd})
			}
		} else {
			fmt.Printf("%s Cached Neural Repository Graph (.codemri) verified!\n", green("✔ [INSTANT CACHE]"))
			fmt.Println("ℹ (To force refresh syntax ASTs, execute: codemri scan .)")
		}

		fmt.Println("\n☕ All analytical structures ready. Launching interactive visual dashboard...")
		time.Sleep(500 * time.Millisecond)

		// 2. Automatically launch serve UI and open browser
		if serveCmd.Run != nil {
			serveCmd.Run(cmd, args)
		}
	},
}

// Execute adds all child commands to the root command and initiates parsing.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Enable detailed verbose debugging logs")

	// Register subcommands
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(doctorCmd)
	rootCmd.AddCommand(scanCmd)
	rootCmd.AddCommand(serveCmd)
}
