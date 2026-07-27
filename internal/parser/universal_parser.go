package parser

import (
	"context"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/KangBasrengg/MRI-Code/internal/graph"
)

// UniversalParser implements a deterministic syntactic tokenizer and AST structural extractor
// for TypeScript, JavaScript, Python, PHP, Java, Rust, and SQL codebases.
type UniversalParser struct {
	importRegexTS  *regexp.Regexp
	importRegexPy  *regexp.Regexp
	funcRegexTS    *regexp.Regexp
	funcRegexPy    *regexp.Regexp
	classRegex     *regexp.Regexp
	sqlTableRegex  *regexp.Regexp
}

// NewUniversalParser instantiates an optimized multi-language parser engine.
func NewUniversalParser() *UniversalParser {
	return &UniversalParser{
		// ES6 / TS Import matches: import { X } from 'path' or import X from "path"
		importRegexTS: regexp.MustCompile(`(?m)^[ \t]*import\s+(?:[^"']+from\s+)?["']([^"']+)["']`),
		// Python matches: import mod or from mod import sym
		importRegexPy: regexp.MustCompile(`(?m)^[ \t]*(?:from|import)\s+([a-zA-Z0-9_\.]+)`),
		// TS / JS / PHP Functions & methods
		funcRegexTS:   regexp.MustCompile(`(?m)^[ \t]*(?:export\s+)?(?:async\s+)?(?:function|const|let|var)\s+([a-zA-Z0-9_]+)\s*(?:=\s*(?:async\s*)?\(|\()`),
		// Python / Rust Functions
		funcRegexPy:   regexp.MustCompile(`(?m)^[ \t]*(?:def|fn|pub\s+fn)\s+([a-zA-Z0-9_]+)\s*\(`),
		// Classes and Interfaces
		classRegex:    regexp.MustCompile(`(?m)^[ \t]*(?:export\s+)?(?:class|interface|struct)\s+([a-zA-Z0-9_]+)`),
		// SQL CREATE TABLE
		sqlTableRegex: regexp.MustCompile(`(?mi)^[ \t]*CREATE\s+(?:TABLE|VIEW)\s+(?:IF\s+NOT\s+EXISTS\s+)?([a-zA-Z0-9_\.]+)`),
	}
}

// SupportedLanguages returns the broad set of syntax engines served by this parser.
func (p *UniversalParser) SupportedLanguages() []Language {
	return []Language{
		LangTypeScript, LangJavaScript, LangPython, LangJava, LangPHP, LangSQL, LangRust,
	}
}

