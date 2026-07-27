// Package performance implements Phase 08 (Velocity) performance intelligence.
// It analyzes import costs, duplicate dependencies, large files, and module complexity
// to surface performance bottlenecks in the repository.
package performance

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// PerformanceReport is the aggregate output of a Velocity performance scan.
type PerformanceReport struct {
	ScannedAt         string             `json:"scanned_at"`
	TotalFiles        int                `json:"total_files_scanned"`
	TotalSizeBytes    int64              `json:"total_size_bytes"`
	PerformanceGrade  string             `json:"performance_grade"`
	LargeFiles        []LargeFile        `json:"large_files"`
	HeavyImports      []HeavyImport      `json:"heavy_imports"`
	DuplicatePackages []DuplicatePackage  `json:"duplicate_packages"`
	Suggestions       []string           `json:"suggestions"`
}

// LargeFile represents a source file that exceeds reasonable size thresholds.
type LargeFile struct {
	Path      string `json:"path"`
	SizeBytes int64  `json:"size_bytes"`
	SizeHuman string `json:"size_human"`
	LineCount int    `json:"line_count"`
	Severity  string `json:"severity"`
}

// HeavyImport represents a file with an unusually high number of import statements.
type HeavyImport struct {
	Path        string `json:"path"`
	ImportCount int    `json:"import_count"`
	Language    string `json:"language"`
}

// DuplicatePackage represents package names that appear in multiple locations.
type DuplicatePackage struct {
	Name       string   `json:"name"`
	Locations  []string `json:"locations"`
	Occurrences int     `json:"occurrences"`
}

var sourceExtensions = map[string]string{
	".go":   "Go",
	".ts":   "TypeScript",
	".tsx":  "TypeScript",
	".js":   "JavaScript",
	".jsx":  "JavaScript",
	".py":   "Python",
	".java": "Java",
	".rs":   "Rust",
	".php":  "PHP",
	".rb":   "Ruby",
	".cs":   "C#",
	".cpp":  "C++",
	".c":    "C",
	".swift": "Swift",
	".kt":  "Kotlin",
}

var skipDirs = map[string]bool{
	".git": true, ".codemri": true, "node_modules": true, "vendor": true,
	"dist": true, "build": true, "__pycache__": true, ".venv": true,
	"bin": true, ".idea": true, ".vscode": true,
}

