package cli

// dashboardHTML is an embedded responsive modern dashboard UI that dynamically fetches
// and renders real analytical data from /api/repository, /api/graph/summary, and /api/pulse.
const dashboardHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>CodeMRI Dashboard — Neural Repository Intelligence (MRI v1.0.0)</title>
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
                <span class="storage-pill" id="storage-engine-badge">SQLite Relational & Full-Spectrum Intelligence Engine</span>
                <div class="status-badge">
                    <span class="pulse"></span>
                    <span id="engine-status">MRI v1.0.0 ACTIVE</span>
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

        <div class="section-wide" style="margin-bottom: 2rem;">
            <div style="display: flex; justify-content: space-between; align-items: center; flex-wrap: wrap; margin-bottom: 1.5rem;">
                <div>
                    <h2>🕸️ Phase 05 ("Vision") Interactive Force-Directed Topology Canvas</h2>
                    <p style="color: var(--text-muted); font-size: 0.9rem;">Click any structural node on the physics canvas to evaluate architectural impact, dependents, and ripple effects.</p>
                </div>
                <div>
                    <span id="canvas-status" style="font-family: monospace; font-size: 0.85rem; padding: 0.4rem 0.8rem; background: rgba(0, 242, 255, 0.1); border: 1px solid var(--accent-cyan); border-radius: 99px; color: var(--accent-cyan);">⚡ 60FPS PHYSICS ENGINE ONLINE</span>
                </div>
            </div>
            <div style="display: grid; grid-template-columns: repeat(auto-fit, minmax(450px, 1fr)); gap: 1.5rem; align-items: start;">
                <div style="position: relative; background: rgba(8, 13, 28, 0.95); border: 1px solid rgba(0, 242, 255, 0.25); border-radius: 16px; overflow: hidden; box-shadow: 0 10px 30px rgba(0,0,0,0.5);">
                    <canvas id="topologyCanvas" width="650" height="520" style="width: 100%; height: 520px; display: block; cursor: pointer;"></canvas>
                    <div style="position: absolute; bottom: 12px; left: 16px; font-size: 0.75rem; color: #94a3b8; font-family: monospace; background: rgba(0,0,0,0.7); padding: 4px 10px; border-radius: 6px;">
                        🖱️ Click node to analyze impact • Nodes: <span id="canvas-node-count" style="color:#38bdf8;">0</span> • Edges: <span id="canvas-edge-count" style="color:#f472b6;">0</span>
                    </div>
                </div>
                <div id="impact-panel" style="background: rgba(15, 23, 42, 0.95); border: 1px solid rgba(236, 72, 153, 0.4); border-radius: 16px; padding: 1.5rem; height: 520px; display: flex; flex-direction: column; justify-content: space-between; overflow-y: auto; box-shadow: 0 10px 30px rgba(0,0,0,0.5);">
                    <div>
                        <div class="panel-title" style="color: var(--accent-pink); font-size: 1rem;">🚨 Instant Architecture Impact Assessment</div>
                        <div id="impact-content">
                            <div style="padding: 3rem 1rem; text-align: center; color: var(--text-muted);">
                                <div style="font-size: 2.5rem; margin-bottom: 1rem;">👆</div>
                                <h3 style="color: var(--text-main); font-size: 1.1rem; margin-bottom: 0.5rem;">No Symbol Node Selected</h3>
                                <p style="font-size: 0.9rem; line-height: 1.6;">Click any glowing node on the left force-directed topology canvas to perform sub-millisecond cascading dependency analysis via SQLite NRG.</p>
                            </div>
                        </div>
                    </div>
                    <div style="border-top: 1px solid rgba(255,255,255,0.08); padding-top: 1rem; font-size: 0.75rem; color: #64748b; display: flex; justify-content: space-between;">
                        <span>Source of Truth: .codemri/graph.db</span>
                        <span style="color: var(--accent-green);">✔ Offline Protected</span>
                    </div>
                </div>
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
            CodeMRI v1.0.0 (MRI) • Licensed under Apache 2.0 for Enterprise Patent Security • Built with ❤️ by Muhammad Nuril
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
                
                // Initialize Phase 5 Vision Interactive Canvas & Physics
                initTopologyCanvas();
            } catch (err) {
                console.error("Failed to sync with local CodeMRI server:", err);
            }
        }

        // Phase 5 Vision: Interactive Force-Directed Physics Engine
        let selectedNode = null;
        let graphNodes = [];
        let graphEdges = [];
        let animFrame = null;

        async function initTopologyCanvas() {
            try {
                const res = await fetch('/api/graph');
                if (!res.ok) return;
                const graphData = await res.json();

                const nodeMap = graphData.nodes || {};
                const edgeList = graphData.edges || [];

                const canvas = document.getElementById('topologyCanvas');
                const ctx = canvas.getContext('2d');
                const width = canvas.width;
                const height = canvas.height;

                const nodeKeys = Object.keys(nodeMap).slice(0, 70); // Optimize for smooth visual clarity
                graphNodes = nodeKeys.map((key, index) => {
                    const n = nodeMap[key];
                    const angle = (index / nodeKeys.length) * Math.PI * 2;
                    const radius = 120 + Math.random() * 80;
                    return {
                        id: n.id,
                        name: n.name || n.id,
                        type: n.type || 'symbol',
                        path: n.path || 'internal/engine',
                        x: width / 2 + Math.cos(angle) * radius + (Math.random() - 0.5) * 40,
                        y: height / 2 + Math.sin(angle) * radius + (Math.random() - 0.5) * 40,
                        vx: 0,
                        vy: 0,
                        radius: n.type === 'package' || n.type === 'function' ? 10 : 7,
                        color: n.type === 'function' ? '#38bdf8' : n.type === 'struct' || n.type === 'class' ? '#f472b6' : '#a78bfa'
                    };
                });

                const validIds = new Set(graphNodes.map(n => n.id));
                graphEdges = edgeList.filter(e => validIds.has(e.source_id) && validIds.has(e.target_id)).slice(0, 100);

                document.getElementById('canvas-node-count').innerText = graphNodes.length;
                document.getElementById('canvas-edge-count').innerText = graphEdges.length;

                // Physics simulation loop
                function stepPhysics() {
                    for (let i = 0; i < graphNodes.length; i++) {
                        let n1 = graphNodes[i];
                        // Centering pull
                        n1.vx += (width / 2 - n1.x) * 0.003;
                        n1.vy += (height / 2 - n1.y) * 0.003;

                        for (let j = i + 1; j < graphNodes.length; j++) {
                            let n2 = graphNodes[j];
                            let dx = n2.x - n1.x;
                            let dy = n2.y - n1.y;
                            let dist = Math.sqrt(dx * dx + dy * dy) || 1;
                            if (dist < 180) {
                                let force = (180 - dist) * 0.03;
                                let fx = (dx / dist) * force;
                                let fy = (dy / dist) * force;
                                n1.vx -= fx;
                                n1.vy -= fy;
                                n2.vx += fx;
                                n2.vy += fy;
                            }
                        }
                    }

                    for (let e of graphEdges) {
                        let src = graphNodes.find(n => n.id === e.source_id);
                        let tgt = graphNodes.find(n => n.id === e.target_id);
                        if (src && tgt) {
                            let dx = tgt.x - src.x;
                            let dy = tgt.y - src.y;
                            let dist = Math.sqrt(dx * dx + dy * dy) || 1;
                            let pull = (dist - 90) * 0.015;
                            let fx = (dx / dist) * pull;
                            let fy = (dy / dist) * pull;
                            src.vx += fx;
                            src.vy += fy;
                            tgt.vx -= fx;
                            tgt.vy -= fy;
                        }
                    }

                    ctx.clearRect(0, 0, width, height);

                    // Draw edges
                    for (let e of graphEdges) {
                        let src = graphNodes.find(n => n.id === e.source_id);
                        let tgt = graphNodes.find(n => n.id === e.target_id);
                        if (src && tgt) {
                            ctx.beginPath();
                            ctx.moveTo(src.x, src.y);
                            ctx.lineTo(tgt.x, tgt.y);
                            ctx.strokeStyle = (selectedNode && (selectedNode.id === src.id || selectedNode.id === tgt.id)) ? '#f472b6' : 'rgba(0, 242, 255, 0.18)';
                            ctx.lineWidth = (selectedNode && (selectedNode.id === src.id || selectedNode.id === tgt.id)) ? 2 : 1;
                            ctx.stroke();
                        }
                    }

                    // Draw nodes
                    for (let n of graphNodes) {
                        n.x += n.vx;
                        n.y += n.vy;
                        n.vx *= 0.82;
                        n.vy *= 0.82;

                        ctx.beginPath();
                        ctx.arc(n.x, n.y, n.radius, 0, Math.PI * 2);
                        ctx.fillStyle = n.color;
                        ctx.shadowBlur = selectedNode && selectedNode.id === n.id ? 20 : 8;
                        ctx.shadowColor = selectedNode && selectedNode.id === n.id ? '#f472b6' : n.color;
                        ctx.fill();
                        ctx.shadowBlur = 0;

                        if (selectedNode && selectedNode.id === n.id) {
                            ctx.lineWidth = 3;
                            ctx.strokeStyle = '#f472b6';
                            ctx.stroke();
                        }

                        ctx.fillStyle = '#cbd5e1';
                        ctx.font = '10px monospace';
                        ctx.fillText(n.name.length > 15 ? n.name.substring(0, 12) + '...' : n.name, n.x + 14, n.y + 4);
                    }

                    animFrame = requestAnimationFrame(stepPhysics);
                }

                if (animFrame) cancelAnimationFrame(animFrame);
                stepPhysics();

                canvas.addEventListener('click', async (evt) => {
                    const rect = canvas.getBoundingClientRect();
                    const scaleX = canvas.width / rect.width;
                    const scaleY = canvas.height / rect.height;
                    const mx = (evt.clientX - rect.left) * scaleX;
                    const my = (evt.clientY - rect.top) * scaleY;

                    for (let n of graphNodes) {
                        let dx = n.x - mx;
                        let dy = n.y - my;
                        if (Math.sqrt(dx * dx + dy * dy) <= n.radius + 10) {
                            selectedNode = n;
                            await loadNodeImpact(n.id);
                            return;
                        }
                    }
                });
            } catch (err) {
                console.error("Canvas topology failed:", err);
            }
        }

        async function loadNodeImpact(nodeId) {
            const container = document.getElementById('impact-content');
            container.innerHTML = '<div style="padding: 2rem; text-align: center; color: var(--text-muted);">🔄 Calculating real-time impact radius via SQLite NRG...</div>';
            try {
                const res = await fetch('/api/graph/impact/' + encodeURIComponent(nodeId));
                if (!res.ok) {
                    container.innerHTML = '<div style="color: #ef4444;">Failed to calculate impact.</div>';
                    return;
                }
                const data = await res.json();
                const node = data.target_node;
                const score = data.impact_score || 0;
                const scoreColor = score >= 70 ? '#ef4444' : score >= 40 ? '#f59e0b' : '#34d399';

                const upCount = (data.upstream_dependents || []).length;
                const downCount = (data.downstream_dependencies || []).length;

                container.innerHTML = '<div style="margin-top: 1rem;">' +
                    '<div style="display: flex; justify-content: space-between; align-items: start;">' +
                        '<div>' +
                            '<h3 style="color: #ffffff; font-size: 1.15rem; margin-bottom: 0.2rem;">' + (node.name || node.id) + '</h3>' +
                            '<span style="font-family: monospace; font-size: 0.75rem; color: #38bdf8; background: rgba(56,189,248,0.1); padding: 2px 6px; border-radius: 4px;">' + node.type + '</span>' +
                        '</div>' +
                        '<div style="text-align: right;">' +
                            '<div style="font-weight: 800; font-size: 1.4rem; color: ' + scoreColor + ';">' + score + '%</div>' +
                            '<div style="font-size: 0.7rem; color: var(--text-muted);">IMPACT RADIUS</div>' +
                        '</div>' +
                    '</div>' +
                    '<div style="margin: 1rem 0; font-family: monospace; font-size: 0.8rem; color: #94a3b8; background: rgba(0,0,0,0.4); padding: 0.6rem; border-radius: 8px;">' +
                        '📂 ' + (node.path || 'root') + ' (L' + (node.start_line || 1) + '-L' + (node.end_line || 50) + ')' +
                    '</div>' +
                    '<div style="padding: 0.8rem; border-left: 3px solid ' + scoreColor + '; background: rgba(255,255,255,0.02); margin-bottom: 1rem;">' +
                        '<div style="font-size: 0.75rem; text-transform: uppercase; color: ' + scoreColor + '; font-weight: bold; margin-bottom: 0.3rem;">⚡ Severity Assessment</div>' +
                        '<div style="font-size: 0.85rem; color: #e2e8f0;">' + data.severity + '</div>' +
                        '<div style="font-size: 0.8rem; color: #94a3b8; margin-top: 0.4rem;">' + data.advice + '</div>' +
                    '</div>' +
                    '<div style="display: grid; grid-template-columns: 1fr 1fr; gap: 0.8rem; text-align: center;">' +
                        '<div style="background: rgba(0,242,255,0.05); border: 1px solid rgba(0,242,255,0.2); padding: 0.8rem; border-radius: 10px;">' +
                            '<div style="font-size: 1.3rem; font-weight: 800; color: #38bdf8;">' + upCount + '</div>' +
                            '<div style="font-size: 0.75rem; color: #94a3b8;">Upstream Callers</div>' +
                        '</div>' +
                        '<div style="background: rgba(236,72,153,0.05); border: 1px solid rgba(236,72,153,0.2); padding: 0.8rem; border-radius: 10px;">' +
                            '<div style="font-size: 1.3rem; font-weight: 800; color: #f472b6;">' + downCount + '</div>' +
                            '<div style="font-size: 0.75rem; color: #94a3b8;">Dependencies</div>' +
                        '</div>' +
                    '</div>' +
                '</div>';
            } catch (err) {
                container.innerHTML = '<div style="color: #ef4444;">Error evaluating node impact.</div>';
            }
        }
    </script>
</body>
</html>`
