# ADR-0002: Local Privacy — Zero Cloud Transmission Policy

**Status:** Accepted  
**Date:** 2026-06-15  
**Author:** Muhammad Nuril (@KangBasrengg)  
**Supersedes:** None

---

## Context

CodeMRI is designed to analyze proprietary enterprise repositories that may contain trade secrets, patented algorithms, confidential business logic, and sensitive authentication infrastructure. Engineers and organizations must trust that running `codemri` on their codebase will not expose any source code, metadata, or analytical results to external servers.

Many competing "AI-powered" code analysis tools require uploading repository contents to cloud endpoints for processing, creating significant legal, compliance, and intellectual property risks — especially in regulated industries (finance, healthcare, defense, government).

## Decision

**All CodeMRI analysis — scanning, parsing, graph indexing, technical debt scoring, topology rendering, and impact assessment — MUST execute entirely on the user's local machine.** No source code, NRG graph data, or analytical output is ever transmitted to external servers during normal CLI operation.

### Binding Rules

1. **Zero Network Calls During Analysis:** The `codemri scan`, `codemri analyze`, and `codemri serve` commands must never initiate outbound HTTP/HTTPS requests. All computation runs against local CPU, memory, and disk I/O.

2. **Embedded Storage Only:** The SQLite relational engine (`.codemri/graph.db`) and all JSON artifacts (`.codemri/repository.json`, `.codemri/pulse.json`) are persisted exclusively on the local filesystem within the scanned repository directory.

3. **No Telemetry or Usage Tracking:** CodeMRI does not collect, aggregate, or transmit any telemetry, usage statistics, crash reports, or system fingerprints. There are no analytics beacons, no phone-home mechanisms.

4. **CDN-Free Dashboard Rendering:** The embedded interactive dashboard served at `localhost:4000` must render using only locally bundled HTML, CSS, and JavaScript. No external CDN scripts (D3.js, Three.js, Google Fonts, etc.) are loaded from the internet.

5. **Optional Cloud Features (Future):** If cloud synchronization features are introduced in future versions (e.g., CodeMRI Cloud, Team Dashboard), they must be:
   - Strictly opt-in (never default)
   - Clearly disclosed before any data leaves the machine
   - Governed by explicit user consent and configuration

### Exceptions

- **`codemri update`** (planned): May check GitHub Releases API for newer binary versions. This transmits only the current version string, never source code or analysis data.
- **AI Reasoning (Phase 6 "Cortex"):** If users choose to connect external LLM providers (OpenAI, Claude, Gemini), only NRG graph summaries — never raw source code — are sent, and only with explicit user configuration of API keys.

## Consequences

### Positive
- Enterprise and government organizations can safely adopt CodeMRI without legal review friction
- Air-gapped environments (no internet) are fully supported
- User trust is established from day one — no "what are they doing with my code?" anxiety
- Competitive differentiation against cloud-dependent analysis tools

### Negative
- Limits ability to provide automatic crash reporting for debugging
- Real-time collaborative features require explicit opt-in architecture
- AI reasoning features must carefully scope what graph metadata is shared with LLM providers

## References

- [ADR-0001: Neural Repository Graph as Single Source of Truth](0001-neural-repository-graph.md)
- [CodeMRI Philosophy Document](../philosophy.md)
- Product Requirement Document: "Offline first. Repository scanning should not require internet." (Principle 3)
