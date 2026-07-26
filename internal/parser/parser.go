// Package parser defines the universal language parser contracts for CodeMRI.
// In Phase 2 (Atlas), these interfaces will be concretely bound to multi-language Tree-sitter bindings.
package parser

import (
	"context"

	"github.com/KangBasrengg/MRI-Code/internal/graph"
)

// Language represents an official programming syntax supported by CodeMRI parsers.
type Language string

const (
	LangGo         Language = "go"
	LangTypeScript Language = "typescript"
	LangJavaScript Language = "javascript"
	LangPython     Language = "python"
	LangJava       Language = "java"
	LangSQL        Language = "sql"
	LangUnknown    Language = "unknown"
)

// FileAST represents an abstract structural syntax outcome resulting from parser execution.
type FileAST struct {
	FilePath string
	Language Language
	Nodes    []*graph.Node
	Edges    []*graph.Edge
	Errors   []string
}

// Parser defines the modular behavior required of any language-specific syntax engine.
// True to principle: "Parser understands code. AI understands meaning."
type Parser interface {
	SupportedLanguage() Language
	ParseFile(ctx context.Context, filePath string, content []byte) (*FileAST, error)
}
