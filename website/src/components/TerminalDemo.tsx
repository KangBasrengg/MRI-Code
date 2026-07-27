import React, { useState, useEffect } from 'react';
import { ReactFlow, Background, Controls } from '@xyflow/react';
import type { Node, Edge } from '@xyflow/react';
import { RotateCcw, Layers } from 'lucide-react';

export const TerminalDemo: React.FC = () => {
  const [step, setStep] = useState<number>(0);
  const [isSimulating, setIsSimulating] = useState<boolean>(false);

  const initialNodes: Node[] = [
    { id: 'cli', data: { label: '💻 codemri scan .' }, position: { x: 280, y: 30 }, style: { background: '#0e1726', border: '2px solid #00f2ff', color: '#00f2ff', fontWeight: 'bold', borderRadius: '10px' } },
    { id: 'tree', data: { label: '⚡ Tree-sitter Parser (Go/TS/Python)' }, position: { x: 250, y: 150 }, style: { background: '#0e1726', border: '2px solid #3b82f6', color: '#60a5fa', fontWeight: 'bold' } },
    { id: 'nrg', data: { label: '🧠 Neural Repository Graph (SQLite)' }, position: { x: 230, y: 280 }, style: { background: '#0e1726', border: '2px solid #8a2be2', color: '#d8b4fe', padding: '12px 20px', fontWeight: 'bold' } },
    { id: 'ui', data: { label: '👁️ Architecture Dashboard' }, position: { x: 60, y: 410 }, style: { background: '#0e1726', border: '1px solid #10b981', color: '#34d399' } },
    { id: 'sec', data: { label: '🛡️ Security Hotspot Check' }, position: { x: 260, y: 410 }, style: { background: '#0e1726', border: '1px solid #f59e0b', color: '#fbbf24' } },
    { id: 'ai', data: { label: '🤖 AI Reasoning Engine' }, position: { x: 470, y: 410 }, style: { background: '#0e1726', border: '1px solid #ef4444', color: '#fca5a5' } },
  ];

  const initialEdges: Edge[] = [
    { id: 'e-1', source: 'cli', target: 'tree', animated: true, style: { stroke: '#00f2ff', strokeWidth: 2 }, label: 'traverses files' },
    { id: 'e-2', source: 'tree', target: 'nrg', animated: true, style: { stroke: '#3b82f6', strokeWidth: 2 }, label: 'generates AST' },
    { id: 'e-3', source: 'nrg', target: 'ui', style: { stroke: '#10b981', strokeWidth: 1.5 } },
    { id: 'e-4', source: 'nrg', target: 'sec', style: { stroke: '#f59e0b', strokeWidth: 1.5 } },
    { id: 'e-5', source: 'nrg', target: 'ai', style: { stroke: '#ef4444', strokeWidth: 1.5 }, label: 'zero token waste' }
  ];

  const runSimulation = () => {
    setIsSimulating(true);
    setStep(1); // Terminal input
    setTimeout(() => setStep(2), 1500); // Parsing files
    setTimeout(() => setStep(3), 3200); // NRG generation
    setTimeout(() => {
      setStep(4); // Dashboard complete!
      setIsSimulating(false);
    }, 5000);
  };

  useEffect(() => {
    runSimulation();
  }, []);

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
      <div className="text-center mb-10">
        <span className="text-xs uppercase font-mono px-3 py-1 bg-cyan-500/10 text-cyan-400 border border-cyan-500/30 rounded-full font-bold">
          🎬 Interactive 15-Second Showcase
        </span>
        <h2 className="text-3xl sm:text-5xl font-black text-white mt-3 mb-4">
          See How Code-MRI Works Live
        </h2>
        <p className="text-slate-400 max-w-2xl mx-auto text-sm sm:text-base">
          Watch how executing one terminal command instantaneously generates an intelligent interactive architecture graph.
        </p>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-12 gap-8 items-stretch">
        
        {/* Left: Terminal Simulation */}
        <div className="lg:col-span-5 bg-[#0b0f19] border border-slate-800 rounded-3xl overflow-hidden shadow-2xl flex flex-col justify-between">
          <div>
            <div className="bg-slate-900 px-4 py-3 border-b border-slate-800 flex items-center justify-between">
              <div className="flex items-center space-x-2">
                <span className="w-3 h-3 rounded-full bg-rose-500 inline-block" />
                <span className="w-3 h-3 rounded-full bg-amber-500 inline-block" />
                <span className="w-3 h-3 rounded-full bg-emerald-500 inline-block" />
                <span className="text-xs font-mono text-slate-400 ml-2">bash — codemri-cli</span>
              </div>
              <button
                onClick={runSimulation}
                disabled={isSimulating}
                className="flex items-center space-x-1 px-2.5 py-1 bg-slate-800 hover:bg-slate-700 text-cyan-400 text-xs font-mono rounded-lg border border-slate-700 transition-all disabled:opacity-50"
              >
                <RotateCcw className={`w-3 h-3 ${isSimulating ? 'animate-spin' : ''}`} />
                <span>Re-run</span>
              </button>
            </div>

            <div className="p-5 font-mono text-xs sm:text-sm text-slate-300 space-y-3 text-left">
              <div className="flex items-center space-x-2 text-cyan-400">
                <span>$</span>
                <span>codemri scan . --verbose</span>
              </div>

              {step >= 2 && (
                <div className="text-slate-400 space-y-1.5 animate-fadeIn">
                  <p className="text-amber-400/90">📡 [CodeMRI] Scanning repository root: /workspace/enterprise-app</p>
                  <p>✔ Discovered 4,210 source files across 48 packages</p>
                  <p>⚡ Engaging deterministic Tree-sitter parsers (Go, TS, SQL)...</p>
                  <p>✔ Extracted 12,480 Abstract Syntax Tree nodes</p>
                </div>
              )}

              {step >= 3 && (
                <div className="text-purple-300 space-y-1.5 animate-fadeIn">
                  <p>🧠 Building Neural Repository Graph (.codemri/graph.db)...</p>
                  <p>✔ Resolved 18,920 cross-file relational edges</p>
                  <p>✔ Computed architecture health score: <span className="text-emerald-400 font-bold">96/100</span></p>
                </div>
              )}

              {step >= 4 && (
                <div className="pt-2 border-t border-slate-800/80 text-emerald-400 space-y-1 font-bold animate-fadeIn">
                  <p>✨ Genesis scan completed in 0.84 seconds!</p>
                  <p className="text-cyan-300">👉 Launching visual server on http://localhost:4000</p>
                </div>
              )}
            </div>
          </div>

          {/* Status badge */}
          <div className="p-4 bg-slate-900/50 border-t border-slate-800/80 flex items-center justify-between font-mono text-xs text-slate-400">
            <span className="flex items-center space-x-2">
              <span className={`w-2 h-2 rounded-full ${step === 4 ? 'bg-emerald-400' : 'bg-amber-400 animate-pulse'}`} />
              <span>{step === 4 ? 'NRG Compiled Successfully' : 'Executing Fast AST Walk...'}</span>
            </span>
            <span className="text-slate-500">v0.4.0 Pulse Engine</span>
          </div>
        </div>

        {/* Right: Live Interactive Architecture Graph */}
        <div className="lg:col-span-7 bg-[#0f1623]/80 border-2 border-slate-800 rounded-3xl overflow-hidden shadow-2xl h-[520px] relative flex flex-col">
          <div className="bg-slate-900/90 px-4 py-3 border-b border-slate-800 flex items-center justify-between z-10">
            <div className="flex items-center space-x-2">
              <Layers className="w-4 h-4 text-cyan-400" />
              <span className="text-xs sm:text-sm font-bold text-white font-sans">Interactive Architecture Visualizer</span>
              <span className="text-[10px] font-mono bg-emerald-500/20 text-emerald-300 px-2 py-0.5 rounded border border-emerald-500/30">
                Live React Flow
              </span>
            </div>
            <span className="text-[11px] text-slate-400 hidden sm:inline">Drag nodes & interact with the graph</span>
          </div>

          <div className="flex-1 w-full h-full relative">
            <ReactFlow
              nodes={initialNodes}
              edges={initialEdges}
              fitView
              attributionPosition="bottom-right"
              proOptions={{ hideAttribution: true }}
            >
              <Background color="#1e293b" gap={24} size={1} />
              <Controls className="bg-slate-900 border border-slate-700 fill-white rounded-xl overflow-hidden shadow-lg" />
            </ReactFlow>

            {/* Overlay instructions */}
            {step < 4 && (
              <div className="absolute inset-0 bg-slate-950/80 backdrop-blur-sm flex flex-col items-center justify-center z-20 space-y-4">
                <div className="w-12 h-12 rounded-full border-4 border-cyan-400 border-t-transparent animate-spin" />
                <p className="font-mono text-cyan-300 font-bold text-sm">
                  {step <= 2 ? '⚡ Tree-sitter Extracting Syntax Trees...' : '🧠 Compiling Neural Repository Graph...'}
                </p>
              </div>
            )}
          </div>
        </div>

      </div>
    </div>
  );
};
