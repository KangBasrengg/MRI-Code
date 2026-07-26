// Package scanner provides high-speed, concurrent directory traversal logic for CodeMRI.
package scanner

import (
	"context"

	"github.com/KangBasrengg/MRI-Code/internal/parser"
)

// ScanConfig defines behavioral boundaries and filter exclusions for repository walks.
type ScanConfig struct {
	RootPath     string
	IgnoreRules  []string
	MaxWorkers   int
	IncludeHidden bool
}

// ScanResult contains aggregate statistics and discovered sources from an execution cycle.
type ScanResult struct {
	TotalFiles  int
	TotalBytes  int64
	Languages   map[parser.Language]int
	DurationSec float64
	Errors      []error
}

// Scanner defines the contractual obligations of any fast filesystem scanning implementation.
type Scanner interface {
	ExecuteScan(ctx context.Context, config ScanConfig) (*ScanResult, error)
}
