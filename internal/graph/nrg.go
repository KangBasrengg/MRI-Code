// Package graph provides structural type abstractions for the Neural Repository Graph (NRG).
// As mandated by ADR-0001, the NRG serves as the Single Source of Truth (SSOT) across all features.
package graph

import (
	"time"
)

// NodeType identifies the semantic architectural role of a codebase symbol or file.
type NodeType string

const (
	NodeModule    NodeType = "MODULE"
	NodePackage   NodeType = "PACKAGE"
	NodeFile      NodeType = "FILE"
	NodeFunction  NodeType = "FUNCTION"
	NodeClass     NodeType = "CLASS"
	NodeRoute     NodeType = "ROUTE"
	NodeDatabase  NodeType = "DATABASE_TABLE"
	NodeService   NodeType = "SERVICE"
)

// EdgeType identifies the directional relational bond between two NRG nodes.
type EdgeType string

const (
	EdgeImports    EdgeType = "IMPORTS"
	EdgeCalls      EdgeType = "CALLS"
	EdgeInherits   EdgeType = "INHERITS"
	EdgeQueries    EdgeType = "QUERIES"
	EdgeExposes    EdgeType = "EXPOSES"
	EdgeDependsOn  EdgeType = "DEPENDS_ON"
)

// Node represents an individual atomic architectural component within the repository.
type Node struct {
	ID          string     `json:"id"`
	Type        NodeType   `json:"type"`
	Name        string     `json:"name"`
	Path        string     `json:"path,omitempty"`
	StartLine   int        `json:"start_line,omitempty"`
	EndLine     int        `json:"end_line,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}

// Edge represents a directed architectural interaction between Source Node and Target Node.
type Edge struct {
	ID          string   `json:"id"`
	SourceID    string   `json:"source_id"`
	TargetID    string   `json:"target_id"`
	Type        EdgeType `json:"type"`
	Weight      float64  `json:"weight,omitempty"`
	Description string   `json:"description,omitempty"`
}

// NeuralRepositoryGraph represents the in-memory or persisted architectural network.
type NeuralRepositoryGraph struct {
	ID        string           `json:"graph_id"`
	RepoRoot  string           `json:"repo_root"`
	Version   string           `json:"version"`
	Nodes     map[string]*Node `json:"nodes"`
	Edges     []*Edge          `json:"edges"`
	UpdatedAt time.Time        `json:"updated_at"`
}

// GraphStorage defines the strict contractual interface required to persist and query an NRG.
// In Phase 3 (Neuron), this interface will be backed by a highly optimized offline SQLite engine.
type GraphStorage interface {
	SaveGraph(graph *NeuralRepositoryGraph) error
	LoadGraph(repoRoot string) (*NeuralRepositoryGraph, error)
	FindNodeByID(id string) (*Node, error)
	FindNeighbors(id string, edgeType EdgeType) ([]*Node, error)
	Close() error
}

// NewNRG instantiates an empty Neural Repository Graph structure.
func NewNRG(repoRoot string, version string) *NeuralRepositoryGraph {
	return &NeuralRepositoryGraph{
		ID:        "nrg-genesis-01",
		RepoRoot:  repoRoot,
		Version:   version,
		Nodes:     make(map[string]*Node),
		Edges:     make([]*Edge, 0),
		UpdatedAt: time.Now(),
	}
}
