package cli

import (
	"fmt"
	"runtime"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const (
	// Version is the current release version of CodeMRI.
	Version = "v0.4.0"
	// Codename represents the thematic title for this structural epoch.
	Codename = "Pulse"
	// Architecture is the platform core structural engine.
	Architecture = "Neural Repository Graph (NRG) - SQLite Relational & Pulse Analytical Engine"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display CodeMRI platform version, codename, and runtime specifications",
	Run: func(cmd *cobra.Command, args []string) {
		cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()

		fmt.Printf("%s %s (%s)\n", cyan("CodeMRI"), green(Version), yellow(fmt.Sprintf("\"%s\"", Codename)))
		fmt.Println("---------------------------------------------------------")
		fmt.Printf("🎯 Core Engine : %s\n", Architecture)
		fmt.Printf("🚀 Go Runtime  : %s (%s/%s)\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
		fmt.Printf("💎 Philosophy  : %s\n", "We don't generate code. We generate understanding.")
	},
}
