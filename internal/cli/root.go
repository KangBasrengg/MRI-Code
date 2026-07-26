// Package cli defines all user-facing subcommands and CLI execution logic for CodeMRI.
package cli

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	// Verbose toggles detailed debugging output across CLI commands.
	Verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "codemri",
	Short: color.CyanString("CodeMRI") + " - Neural Repository Intelligence Platform",
	Long: color.CyanString("🧠 CodeMRI - Neural Repository Intelligence Platform\n") +
		"\"GitHub shows your files. CodeMRI shows how your software actually works.\"\n\n" +
		"CodeMRI acts like an MRI machine for modern software codebases, scanning project\n" +
		"architectures in under 60 seconds and structuring them into an offline-first\n" +
		"Neural Repository Graph (NRG) for interactive exploration and AI reasoning.",
	SilenceErrors: true,
	SilenceUsage:  true,
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
