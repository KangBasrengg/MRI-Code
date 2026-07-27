<div align="center">

# 🧠 CodeMRI
### The Offline-First Neural Repository Intelligence Platform

*"GitHub shows your syntax and file lists. CodeMRI reveals how your software actually thinks and executes."*

[![Version](https://img.shields.io/badge/Release-v0.3.0__Neuron-00F2FF?style=for-the-badge&logo=go&logoColor=black&labelColor=080b12)](https://github.com/KangBasrengg/MRI-Code/releases)
[![Engine: SQLite NRG](https://img.shields.io/badge/Engine-SQLite__Relational__Graph-10B981?style=for-the-badge&logo=sqlite&logoColor=white&labelColor=080b12)](docs/phases/phase-03-neuron.md)
[![Privacy: 100% Offline](https://img.shields.io/badge/Privacy-100%25__Offline__CPU-3B82F6?style=for-the-badge&logo=security&logoColor=white&labelColor=080b12)](docs/philosophy.md)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-F59E0B.svg?style=for-the-badge&labelColor=080b12)](https://opensource.org/licenses/Apache-2.0)

<br>

<p align="center">
  <b>Don't clone our tool. Just install once, navigate into any complex repository, and type one word:</b><br>
  <code style="font-size: 1.3em; color: #00f2ff; background: #0f172a; padding: 6px 12px; border-radius: 6px;">codemri</code>
</p>

</div>

---

## ⚡ What is CodeMRI?

Modern enterprise software repositories contain thousands of files, hidden dependencies, undocumented architectures, and tangled execution flows. Engineers waste **hours or days** attempting to comprehend unfamiliar repositories before writing a single line of code. Traditional AI assistants act as blind text predictors—generating boilerplate without understanding your architecture as a living system.

**CodeMRI works like an automated medical MRI scan for your codebase.**
In under **60 seconds** and entirely offline, it compiles your raw code into an indexed **Neural Repository Graph (NRG)** powered by an embedded pure-Go SQLite relational engine. It launches an interactive visual dashboard in your browser instantly—giving you superpowers to understand architectures, dependency bottlenecks, and AI reasoning insights without leaking a single byte of code to cloud endpoints.

---

## 📦 Zero-Friction Installation & Quickstart

We deliberately chose a **Native Go Binary** distribution over cumbersome NPM global installations or complex repository clones. No Node.js required. No JVM required. Launches in less than 5 milliseconds.

### 1. Install Global Binary (Pick Your OS)

| Platform | Recommended Command | Alternative |
| :--- | :--- | :--- |
| **🪟 Windows** | `winget install codemri` | `scoop install codemri` |
| **🍏 macOS** | `brew install codemri` | `curl -fsSL https://codemri.dev/install.sh \| sh` |
| **🐧 Linux** | `curl -fsSL https://codemri.dev/install.sh \| sh` | Download `.tar.gz` from GitHub Releases |
| **🐹 Go Developers** | `go install github.com/KangBasrengg/MRI-Code/cmd/codemri@latest` | Requires Go 1.22+ |

### 2. The "One-Word" Killer Workflow

Once installed, you **never** need to configure complex YAML files or clone tools again. Whenever you encounter an unfamiliar codebase or want to inspect your current software, follow this 3-step muscle memory:

```bash
# 1. Navigate to ANY target project directory on your computer
cd my-unfamiliar-monorepo

# 2. Execute the one-word intelligence engine
codemri

# 3. ☕ Sit back! CodeMRI scans syntax, indexes SQLite dependencies, and launches your visual dashboard automatically!
```

---

## 💎 Features & Architecture Highlights (v0.3.0 Neuron)

- **🚀 Sub-Millisecond Relational Engine:** Transitioned in Phase 3 from static flat JSON files to a pure-Go embedded SQLite graph database (`.codemri/graph.db`). Query dependencies and architectural bonds across 100k+ lines of code instantly.
- **🛡️ 100% Local & Offline Privacy:** All parsing and graph structural indexing happens on your local CPU and Disk. Your proprietary code is never modified, uploaded to third-party servers, or exposed.
- **🎨 Reactive Interactive Dashboard:** When you run `codemri` (or `codemri serve`), it spins up a local high-speed Fiber server on port 4000 and displays an ultra-premium dark-mode interactive dashboard directly in your browser.
- **🔍 Universal Deterministic Syntax Parser:** Built-in multi-language AST classification supporting Go, TypeScript, JavaScript, Python, Java, PHP, Rust, SQL, HTML, CSS, and Docker without external runtime SDKs.
- **🩺 Environmental Health Diagnostics:** Verify your local machine readiness anytime with a single diagnostic run:
  ```bash
  codemri doctor
  ```

---

## 🗺️ Engineering Roadmap & Release Progress

CodeMRI follows strict iterative engineering discipline where every release represents a monumental structural breakthrough:

| Version | Codename | Primary Focus | Status | Specification |
| :---: | :---: | :--- | :---: | :---: |
| **v0.1.0** | 🌟 **Genesis** | Monorepo architecture, CLI foundational engine, interface contracts | **✅ RELEASED** | [Read Spec](docs/phases/phase-01-genesis.md) |
| **v0.2.0** | 🗺️ **Atlas** | High-speed concurrent worker pools, multi-language syntax AST parsing | **✅ RELEASED** | [Read Spec](docs/phases/phase-02-atlas.md) |
| **v0.3.0** | 🧠 **Neuron** | CGO-free embedded SQLite relational graph indexing (`graph.db`), API querying | **✅ RELEASED** | [Read Spec](docs/phases/phase-03-neuron.md) |
| **v0.4.0** | 💓 **Pulse** | Incremental Git history intelligence, technical debt & complexity algorithms | *In Progress* | — |
| **v0.5.0** | 👁️ **Vision** | Interactive visual web canvas (Force-directed graph 2D/3D topology nodes) | *Planned* | — |
| **v0.6.0** | ⚡ **Cortex** | Offline & targeted AI reasoning endpoints over NRG nodes (Ollama/Claude/GPT) | *Planned* | — |
| **v0.7.0** | 🛡️ **Shield** | Structural security intelligence (Secrets detection, dependency CVE verification) | *Planned* | — |
| **v0.8.0** | 🚀 **Velocity** | Performance compilation insights & bundling bottleneck diagnostics | *Planned* | — |
| **v0.9.0** | 🔨 **Forge** | Extensible plugin marketplace & universal community SDK | *Planned* | — |
| **v1.0.0** | 🎉 **MRI** | Stable worldwide production enterprise release | *Target Goal* | — |

---

## 🏛️ Documentation & Manifesto

We believe in absolute transparency and building developer understanding over mindless code generation:
- **[📜 The CodeMRI Product Philosophy & Adoption Manifesto](docs/philosophy.md)**
- **[📐 ADR 0001: Neural Repository Graph (NRG) as Single Source of Truth](docs/ADR/0001-neural-repository-graph.md)**
- **[🌐 Official Web Landing Page & Documentation Portal](website/)**

---

## 🤝 Contributing & Community

CodeMRI is licensed under the commercial-friendly **Apache 2.0 License**, guaranteeing complete patent security and free forever usage of the core CLI engine. Contributions, feature suggestions, and bug reports are warmly welcomed!

1. Check our [Pull Request Guidelines](.github/pull_request_template.md).
2. Fork the repository and build from source using `make build` or `go test ./...`.
3. Submit your pull request adhering to Conventional Commits!

<div align="center">
  <br>
  Made with passion and zero-dependency discipline by <b>Muhammad Nuril (KangBasrengg)</b> and the CodeMRI Open Source Community.
</div>
