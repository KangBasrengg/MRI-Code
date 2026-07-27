// Package graph implements the structural types and persistence layers for the Neural Repository Graph (NRG).
package graph

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "github.com/glebarez/go-sqlite" // Pure Go SQLite driver (CGO-free)
)

// SQLiteStorage implements the GraphStorage interface using an embedded SQLite database.
// Designed for Phase 03 (Neuron), enabling sub-millisecond indexed relational graph queries.
type SQLiteStorage struct {
	db     *sql.DB
	dbPath string
}

// NewSQLiteStorage opens or creates the SQLite database at dbPath and initializes tables & indexes.
func NewSQLiteStorage(dbPath string) (*SQLiteStorage, error) {
	// Ensure parent directory exists
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}

	db, err := sql.Open("sqlite", dbPath+"?_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to SQLite storage: %w", err)
	}

	storage := &SQLiteStorage{
		db:     db,
		dbPath: dbPath,
	}

	if err := storage.initSchema(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to initialize SQLite schema: %w", err)
	}

	return storage, nil
}

// initSchema creates the foundational relational tables and performance indexes for the NRG.
func (s *SQLiteStorage) initSchema() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS graphs (
			id TEXT PRIMARY KEY,
			repo_root TEXT NOT NULL,
			version TEXT NOT NULL,
			updated_at DATETIME NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS nodes (
			id TEXT PRIMARY KEY,
			graph_id TEXT NOT NULL,
			type TEXT NOT NULL,
			name TEXT NOT NULL,
			path TEXT,
			start_line INTEGER,
			end_line INTEGER,
			metadata TEXT,
			created_at DATETIME NOT NULL,
			FOREIGN KEY(graph_id) REFERENCES graphs(id) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS edges (
			id TEXT PRIMARY KEY,
			graph_id TEXT NOT NULL,
			source_id TEXT NOT NULL,
			target_id TEXT NOT NULL,
			type TEXT NOT NULL,
			weight REAL,
			description TEXT,
			FOREIGN KEY(graph_id) REFERENCES graphs(id) ON DELETE CASCADE
		);`,
		`CREATE INDEX IF NOT EXISTS idx_nodes_graph_id ON nodes(graph_id);`,
		`CREATE INDEX IF NOT EXISTS idx_nodes_type ON nodes(type);`,
		`CREATE INDEX IF NOT EXISTS idx_nodes_path ON nodes(path);`,
		`CREATE INDEX IF NOT EXISTS idx_edges_graph_id ON edges(graph_id);`,
		`CREATE INDEX IF NOT EXISTS idx_edges_source ON edges(source_id);`,
		`CREATE INDEX IF NOT EXISTS idx_edges_target ON edges(target_id);`,
		`CREATE INDEX IF NOT EXISTS idx_edges_type ON edges(type);`,
	}

	for _, q := range queries {
		if _, err := s.db.Exec(q); err != nil {
			return err
		}
	}
	return nil
}

// SaveGraph persists the NeuralRepositoryGraph into the SQLite database inside an atomic transaction.
func (s *SQLiteStorage) SaveGraph(graph *NeuralRepositoryGraph) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 1. Insert or replace graph record
	_, err = tx.Exec(`INSERT OR REPLACE INTO graphs (id, repo_root, version, updated_at) VALUES (?, ?, ?, ?)`,
		graph.ID, graph.RepoRoot, graph.Version, graph.UpdatedAt.Format(time.RFC3339))
	if err != nil {
		return fmt.Errorf("failed to save graph metadata: %w", err)
	}

	// 2. Clear existing nodes & edges for this graph (support clean re-indexing)
	_, err = tx.Exec(`DELETE FROM edges WHERE graph_id = ?`, graph.ID)
	if err != nil {
		return err
	}
	_, err = tx.Exec(`DELETE FROM nodes WHERE graph_id = ?`, graph.ID)
	if err != nil {
		return err
	}

	// 3. Prepare node insert statement
	nodeStmt, err := tx.Prepare(`INSERT OR REPLACE INTO nodes (id, graph_id, type, name, path, start_line, end_line, metadata, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer nodeStmt.Close()

	for _, node := range graph.Nodes {
		metaBytes, _ := json.Marshal(node.Metadata)
		_, err = nodeStmt.Exec(node.ID, graph.ID, string(node.Type), node.Name, node.Path, node.StartLine, node.EndLine, string(metaBytes), node.CreatedAt.Format(time.RFC3339))
		if err != nil {
			return fmt.Errorf("failed to insert node %s: %w", node.ID, err)
		}
	}

	// 4. Prepare edge insert statement
	edgeStmt, err := tx.Prepare(`INSERT OR REPLACE INTO edges (id, graph_id, source_id, target_id, type, weight, description) VALUES (?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer edgeStmt.Close()

	for _, edge := range graph.Edges {
		_, err = edgeStmt.Exec(edge.ID, graph.ID, edge.SourceID, edge.TargetID, string(edge.Type), edge.Weight, edge.Description)
		if err != nil {
			return fmt.Errorf("failed to insert edge %s: %w", edge.ID, err)
		}
	}

	return tx.Commit()
}

// LoadGraph loads the active NeuralRepositoryGraph from SQLite storage for a given repository root.
func (s *SQLiteStorage) LoadGraph(repoRoot string) (*NeuralRepositoryGraph, error) {
	var id, root, version, updatedStr string
	err := s.db.QueryRow(`SELECT id, repo_root, version, updated_at FROM graphs WHERE repo_root = ? ORDER BY updated_at DESC LIMIT 1`, repoRoot).Scan(&id, &root, &version, &updatedStr)
	if err != nil {
		return nil, fmt.Errorf("no cached graph found for repository %s: %w", repoRoot, err)
	}

	updatedAt, _ := time.Parse(time.RFC3339, updatedStr)
	nrg := &NeuralRepositoryGraph{
		ID:        id,
		RepoRoot:  root,
		Version:   version,
		Nodes:     make(map[string]*Node),
		Edges:     make([]*Edge, 0),
		UpdatedAt: updatedAt,
	}

	// Load Nodes
	rows, err := s.db.Query(`SELECT id, type, name, path, start_line, end_line, metadata, created_at FROM nodes WHERE graph_id = ?`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var nodeID, nodeType, name, path, metaStr, createdStr string
		var startLine, endLine int
		if err := rows.Scan(&nodeID, &nodeType, &name, &path, &startLine, &endLine, &metaStr, &createdStr); err != nil {
			return nil, err
		}
		var metadata map[string]string
		json.Unmarshal([]byte(metaStr), &metadata)
		createdAt, _ := time.Parse(time.RFC3339, createdStr)

		nrg.Nodes[nodeID] = &Node{
			ID:        nodeID,
			Type:      NodeType(nodeType),
			Name:      name,
			Path:      path,
			StartLine: startLine,
			EndLine:   endLine,
			Metadata:  metadata,
			CreatedAt: createdAt,
		}
	}

	// Load Edges
	edgeRows, err := s.db.Query(`SELECT id, source_id, target_id, type, weight, description FROM edges WHERE graph_id = ?`, id)
	if err != nil {
		return nil, err
	}
	defer edgeRows.Close()

	for edgeRows.Next() {
		var edgeID, sourceID, targetID, edgeType, desc string
		var weight float64
		if err := edgeRows.Scan(&edgeID, &sourceID, &targetID, &edgeType, &weight, &desc); err != nil {
			return nil, err
		}
		nrg.Edges = append(nrg.Edges, &Edge{
			ID:          edgeID,
			SourceID:    sourceID,
			TargetID:    targetID,
			Type:        EdgeType(edgeType),
			Weight:      weight,
			Description: desc,
		})
	}

	return nrg, nil
}

// FindNodeByID executes an indexed lookup for a specific node in SQLite storage.
func (s *SQLiteStorage) FindNodeByID(id string) (*Node, error) {
	var nodeID, nodeType, name, path, metaStr, createdStr string
	var startLine, endLine int
	err := s.db.QueryRow(`SELECT id, type, name, path, start_line, end_line, metadata, created_at FROM nodes WHERE id = ?`, id).
		Scan(&nodeID, &nodeType, &name, &path, &startLine, &endLine, &metaStr, &createdStr)
	if err != nil {
		return nil, fmt.Errorf("node %s not found: %w", id, err)
	}

	var metadata map[string]string
	json.Unmarshal([]byte(metaStr), &metadata)
	createdAt, _ := time.Parse(time.RFC3339, createdStr)

	return &Node{
		ID:        nodeID,
		Type:      NodeType(nodeType),
		Name:      name,
		Path:      path,
		StartLine: startLine,
		EndLine:   endLine,
		Metadata:  metadata,
		CreatedAt: createdAt,
	}, nil
}

// FindNeighbors performs a relational join to discover all Target nodes connected from Source ID via edgeType.
func (s *SQLiteStorage) FindNeighbors(id string, edgeType EdgeType) ([]*Node, error) {
	query := `
		SELECT n.id, n.type, n.name, n.path, n.start_line, n.end_line, n.metadata, n.created_at 
		FROM nodes n 
		JOIN edges e ON n.id = e.target_id 
		WHERE e.source_id = ? AND e.type = ?`

	rows, err := s.db.Query(query, id, string(edgeType))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	neighbors := make([]*Node, 0)
	for rows.Next() {
		var nodeID, nodeType, name, path, metaStr, createdStr string
		var startLine, endLine int
		if err := rows.Scan(&nodeID, &nodeType, &name, &path, &startLine, &endLine, &metaStr, &createdStr); err != nil {
			return nil, err
		}
		var metadata map[string]string
		json.Unmarshal([]byte(metaStr), &metadata)
		createdAt, _ := time.Parse(time.RFC3339, createdStr)

		neighbors = append(neighbors, &Node{
			ID:        nodeID,
			Type:      NodeType(nodeType),
			Name:      name,
			Path:      path,
			StartLine: startLine,
			EndLine:   endLine,
			Metadata:  metadata,
			CreatedAt: createdAt,
		})
	}
	return neighbors, nil
}

// GetTopologySummary executes ultra-fast SQL aggregation queries without loading the entire graph into RAM.
func (s *SQLiteStorage) GetTopologySummary() (map[string]interface{}, error) {
	var totalNodes, totalEdges int
	_ = s.db.QueryRow(`SELECT COUNT(*) FROM nodes`).Scan(&totalNodes)
	_ = s.db.QueryRow(`SELECT COUNT(*) FROM edges`).Scan(&totalEdges)

	rows, err := s.db.Query(`SELECT type, COUNT(*) FROM nodes GROUP BY type`)
	nodeTypes := make(map[string]int)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var nType string
			var count int
			if err := rows.Scan(&nType, &count); err == nil {
				nodeTypes[nType] = count
			}
		}
	}

	edgeRows, err := s.db.Query(`SELECT type, COUNT(*) FROM edges GROUP BY type`)
	edgeTypes := make(map[string]int)
	if err == nil {
		defer edgeRows.Close()
		for edgeRows.Next() {
			var eType string
			var count int
			if err := edgeRows.Scan(&eType, &count); err == nil {
				edgeTypes[eType] = count
			}
		}
	}

	return map[string]interface{}{
		"total_nodes":    totalNodes,
		"total_edges":    totalEdges,
		"node_types":     nodeTypes,
		"edge_types":     edgeTypes,
		"storage_engine": "SQLite Relational Engine (Phase 3 Neuron)",
	}, nil
}

// Close gracefully closes the underlying SQLite database connection.
func (s *SQLiteStorage) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}

// Verify that SQLiteStorage satisfies the GraphStorage interface at compile time.
var _ GraphStorage = (*SQLiteStorage)(nil)
