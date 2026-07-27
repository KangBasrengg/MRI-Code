import React, { useState } from 'react';
import { Terminal, CheckCircle2, Sparkles, FolderGit2, Layers, Globe } from 'lucide-react';
import { ReactFlow, Background, Controls } from '@xyflow/react';
import type { Node, Edge } from '@xyflow/react';

export const LiveScanner: React.FC = () => {
  const [url, setUrl] = useState<string>('https://github.com/torvalds/linux');
  const [scanning, setScanning] = useState<boolean>(false);
  const [scanComplete, setScanComplete] = useState<boolean>(false);
  const [logs, setLogs] = useState<string[]>([]);

  const sampleUrls = [
    'https://github.com/golang/go',
    'https://github.com/tailwindlabs/tailwindcss',
    'https://github.com/gofiber/fiber',
    'https://github.com/sveltejs/svelte'
  ];

  const mockNodes: Node[] = [
    { id: '1', data: { label: '🌐 Target Repo: ' + url.split('/').pop() }, position: { x: 250, y: 30 }, style: { background: '#0e1726', border: '2px solid #00f2ff', color: '#00f2ff', fontWeight: 'bold' } },
    { id: '2', data: { label: '📦 Core Engine & AST Trees' }, position: { x: 250, y: 160 }, style: { background: '#0e1726', border: '1px solid #3b82f6', color: '#60a5fa', fontWeight: 'bold' } },
    { id: '3', data: { label: '⚙️ API & Network Layer' }, position: { x: 90, y: 290 }, style: { background: '#0e1726', border: '1px solid #8a2be2', color: '#c084fc' } },
    { id: '4', data: { label: '🗃️ Storage & Cache Sinks' }, position: { x: 410, y: 290 }, style: { background: '#0e1726', border: '1px solid #f59e0b', color: '#fbbf24' } },
    { id: '5', data: { label: '🛡️ Verified Security Boundaries' }, position: { x: 250, y: 410 }, style: { background: '#0e1726', border: '1px solid #10b981', color: '#34d399' } }
  ];

  const mockEdges: Edge[] = [
    { id: 'e1-2', source: '1', target: '2', animated: true, style: { stroke: '#00f2ff', strokeWidth: 2 } },
    { id: 'e2-3', source: '2', target: '3', style: { stroke: '#8a2be2' } },
    { id: 'e2-4', source: '2', target: '4', style: { stroke: '#f59e0b' } },
    { id: 'e3-5', source: '3', target: '5', animated: true, style: { stroke: '#10b981' } },
    { id: 'e4-5', source: '4', target: '5', animated: true, style: { stroke: '#10b981' } }
  ];

  const triggerScan = (target: string = url) => {
    setUrl(target);
    setScanning(true);
    setScanComplete(false);
    setLogs([
      '💓 Executing Pulse architectural intelligence scan (v0.4.0)...',
      '📂 Initializing universal AST syntax classification parser...'
    ]);

    setTimeout(() => {
      setLogs((prev) => [...prev, '✔ Downloaded repository metadata and directory trees', '⚡ Running deterministic Tree-sitter parsers...']);
    }, 1200);

    setTimeout(() => {
      setLogs((prev) => [...prev, '🧠 Building Neural Repository Graph (NRG) across 2,450 discovered modules...', '✔ Computed circular dependency count: 0']);
    }, 2800);

    setTimeout(() => {
      setLogs((prev) => [...prev, '✨ Simulation complete! Health Score: 97/100. Rendering Interactive Architecture Blueprint...']);
      setScanning(false);
      setScanComplete(true);
    }, 4200);
  };

  return (
    <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 py-12 text-left">
      <div className="text-center mb-10">
        <div className="inline-flex items-center space-x-2 text-cyan-400 font-mono text-xs uppercase px-3 py-1 bg-cyan-500/10 border border-cyan-500/30 rounded-full font-bold mb-3">
          <Sparkles className="w-4 h-4 text-amber-400" />
          <span>Interactive Online Sandbox</span>
        </div>
        <h2 className="text-3xl sm:text-5xl font-black text-white">
          Test Scan Any Public Repository
        </h2>
        <p className="text-slate-400 max-w-2xl mx-auto text-sm sm:text-base mt-2">
          Paste a valid GitHub Repository URL or choose a famous open-source example to watch CodeMRI simulate generating its Neural Repository Graph on the fly!
        </p>
      </div>

      {/* Input Form */}
      <div className="glass-panel p-6 sm:p-8 rounded-3xl border-slate-800 shadow-2xl mb-8">
        <div className="flex flex-col sm:flex-row items-center gap-4">
          <div className="relative flex-1 w-full">
            <FolderGit2 className="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-500" />
            <input
              type="text"
              value={url}
              onChange={(e) => setUrl(e.target.value)}
              placeholder="https://github.com/username/repository"
              className="w-full bg-slate-950 text-white font-mono text-sm pl-12 pr-4 py-4 rounded-2xl border border-slate-800 focus:border-cyan-500 focus:outline-none transition-all shadow-inner"
            />
          </div>
          <button
            onClick={() => triggerScan(url)}
            disabled={scanning}
            className="w-full sm:w-auto px-8 py-4 bg-gradient-to-r from-cyan-400 via-blue-500 to-purple-600 hover:from-cyan-300 text-slate-950 font-extrabold text-base rounded-2xl shadow-lg shadow-cyan-500/30 hover:scale-105 transition-all flex items-center justify-center space-x-2 disabled:opacity-50"
          >
            <Terminal className="w-5 h-5 text-slate-950" />
            <span>{scanning ? 'Scanning...' : 'Run Scan Now'}</span>
          </button>
        </div>

        {/* Sample presets */}
        <div className="mt-6 flex flex-wrap items-center gap-2 text-xs font-mono text-slate-400">
          <span>Try famous examples:</span>
          {sampleUrls.map((s, idx) => (
            <button
              key={idx}
              onClick={() => triggerScan(s)}
              className="px-3 py-1.5 rounded-lg bg-slate-900 border border-slate-800 hover:border-cyan-500/50 text-slate-300 hover:text-white transition-all"
            >
              {s.split('/').slice(-2).join('/')}
            </button>
          ))}
        </div>
      </div>

      {/* Output Console & Interactive Graph */}
      <div className="grid grid-cols-1 lg:grid-cols-12 gap-8 items-stretch">
        
        {/* Terminal Logs */}
        <div className="lg:col-span-5 bg-slate-950 border border-slate-800 rounded-3xl p-6 font-mono text-xs sm:text-sm text-slate-300 shadow-xl flex flex-col justify-between h-[500px]">
          <div>
            <div className="flex items-center justify-between pb-4 mb-4 border-b border-slate-800/80">
              <span className="font-bold text-cyan-400 flex items-center space-x-2">
                <Terminal className="w-4 h-4" />
                <span>Execution Logs</span>
              </span>
              <span className="text-[11px] text-slate-500">Under 60s target</span>
            </div>
            <div className="space-y-2 overflow-y-auto max-h-[380px] text-left">
              {logs.length === 0 ? (
                <p className="text-slate-600 italic">Click 'Run Scan Now' above or pick an example to launch simulation.</p>
              ) : (
                logs.map((line, i) => (
                  <p key={i} className={line.startsWith('✨') ? 'text-emerald-400 font-bold' : line.startsWith('🧠') ? 'text-purple-300' : 'text-slate-300'}>
                    {line}
                  </p>
                ))
              )}
            </div>
          </div>
          {scanComplete && (
            <div className="mt-4 p-3.5 bg-emerald-950/40 border border-emerald-500/40 rounded-2xl text-emerald-300 text-xs font-sans flex items-center justify-between">
              <span className="font-bold flex items-center space-x-2">
                <CheckCircle2 className="w-4 h-4 text-emerald-400" />
                <span>NRG Generated</span>
              </span>
              <span className="font-mono bg-emerald-500/20 px-2 py-0.5 rounded font-bold">Health: 97/100</span>
            </div>
          )}
        </div>

        {/* Graph Preview */}
        <div className="lg:col-span-7 bg-[#0b0f19] border-2 border-slate-800 rounded-3xl overflow-hidden h-[500px] relative">
          <div className="bg-slate-900 px-4 py-3 border-b border-slate-800 flex items-center justify-between">
            <span className="text-xs sm:text-sm font-bold text-white flex items-center space-x-2">
              <Layers className="w-4 h-4 text-cyan-400" />
              <span>Simulated NRG Blueprint</span>
            </span>
            <span className="text-[11px] font-mono text-slate-400">Offline-first standard</span>
          </div>

          {scanComplete ? (
            <div className="w-full h-full pb-10">
              <ReactFlow
                nodes={mockNodes}
                edges={mockEdges}
                fitView
                key={url}
                proOptions={{ hideAttribution: true }}
              >
                <Background color="#1e293b" gap={20} size={1} />
                <Controls className="bg-slate-900 border border-slate-700 fill-white rounded-xl" />
              </ReactFlow>
            </div>
          ) : (
            <div className="absolute inset-0 flex flex-col items-center justify-center text-slate-500 p-6 text-center space-y-4">
              {scanning ? (
                <>
                  <div className="w-12 h-12 rounded-full border-4 border-cyan-400 border-t-transparent animate-spin" />
                  <p className="font-mono text-cyan-300 font-bold text-sm">Traversing files & building SQLite NRG Graph...</p>
                </>
              ) : (
                <>
                  <Globe className="w-16 h-16 text-slate-800 stroke-[1]" />
                  <p className="text-sm font-sans text-slate-400 max-w-sm">
                    Enter a repository link above to view an interactive visual map of its code execution architecture!
                  </p>
                </>
              )}
            </div>
          )}
        </div>

      </div>

    </div>
  );
};
