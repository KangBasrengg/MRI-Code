# Phase 6: Cortex — Offline AI Reasoning Engine

**Codename:** Cortex  
**Version:** v0.6.0  
**Status:** COMPLETED  

## Summary
Phase 6 introduces the `codemri explain` command, enabling developers to ask natural language questions about their codebase architecture. The Cortex engine traverses the Neural Repository Graph using keyword extraction and relevance scoring — entirely offline, without requiring external LLM API keys.

## Key Features
- Natural language query parsing with stop-word removal
- NRG node search with weighted relevance scoring (name match: 10pts, path match: 5pts)
- Connection analysis (upstream callers, downstream dependencies)
- Automatic identification of critical architectural hubs
- Pulse health data integration for context enrichment

## CLI Usage
```bash
codemri explain "How does authentication work?"
codemri explain "What are the main packages?"
codemri explain "Which functions have the most dependencies?"
```