// ParseFile parses multi-language syntax into structured Neural Repository Graph symbols and relations.
func (p *UniversalParser) ParseFile(ctx context.Context, filePath string, content []byte) (*FileAST, error) {
	lang := DetectLanguage(filePath)
	result := &FileAST{
		FilePath: filePath,
		Language: lang,
		Nodes:    make([]*graph.Node, 0),
		Edges:    make([]*graph.Edge, 0),
		Errors:   make([]string, 0),
	}

	strContent := string(content)
	lines := strings.Split(strContent, "\n")
	for _, l := range lines {
		trimmed := strings.TrimSpace(l)
		if trimmed == "" {
			result.BlankLines++
		} else if strings.HasPrefix(trimmed, "//") || strings.HasPrefix(trimmed, "#") || strings.HasPrefix(trimmed, "/*") || strings.HasPrefix(trimmed, "--") {
			result.Comments++
		} else {
			result.LinesOfCode++
		}
	}

	relName := filepath.Base(filePath)
	fileNodeID := fmt.Sprintf("file:%s", filePath)

	fileNode := &graph.Node{
		ID:        fileNodeID,
		Type:      graph.NodeFile,
		Name:      relName,
		Path:      filePath,
		StartLine: 1,
		EndLine:   len(lines),
		Metadata: map[string]string{
			"lang": string(lang),
		},
		CreatedAt: time.Now(),
	}
	result.Nodes = append(result.Nodes, fileNode)

	// 1. Extract Database Table Schemas (SQL or ORM definitions)
	if lang == LangSQL {
		matches := p.sqlTableRegex.FindAllStringSubmatchIndex(strContent, -1)
		for _, idxs := range matches {
			tableName := strContent[idxs[2]:idxs[3]]
			tableID := fmt.Sprintf("db:%s", tableName)
			lineNum := strings.Count(strContent[:idxs[0]], "\n") + 1

			tNode := &graph.Node{
				ID:        tableID,
				Type:      graph.NodeDatabase,
				Name:      tableName,
				Path:      filePath,
				StartLine: lineNum,
				EndLine:   lineNum + 15, // estimated scope or statement terminator
				Metadata:  map[string]string{"schema": "relational_table"},
				CreatedAt: time.Now(),
			}
			result.Nodes = append(result.Nodes, tNode)

			result.Edges = append(result.Edges, &graph.Edge{
				ID:       fmt.Sprintf("e:%s:defines:%s", fileNodeID, tableID),
				SourceID: fileNodeID,
				TargetID: tableID,
				Type:     graph.EdgeExposes,
				Weight:   1.0,
			})
		}
		return result, nil
	}

	// 2. Extract Imports & Dependency references
	var importMatches [][]string
	if lang == LangPython {
		importMatches = p.importRegexPy.FindAllStringSubmatch(strContent, -1)
	} else {
		importMatches = p.importRegexTS.FindAllStringSubmatch(strContent, -1)
	}
	for _, match := range importMatches {
		if len(match) > 1 {
			targetMod := strings.TrimSpace(match[1])
			modID := fmt.Sprintf("mod:%s", targetMod)
			result.Edges = append(result.Edges, &graph.Edge{
				ID:          fmt.Sprintf("e:%s:imports:%s", fileNodeID, modID),
				SourceID:    fileNodeID,
				TargetID:    modID,
				Type:        graph.EdgeImports,
				Weight:      1.0,
				Description: fmt.Sprintf("Imports module %s", targetMod),
			})
		}
	}

	// 3. Extract Classes, Structs, and Interfaces
	classMatches := p.classRegex.FindAllStringSubmatchIndex(strContent, -1)
	for _, idxs := range classMatches {
		className := strContent[idxs[2]:idxs[3]]
		classID := fmt.Sprintf("class:%s:%s", filePath, className)
		lineNum := strings.Count(strContent[:idxs[0]], "\n") + 1

		cNode := &graph.Node{
			ID:        classID,
			Type:      graph.NodeClass,
			Name:      className,
			Path:      filePath,
			StartLine: lineNum,
			EndLine:   lineNum,
			Metadata:  map[string]string{"lang": string(lang)},
			CreatedAt: time.Now(),
		}
		result.Nodes = append(result.Nodes, cNode)

		result.Edges = append(result.Edges, &graph.Edge{
			ID:       fmt.Sprintf("e:%s:exposes:%s", fileNodeID, classID),
			SourceID: fileNodeID,
			TargetID: classID,
			Type:     graph.EdgeExposes,
			Weight:   1.0,
		})
	}

	// 4. Extract Functions and Method blocks
	var fnMatches [][]int
	if lang == LangPython || lang == LangRust {
		fnMatches = p.funcRegexPy.FindAllStringSubmatchIndex(strContent, -1)
	} else {
		fnMatches = p.funcRegexTS.FindAllStringSubmatchIndex(strContent, -1)
	}
	for _, idxs := range fnMatches {
		fnName := strContent[idxs[2]:idxs[3]]
		fnID := fmt.Sprintf("fn:%s:%s", filePath, fnName)
		lineNum := strings.Count(strContent[:idxs[0]], "\n") + 1

		fnNode := &graph.Node{
			ID:        fnID,
			Type:      graph.NodeFunction,
			Name:      fnName,
			Path:      filePath,
			StartLine: lineNum,
			EndLine:   lineNum,
			Metadata:  map[string]string{"lang": string(lang)},
			CreatedAt: time.Now(),
		}
		result.Nodes = append(result.Nodes, fnNode)

		result.Edges = append(result.Edges, &graph.Edge{
			ID:       fmt.Sprintf("e:%s:exposes:%s", fileNodeID, fnID),
			SourceID: fileNodeID,
			TargetID: fnID,
			Type:     graph.EdgeExposes,
			Weight:   1.0,
		})
	}

	return result, nil
}
