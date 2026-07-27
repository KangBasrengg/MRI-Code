import React, { useState } from 'react';
import { HALL_OF_FAME_DATA } from '../data/hallOfFame';
import type { RepositoryData } from '../types';
import { ReactFlow, Background, Controls } from '@xyflow/react';
import { ShieldCheck, Activity, Layers, Award, Zap } from 'lucide-react';

export const HallOfFame: React.FC = () => {
  const [selectedRepo, setSelectedRepo] = useState<RepositoryData>(HALL_OF_FAME_DATA[0]);

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12 text-left">
      
      {/* Header Banner */}
      <div className="bg-gradient-to-r from-amber-500/10 via-slate-900 to-purple-600/10 border border-amber-500/30 rounded-3xl p-8 mb-10 relative overflow-hidden shadow-2xl">
        <div className="flex flex-col md:flex-row md:items-center justify-between gap-6">
          <div>
            <div className="inline-flex items-center space-x-2 text-amber-400 text-xs font-mono font-bold mb-2">
              <Award className="w-4 h-4" />
              <span>THE VIRAL EXHIBITION — POWERED BY CODEMRI ENGINE</span>
            </div>
            <h1 className="text-3xl sm:text-5xl font-black text-white tracking-tight">
              Repository Hall of Fame 🏆
            </h1>
            <p className="text-slate-300 text-sm sm:text-base mt-2 max-w-3xl leading-relaxed">
              Explore live, pre-scanned architectural visualisations of the world's most popular open-source repositories without installing anything. Click any project below to inspect its internal Neural Repository Graph!
            </p>
          </div>
          <div className="flex flex-col sm:flex-row items-start sm:items-center gap-3 bg-slate-900/90 p-4 rounded-2xl border border-slate-800 flex-shrink-0">
            <div className="w-10 h-10 rounded-xl bg-emerald-500/20 text-emerald-400 flex items-center justify-center font-black text-lg border border-emerald-500/40">
              98
            </div>
            <div>
              <p className="text-xs font-mono font-bold text-white">Repository Intelligence Badge</p>
              <p className="text-[11px] text-slate-400">Embed verified scores on any GitHub README!</p>
            </div>
          </div>
        </div>
      </div>

      {/* Repo Selection Tabs */}
      <div className="grid grid-cols-2 md:grid-cols-4 gap-4 mb-8">
        {HALL_OF_FAME_DATA.map((repo) => {
          const isSelected = selectedRepo.id === repo.id;
          return (
            <button
              key={repo.id}
              onClick={() => setSelectedRepo(repo)}
              className={`p-5 rounded-2xl border transition-all text-left group flex flex-col justify-between relative overflow-hidden ${
                isSelected
                  ? 'bg-gradient-to-b from-cyan-950/40 to-slate-900 border-cyan-500 shadow-lg shadow-cyan-500/10 scale-[1.02]'
                  : 'bg-slate-900/60 border-slate-800 hover:border-slate-700 hover:bg-slate-800/50'
              }`}
            >
              <div>
                <div className="flex items-center justify-between mb-2">
                  <span className="font-mono text-sm sm:text-base font-bold text-white group-hover:text-cyan-300 transition-colors">
                    {repo.name}
                  </span>
                  <span className="text-xs font-mono bg-slate-800 text-amber-300 px-2 py-0.5 rounded-full font-semibold border border-amber-500/30">
                    {repo.stars}
                  </span>
                </div>
                <p className="text-xs text-slate-400 line-clamp-2">{repo.description}</p>
              </div>
              <div className="mt-4 pt-3 border-t border-slate-800/80 flex items-center justify-between text-[11px] font-mono text-slate-400">
                <span>{repo.language}</span>
                <span className="text-emerald-400 font-bold">Health: {repo.healthScore}/100</span>
              </div>
            </button>
          );
        })}
      </div>

      {/* Main Exhibition Grid */}
      <div className="grid grid-cols-1 lg:grid-cols-12 gap-8 items-start">
        
        {/* Left: Interactive Architecture Graph (React Flow) */}
        <div className="lg:col-span-8 bg-slate-900/90 border-2 border-slate-800 rounded-3xl overflow-hidden shadow-2xl h-[560px] flex flex-col">
          <div className="bg-slate-950 px-5 py-3.5 border-b border-slate-800 flex items-center justify-between">
            <div className="flex items-center space-x-3">
              <Layers className="w-5 h-5 text-cyan-400" />
              <span className="font-bold text-sm sm:text-base text-white">{selectedRepo.name} — Architectural Neural Graph</span>
            </div>
            <span className="text-xs font-mono bg-purple-500/20 text-purple-300 px-2.5 py-1 rounded border border-purple-500/30">
              {selectedRepo.modulesCount} Modules Mapped
            </span>
          </div>
          <div className="flex-1 w-full h-full relative">
            <ReactFlow
              nodes={selectedRepo.nodes}
              edges={selectedRepo.edges}
              fitView
              key={selectedRepo.id} // Re-render flow on switch
              proOptions={{ hideAttribution: true }}
            >
              <Background color="#1e293b" gap={20} size={1} />
              <Controls className="bg-slate-900 border border-slate-700 fill-white rounded-xl overflow-hidden" />
            </ReactFlow>
          </div>
        </div>

        {/* Right: AI Intelligence Insights & Scorecards */}
        <div className="lg:col-span-4 space-y-6">
          
          {/* Score Card */}
          <div className="glass-panel p-6 rounded-3xl border-slate-800 shadow-2xl">
            <h3 className="text-lg font-bold text-white mb-4 flex items-center justify-between">
              <span className="flex items-center space-x-2">
                <Activity className="w-5 h-5 text-emerald-400" />
                <span>Repository Intelligence Score</span>
              </span>
              <span className="text-2xl font-black font-mono text-emerald-400">{selectedRepo.healthScore}/100</span>
            </h3>

            <div className="space-y-3 font-mono text-xs sm:text-sm">
              <div className="flex justify-between py-2 border-b border-slate-800">
                <span className="text-slate-400">Scanned Source Files:</span>
                <span className="text-white font-bold">{selectedRepo.filesCount.toLocaleString()} files</span>
              </div>
              <div className="flex justify-between py-2 border-b border-slate-800">
                <span className="text-slate-400">Circular Dependencies:</span>
                <span className="text-emerald-400 font-bold">{selectedRepo.circularDeps} detected</span>
              </div>
              <div className="flex justify-between py-2">
                <span className="text-slate-400">NRG Index Status:</span>
                <span className="text-cyan-400 font-bold">100% Deterministic AST</span>
              </div>
            </div>
          </div>

          {/* AI Reasoning Insights Card */}
          <div className="glass-panel p-6 rounded-3xl border-slate-800 shadow-2xl space-y-4">
            <div className="flex items-center space-x-2 text-amber-400 text-xs font-mono font-bold">
              <Zap className="w-4 h-4" />
              <span>AI REASONING ENGINE FINDINGS</span>
            </div>

            <div>
              <h4 className="text-xs uppercase font-mono font-bold text-slate-300 mb-1">📐 Architecture Pattern</h4>
              <p className="text-sm text-slate-400 leading-relaxed bg-slate-950/60 p-3.5 rounded-xl border border-slate-800/80">
                {selectedRepo.insights.architecture}
              </p>
            </div>

            <div>
              <h4 className="text-xs uppercase font-mono font-bold text-slate-300 mb-1">🛠️ Technical Debt</h4>
              <p className="text-sm text-slate-400 leading-relaxed bg-slate-950/60 p-3.5 rounded-xl border border-slate-800/80">
                {selectedRepo.insights.techDebt}
              </p>
            </div>

            <div>
              <h4 className="text-xs uppercase font-mono font-bold text-slate-300 mb-1">🛡️ Security Hotspots</h4>
              <p className="text-sm text-emerald-400 leading-relaxed bg-slate-950/60 p-3.5 rounded-xl border border-slate-800/80 flex items-start space-x-2">
                <ShieldCheck className="w-4 h-4 flex-shrink-0 mt-0.5" />
                <span>{selectedRepo.insights.security}</span>
              </p>
            </div>
          </div>

        </div>

      </div>

    </div>
  );
};