// RunPerformanceScan performs a full Velocity performance analysis.
func RunPerformanceScan(rootDir string) (*PerformanceReport, error) {
	report := &PerformanceReport{
		ScannedAt:         time.Now().Format(time.RFC3339),
		LargeFiles:        make([]LargeFile, 0),
		HeavyImports:      make([]HeavyImport, 0),
		DuplicatePackages: make([]DuplicatePackage, 0),
		Suggestions:       make([]string, 0),
	}

	packageLocations := make(map[string][]string)

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			if skipDirs[info.Name()] {
				return filepath.SkipDir
			}
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		lang, isSource := sourceExtensions[ext]
		if !isSource {
			return nil
		}

		report.TotalFiles++
		report.TotalSizeBytes += info.Size()

		relPath, _ := filepath.Rel(rootDir, path)
		relPath = filepath.ToSlash(relPath)

		// Detect large files
		if info.Size() > 50*1024 { // > 50KB
			severity := "WARNING"
			if info.Size() > 200*1024 {
				severity = "CRITICAL"
			} else if info.Size() > 100*1024 {
				severity = "HIGH"
			}

			data, readErr := os.ReadFile(path)
			lineCount := 0
			if readErr == nil {
				lineCount = strings.Count(string(data), "\n") + 1
			}

			report.LargeFiles = append(report.LargeFiles, LargeFile{
				Path:      relPath,
				SizeBytes: info.Size(),
				SizeHuman: humanizeBytes(info.Size()),
				LineCount: lineCount,
				Severity:  severity,
			})
		}

		// Count imports
		data, readErr := os.ReadFile(path)
		if readErr != nil {
			return nil
		}
		content := string(data)
		lines := strings.Split(content, "\n")

		importCount := 0
		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed, "import ") ||
				strings.HasPrefix(trimmed, "from ") ||
				strings.HasPrefix(trimmed, "require(") ||
				strings.HasPrefix(trimmed, "use ") {
				importCount++
			}
		}

		if importCount > 15 {
			report.HeavyImports = append(report.HeavyImports, HeavyImport{
				Path:        relPath,
				ImportCount: importCount,
				Language:    lang,
			})
		}

		// Track package names for duplicate detection
		baseName := strings.TrimSuffix(info.Name(), ext)
		packageLocations[baseName] = append(packageLocations[baseName], relPath)

		return nil
	})
	if err != nil {
		return nil, err
	}

	// Detect duplicate package names across different directories
	for name, locations := range packageLocations {
		if len(locations) > 2 && len(name) > 2 {
			dirs := make(map[string]bool)
			for _, loc := range locations {
				dirs[filepath.Dir(loc)] = true
			}
			if len(dirs) > 1 {
				report.DuplicatePackages = append(report.DuplicatePackages, DuplicatePackage{
					Name:        name,
					Locations:   locations,
					Occurrences: len(locations),
				})
			}
		}
	}

	// Sort by impact
	sort.Slice(report.LargeFiles, func(i, j int) bool {
		return report.LargeFiles[i].SizeBytes > report.LargeFiles[j].SizeBytes
	})
	sort.Slice(report.HeavyImports, func(i, j int) bool {
		return report.HeavyImports[i].ImportCount > report.HeavyImports[j].ImportCount
	})

	// Generate suggestions
	if len(report.LargeFiles) > 0 {
		report.Suggestions = append(report.Suggestions,
			fmt.Sprintf("📦 %d files exceed 50KB. Consider splitting large modules into smaller, focused units.", len(report.LargeFiles)))
	}
	if len(report.HeavyImports) > 0 {
		report.Suggestions = append(report.Suggestions,
			fmt.Sprintf("📥 %d files have >15 imports. High import counts indicate tightly coupled modules. Consider dependency injection or facade patterns.", len(report.HeavyImports)))
	}
	if len(report.DuplicatePackages) > 0 {
		report.Suggestions = append(report.Suggestions,
			fmt.Sprintf("♻️ %d package names appear in multiple directories. Consolidate or namespace to reduce confusion.", len(report.DuplicatePackages)))
	}
	if report.TotalSizeBytes > 5*1024*1024 {
		report.Suggestions = append(report.Suggestions,
			fmt.Sprintf("⚠️ Total source code size is %s. Large codebases benefit from modular architecture and lazy loading.", humanizeBytes(report.TotalSizeBytes)))
	}

	report.PerformanceGrade = calculatePerfGrade(report)

	return report, nil
}

// SaveReport persists the performance report as JSON.
func SaveReport(report *PerformanceReport, outputPath string) error {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(outputPath, data, 0644)
}

func calculatePerfGrade(r *PerformanceReport) string {
	score := 100
	for _, lf := range r.LargeFiles {
		if lf.Severity == "CRITICAL" {
			score -= 15
		} else if lf.Severity == "HIGH" {
			score -= 8
		} else {
			score -= 3
		}
	}
	score -= len(r.HeavyImports) * 5
	score -= len(r.DuplicatePackages) * 3

	if score < 0 {
		score = 0
	}

	switch {
	case score >= 90:
		return "A+ — Excellent Performance Profile"
	case score >= 80:
		return "A — Strong Performance"
	case score >= 70:
		return "B — Good Performance"
	case score >= 60:
		return "C — Moderate Performance Concerns"
	case score >= 40:
		return "D — Significant Performance Issues"
	default:
		return "F — Critical Performance Bottlenecks"
	}
}

func humanizeBytes(b int64) string {
	switch {
	case b >= 1024*1024:
		return fmt.Sprintf("%.1f MB", float64(b)/(1024*1024))
	case b >= 1024:
		return fmt.Sprintf("%.1f KB", float64(b)/1024)
	default:
		return fmt.Sprintf("%d B", b)
	}
}
