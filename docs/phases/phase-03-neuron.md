# Phase 03 — v0.3.0 "Neuron" Specification & Implementation Plan

> **Codename:** Neuron  
> **Status:** Completed & Active Runtime  
> **Target Release:** v0.3.0  
> **Author:** Muhammad Nuril  

---

## 1. Phase Vision
Phase 03 ("Neuron") marks an unprecedented architectural leap for the **CodeMRI Repository Intelligence Platform**. In previous phases, architectural knowledge was stored in static, flat-file JSON formats. While human-readable, flat files encounter heavy RAM bottlenecks when traversing thousands of interconnected files, functions, and relational modules. 

**Neuron** transforms CodeMRI into an enterprise-grade **Relational Graph Query Engine**. By embedding a high-speed, zero-dependency **SQLite database** directly inside `.codemri/graph.db`, CodeMRI achieves sub-millisecond relational joins over dependency trees without consuming excessive system RAM. This phase cements CodeMRI as the ultimate local foundation for instant offline codebase comprehension and lightning-fast AI reasoning queries.

---

## 2. Technical Rationale & Architecture Decisions
- **CGO-Free Embedded SQLite Storage Engine (`modernc/sqlite` & `glebarez/go-sqlite`):** To preserve CodeMRI's primary manifesto requirement of painless cross-platform distribution (via WinGet, Homebrew, Scoop, and simple `curl-to-sh`), we strictly shun traditional CGO SQLite bindings. Using pure-Go compiled SQLite engines ensures seamless machine-code binaries across Windows, Linux, and Apple Silicon with zero system-level GCC compiler requirements.
- **Relational Neural Repository Schema (`graphs`, `nodes`, `edges`):**
  - **`graphs` Table:** Tracks structural epochs, repository root paths, and modification timestamps.
  - **`nodes` Table:** Indexed by `id`, `type`, and `path`, enabling microsecond lookups for specific classes, services, functions, and files.
  - **`edges` Table:** Indexed by `source_id`, `target_id`, and `type` (`IMPORTS`, `CALLS`, `DEPENDS_ON`), making complex bi-directional dependency tracing (such as finding all callers of a critical security API) instant and reliable.
- **Hybrid Storage & Backwards Compatibility (SSOT):** Every scan concurrently exports both `.codemri/graph.db` (for high-speed query indexing and API querying) and `.codemri/graph.json` (as a human-readable diff backup), preserving transparency while supercharging interactive tools.
- **Dynamic Relational API Endpoints:** Fiber endpoints (`/api/graph/summary`, `/api/graph/node/:id`, and `/api/graph/neighbors/:id/:edge`) query directly from SQLite via prepared SQL aggregations and joins.

---

## 3. Deliverables
1. **SQLite Storage Layer Implementation:**
   - Contract implementation in `internal/graph/sqlite_storage.go` adhering to the `GraphStorage` interface.
   - Atomic SQLite transactions with `INSERT OR REPLACE INTO` deduplication logic to effortlessly process multi-line import repetitions.
2. **CLI Engine Upgrades (v0.3.0 Neuron):**
   - Updated `internal/cli/scan.go` to construct and verify SQLite indexes in real time during repository scans.
   - Enhanced `internal/cli/doctor.go` with specialized diagnostic checks for local database instantiation and storage engine health.
   - Upgraded version descriptors in `internal/cli/version.go` to publish `v0.3.0 ("Neuron")`.
3. **Interactive Reactive Dashboard & Query UI:**
   - Enhanced embedded real-time web UI (`internal/cli/dashboard_html.go`) to display active SQLite relational metrics, edge dependency distributions, and engine status indicators directly inside the developer's web browser.

---

## 4. Acceptance Criteria
- [x] Executing `codemri version` clearly displays: `CodeMRI v0.3.0 ("Neuron") | Core Engine: Neural Repository Graph (NRG) - SQLite Relational Engine`.
- [x] Running `codemri doctor` passes all environment tests and explicitly verifies: `✔ SQLite Relational Engine (Neuron): ONLINE (CGO-free embedded graph querying active)`.
- [x] Running `codemri scan .` compiles AST structures and indexes both `.codemri/graph.db` and `.codemri/graph.json` in under 50 milliseconds without UNIQUE database constraints failing.
- [x] The interactive visualization dashboard at `http://localhost:4000` immediately reflects live SQLite analytical aggregation data without lagging or requiring Node.js servers.

---

## 5. Folder Structure (Phase 03 Snapshot)
```text
.
├── cmd/
│   └── codemri/
│       └── main.go              # CLI Entrypoint
├── docs/
│   ├── ADR/
│   │   └── 0001-neural-repository-graph.md
│   ├── phases/
│   │   ├── phase-01-genesis.md
│   │   ├── phase-02-atlas.md
│   │   └── phase-03-neuron.md    # Official Neuron Specification
│   └── philosophy.md
├── internal/
│   ├── cli/
│   │   ├── dashboard_html.go     # Reactive real-time UI connected to SQLite APIs
│   │   ├── doctor.go             # Diagnostics including SQLite storage verification
│   │   ├── root.go               # 1-word killer command
│   │   ├── scan.go               # Conveys NRG into SQLite database
│   │   ├── serve.go              # Relational Graph query endpoints
│   │   └── version.go            # v0.3.0 Neuron identifiers
│   ├── graph/
│   │   ├── nrg.go                # SSOT Structural graph interfaces
│   │   └── sqlite_storage.go     # CGO-free embedded SQLite Engine implementation
│   ├── parser/
│   │   └── parser.go
│   └── scanner/
│       ├── engine.go
│       └── stats.go
└── README.md
```

---

## 6. Coding & Git Standards
- **Database Safety & Connection Management:** Always close `sql.Rows` and `sql.Stmt` structures immediately using standard deferral routines to prevent database lockups (`SQLITE_BUSY`) during rapid iterative development.
- **Conventional Commits:** Adhere strictly to professional commit messages (e.g., `feat(graph): implement CGO-free sqlite storage engine for neuron`).

---

## 7. AI Implementation Instructions

```
You are implementing Phase 03 of CodeMRI.
Strictly follow this PRD and architecture decision rules.
Do NOT use CGO-dependent libraries (such as mattn/go-sqlite3); you must strictly utilize pure-Go SQLite implementations to guarantee zero-config cross-compilation.
Ensure database transactions wrap all bulk node and edge insertions to preserve sub-50ms scan execution times.
Do NOT implement features outside this phase (e.g., do not build AI chat LLM interfaces or 3D canvas renderers yet; those belong to Cortex and Vision).
When in doubt, choose maintainability, zero-runtime dependency, and speed over cleverness.
```
