# Phase 01 — v0.1.0 "Genesis" Specification & Implementation Plan

> **Codename:** Genesis  
> **Status:** Active Implementation  
> **Target Release:** v0.1.0  
> **Author:** Muhammad Nuril  

---

## 1. Phase Vision
Phase 01 ("Genesis") focuses on establishing the absolute structural and architectural foundation of the **CodeMRI Repository Intelligence Platform**. 
In this phase, we transform an empty directory into a professional, enterprise-grade monorepo ready for advanced syntactic scanning and interactive visualization. Every folder, interface, and structural standard established here serves as the solid rock upon which all future phases (Atlas, Neuron, Pulse, Vision, Cortex) will rely.

---

## 2. Technical Rationale & Architecture Decisions
- **Core Engine in Go (Golang):** Chosen for optimal concurrency, near-zero startup overhead, native compilation to single binaries, and extreme memory efficiency (< 1GB RAM target for 10k files).
- **Monorepo Structure:** We unite the CLI engine (`/cmd/codemri`), internal logic (`/internal`), analysis systems, and UI dashboard skeleton (`/dashboard`) in a single organized repository to ensure consistent versioning and unified CI/CD workflows.
- **Strict Separation of Concerns:**
  - `scanner`: Dedicated to ultra-fast directory traversal and ignoring rules (like `.gitignore`).
  - `parser`: Language abstraction layer (preps for Tree-sitter in Phase 2).
  - `graph`: Skeleton definitions for the future Neural Repository Graph (NRG).
  - `server`: Local Go Fiber HTTP backend to serve the interactive web UI.
  - `doctor`: Environmental validation and health diagnosis tool for users.

---

## 3. Deliverables
1. **Repository Identity & Documentation:**
   - Philosophy document (`docs/philosophy.md`).
   - Architecture Decision Records (`docs/ADR/0001-neural-repository-graph.md`).
   - GitHub community templates (Issues, Pull Requests, Contribution guidelines).
2. **Go Module Setup & CLI Foundation:**
   - `go mod` initialization under `github.com/KangBasrengg/MRI-Code`.
   - Powerful CLI application powered by SPF13 Cobra with core subcommands:
     - `codemri scan <dir>` (Initial skeleton syntax with dry-run/mock AST summary).
     - `codemri serve --port 4000` (Local embedded server framework).
     - `codemri version` (Displays Version, Codename, Build details, and Go runtime).
     - `codemri doctor` (Checks local environment stability, Git availability, port availability).
3. **Internal Interface Skeletons:**
   - Modular interface contracts for `Scanner`, `Parser`, and `NRG Storage` in `/internal`.

---

## 4. Acceptance Criteria
- [x] Project cleanly compiles via standard `go build ./cmd/codemri` without warnings or errors.
- [x] Running `codemri version` prints precisely: `CodeMRI v0.1.0 "Genesis" (Repository Intelligence Platform)`.
- [x] Running `codemri doctor` accurately reports environment health (Go runtime, disk write permissions, TCP port 4000 availability).
- [x] Running `codemri scan .` creates an initial `.codemri` directory structure with skeleton metadata (`repository.json`).
- [x] Clean directory organization with NO circular dependencies between internal packages.

---

## 5. Folder Structure (Phase 01 Snapshot)
```text
.
├── .github/
│   ├── ISSUE_TEMPLATE/
│   └── pull_request_template.md
├── cmd/
│   └── codemri/
│       └── main.go              # Entrypoint
├── docs/
│   ├── ADR/
│   ├── phases/
│   └── philosophy.md
├── internal/
│   ├── cli/                   # Cobra command definitions
│   ├── doctor/                # System diagnostics & health check
│   ├── graph/                 # Neural Repository Graph core types & definitions
│   ├── parser/                # Universal parser abstraction layer
│   ├── scanner/               # High-speed repository traversal engine
│   └── server/                # Go Fiber local server engine
├── .gitignore
├── go.mod
├── go.sum
├── Makefile                   # Build, clean, and test automation
└── README.md                  # Public production-grade readme
```

---

## 6. Coding & Git Standards
- **Coding Style (Go):** Fully adhere to `gofmt` and `golangci-lint` guidelines. Write explicit Go doc comments for all exported types and functions. Choose clarity and maintainability over clever micro-optimizations.
- **Git Commit Rules:** Strict adherence to Conventional Commits:
  - `feat:` for new capabilities (e.g., `feat(cli): implement basic version command`).
  - `fix:` for bug fixes.
  - `docs:` for documentation updates.
  - `refactor:` for structural refactoring without breaking changes.
  - `chore:` for build scripts and maintenance.

---

## 7. AI Implementation Instructions

```
You are implementing Phase 01 of CodeMRI.
Strictly follow this PRD.
Do NOT implement any feature outside this phase.
Do NOT create placeholder code for future phases (e.g., do not implement actual AI calls or SQLite querying in Genesis).
The output must be production-ready, cleanly architecture, and compiles error-free.
Every architecture decision must follow this document and our philosophy.
When in doubt, choose maintainability over cleverness.
```
