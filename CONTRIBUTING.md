# Contributing to CodeMRI

Thank you for your interest in contributing to **CodeMRI** — the Offline-First Neural Repository Intelligence Platform! Every contribution, whether a typo fix or a major feature, makes a real difference.

---

## 📜 Code of Conduct

By participating in this project, you agree to abide by our [Code of Conduct](https://www.contributor-covenant.org/version/2/1/code_of_conduct/). Be kind, inclusive, and respectful.

---

## 🚀 Getting Started

### Prerequisites

- **Go 1.22+** (required to build from source)
- **Git** (for cloning and submitting pull requests)

### Build from Source

```bash
# Clone the repository
git clone https://github.com/KangBasrengg/MRI-Code.git
cd MRI-Code

# Run tests
go test ./...

# Build the binary
go build -o bin/codemri ./cmd/codemri

# Or install globally
go install ./cmd/codemri
```

### Verify Your Setup

```bash
codemri version
codemri doctor
```

---

## 🔧 How to Contribute

### 1. Reporting Bugs

Use our [Bug Report Template](https://github.com/KangBasrengg/MRI-Code/issues/new?template=bug_report.md). Include:
- Steps to reproduce
- Expected vs actual behavior
- OS and CodeMRI version (`codemri version`)

### 2. Suggesting Features

Use our [Feature Request Template](https://github.com/KangBasrengg/MRI-Code/issues/new?template=feature_request.md). Check the [Roadmap](README.md#-engineering-roadmap--release-progress) first to see if your idea aligns with a planned phase.

### 3. Submitting Code

1. **Fork** the repository
2. **Create a branch** from `main`:
   ```bash
   git checkout -b feat/my-awesome-feature
   ```
3. **Make your changes** following the coding standards below
4. **Run tests** to ensure nothing breaks:
   ```bash
   go test ./...
   ```
5. **Commit** using [Conventional Commits](https://www.conventionalcommits.org/):
   ```bash
   git commit -m "feat(scanner): add Ruby AST parser support"
   ```
6. **Push** and open a **Pull Request** against `main`

---

## 📐 Coding Standards

### Go Code Style

- Follow standard `gofmt` formatting
- Use meaningful variable and function names
- Add doc comments (`//`) to all exported types and functions
- Keep functions focused — prefer small, testable functions over monoliths
- Place business logic in `/internal/` packages, never in `/cmd/`

### Commit Message Convention

We use **Conventional Commits** strictly:

| Prefix | Purpose | Example |
| :--- | :--- | :--- |
| `feat` | New feature | `feat(parser): add Rust syntax classifier` |
| `fix` | Bug fix | `fix(scan): handle symlink loops gracefully` |
| `refactor` | Code restructuring | `refactor(graph): extract edge builder` |
| `docs` | Documentation only | `docs: update README badges to v0.5.0` |
| `test` | Adding/updating tests | `test(analyzer): add dead code edge cases` |
| `chore` | Build/CI/tooling | `chore: update goreleaser to v2 syntax` |

### Branch Naming

- `feat/description` — New features
- `fix/description` — Bug fixes
- `docs/description` — Documentation changes
- `refactor/description` — Code refactoring

---

## 🏗️ Project Architecture

```
cmd/codemri/          → CLI entrypoint
internal/
  cli/                → Cobra commands, dashboard HTML, Fiber server
  scanner/            → High-speed concurrent filesystem walker
  parser/             → Multi-language AST syntax classifiers
  graph/              → NRG types, SQLite storage engine
  analyzer/           → Pulse health scoring, dead code, circular deps
docs/
  ADR/                → Architecture Decision Records
  phases/             → Per-version technical specifications
  release_notes/      → GitHub release notes
website/              → React + Vite landing page & documentation portal
```

### Key Design Decisions

- **ADR-0001:** The Neural Repository Graph (NRG) is the Single Source of Truth for all features
- **ADR-0002:** All analysis runs 100% offline — zero source code is transmitted to external servers
- **Zero CGO:** The entire binary compiles without C dependencies for maximum portability

---

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run a specific package's tests
go test ./internal/analyzer/...
```

We welcome test contributions! Areas that need more coverage:
- `internal/parser/` — Language-specific AST classification
- `internal/graph/` — SQLite storage operations
- `internal/scanner/` — Filesystem walking edge cases

---

## 📝 Documentation Contributions

Documentation is just as important as code. You can help by:
- Improving existing docs in `/docs/`
- Adding examples to the README
- Writing Architecture Decision Records (ADRs) for significant changes
- Translating documentation to other languages

---

## 🎯 Roadmap & Phase Alignment

Before starting work on a large feature, check which [roadmap phase](README.md#-engineering-roadmap--release-progress) it belongs to. If you're unsure, open a Discussion or Issue first to align with the maintainers.

---

## ❤️ Thank You

Every contribution to CodeMRI helps build the future of repository intelligence. Whether you report a bug, improve documentation, or ship a feature — **you are part of this journey**.

*Licensed under Apache 2.0 • Built with ❤️ by Muhammad Nuril ([@KangBasrengg](https://github.com/KangBasrengg)) and the CodeMRI community.*
