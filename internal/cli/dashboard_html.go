package cli

// dashboardHTML is an embedded responsive modern dashboard UI that dynamically fetches
// and renders real analytical data from /api/repository, /api/graph/summary, and /api/pulse.
const dashboardHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>CodeMRI Dashboard — Neural Repository Intelligence (Pulse v0.4.0)</title>
    <style>
        :root {
            --bg-base: #080b12;
            --bg-panel: rgba(15, 22, 35, 0.8);
            --border-color: rgba(255, 255, 255, 0.1);
            --accent-cyan: #00f2ff;
            --accent-blue: #3b82f6;
            --accent-emerald: #10b981;
            --accent-pink: #ec4899;
            --accent-amber: #f59e0b;
            --text-main: #e2e8f0;
            --text-muted: #94a3b8;
        }

        * {
            box-sizing: border-box;
            margin: 0;
            padding: 0;
        }

        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Inter', 'Segoe UI', Roboto, sans-serif;
            background: linear-gradient(135deg, #05080f 0%, #0b111e 50%, #080c16 100%);
            color: var(--text-main);
            min-height: 100vh;
            padding: 2rem;
            line-height: 1.6;
        }

        .container {
            max-width: 1200px;
            margin: 0 auto;
        }

        header {
            display: flex;
            justify-content: space-between;
            align-items: center;
            padding-bottom: 2rem;
            border-bottom: 1px solid var(--border-color);
            margin-bottom: 2.5rem;
            flex-wrap: wrap;
            gap: 1rem;
        }

        .logo-area h1 {
            font-size: 1.8rem;
            font-weight: 900;
            background: linear-gradient(90deg, #00f2ff, #38bdf8, #ec4899);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
        }

        .logo-area p {
            color: var(--text-muted);
            font-size: 0.9rem;
            margin-top: 0.2rem;
        }

        .status-badge {
            background: rgba(236, 72, 153, 0.12);
            border: 1px solid rgba(236, 72, 153, 0.35);
            color: #f472b6;
            padding: 0.5rem 1.2rem;
            border-radius: 9999px;
            font-weight: 700;
            font-size: 0.85rem;
            display: inline-flex;
            align-items: center;
            gap: 0.6rem;
            box-shadow: 0 0 25px rgba(236, 72, 153, 0.25);
        }

        .pulse {
            width: 9px;
            height: 9px;
            background-color: #f472b6;
            border-radius: 50%;
            display: inline-block;
            box-shadow: 0 0 10px #f472b6;
        }

        .grid {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
            gap: 1.5rem;
            margin-bottom: 2.5rem;
        }

        .panel {
            background: var(--bg-panel);
            border: 1px solid var(--border-color);
            border-radius: 1.25rem;
            padding: 1.75rem;
            box-shadow: 0 20px 40px rgba(0, 0, 0, 0.4);
            backdrop-filter: blur(16px);
            display: flex;
            flex-direction: column;
            justify-content: space-between;
        }

        .panel-title {
            font-size: 0.9rem;
            text-transform: uppercase;
            letter-spacing: 0.08em;
            color: var(--accent-cyan);
            font-weight: 800;
            margin-bottom: 1.25rem;
            display: flex;
            align-items: center;
            gap: 0.5rem;
        }

        .metric-big {
            font-size: 2.75rem;
            font-weight: 900;
            color: #ffffff;
            line-height: 1;
            margin-bottom: 0.5rem;
        }

        .metric-sub {
            color: var(--text-muted);
            font-size: 0.85rem;
        }

        .lang-list, .node-list {
            list-style: none;
            display: flex;
            flex-direction: column;
            gap: 0.8rem;
        }

        .list-item {
            display: flex;
            justify-content: space-between;
            align-items: center;
            padding-bottom: 0.6rem;
            border-bottom: 1px dashed rgba(255, 255, 255, 0.08);
            font-size: 0.95rem;
        }

        .list-item:last-child {
            border-bottom: none;
            padding-bottom: 0;
        }

        .tag {
            background: rgba(255, 255, 255, 0.06);
            padding: 0.2rem 0.6rem;
            border-radius: 6px;
            font-family: monospace;
            font-size: 0.85rem;
            color: #7dd3fc;
        }

        .section-wide {
            background: var(--bg-panel);
            border: 1px solid var(--border-color);
            border-radius: 1.25rem;
            padding: 2.5rem;
            box-shadow: 0 25px 50px rgba(0, 0, 0, 0.5);
            margin-bottom: 2.5rem;
        }

        .section-wide h2 {
            font-size: 1.5rem;
            margin-bottom: 1rem;
            color: #ffffff;
            display: flex;
            align-items: center;
            gap: 0.5rem;
        }

        .pulse-banner {
            background: rgba(236, 72, 153, 0.08);
            border: 1px solid rgba(236, 72, 153, 0.25);
            border-radius: 1rem;
            padding: 1.5rem;
            color: #fbcfe8;
            margin-top: 1.5rem;
            font-size: 0.95rem;
            display: flex;
            gap: 1.5rem;
            align-items: center;
        }

        .storage-pill {
            background: rgba(59, 130, 246, 0.15);
            border: 1px solid rgba(59, 130, 246, 0.4);
            color: #60a5fa;
            font-size: 0.75rem;
            padding: 0.2rem 0.6rem;
            border-radius: 6px;
            font-family: monospace;
            font-weight: 700;
        }

        .suggestion-box {
            background: rgba(0, 0, 0, 0.3);
            border-left: 4px solid var(--accent-cyan);
            padding: 0.8rem 1.2rem;
            border-radius: 0 8px 8px 0;
            margin-top: 0.8rem;
            font-size: 0.9rem;
            color: #e2e8f0;
        }

        footer {
            text-align: center;
            color: #64748b;
            font-size: 0.85rem;
            padding-top: 2rem;
            border-top: 1px solid var(--border-color);
        }
    </style>
</head>
<body>
    <div class="container">
        <header>
            <div class="logo-area">
                <h1>🧠 CodeMRI Dashboard</h1>
                <p>Offline-First Neural Repository Intelligence Platform</p>
            </div>
            <div style="display: flex; gap: 0.8rem; align-items: center; flex-wrap: wrap;">
                <span class="storage-pill" id="storage-engine-badge">SQLite Relational & Pulse Engine</span>
                <div class="status-badge">
                    <span class="pulse"></span>
                    <span id="engine-status">PULSE v0.4.0 ACTIVE</span>
                </div>
            </div>
        </header>

        <div class="grid">
            <div class="panel" style="border-color: rgba(236, 72, 153, 0.3);">
                <div>
                    <div class="panel-title" style="color: var(--accent-pink);">💓 Repository Health Score</div>
                    <div class="metric-big" id="health-score">100 / 100</div>
                </div>
                <div class="metric-sub"><span id="health-grade-text" style="font-weight: bold; color: #34d399;">Grade A+</span> • <span id="health-debt">Evaluating architectural technical debt...</span></div>
            </div>
            <div class="panel">
                <div>
                    <div class="panel-title">🧟 Dead Code & Isolation</div>
                    <div class="metric-big" id="dead-code-count">0</div>
                </div>
                <div class="metric-sub">Symbols receiving zero incoming relational bindings</div>
            </div>
            <div class="panel">
                <div>
                    <div class="panel-title">🔁 Circular Dependency Loops</div>
                    <div class="metric-big" id="circular-count">0</div>
                </div>
                <div class="metric-sub">Cyclical import bindings in SQLite NRG</div>
            </div>
        </div>

        <div class="grid">
            <div class="panel">
                <div>
                    <div class="panel-title">⚡ Repository Scale</div>
                    <div class="metric-big" id="loc-count">...</div>
                </div>
                <div class="metric-sub"><span id="files-count">0</span> Files Analyzed via Single Source of Truth</div>
            </div>
            <div class="panel">
                <div>
                    <div class="panel-title">🛡️ Offline Privacy & Engine</div>
                    <div class="metric-big" style="color: #34d399;">100% Offline</div>
                </div>
                <div class="metric-sub">SQLite Indexed Execution • Zero Raw Code Transmitted</div>
            </div>
            <div class="panel">
                <div>
                    <div class="panel-title">🧭 Neural Graph Scale</div>
                    <div class="metric-big" id="nodes-count">...</div>
                </div>
                <div class="metric-sub"><span id="edges-count">0</span> Relational Dependency Edges Mapped</div>
            </div>
        </div>

        <div class="grid" style="grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));">
            <div class="panel">
                <div class="panel-title">📊 Language Distribution</div>
                <ul class="lang-list" id="lang-distribution">
                    <li class="list-item"><span>Loading repository metadata...</span></li>
                </ul>
            </div>
            <div class="panel">
                <div class="panel-title">🏗️ Neural Symbol Topology</div>
                <ul class="node-list" id="node-distribution">
                    <li class="list-item"><span>Querying SQLite structural nodes...</span></li>
                </ul>
            </div>
            <div class="panel">
                <div class="panel-title">⚡ Relational Dependency Edges</div>
                <ul class="node-list" id="edge-distribution">
                    <li class="list-item"><span>Querying relational bonds...</span></li>
                </ul>
            </div>
        </div>

        <div class="section-wide">
            <h2>💎 Phase 04 ("Pulse") Architectural Technical Debt & Health Online</h2>
            <p style="color: var(--text-muted); line-height: 1.8;">
                In Phase 4, CodeMRI integrates advanced technical debt algorithms directly over the <strong>Neural Repository Graph (NRG)</strong>. By treating the SQLite graph database at 
                <code style="color: #f472b6; background: rgba(236, 72, 153, 0.1); padding: 0.2rem 0.5rem; border-radius: 4px;">.codemri/pulse.json</code> as the Single Source of Truth, CodeMRI instantly calculates cyclomatic symbol density, discovers unreferenced dead code, and exposes architectural circular dependency loops without cloud latency.
            </p>
            <div id="suggestions-area" style="margin-top: 1.5rem;">
                <div class="suggestion-box">✨ Initializing automated AI reasoning structural recommendations...</div>
            </div>
        </div>

        <footer>
            CodeMRI v0.4.0 (Pulse) • Licensed under Apache 2.0 for Enterprise Patent Security • Built with ❤️ by Muhammad Nuril
        </footer>
    </div>

    <script>
        // Fetch Real Analytics from Local CodeMRI SQLite/Fiber API
        async function initDashboard() {
            try {
                // Fetch Repository Metadata (.codemri/repository.json)
                const repoRes = await fetch('/api/repository');
                if (repoRes.ok) {
                    const repoData = await repoRes.json();
                    document.getElementById('loc-count').innerText = (repoData.total_loc || 0).toLocaleString();
                    document.getElementById('files-count').innerText = (repoData.total_files || 0).toLocaleString();
                    
                    const langUl = document.getElementById('lang-distribution');
                    langUl.innerHTML = '';
                    const langs = repoData.language_stats || repoData.languages || {};
                    if (Object.keys(langs).length === 0) {
                        langUl.innerHTML = '<li class="list-item"><span>No language data found</span></li>';
                    } else {
                        for (const [lang, data] of Object.entries(langs)) {
                            const loc = typeof data === 'number' ? data : (data.Lines || 0);
                            const li = document.createElement('li');
                            li.className = 'list-item';
                            li.innerHTML = '<span>' + lang + '</span><span class="tag">' + loc.toLocaleString() + ' LOC</span>';
                            langUl.appendChild(li);
                        }
                    }
                }

                // Fetch NRG Graph Topology Summary from SQLite (.codemri/graph.db)
                const graphRes = await fetch('/api/graph/summary');
                if (graphRes.ok) {
                    const graphData = await graphRes.json();
                    document.getElementById('nodes-count').innerText = (graphData.total_nodes || 0).toLocaleString() + ' Nodes';
                    document.getElementById('edges-count').innerText = (graphData.total_edges || 0).toLocaleString();
                    
                    if (graphData.storage_engine) {
                        document.getElementById('storage-engine-badge').innerText = graphData.storage_engine;
                    }

                    // Render Node Types
                    const nodeUl = document.getElementById('node-distribution');
                    nodeUl.innerHTML = '';
                    const nodeTypes = graphData.node_types || {};
                    if (Object.keys(nodeTypes).length === 0) {
                        nodeUl.innerHTML = '<li class="list-item"><span>No structural symbols found</span></li>';
                    } else {
                        for (const [type, count] of Object.entries(nodeTypes)) {
                            const li = document.createElement('li');
                            li.className = 'list-item';
                            li.innerHTML = '<span>' + type + '</span><span class="tag" style="color: #34d399;">' + count.toLocaleString() + ' Symbols</span>';
                            nodeUl.appendChild(li);
                        }
                    }

                    // Render Edge Types
                    const edgeUl = document.getElementById('edge-distribution');
                    edgeUl.innerHTML = '';
                    const edgeTypes = graphData.edge_types || {};
                    if (Object.keys(edgeTypes).length === 0) {
                        edgeUl.innerHTML = '<li class="list-item"><span>No relational edges categorized yet</span></li>';
                    } else {
                        for (const [type, count] of Object.entries(edgeTypes)) {
                            const li = document.createElement('li');
                            li.className = 'list-item';
                            li.innerHTML = '<span>' + type + '</span><span class="tag" style="color: #60a5fa;">' + count.toLocaleString() + ' Bonds</span>';
                            edgeUl.appendChild(li);
                        }
                    }
                }

                // Fetch Pulse Analytics & Technical Debt (.codemri/pulse.json)
                const pulseRes = await fetch('/api/pulse');
                if (pulseRes.ok) {
                    const pulseData = await pulseRes.json();
                    if (pulseData.health) {
                        const h = pulseData.health;
                        document.getElementById('health-score').innerText = h.overall_score + ' / 100';
                        document.getElementById('health-grade-text').innerText = 'Grade ' + h.grade;
                        document.getElementById('health-debt').innerText = h.debt_status;
                        
                        if (h.overall_score < 65) {
                            document.getElementById('health-score').style.color = '#ef4444';
                        } else if (h.overall_score < 80) {
                            document.getElementById('health-score').style.color = '#f59e0b';
                        }
                    }
                    
                    const deadCount = (pulseData.dead_code_issues || []).length;
                    const circCount = (pulseData.circular_dependencies || []).length;
                    document.getElementById('dead-code-count').innerText = deadCount.toLocaleString();
                    document.getElementById('circular-count').innerText = circCount.toLocaleString();

                    if (deadCount > 0) document.getElementById('dead-code-count').style.color = '#f59e0b';
                    if (circCount > 0) document.getElementById('circular-count').style.color = '#ef4444';

                    // Render AI Reasoning Suggestions
                    const suggArea = document.getElementById('suggestions-area');
                    if (pulseData.health && pulseData.health.suggestions && pulseData.health.suggestions.length > 0) {
                        suggArea.innerHTML = '<h3 style="font-size: 1.1rem; color: #38bdf8; margin-bottom: 0.5rem;">💡 Actionable Architectural Advice:</h3>';
                        for (const s of pulseData.health.suggestions) {
                            const div = document.createElement('div');
                            div.className = 'suggestion-box';
                            div.innerText = s;
                            suggArea.appendChild(div);
                        }
                    }
                }
            } catch (err) {
                console.error("Failed to sync with local CodeMRI server:", err);
            }
        }

        window.addEventListener('DOMContentLoaded', initDashboard);
    </script>
</body>
</html>`
