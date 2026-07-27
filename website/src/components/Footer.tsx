import React from 'react';
import { Disc as Discord, Shield, Globe, GitBranch } from 'lucide-react';
import codemriLogo from '../assets/codemri.png';

export const Footer: React.FC = () => {
  return (
    <footer className="bg-black/60 backdrop-blur-xl border-t border-white/10 pt-12 pb-8 font-sans text-slate-300 text-xs sm:text-sm">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        
        <div className="grid grid-cols-1 md:grid-cols-4 gap-8 mb-10 text-left">
          
          {/* Col 1: Brand */}
          <div className="md:col-span-2 space-y-3">
            <div className="flex items-center space-x-2.5 text-white font-sans font-black text-lg">
              <div className="w-8 h-8 rounded-lg overflow-hidden border border-cyan-500/50 shadow-md flex items-center justify-center bg-black">
                <img src={codemriLogo} alt="CodeMRI Logo" className="w-full h-full object-cover" />
              </div>
              <span>Code-MRI Platform</span>
            </div>
            <p className="text-xs sm:text-sm text-slate-300 max-w-sm leading-relaxed">
              "GitHub shows your files. Code-MRI shows how your software actually works."
              The world's first offline-first Neural Repository Intelligence engine.
            </p>
            <div className="inline-flex items-center space-x-2 text-[11px] font-mono bg-black/60 border border-emerald-500/40 px-3 py-1.5 rounded-full text-emerald-400 font-bold shadow-md">
              <Shield className="w-3.5 h-3.5" />
              <span>Licensed under Apache 2.0 — Safe for Enterprise & Commercial use</span>
            </div>
          </div>

          {/* Col 2: Navigation Links */}
          <div className="space-y-2">
            <h4 className="text-white font-extrabold font-mono uppercase text-xs mb-3 text-cyan-400">Ecosystem</h4>
            <p><a href="https://github.com/KangBasrengg/MRI-Code" target="_blank" rel="noopener noreferrer" className="hover:text-cyan-300 transition-colors">GitHub Repository</a></p>
            <p><a href="https://github.com/KangBasrengg/MRI-Code/blob/main/docs/philosophy.md" target="_blank" rel="noopener noreferrer" className="hover:text-cyan-300 transition-colors">Philosophy & Beliefs</a></p>
            <p><a href="https://github.com/KangBasrengg/MRI-Code/blob/main/docs/ADR/0001-neural-repository-graph.md" target="_blank" rel="noopener noreferrer" className="hover:text-cyan-300 transition-colors">ADR 0001: NRG Core</a></p>
            <p><a href="https://github.com/KangBasrengg/MRI-Code/issues" target="_blank" rel="noopener noreferrer" className="hover:text-cyan-300 transition-colors">Issue Tracker & RFCs</a></p>
          </div>

          {/* Col 3: Community */}
          <div className="space-y-2">
            <h4 className="text-white font-extrabold font-mono uppercase text-xs mb-3 text-cyan-400">Community</h4>
            <p><a href="https://github.com/KangBasrengg/MRI-Code/discussions" target="_blank" rel="noopener noreferrer" className="hover:text-cyan-300 transition-colors flex items-center space-x-2"><GitBranch className="w-4 h-4 text-cyan-400" /> <span>GitHub Discussions</span></a></p>
            <p><span className="text-slate-500 cursor-not-allowed flex items-center space-x-2"><Discord className="w-4 h-4 text-slate-500" /> <span>Discord Server (Soon)</span></span></p>
            <p><span className="text-slate-500 cursor-not-allowed flex items-center space-x-2"><Globe className="w-4 h-4 text-slate-500" /> <span>X / Twitter</span></span></p>
          </div>

        </div>

        {/* Bottom Bar */}
        <div className="pt-8 border-t border-white/10 flex flex-col sm:flex-row items-center justify-between gap-4 text-xs font-mono text-slate-400">
          <div>
            © 2026 <strong className="text-white">Muhammad Nuril (KangBasrengg)</strong> & Code-MRI Open-Source Contributors.
          </div>
          <div className="flex items-center space-x-1">
            <span>Crafted with engineering discipline for software developers worldwide.</span>
          </div>
        </div>

      </div>
    </footer>
  );
};
