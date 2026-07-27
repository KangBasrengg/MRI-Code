# Phase 8: Velocity — Performance Intelligence Engine

**Codename:** Velocity  
**Version:** v0.8.0  
**Status:** COMPLETED  

## Summary
Phase 8 surfaces performance bottlenecks by analyzing file sizes, import density, and duplicate package distribution. Velocity helps developers identify modules that may cause slow builds, high memory usage, or excessive coupling.

## Key Features
- Large file detection (>50KB warning, >100KB high, >200KB critical)
- Heavy import analysis (files with >15 import statements)
- Duplicate package name detection across directories
- Performance grade scoring (A+ to F)
- JSON output at `.codemri/performance.json`
- Integrated into `codemri analyze` command

## Package
`internal/performance/analyzer.go`
