package parser

import (
	"context"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
	"time"

	"github.com/KangBasrengg/MRI-Code/internal/graph"
)

// GoParser implements deterministic Abstract Syntax Tree parsing for Go source files.
type GoParser struct{}

// NewGoParser initializes a Go language parser instance.
func NewGoParser() *GoParser {
	return &GoParser{}
}

// SupportedLanguages returns Go language signature.
func (p *GoParser) SupportedLanguages() []Language {
	return []Language{LangGo}
}

// ParseFile parses a Go source code file into standardized Neural Repository Graph symbols and relations.
func (p *GoParser) ParseFile(ctx context.Context, filePath string, content []byte) (*FileAST, error) {
	result := &FileAST{
		FilePath: filePath,
		Language: LangGo,
		Nodes:    make([]*graph.Node, 0),
		Edges:    make([]*graph.Edge, 0),
		Errors:   make([]string, 0),
	}

	// Calculate Line counts
	lines := strings.Split(string(content), "\n")
	for _, l := range lines {
		trimmed := strings.TrimSpace(l)
		if trimmed == "" {
			result.BlankLines++
		} else if strings.HasPrefix(trimmed, "//") || strings.HasPrefix(trimmed, "/*") || strings.HasPrefix(trimmed, "*") {
			result.Comments++
		} else {
			result.LinesOfCode++
		}
	}

	fset := token.NewFileSet()
	fileNode, err := parser.ParseFile(fset, filePath, content, parser.ParseComments)
	if err != nil {
		result.Errors = append(result.Errors, fmt.Sprintf("Go syntax parse error: %v", err))
		return result, nil
	}

	relName := filepath.Base(filePath)
	fileNodeID := fmt.Sprintf("file:%s", filePath)

	// Create root Node for the file itself
	fNode := &graph.Node{
		ID:        fileNodeID,
		Type:      graph.NodeFile,
		Name:      relName,
		Path:      filePath,
		StartLine: 1,
		EndLine:   len(lines),
		Metadata: map[string]string{
			"package": fileNode.Name.Name,
			"lang":    string(LangGo),
		},
		CreatedAt: time.Now(),
	}
	result.Nodes = append(result.Nodes, fNode)

	// Extract Imports and create EdgeImports relations
	for _, imp := range fileNode.Imports {
		impPath := strings.Trim(imp.Path.Value, `"`)
		impID := fmt.Sprintf("pkg:%s", impPath)
		
		edge := &graph.Edge{
			ID:          fmt.Sprintf("e:%s:imports:%s", fileNodeID, impID),
			SourceID:    fileNodeID,
			TargetID:    impID,
			Type:        graph.EdgeImports,
			Weight:      1.0,
			Description: fmt.Sprintf("Imports %s", impPath),
		}
		result.Edges = append(result.Edges, edge)
	}

	// Traverse AST declarations for Functions, Methods, Structs, and Interfaces
	for _, decl := range fileNode.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			fnName := d.Name.Name
			if d.Recv != nil && len(d.Recv.List) > 0 {
				// It's a method on a struct/type
				recvType := ""
				switch r := d.Recv.List[0].Type.(type) {
				case *ast.Ident:
					recvType = r.Name
				case *ast.StarExpr:
					if id, ok := r.X.(*ast.Ident); ok {
						recvType = id.Name
					}
				}
				if recvType != "" {
					fnName = fmt.Sprintf("%s.%s", recvType, d.Name.Name)
				}
			}

			start := fset.Position(d.Pos()).Line
			end := fset.Position(d.End()).Line
			fnID := fmt.Sprintf("fn:%s:%s", filePath, fnName)

			fnNode := &graph.Node{
				ID:        fnID,
				Type:      graph.NodeFunction,
				Name:      fnName,
				Path:      filePath,
				StartLine: start,
				EndLine:   end,
				Metadata: map[string]string{
					"exported": fmt.Sprintf("%v", d.Name.IsExported()),
				},
				CreatedAt: time.Now(),
			}
			result.Nodes = append(result.Nodes, fnNode)

			// Edge: File EXPOSES Function
			result.Edges = append(result.Edges, &graph.Edge{
				ID:          fmt.Sprintf("e:%s:exposes:%s", fileNodeID, fnID),
				SourceID:    fileNodeID,
				TargetID:    fnID,
				Type:        graph.EdgeExposes,
				Weight:      1.0,
			})

		case *ast.GenDecl:
			if d.Tok == token.TYPE {
				for _, spec := range d.Specs {
					if typeSpec, ok := spec.(*ast.TypeSpec); ok {
						typeName := typeSpec.Name.Name
						typeID := fmt.Sprintf("type:%s:%s", filePath, typeName)
						start := fset.Position(typeSpec.Pos()).Line
						end := fset.Position(typeSpec.End()).Line

						nodeType := graph.NodeClass
						metaType := "struct"
						if _, isIntf := typeSpec.Type.(*ast.InterfaceType); isIntf {
							metaType = "interface"
						}

						tNode := &graph.Node{
							ID:        typeID,
							Type:      nodeType,
							Name:      typeName,
							Path:      filePath,
							StartLine: start,
							EndLine:   end,
							Metadata: map[string]string{
								"kind":     metaType,
								"exported": fmt.Sprintf("%v", typeSpec.Name.IsExported()),
							},
							CreatedAt: time.Now(),
						}
						result.Nodes = append(result.Nodes, tNode)

						// Edge: File EXPOSES Type/Struct
						result.Edges = append(result.Edges, &graph.Edge{
							ID:       fmt.Sprintf("e:%s:exposes:%s", fileNodeID, typeID),
							SourceID: fileNodeID,
							TargetID: typeID,
							Type:     graph.EdgeExposes,
							Weight:   1.0,
						})
					}
				}
			}
		}
	}

	return result, nil
}
