# Phase 9: Forge — Plugin & Extensibility Foundation

**Codename:** Forge  
**Version:** v0.9.0  
**Status:** COMPLETED  

## Summary
Phase 9 establishes the extensibility foundation for CodeMRI's plugin ecosystem. By exposing structured JSON APIs, terminal query interfaces, and well-defined analytical output contracts, Forge enables community developers and enterprise teams to build custom integrations on top of the NRG.

## Key Features
- `codemri graph [query]` — terminal-based NRG query interface with search and `--json` output
- Structured JSON contracts: `pulse.json`, `security.json`, `performance.json`
- REST API endpoints for programmatic access (`/api/graph`, `/api/pulse`, `/api/graph/impact/:id`)
- Foundation for future Plugin SDK and marketplace

## Package
`internal/cli/graph.go`
