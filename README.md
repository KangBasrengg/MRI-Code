<div align="center">

# 🧠 CodeMRI
### The Offline-First Neural Repository Intelligence Platform

*"GitHub shows your syntax and file lists. CodeMRI reveals how your software actually thinks and executes."*

[![Version](https://img.shields.io/badge/Release-v1.0.0__MRI-10B981?style=for-the-badge&logo=go&logoColor=white&labelColor=080b12)](https://github.com/KangBasrengg/MRI-Code/releases)
[![Engine: Full Spectrum](https://img.shields.io/badge/Engine-Full__Spectrum__Intelligence-6366F1?style=for-the-badge&logo=sqlite&logoColor=white&labelColor=080b12)](docs/release_notes/v1.0.0-mri.md)
[![Privacy: 100% Offline](https://img.shields.io/badge/Privacy-100%25__Offline__CPU-3B82F6?style=for-the-badge&logo=security&logoColor=white&labelColor=080b12)](docs/philosophy.md)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-F59E0B.svg?style=for-the-badge&labelColor=080b12)](https://opensource.org/licenses/Apache-2.0)
[![Support & Donate: Saweria](https://img.shields.io/badge/Support-Saweria%20Donation-FFB600?style=for-the-badge&logo=ko-fi&logoColor=black&labelColor=080b12)](https://saweria.co/Hoodtech)


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

## 💎 Features & Architecture Highlights (v1.0.0 MRI Enterprise Release)

- **🤖 Cortex AI Copilot (Offline vs. Online AI Mode):** Ask natural language questions about your repository structure using `codemri explain` or via the persistent **Floating Chat Bubble** inside the analyzer web UI.
  - **🌐 Online AI Model Enrichment (`--online`):** Natively integrates with `freemodel.dev` API models to provide deep conversational architectural analysis enriched by local graph structural context.
  - **🔒 Strict Zero-Cloud Privacy Lock (`--offline`):** Enforced local processing with `codemri --offline`. Disables all cloud transmissions and computes answers exclusively over local SQLite structural topology (ADR-0002 compliant).
- **🕸️ Interactive Force-Directed Topology Canvas:** Introduces a 60FPS physics-driven graph visualization at `localhost:4000`. Click any symbol node to instantly compute its architectural impact radius, upstream callers, and downstream dependencies.
- **💓 Pulse — Tech Debt & Health Scoring:** Algorithmically computes an overall repository letter grade (A+ to F), isolates uncalled dead code symbols, and exposes circular package import chains instantly using `codemri analyze .`.
- **🛡️ Shield — Security & Secret Diagnostics:** Multi-layered security intelligence scanning for hardcoded API keys, JWTs, XSS vulnerabilities (`innerHTML`), SQL injections, and dependency risks with actionable remediation advice.
- **🚀 Velocity — Performance & Footprint Profiling:** Identifies packaging compilation bottlenecks, bloated source files (>50KB), heavy standard library dependency chains, and optimization targets.
- **💾 Sub-Millisecond Relational SQLite Engine:** Powered by a pure-Go embedded SQLite graph database (`.codemri/graph.db`). Query dependencies and architectural bonds across 100k+ lines of code in sub-millisecond latencies.
- **🔍 Universal Deterministic Syntax Parser:** Multi-language AST classification supporting Go, TypeScript, JavaScript, Python, Java, PHP, Rust, SQL, HTML, CSS, and Docker without external runtime SDKs.
- **🩺 Environmental Diagnostics & Port Guard:** Verify system readiness anytime with a single diagnostic run:
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
| **v0.4.0** | 💓 **Pulse** | Dead code discovery, circular dependency loops, complexity & repository health scoring | **✅ RELEASED** | [Read Spec](docs/phases/phase-04-pulse.md) |
| **v0.5.0** | 👁️ **Vision** | Interactive force-directed topology canvas & instant impact analysis engine | **✅ RELEASED** | [Read Spec](docs/phases/phase-05-vision.md) |
| **v0.6.0** | ⚡ **Cortex** | Offline AI reasoning engine — `codemri explain` natural language queries | **✅ RELEASED** | [Read Spec](docs/phases/phase-06-cortex.md) |
| **v0.7.0** | 🛡️ **Shield** | Security intelligence — secrets detection, SQLi/XSS/CSRF scanning | **✅ RELEASED** | [Read Spec](docs/phases/phase-07-shield.md) |
| **v0.8.0** | 🚀 **Velocity** | Performance intelligence — large files, heavy imports, duplicate packages | **✅ RELEASED** | [Read Spec](docs/phases/phase-08-velocity.md) |
| **v0.9.0** | 🔨 **Forge** | Plugin extensibility foundation — terminal graph queries, JSON API contracts | **✅ RELEASED** | [Read Spec](docs/phases/phase-09-forge.md) |
| **v1.0.0** | 🎉 **MRI** | Stable worldwide production enterprise release | **🌟 STABLE RELEASE** | [Read Spec](docs/release_notes/v1.0.0-mri.md) |

---

## 🏛️ Documentation & Manifesto

We believe in absolute transparency and building developer understanding over mindless code generation:
- **[📜 The CodeMRI Product Philosophy & Adoption Manifesto](docs/philosophy.md)**
- **[📐 ADR-0001: Neural Repository Graph (NRG) as Single Source of Truth](docs/ADR/0001-neural-repository-graph.md)**
- **[🛡️ ADR-0002: Local Privacy — Zero Cloud Transmission Policy](docs/ADR/0002-local-privacy.md)**
- **[🌐 Official Web Landing Page & Documentation Portal](website/)**

---

## 🤝 Contributing & Community

CodeMRI is licensed under the commercial-friendly **Apache 2.0 License**, guaranteeing complete patent security and free forever usage of the core CLI engine. Contributions, feature suggestions, and bug reports are warmly welcomed!

1. Read our **[Contributing Guide](CONTRIBUTING.md)** for setup, coding standards, and commit conventions.
2. Use our [Issue Templates](.github/ISSUE_TEMPLATE/) to report bugs, request features, or ask questions.
3. Check the [Pull Request Guidelines](.github/pull_request_template.md).
4. Fork the repository, build from source with `make build` or `go test ./...`, and submit your PR!

---

## 💖 Sponsor & Support the Project

CodeMRI is engineered with intense dedication, thousands of lines of CGO-free performance discipline, and an unconditional commitment to developers worldwide. If this tool saves your engineering team time, simplifies complex codebase analysis, or enhances your architecture workflows, please consider showing your appreciation to **Muhammad Nuril (KangBasrengg)** via Saweria:

<div align="center">
  <br>
  <a href="https://saweria.co/Hoodtech">
    <img src="https://img.shields.io/badge/☕_Dukung_&_Donasi_via-Saweria.co%2FHoodtech-FFC800?style=for-the-badge&logo=buymeacoffee&logoColor=black&labelColor=111827" alt="Support & Donate on Saweria" height="44">
  </a>
  <br><br>
  <i>Every contribution directly catalyzes future open-source innovation, new structural AI enhancements, and ongoing project maintenance. Thank you deeply for your support! ❤️</i>
</div>

<div align="center">
  <br><hr><br>
  Made with passion and zero-dependency discipline by <b>Muhammad Nuril (KangBasrengg)</b> and the CodeMRI Open Source Community.
</div>
