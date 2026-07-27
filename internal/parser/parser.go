// Package parser defines the universal language parser contracts for CodeMRI.
// In Phase 2 (Atlas), these interfaces are bound to native syntax analyzers and AST tokenizers.
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
	LangPHP        Language = "php"
	LangSQL        Language = "sql"
	LangRust       Language = "rust"
	LangHTML       Language = "html"
	LangCSS        Language = "css"
	LangMarkdown   Language = "markdown"
	LangJSON       Language = "json"
	LangYAML       Language = "yaml"
	LangShell      Language = "shell"
	LangUnknown    Language = "unknown"
)

// FileAST represents an abstract structural syntax outcome resulting from parser execution.
type FileAST struct {
	FilePath    string
	Language    Language
	LinesOfCode int
	BlankLines  int
	Comments    int
	Nodes       []*graph.Node
	Edges       []*graph.Edge
	Errors      []string
}

// Parser defines the modular behavior required of any language-specific syntax engine.
// True to principle: "Parser understands code. AI understands meaning."
type Parser interface {
	SupportedLanguages() []Language
	ParseFile(ctx context.Context, filePath string, content []byte) (*FileAST, error)
}
