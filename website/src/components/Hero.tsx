import React, { useState } from 'react';
import { Terminal, Copy, Check, ArrowRight, Award, Cpu, Shield, Sparkles } from 'lucide-react';

interface HeroProps {
  onInstallClick: () => void;
  onOpenHallOfFame: () => void;
}

export const Hero: React.FC<HeroProps> = ({ onInstallClick, onOpenHallOfFame }) => {
  const [copied, setCopied] = useState(false);
  const killerCmd = "brew install codemri && cd my-unfamiliar-repo && codemri";

  const copyCmd = () => {
    navigator.clipboard.writeText(killerCmd);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };

  return (
    <div className="relative overflow-hidden pt-12 pb-20 text-center">
      
      {/* Subtle Background Glow Gradients */}
      <div className="absolute top-0 left-1/2 -translate-x-1/2 w-[650px] sm:w-[850px] h-[350px] bg-gradient-to-tr from-cyan-500/20 via-blue-600/15 to-purple-600/15 blur-[140px] rounded-full pointer-events-none -z-10" />

      <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8">
        
        {/* Release Badge */}
        <button onClick={() => {}} className="inline-flex items-center space-x-2 px-4 py-1.5 rounded-full bg-black/60 border border-cyan-500/40 text-xs text-cyan-300 font-mono font-extrabold mb-8 shadow-xl hover:border-cyan-400 transition-all cursor-pointer backdrop-blur-md group">
          <span className="w-2 h-2 rounded-full bg-cyan-400 animate-ping mr-2" />
          <span>v0.5.0 Vision Released — Interactive Force-Directed Topology Canvas Online 🕸️</span>
          <ArrowRight className="w-3.5 h-3.5 ml-1 group-hover:translate-x-1 transition-transform" />
        </button>

        {/* Headline */}
        <h1 className="text-4xl sm:text-6xl lg:text-7xl font-black text-white tracking-tight leading-none mb-6">
          Understand any codebase <br className="hidden sm:block" />
          in <span className="bg-gradient-to-r from-cyan-400 via-blue-400 to-emerald-400 bg-clip-text text-transparent">under 60 seconds.</span>
        </h1>

        {/* Subtitle */}
        <p className="max-w-3xl mx-auto text-base sm:text-xl text-slate-200 font-normal leading-relaxed mb-10">
          The offline-first Neural Repository Intelligence platform. When you clone an unfamiliar software repository, don&apos;t read 3,000 files manually. Install once via WinGet or Homebrew, enter the directory, and type one word: <strong>codemri</strong>.
        </p>

        {/* Primary CTA Buttons */}
        <div className="flex flex-col sm:flex-row items-center justify-center gap-4 mb-12">
          <button
            onClick={onInstallClick}
            className="w-full sm:w-auto px-8 py-4 bg-gradient-to-r from-cyan-400 via-blue-500 to-blue-600 hover:from-cyan-300 hover:to-blue-500 text-slate-950 font-black text-base rounded-2xl shadow-xl shadow-cyan-500/25 hover:shadow-cyan-500/40 hover:scale-[1.02] transition-all flex items-center justify-center space-x-2"
          >
            <Terminal className="w-5 h-5 text-slate-950" />
            <span>Install Binary &amp; Quickstart</span>
          </button>

          <button
            onClick={onOpenHallOfFame}
            className="w-full sm:w-auto px-8 py-4 bg-black/60 hover:bg-white/10 text-white font-extrabold text-base rounded-2xl border border-white/20 hover:border-white/40 transition-all flex items-center justify-center space-x-2 shadow-lg backdrop-blur-md group"
          >
            <Award className="w-5 h-5 text-amber-400 group-hover:scale-110 transition-transform" />
            <span>Repository Hall of Fame 🏆</span>
          </button>
        </div>

        {/* One-Click Terminal Copy Box */}
        <div className="max-w-3xl mx-auto bg-black/70 backdrop-blur-xl border-2 border-white/20 rounded-2xl p-4 sm:p-5 flex flex-col sm:flex-row items-center justify-between gap-4 shadow-2xl mb-16 text-left">
          <div className="flex items-center space-x-3 w-full sm:w-auto overflow-hidden">
            <div className="flex space-x-1.5 flex-shrink-0">
              <span className="w-3 h-3 rounded-full bg-rose-500/90" />
              <span className="w-3 h-3 rounded-full bg-amber-500/90" />
              <span className="w-3 h-3 rounded-full bg-emerald-500/90" />
            </div>
            <code className="font-mono text-xs sm:text-sm text-cyan-300 truncate font-bold">
              <span className="text-slate-400 mr-2">$</span>
              {killerCmd}
            </code>
          </div>

          <button
            onClick={copyCmd}
            className="w-full sm:w-auto flex items-center justify-center space-x-1.5 px-4 py-2 rounded-xl bg-slate-800 hover:bg-slate-700 text-slate-100 text-xs font-mono border border-slate-600 flex-shrink-0 transition-all shadow-sm"
          >
            {copied ? <Check className="w-4 h-4 text-emerald-400" /> : <Copy className="w-4 h-4 text-slate-300" />}
            <span className="font-bold">{copied ? 'Copied!' : 'Copy Killer Command'}</span>
          </button>
        </div>

        {/* Three Core Pillars Grid */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6 text-left mt-10">
          
          <div className="glass-panel p-7 rounded-3xl border border-white/15 hover:border-cyan-400/40 transition-all shadow-xl space-y-3">
            <div className="w-12 h-12 rounded-2xl bg-cyan-500/20 border border-cyan-500/40 flex items-center justify-center text-cyan-400 mb-2 shadow-sm">
              <Sparkles className="w-6 h-6 text-cyan-300" />
            </div>
            <h3 className="text-lg font-extrabold text-white">The "One-Word" Habit</h3>
            <p className="text-sm text-slate-300 leading-relaxed font-sans">
              "Whenever you clone a new repository, just run <code className="text-cyan-300 font-mono font-bold">codemri</code> first." Our engine detects your stack, parses syntax ASTs, and opens an architectural UI automatically in your browser.
            </p>
          </div>

          <div className="glass-panel p-7 rounded-3xl border border-white/15 hover:border-blue-400/40 transition-all shadow-xl space-y-3">
            <div className="w-12 h-12 rounded-2xl bg-blue-500/20 border border-blue-500/40 flex items-center justify-center text-blue-400 mb-2 shadow-sm">
              <Cpu className="w-6 h-6 text-blue-300" />
            </div>
            <h3 className="text-lg font-extrabold text-white">Zero Source Untouched</h3>
            <p className="text-sm text-slate-300 leading-relaxed font-sans">
              All parsed structural symbols and relational graphs reside strictly inside your target project in a local <code className="text-blue-300 font-mono font-bold">.codemri/</code> folder. Your actual application code is never modified.
            </p>
          </div>

          <div className="glass-panel p-7 rounded-3xl border border-white/15 hover:border-purple-400/40 transition-all shadow-xl space-y-3">
            <div className="w-12 h-12 rounded-2xl bg-purple-500/20 border border-purple-500/40 flex items-center justify-center text-purple-400 mb-2 shadow-sm">
              <Shield className="w-6 h-6 text-purple-300" />
            </div>
            <h3 className="text-lg font-extrabold text-white">Offline &amp; Online Modes</h3>
            <p className="text-sm text-slate-300 leading-relaxed font-sans">
              <strong>Offline Mode</strong> operates 100% locally with zero internet. When invoking <strong>Online Mode</strong> (`codemri explain`), only the abstract relational graph is shared with AI—protecting raw intellectual property.
            </p>
          </div>

        </div>

      </div>

    </div>
  );
};
