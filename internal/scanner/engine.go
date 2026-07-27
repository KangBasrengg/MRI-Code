package scanner

import (
	"context"
	"fmt"
	"io/fs"
	"math"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/KangBasrengg/MRI-Code/internal/parser"
)

// AtlasEngine implements high-speed directory walking, AST extraction, and language detection.
type AtlasEngine struct {
	registry *parser.Registry
}

// NewAtlasEngine initializes the Phase 2 scanner with an active parser registry.
func NewAtlasEngine() *AtlasEngine {
	return &AtlasEngine{
		registry: parser.NewRegistry(),
	}
}

// defaultExcludes represents system and dependency directories always skipped by default.
var defaultExcludes = map[string]bool{
	".git":          true,
	".codemri":      true,
	"node_modules":  true,
	"vendor":        true,
	"dist":          true,
	"build":         true,
	"bin":           true,
	".next":         true,
	".vercel":       true,
	"__pycache__":   true,
	".idea":         true,
	".vscode":       true,
}

// ExecuteScan traverses the target repository root, extracts ASTs, and aggregates language statistics.
func (e *AtlasEngine) ExecuteScan(ctx context.Context, config ScanConfig) (*ScanResult, error) {
	start := time.Now()
	result := &ScanResult{
		Languages:     make(map[parser.Language]int),
		LanguageStats: make(map[parser.Language]*LanguageSummary),
		ParsedASTs:    make([]*parser.FileAST, 0),
		Errors:        make([]error, 0),
	}

	// Prepare exclusion check set
	excludes := make(map[string]bool)
	for k, v := range defaultExcludes {
		excludes[k] = v
	}
	for _, rule := range config.IgnoreRules {
		excludes[strings.ToLower(rule)] = true
	}

	// Discover candidate files
	var filePaths []string
	var totalBytes int64

	err := filepath.WalkDir(config.RootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}

		name := d.Name()
		if d.IsDir() {
			if !config.IncludeHidden && strings.HasPrefix(name, ".") && name != "." && name != ".." && path != config.RootPath {
				return filepath.SkipDir
			}
			if excludes[strings.ToLower(name)] {
				return filepath.SkipDir
			}
			return nil
		}

		if !config.IncludeHidden && strings.HasPrefix(name, ".") {
			return nil
		}

		info, err := d.Info()
		if err == nil {
			totalBytes += info.Size()
		}

		filePaths = append(filePaths, path)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("repository traversal failed: %w", err)
	}

	result.TotalFiles = len(filePaths)
	result.TotalBytes = totalBytes

	// Parse syntax concurrently across workers
	workers := config.MaxWorkers
	if workers <= 0 {
		workers = 4
	}
	jobs := make(chan string, len(filePaths))
	resultsChan := make(chan *parser.FileAST, len(filePaths))

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for p := range jobs {
				select {
				case <-ctx.Done():
					return
				default:
					ast, err := e.registry.Parse(ctx, p)
					if err == nil && ast != nil {
						resultsChan <- ast
					}
				}
			}
		}()
	}

	for _, p := range filePaths {
		jobs <- p
	}
	close(jobs)
	wg.Wait()
	close(resultsChan)

	// Aggregate metrics
	for ast := range resultsChan {
		result.ParsedASTs = append(result.ParsedASTs, ast)
		result.TotalLOC += ast.LinesOfCode
		result.TotalBlank += ast.BlankLines
		result.TotalComments += ast.Comments
		result.Languages[ast.Language]++

		summary, exists := result.LanguageStats[ast.Language]
		if !exists {
			summary = &LanguageSummary{}
			result.LanguageStats[ast.Language] = summary
		}
		summary.FileCount++
		summary.Lines += ast.LinesOfCode + ast.Comments + ast.BlankLines
	}

	// Calculate percentage distributions
	totalLines := float64(result.TotalLOC + result.TotalBlank + result.TotalComments)
	if totalLines > 0 {
		for _, summary := range result.LanguageStats {
			ratio := (float64(summary.Lines) / totalLines) * 100.0
			summary.Percentage = math.Round(ratio*100) / 100
		}
	}

	result.DurationSec = math.Round(time.Since(start).Seconds()*1000) / 1000
	return result, nil
}
