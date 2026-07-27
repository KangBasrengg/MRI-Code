// Package scanner provides high-speed directory traversal logic for CodeMRI.
// Phase 2 Atlas engine implements Tree-sitter style AST symbol extraction and metadata compilation.
package scanner

import (
	"context"

	"github.com/KangBasrengg/MRI-Code/internal/parser"
)

// ScanConfig defines behavioral boundaries and filter exclusions for repository walks.
type ScanConfig struct {
	RootPath      string
	IgnoreRules   []string
	MaxWorkers    int
	IncludeHidden bool
}

// LanguageSummary encapsulates line metrics and ratio statistics per programming language.
type LanguageSummary struct {
	FileCount int     `json:"file_count"`
	Lines     int     `json:"lines"`
	Percentage float64 `json:"percentage"`
}

// ScanResult contains aggregate statistics and discovered sources from an execution cycle.
type ScanResult struct {
	TotalFiles    int                                `json:"total_files"`
	TotalBytes    int64                              `json:"total_bytes"`
	TotalLOC      int                                `json:"total_loc"`
	TotalBlank    int                                `json:"total_blank"`
	TotalComments int                                `json:"total_comments"`
	Languages     map[parser.Language]int            `json:"-"`
	LanguageStats map[parser.Language]*LanguageSummary `json:"language_stats"`
	ParsedASTs    []*parser.FileAST                  `json:"-"`
	DurationSec   float64                            `json:"duration_sec"`
	Errors        []error                            `json:"-"`
}

// Scanner defines the contractual obligations of any fast filesystem scanning implementation.
type Scanner interface {
	ExecuteScan(ctx context.Context, config ScanConfig) (*ScanResult, error)
}
