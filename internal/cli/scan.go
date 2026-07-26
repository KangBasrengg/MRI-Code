package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type skeletonMeta struct {
	Platform    string    `json:"platform"`
	Version     string    `json:"version"`
	Codename    string    `json:"codename"`
	ScannedRoot string    `json:"scanned_root"`
	Timestamp   time.Time `json:"timestamp"`
	Status      string    `json:"status"`
}

var scanCmd = &cobra.Command{
	Use:   "scan [repository_path]",
	Short: "Scan target repository and initialize Neural Repository Graph (NRG)",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
		green := color.New(color.FgGreen, color.Bold).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()

		targetDir := "."
		if len(args) > 0 {
			targetDir = args[0]
		}

		absPath, err := filepath.Abs(targetDir)
		if err != nil {
			fmt.Printf("Error resolving path: %v\n", err)
			return
		}

		fmt.Printf("%s Starting Repository Scan on: %s\n", cyan("📡 [CodeMRI]"), yellow(absPath))
		startTime := time.Now()

		// Step 1: Initialize .codemri repository data directory
		dotDir := filepath.Join(absPath, ".codemri")
		if err := os.MkdirAll(dotDir, 0755); err != nil {
			fmt.Printf("❌ Failed to initialize local intelligence workspace (.codemri): %v\n", err)
			return
		}

		// Step 2: Generate Genesis repository skeleton file
		meta := skeletonMeta{
			Platform:    "CodeMRI Neural Repository Intelligence",
			Version:     Version,
			Codename:    Codename,
			ScannedRoot: absPath,
			Timestamp:   time.Now(),
			Status:      "GENESIS_INITIALIZED (Tree-sitter parser & NRG indexing queued for Phase 2 Atlas)",
		}

		metaBytes, _ := json.MarshalIndent(meta, "", "  ")
		metaPath := filepath.Join(dotDir, "repository.json")
		if err := os.WriteFile(metaPath, metaBytes, 0644); err != nil {
			fmt.Printf("❌ Failed to write skeleton metadata: %v\n", err)
			return
		}

		elapsed := time.Since(startTime)
		fmt.Printf("✔ Local intelligence repository created at: %s\n", green(dotDir))
		fmt.Printf("✔ Skeleton metadata generated: %s\n", green("repository.json"))
		fmt.Println("---------------------------------------------------------")
		fmt.Printf("✨ Genesis scan check complete in %v! (Ready for Phase 02 Tree-sitter & AST Pipeline)\n", elapsed)
	},
}
