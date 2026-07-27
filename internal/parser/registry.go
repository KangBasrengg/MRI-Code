package parser

import (
	"context"
	"fmt"
	"os"

	"github.com/KangBasrengg/MRI-Code/internal/graph"
)

// Registry coordinates language detection and syntax extraction across multi-language projects.
type Registry struct {
	parsers map[Language]Parser
}

// NewRegistry initializes the unified syntax registry with all Phase 2 native parsers.
func NewRegistry() *Registry {
	r := &Registry{
		parsers: make(map[Language]Parser),
	}
	r.Register(NewGoParser())
	r.Register(NewUniversalParser())
	return r
}

// Register registers a concrete syntax language parser in the lookup map.
func (r *Registry) Register(p Parser) {
	for _, lang := range p.SupportedLanguages() {
		r.parsers[lang] = p
	}
}

// Parse performs language detection and routes source code to its matching syntax engine.
func (r *Registry) Parse(ctx context.Context, filePath string) (*FileAST, error) {
	lang := DetectLanguage(filePath)

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read source file %s: %w", filePath, err)
	}

	// Route to specialized Parser if supported
	if p, ok := r.parsers[lang]; ok {
		return p.ParseFile(ctx, filePath, content)
	}

	// Fallback for documentation or config files (Count lines without breaking AST)
	result := &FileAST{
		FilePath: filePath,
		Language: lang,
		Nodes:    make([]*graph.Node, 0),
		Edges:    make([]*graph.Edge, 0),
		Errors:   make([]string, 0),
	}
	
	lines := len(content)
	if lines > 0 {
		result.LinesOfCode = 1
	}
	return result, nil
}
