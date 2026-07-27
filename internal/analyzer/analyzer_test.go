package analyzer

import (
	"testing"
	"time"

	"github.com/KangBasrengg/MRI-Code/internal/graph"
)

func TestPulseAnalysisEngine(t *testing.T) {
	nrg := graph.NewNRG("/tmp/test_repo", "v0.4.0")

	// Create test nodes
	node1 := &graph.Node{ID: "f_1", Name: "calculateTax", Type: graph.NodeFunction, Path: "pkg/calc/tax.go", StartLine: 10, EndLine: 30, CreatedAt: time.Now()}
	node2 := &graph.Node{ID: "f_2", Name: "oldUnusedHelper", Type: graph.NodeFunction, Path: "pkg/calc/tax.go", StartLine: 35, EndLine: 50, CreatedAt: time.Now()}
	node3 := &graph.Node{ID: "f_main", Name: "main", Type: graph.NodeFunction, Path: "cmd/main.go", StartLine: 5, EndLine: 20, CreatedAt: time.Now()}

	nrg.Nodes["f_1"] = node1
	nrg.Nodes["f_2"] = node2
	nrg.Nodes["f_main"] = node3

	// Create call edge: main calls calculateTax
	edge1 := &graph.Edge{ID: "e_1", SourceID: "f_main", TargetID: "f_1", Type: graph.EdgeCalls}
	// Create circular dependency edge between two files/symbols
	edgeCirc1 := &graph.Edge{ID: "e_2", SourceID: "f_1", TargetID: "f_2", Type: graph.EdgeImports}
	edgeCirc2 := &graph.Edge{ID: "e_3", SourceID: "f_2", TargetID: "f_1", Type: graph.EdgeImports}

	nrg.Edges = append(nrg.Edges, edge1, edgeCirc1, edgeCirc2)

	// Execute Pulse Analyzer
	report := Analyze(nrg, 100, 15)

	if report.Codename != "Pulse" {
		t.Errorf("expected codename Pulse, got %s", report.Codename)
	}

	// Verify dead code detection: oldUnusedHelper is imported by f_1 in our cycle test, let's test isolation
	// Since e_2 imports f_2, it is counted as referenced. Let's add a truly isolated node
	nodeIsolated := &graph.Node{ID: "f_dead", Name: "forgottenLegacyCode", Type: graph.NodeFunction, Path: "pkg/legacy/old.go", StartLine: 1, EndLine: 100, CreatedAt: time.Now()}
	nrg.Nodes["f_dead"] = nodeIsolated

	report2 := Analyze(nrg, 200, 20)
	if len(report2.DeadCodeIssues) == 0 {
		t.Errorf("expected dead code issues to be detected for forgottenLegacyCode, got 0")
	}

	if len(report2.CircularDeps) == 0 {
		t.Errorf("expected circular dependency loop between f_1 and f_2 to be detected, got 0")
	}

	if report2.Health.OverallScore < 0 || report2.Health.OverallScore > 100 {
		t.Errorf("health score out of bounds: %d", report2.Health.OverallScore)
	}

	t.Logf("Pulse Report verification passed! Health Score: %d (%s) - %s", report2.Health.OverallScore, report2.Health.Grade, report2.Health.DebtStatus)
}
