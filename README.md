<div align="center">

# 🧠 Code-MRI
### The World's First Neural Repository Intelligence Platform

*"GitHub shows your files. CodeMRI shows how your software actually works."*

[![Version](https://img.shields.io/badge/Version-v0.1.0--Genesis-blue?style=for-the-badge&logo=go)](https://github.com/KangBasrengg/MRI-Code)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-orange.svg?style=for-the-badge)](https://opensource.org/licenses/Apache-2.0)
[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go)](https://golang.org)
[![Repository Intelligence](https://img.shields.io/badge/Repository%20Intelligence-98%2F100-success?style=for-the-badge&logo=github)](docs/ADR/0001-neural-repository-graph.md)

</div>

---

## ⚡ Why Code-MRI?

Modern software repositories contain thousands of files, hidden dependencies, undocumented architectures, and tangled execution flows. Engineers waste **hours or days** attempting to comprehend unfamiliar repositories before writing a single line of code. Current AI coding assistants act as blind text predictors—they generate boilerplate without understanding your architecture as a living system.

**CodeMRI functions like an MRI scan for your software.**
In less than **60 seconds**, it translates your raw code into an interactive, offline-first **Neural Repository Graph (NRG)**—empowering developers with instant visualizations of architectural flows, dependencies, technical debt, security boundaries, and deep systemic AI reasoning.

---

## 💎 Our Philosophy & Adoption Strategy

Read our definitive product manifesto: **[The CodeMRI Philosophy Document](docs/philosophy.md)**

1. **We generate understanding, not just code.**
2. **The CLI is 100% Free Forever.** We never monetize the core CLI engine; developers deserve frictionless open-source intelligence under the enterprise-safe **Apache 2.0 License**.
3. **Parser understands syntax; AI understands semantic meaning.**
4. **The Neural Repository Graph (NRG) is our single source of truth.**
5. **Offline by default, instant execution (<60s), zero mandatory config.**

### 🔥 The Viral Intelligence Badge
Imagine opening any open-source or enterprise GitHub repository and instantly spotting:
```markdown
[![Repository Intelligence](https://img.shields.io/badge/Repository%20Intelligence-98%2F100-success?style=for-the-badge)](https://github.com/KangBasrengg/MRI-Code)
```
Clicking the badge instantly transports engineers into the live CodeMRI visual dashboard—exploring architecture, technical debt, and AI insights without manual code auditing.

---

## 🗺️ Release Roadmap & Codenames

CodeMRI is crafted with strict engineering discipline, where every version represents a monumental structural evolution:

| Version | Codename | Primary Focus | Status & License Model |
| :---: | :---: | :--- | :---: |
| **v0.1** | 🌟 **Genesis** | Foundation monorepo, CLI engine, skeleton interfaces | **ACTIVE** *(Free / Apache 2.0)* |
| **v0.2** | 🗺️ **Atlas** | High-speed scanner, Tree-sitter integration, AST pipelines | *Planned* *(Free / Apache 2.0)* |
| **v0.3** | 🧠 **Neuron** | Neural Repository Graph (NRG), SQLite indexing, dependency map | *Planned* *(Free / Apache 2.0)* |
| **v0.4** | 💓 **Pulse** | Repository health, technical debt, complexity calculation | *Planned* *(Free / Apache 2.0)* |
| **v0.5** | 👁️ **Vision** | Interactive visual web dashboard (Next.js + React Flow) | *Planned* *(Free / Apache 2.0)* |
| **v0.6** | ⚡ **Cortex** | AI Reasoning Engine over NRG nodes (Ollama / OpenAI / Claude) | *Planned* *(Free / Apache 2.0)* |
| **v0.7** | 🛡️ **Shield** | Security intelligence (Secrets, SQLi, package vuln scanning) | *Planned* *(Free / Apache 2.0)* |
| **v0.8** | 🚀 **Velocity** | Performance insights, bundle size analysis, compilation bottlenecks | *Planned* *(Free / Apache 2.0)* |
| **v0.9** | 🔨 **Forge** | Extensible SDK & third-party plugin marketplace ecosystem | *Planned* *(Free / Apache 2.0)* |
| **v1.0** | 🎉 **MRI** | Stable public production release | *Planned* *(Free / Apache 2.0)* |
| **v1.5+** | ☁️ **Cloud** | CodeMRI Cloud Beta, Team Dashboards, AI PR Reviewer | *Future Ecosystem (Freemium/Sub)* |

---

## 📦 Quick Start (v0.1.0 Genesis)

### 1. Build from Source
Ensure you have Go installed (1.22+ recommended), then compile the CLI binary:
```bash
# Clone the repository
git clone https://github.com/KangBasrengg/MRI-Code.git
cd MRI-Code

# Compile binary
go build -o bin/codemri ./cmd/codemri
```

### 2. Available Commands
```bash
# Display version and release identity
./bin/codemri version

# Run structural diagnostic check on your machine
./bin/codemri doctor

# Scan current directory and initialize Neural Repository Graph skeleton
./bin/codemri scan .

# Launch interactive visualization server and architecture dashboard (default: port 4000)
./bin/codemri serve --port 4000
```

---

## 🏛️ Architecture Decision Records (ADR)
We believe in absolute engineering transparency and maintainability over cleverness:
- **[ADR 0001: Neural Repository Graph (NRG) as Single Source of Truth](docs/ADR/0001-neural-repository-graph.md)**

---

<div align="center">

Licensed under **Apache 2.0**. Made with passion by **Muhammad Nuril (KangBasrengg)** and the Code-MRI Community.

</div>
