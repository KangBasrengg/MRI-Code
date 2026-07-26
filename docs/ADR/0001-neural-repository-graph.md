# ADR 0001: Neural Repository Graph (NRG) as Single Source of Truth

**Date:** 2026-07-26  
**Status:** Accepted  
**Author:** Muhammad Nuril  

---

## Context
When developers or traditional AI coding tools attempt to understand or interrogate large software repositories (1,000 to 10,000+ files), the industry standard approach relies on two flawed methods:
1. **Raw Text / Regex Scanning:** Searching files line-by-line via tools like grep or regex matches. This ignores contextual semantics, scopes, call hierarchies, and cross-file dependency chains.
2. **LLM Direct File Prompting ("AI Wrappers"):** Stuffing entire source files or arbitrary semantic search chunks into Large Language Model (LLM) context windows. This approach is costly in tokens, extremely slow, prone to hallucinations, and fails to comprehend systemic architecture.

---

## Decision
We will establish a standardized internal structural pattern called the **Neural Repository Graph (NRG)**.
1. **Deterministic Parsing:** Instead of relying on heuristic AI text reading, we will use highly optimized, deterministic parsers (such as Tree-sitter) to extract Abstract Syntax Trees (ASTs), symbols, function calls, variable scopes, and module dependencies.
2. **Graph Representation:** All parsed code elements become **Nodes** (modules, packages, classes, functions, routes, tables), and their interactions become **Edges** (imports, function calls, inheritances, queries).
3. **Single Source of Truth (SSOT):** This structural graph is stored locally in an efficient offline format (`.codemri/graph.db` using SQLite). **Once generated, no feature in CodeMRI will directly re-scan raw source code.**
4. **Universal Consumption:** All higher-level intelligence capabilities—including Interactive Architecture Visualizations, Call Graph Exploration, Technical Debt Calculation, Security Vulnerability Tracing, and the AI Reasoning Engine—will operate exclusively by traversing and querying the NRG.

---

## Consequences
### Positive
- **Instantaneous Speed:** Traversing structured relational edges in SQLite takes milliseconds, fulfilling our `< 60 seconds` repository understanding requirement.
- **Accuracy & Zero Hallucination (Structural):** AI responses are grounded in absolute, mathematically proven code execution paths derived from ASTs rather than probable token guessing.
- **Memory Efficiency:** Instead of loading gigabytes of textual files into memory, applications only query required subgraph nodes.
- **Strong Technical Identity:** CodeMRI clearly distinguishes itself from generic AI wrapper tools by standing firmly as a dedicated **Repository Intelligence Platform**.

### Negative / Trade-offs
- **Parser Complexity:** We must build and maintain language-specific AST mapper bridges (Go, TypeScript, Python, etc.) to normalize disparate language syntax into uniform NRG schema nodes.
- **Storage Overhead:** Each scanned project will generate a small local hidden directory (`.codemri/`) containing the compiled database.
