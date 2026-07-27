# Phase 5: Vision — Interactive Force-Directed Topology & Impact Engine

**Codename:** Vision  
**Version:** v0.5.0  
**Status:** COMPLETED  
**Architecture:** Neural Repository Graph (NRG) - Embedded SQLite & Force Physics Engine

---

## Executive Summary

Phase 5 (**"Vision"**) marks the grand realization of CodeMRI's primary visual architecture mandate in our Product Requirement Document (PRD). Moving beyond static metrics and text-based AST logs, Vision empowers engineers, software architects, and tech leads to physically interact with their repository's structural ecosystem through a high-performance **2D Force-Directed Topology Canvas**.

By treating the embedded SQLite relational index (`.codemri/graph.db`) as the Single Source of Truth (**ADR-0001**), Vision performs microsecond dependency calculations without ever transmitting source code over the network (**ADR-0002**).

---

## 🌟 Key Architectural Innovations

### 1. 60FPS Interactive Force-Directed Physics Canvas
* **Pure Zero-CDN Execution:** To preserve strict 100% offline security, the physics calculation loop is built natively into our embedded Fiber HTML server without relying on external internet CDN dependencies (such as remote D3.js or Three.js CDN scripts).
* **Hooke's Spring Attraction & Coulomb Repulsion:** Node entities (functions, structs, classes, interfaces, imports) organically stabilize into intuitive architectural clusters, with relational bonds acting as attractive structural edges.

### 2. Sub-millisecond Architectural Impact Analysis (`/api/graph/impact/:id`)
* When an engineer clicks any symbol node directly on the physics canvas, CodeMRI instantly calculates an authoritative **Impact Radius & Blast Score (0-100%)**.
* **Upstream Caller Discovery:** Identifies every dependent symbol across all repository packages that relies on the target function or structural interface.
* **Downstream Dependency Mapping:** Outlines foundational internal dependencies required by the selected node.
* **Severity Grading & Refactoring Advice:** Categorizes structural changes into *Low Impact*, *Moderate Ripple Effect*, or *Critical Architectural Bottleneck*.

---

## 📡 API Specification

| Endpoint | Method | Response Description |
| :--- | :--- | :--- |
| `/api/graph/topology` | `GET` | Returns full node network and relational edge bonds for canvas mapping |
| `/api/graph/impact/:id` | `GET` | Evaluates blast radius, upstream callers, and downstream dependencies for symbol ID |

---

## 🛡️ Zero Cloud Transmissions & Enterprise Governance
Governed by **ADR-0002 (Local Privacy)** and **ADR-0001 (Single Source of Truth)**, all analytical processing and visual simulation logic run exclusively inside your local machine memory and SQLite WAL storage. No proprietary source code is ever shared with third-party servers.

---
*Licensed under Apache 2.0 • Built with ❤️ by Muhammad Nuril (@KangBasrengg)*
