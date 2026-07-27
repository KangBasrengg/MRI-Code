package cli

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/KangBasrengg/MRI-Code/internal/graph"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Run structural diagnostic checks on system environment, storage engines, and tools",
	Run: func(cmd *cobra.Command, args []string) {
		green := color.New(color.FgGreen, color.Bold).SprintFunc()
		red := color.New(color.FgRed, color.Bold).SprintFunc()
		cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
		yellow := color.New(color.FgYellow, color.Bold).SprintFunc()

		fmt.Println(cyan(fmt.Sprintf("🔍 Running CodeMRI Diagnostics (%s %s)...", Version, Codename)))
		fmt.Println("---------------------------------------------------------")

		allPassed := true

		// 1. Check Go Runtime Version
		fmt.Printf("✔ Go Runtime Check: %s (%s/%s)\n", green(runtime.Version()), runtime.GOOS, runtime.GOARCH)

		// 2. Check Git Installation
		if path, err := exec.LookPath("git"); err == nil {
			fmt.Printf("✔ Git CLI available at: %s\n", green(path))
		} else {
			fmt.Printf("✖ Git CLI Check: %s (Recommended for git history & author intelligence)\n", red("NOT FOUND"))
			allPassed = false
		}

		// 3. Check Workspace Write Permissions (.codemri)
		testDir := ".codemri_test_doctor"
		err := os.MkdirAll(testDir, 0755)
		if err != nil {
			fmt.Printf("✖ Disk Permission Check: %s (%v)\n", red("FAILED"), err)
			allPassed = false
		} else {
			// 4. Check Phase 04 Pulse SQLite Embedded Storage Engine
			testDbPath := filepath.Join(testDir, "test_pulse.db")
			sqliteStore, dbErr := graph.NewSQLiteStorage(testDbPath)
			if dbErr != nil {
				fmt.Printf("✖ SQLite Relational Engine Check: %s (%v)\n", red("FAILED"), dbErr)
				allPassed = false
			} else {
				sqliteStore.Close()
				fmt.Printf("✔ SQLite Relational & Full-Spectrum Engine (v1.0.0 MRI): %s (CGO-free embedded graph & intelligence active)\n", green("ONLINE"))
			}

			_ = os.RemoveAll(testDir)
			fmt.Printf("✔ Workspace Write Permission: %s (Can initialize and persist .codemri databases)\n", green("OK"))
		}

		// 5. Check TCP Port 4000 (Default Dashboard Port)
		port := "4000"
		ln, err := net.Listen("tcp", ":"+port)
		if err != nil {
			fmt.Printf("⚠ TCP Port %s Status: %s (Port occupied; serve auto-resolution will seamlessly switch to port 4001)\n", port, yellow("IN USE"))
		} else {
			_ = ln.Close()
			fmt.Printf("✔ TCP Port %s Availability: %s (Ready to serve local interactive dashboard)\n", port, green("OPEN"))
		}

		fmt.Println("──────────────────────────────────────────────────────────────────────────")
		if allPassed {
			fmt.Println(green("🎉 All diagnostic tests passed! CodeMRI intelligence engine (v1.0.0 MRI) is 100% operational."))
		} else {
			fmt.Println(color.YellowString("⚠️ Some diagnostic tests raised warnings. See details above."))
		}
	},
}
