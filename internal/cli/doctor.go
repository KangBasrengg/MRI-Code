package cli

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Run structural diagnostic checks on system environment, permissions, and tools",
	Run: func(cmd *cobra.Command, args []string) {
		green := color.New(color.FgGreen, color.Bold).SprintFunc()
		red := color.New(color.FgRed, color.Bold).SprintFunc()
		cyan := color.New(color.FgCyan, color.Bold).SprintFunc()

		fmt.Println(cyan("🔍 Running CodeMRI Diagnostics (v0.1.0 Genesis)..."))
		fmt.Println("---------------------------------------------------------")

		allPassed := true

		// 1. Check Go Version
		fmt.Printf("✔ Go Runtime Check: %s\n", green(runtime.Version()))

		// 2. Check Git Installation
		if path, err := exec.LookPath("git"); err == nil {
			fmt.Printf("✔ Git CLI available at: %s\n", green(path))
		} else {
			fmt.Printf("✖ Git CLI Check: %s (Recommended for git history intelligence)\n", red("NOT FOUND"))
			allPassed = false
		}

		// 3. Check Workspace Write Permissions (.codemri)
		testDir := ".codemri_test_doctor"
		err := os.MkdirAll(testDir, 0755)
		if err != nil {
			fmt.Printf("✖ Disk Permission Check: %s (%v)\n", red("FAILED"), err)
			allPassed = false
		} else {
			_ = os.RemoveAll(testDir)
			fmt.Printf("✔ Workspace Write Permission: %s (Can create .codemri data graphs)\n", green("OK"))
		}

		// 4. Check TCP Port 4000 (Default Dashboard Port)
		port := "4000"
		ln, err := net.Listen("tcp", "127.0.0.1:"+port)
		if err != nil {
			fmt.Printf("✖ TCP Port %s Check: %s (Already in use or inaccessible)\n", port, red("BLOCKED"))
			allPassed = false
		} else {
			_ = ln.Close()
			fmt.Printf("✔ TCP Port %s Availability: %s (Ready to serve local dashboard)\n", port, green("OPEN"))
		}

		fmt.Println("---------------------------------------------------------")
		if allPassed {
			fmt.Println(green("🎉 All diagnostic tests passed! CodeMRI engine is 100% operational."))
		} else {
			fmt.Println(color.YellowString("⚠️ Some non-critical diagnostic tests raised warnings. See details above."))
		}
	},
}
