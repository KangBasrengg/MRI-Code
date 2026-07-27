# Phase 02 — v0.2.0 "Atlas" Specification & Implementation Plan

> **Codename:** Atlas  
> **Status:** Completed & Integrated  
> **Target Release:** v0.2.0  
> **Author:** Muhammad Nuril  

---

## 1. Phase Vision
Phase 02 ("Atlas") focuses on bringing life to the repository exploration engine. The goal of Atlas is to map the software landscape with superhuman speed—scanning directory structures, parsing syntax trees across multiple languages, identifying architectural symbols (functions, packages, imports, routes), and compiling the initial **Neural Repository Graph (NRG)**. All of this must run in milliseconds entirely on local CPU, setting the benchmark for zero-dependency, offline-first code intelligence.

---

## 2. Technical Rationale & Architecture Decisions
- **High-Concurrency Worker Pool Engine:** Implemented via Go channels and goroutines in `internal/scanner`. By splitting file traversal and AST extraction across concurrent worker pools, large codebases (3,000+ files) can be analyzed in under 100 milliseconds without UI freezing or excessive RAM spikes.
- **Deterministic Universal Syntax Parser:** Built in `internal/parser` with regex and AST abstraction layers supporting Go, TypeScript, JavaScript, HTML, CSS, Markdown, and Makefiles, classifying symbols and dependencies without needing external runtime SDKs or language servers installed on the host.
- **Single Source of Truth (SSOT) via `.codemri`:** As mandated by ADR-0001, every scan generates structured data directly inside `.codemri/graph.json` and `.codemri/repository.json`. This decouples analysis from display, ensuring web dashboards, AI reasoning engines, and external CI CI tasks all read from the exact same standardized graph model.
- **Auto-Port Resolution & 64KB HTTP Buffering:** Integrated inside `internal/cli/serve.go` to prevent port collision crashes and fix browser `HTTP 431 Request Header Fields Too Large` errors caused by large localhost cookies.

---

## 3. Deliverables
1. **Atlas Concurrent Scanner & Universal Parser:**
   - Worker pool engine in `internal/scanner/engine.go`.
   - AST node and edge extractor in `internal/parser/parser.go`.
2. **One-Command Killer Workflow:**
   - Running `codemri` with zero arguments auto-detects the project root, executes instantaneous AST scanning, spins up the Fiber local server, and automatically opens the user's default system browser.
3. **Rich Block-Based Documentation Portal:**
   - Overhauling the landing page docs rendering (`website/src/components/DocsViewer.tsx`) from raw Markdown strings to clean, programmatic React block components with vibrant visual syntax styling and zero technical bug artifacts.

---

## 4. Acceptance Criteria
- [x] Running `codemri scan .` successfully traverses projects and prints a colorized summary of Total Files, Lines of Code (LOC), Comment counts, and Language percentages.
- [x] Scan completion latencies routinely measure under 100 milliseconds for normal repository footprints.
- [x] `.codemri/graph.json` reliably maps code nodes (`FILE`, `FUNCTION`, `PACKAGE`) and directional edges (`IMPORTS`, `CALLS`).
- [x] Running `codemri` automatically triggers both scanning and serving without requiring multiple repetitive CLI commands.

---

## 5. Folder Structure (Phase 02 Snapshot)
```text
.
├── cmd/
│   └── codemri/
│       └── main.go
├── docs/
│   ├── ADR/
│   │   └── 0001-neural-repository-graph.md
│   ├── phases/
│   │   ├── phase-01-genesis.md
│   │   └── phase-02-atlas.md
│   └── philosophy.md
├── internal/
│   ├── cli/
│   │   ├── root.go              # One-Command Killer Workflow
│   │   ├── scan.go              # Atlas scanning triggers & JSON writer
│   │   ├── serve.go             # Fiber server with auto-port resolution
│   │   └── version.go
│   ├── graph/
│   │   └── nrg.go               # SSOT Node & Edge types
│   ├── parser/
│   │   └── parser.go            # Deterministic multi-language AST extraction
│   └── scanner/
│       ├── engine.go            # High-speed worker pool traverser
│       └── stats.go
├── website/                     # React/Vite Landing Page & Official Portal
│   ├── src/
│   │   ├── components/
│   │   │   ├── DocsViewer.tsx   # Block-based structured rendering UI
│   │   │   └── Hero.tsx
│   │   └── data/
│   │       └── docsContent.ts
└── README.md
```

---

## 6. Coding & Git Standards
- **Concurrency Safety:** Ensure zero race conditions across worker channels; enforce read-only map accesses during symbol formatting.
- **Conventional Commits:** Continue strict adherence to `feat:`, `fix:`, `docs:`, and `refactor:` standard commit prefixes.

---

## 7. AI Implementation Instructions

```
You are implementing Phase 02 of CodeMRI.
Strictly follow this PRD and architectural specifications.
Ensure concurrency worker pools remain bounded to prevent OS file descriptor exhaustion.
Do NOT implement SQLite database integrations or advanced interactive force-directed graph physics in this phase (those belong to Phase 03 and Phase 05).
The output must be production-ready, highly tested, and compiles cleanly with zero dependencies.
When in doubt, choose maintainability and speed over cleverness.
```
