# Phase 04 ("Pulse") — Architectural Technical Debt & Health Intelligence
**Version:** v0.4.0  
**Status:** Completed & Integrated  
**Core Innovation:** Deep structural technical debt analysis over the SQLite Neural Repository Graph (NRG)

---

## 1. Overview
In Phase 04 ("Pulse"), CodeMRI transitions from passive structural indexing into an active **Architectural Diagnostic & Reasoning Engine**. Rather than relying on cloud LLM inference or fragile linters to diagnose maintainability issues, Pulse leverages algorithmic analyses over the local SQLite graph database (`.codemri/graph.db`).

---

## 2. Core Analytical Engines

### A. Dead Code & Isolated Symbol Discovery (`internal/analyzer/dead_code.go`)
- Traverses incoming dependency edges (`CALLS`, `IMPORTS`, `DEPENDS_ON`, `QUERIES`).
- Identifies internal or private functions, methods, and data classes that receive zero structural invocations across the entire workspace.
- Exempts public exported APIs (capitalized syntax symbols) and framework lifecycle entrypoints (`main`, `init`, tests, react root hooks).

### B. Circular Dependency Detection (`internal/analyzer/circular.go`)
- Executes Depth-First Search (DFS) state tracking with recursion stacks across package import topologies.
- Immediately exposes cyclical architectural bonds (e.g., `package A -> package B -> package A`) which cause spaghetti coupling and maintenance deadlocks.

### C. Structural Complexity Hotspots (`internal/analyzer/complexity.go`)
- Computes an architectural complexity heuristic combining function density, file volume (LOC), and outgoing dependency branching.
- Categorizes code files into severity ratings: *Low*, *Moderate*, *High*, or *Extreme Technical Debt*.

### D. Authoritative Health Scoring (`internal/analyzer/health.go`)
- Synthesizes dead code rates, circular loops, documentation comment ratio, and structural complexity into a 0 to 100 benchmark.
- Assigns industry standard letter grades (**A+, A, B, C, D, F**) and furnishes engineering teams with prioritized, human-readable structural recommendations.

---

## 3. Command Usage & Integration

```bash
# Execute standalone technical debt and health diagnostics
codemri analyze .
# or alias:
codemri pulse .

# Run universal one-command workflow (automatically indexes Pulse diagnostics & launches Web UI)
codemri
```

The analytical output is persisted to `.codemri/pulse.json` as the single source of truth and exposed over HTTP via `/api/pulse` for reactive frontend visual dashboards.
