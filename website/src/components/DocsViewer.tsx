import React, { useState } from 'react';
import { DOCS_DATA } from '../data/docsContent';
import type { DocSection, DocContentBlock } from '../types';
import { BookOpen, Terminal, Copy, Check, ChevronRight, Zap, Info, AlertTriangle, CheckCircle } from 'lucide-react';

// ─────────────────────────────────────────────────
// Block Renderer: Transforms structured content blocks
// into rich, beautiful React elements (no raw markdown!)
// ─────────────────────────────────────────────────
const RenderBlock: React.FC<{ block: DocContentBlock; index: number }> = ({ block, index }) => {
  const [blockCopied, setBlockCopied] = useState(false);

  const handleBlockCopy = (code: string) => {
    navigator.clipboard.writeText(code);
    setBlockCopied(true);
    setTimeout(() => setBlockCopied(false), 2000);
  };

  switch (block.type) {
    case 'paragraph':
      return (
        <p key={index} className="text-slate-300 text-sm sm:text-[15px] leading-[1.8] font-sans">
          {block.text}
        </p>
      );

    case 'heading':
      return (
        <h3 key={index} className="text-lg sm:text-xl font-extrabold text-white mt-8 mb-3 flex items-center space-x-2 border-b border-slate-800/60 pb-3">
          <span className="w-1.5 h-6 rounded-full bg-gradient-to-b from-cyan-400 to-blue-500 inline-block mr-1" />
          <span>{block.text}</span>
        </h3>
      );

    case 'code':
      return (
        <div key={index} className="my-5 bg-slate-950/90 border border-slate-800 rounded-2xl overflow-hidden shadow-xl">
          <div className="bg-slate-900/90 px-4 py-2.5 border-b border-slate-800 flex items-center justify-between">
            <span className="text-[11px] font-mono text-cyan-400 font-bold flex items-center space-x-2 uppercase tracking-wider">
              <Terminal className="w-3.5 h-3.5 text-cyan-400" />
              <span>{block.language || 'shell'}</span>
            </span>
            <button
              onClick={() => handleBlockCopy(block.code || '')}
              className="flex items-center space-x-1.5 px-2.5 py-1 bg-slate-800 hover:bg-slate-700 text-slate-300 text-[11px] font-mono rounded-lg transition-all border border-slate-700"
            >
              {blockCopied ? <Check className="w-3 h-3 text-emerald-400" /> : <Copy className="w-3 h-3" />}
              <span className="font-bold">{blockCopied ? 'Copied!' : 'Copy'}</span>
            </button>
          </div>
          <pre className="p-5 font-mono text-xs sm:text-[13px] text-emerald-300/90 overflow-x-auto leading-relaxed">
            <code>{block.code}</code>
          </pre>
        </div>
      );

    case 'list':
      return (
        <ul key={index} className="my-4 space-y-2.5 pl-1">
          {(block.items as string[])?.map((item, i) => (
            <li key={i} className="flex items-start space-x-3 text-sm text-slate-300 leading-relaxed font-sans">
              <span className="mt-1.5 w-2 h-2 rounded-full bg-gradient-to-br from-cyan-400 to-blue-500 flex-shrink-0 shadow-sm shadow-cyan-400/30" />
              <span>{item}</span>
            </li>
          ))}
        </ul>
      );

    case 'callout': {
      const variants = {
        info: {
          bg: 'bg-blue-500/10',
          border: 'border-blue-500/30',
          icon: <Info className="w-5 h-5 text-blue-400" />,
          text: 'text-blue-200',
          label: 'Note',
          labelColor: 'text-blue-400',
        },
        warning: {
          bg: 'bg-amber-500/10',
          border: 'border-amber-500/30',
          icon: <AlertTriangle className="w-5 h-5 text-amber-400" />,
          text: 'text-amber-200',
          label: 'Important',
          labelColor: 'text-amber-400',
        },
        success: {
          bg: 'bg-emerald-500/10',
          border: 'border-emerald-500/30',
          icon: <CheckCircle className="w-5 h-5 text-emerald-400" />,
          text: 'text-emerald-200',
          label: 'Key Insight',
          labelColor: 'text-emerald-400',
        },
      };
      const v = variants[block.variant || 'info'];
      return (
        <div key={index} className={`my-6 ${v.bg} border ${v.border} rounded-2xl p-5 flex items-start space-x-4`}>
          <div className="flex-shrink-0 mt-0.5">{v.icon}</div>
          <div>
            <span className={`text-[11px] font-mono font-extrabold uppercase tracking-wider ${v.labelColor} block mb-1.5`}>{v.label}</span>
            <p className={`text-sm ${v.text} leading-relaxed font-sans`}>{block.text}</p>
          </div>
        </div>
      );
    }

    case 'steps': {
      const stepItems = block.items as { label: string; description: string }[];
      return (
        <div key={index} className="my-6 space-y-0 relative">
          {/* Vertical connector line */}
          <div className="absolute left-[18px] top-6 bottom-6 w-px bg-gradient-to-b from-cyan-500/60 via-blue-500/40 to-purple-500/30" />
          {stepItems?.map((step, i) => (
            <div key={i} className="flex items-start space-x-4 relative py-4">
              <div className="w-9 h-9 rounded-xl bg-gradient-to-br from-cyan-500/30 to-blue-600/30 border border-cyan-500/40 flex items-center justify-center font-mono font-black text-sm text-cyan-300 flex-shrink-0 shadow-md z-10">
                {i + 1}
              </div>
              <div className="pt-1">
                <h4 className="text-sm font-extrabold text-white mb-1">{step.label}</h4>
                <p className="text-sm text-slate-300 leading-relaxed font-sans">{step.description}</p>
              </div>
            </div>
          ))}
        </div>
      );
    }

    default:
      return null;
  }
};

// ─────────────────────────────────────────────────
// Main DocsViewer Component
// ─────────────────────────────────────────────────
export const DocsViewer: React.FC = () => {
  const [activeDoc, setActiveDoc] = useState<DocSection>(DOCS_DATA[0]);

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-10 text-left font-sans">
      <div className="grid grid-cols-1 md:grid-cols-12 gap-8 items-start">
        
        {/* Sidebar Navigation */}
        <div className="md:col-span-4 lg:col-span-3 bg-[#0b0f19]/90 backdrop-blur-xl border border-slate-800/80 rounded-3xl p-5 md:sticky md:top-24 shadow-xl">
          <div className="flex items-center space-x-2 text-cyan-400 font-bold font-mono text-xs uppercase mb-6 pb-4 border-b border-slate-800">
            <BookOpen className="w-4 h-4" />
            <span>Official Documentation</span>
          </div>
          <nav className="space-y-1.5">
            {DOCS_DATA.map((section) => {
              const isActive = activeDoc.id === section.id;
              return (
                <button
                  key={section.id}
                  onClick={() => setActiveDoc(section)}
                  className={`w-full text-left px-3.5 py-3 rounded-2xl transition-all flex items-center justify-between text-xs sm:text-sm font-medium ${
                    isActive
                      ? 'bg-gradient-to-r from-cyan-500/20 via-blue-500/10 to-transparent text-cyan-300 border-l-2 border-cyan-400 font-bold'
                      : 'text-slate-400 hover:text-white hover:bg-slate-900'
                  }`}
                >
                  <div>
                    <span className="text-[10px] uppercase font-mono block text-slate-500 mb-0.5">{section.category}</span>
                    <span>{section.title}</span>
                  </div>
                  <ChevronRight className={`w-4 h-4 transition-transform ${isActive ? 'text-cyan-400 translate-x-1' : 'text-slate-600'}`} />
                </button>
              );
            })}
          </nav>

          <div className="mt-8 pt-6 border-t border-slate-800/80 text-[11px] text-slate-500 space-y-2">
            <p className="font-mono">Standard: Fumadocs UI Theme</p>
            <p>Licensed under <span className="text-slate-400 font-bold">Apache 2.0</span> for enterprise patent security.</p>
          </div>
        </div>

        {/* Main Doc Article */}
        <div className="md:col-span-8 lg:col-span-9 bg-[#0f1623]/70 backdrop-blur-xl border border-slate-800 rounded-3xl p-6 sm:p-10 shadow-2xl min-h-[600px]">
          
          {/* Article Header */}
          <div className="border-b border-slate-800 pb-6 mb-8">
            <span className="text-xs uppercase font-mono px-3 py-1 bg-blue-500/10 text-blue-400 border border-blue-500/30 rounded-full font-bold">
              {activeDoc.category}
            </span>
            <h1 className="text-3xl sm:text-5xl font-black text-white mt-4 mb-3 tracking-tight leading-tight">
              {activeDoc.title}
            </h1>
            <p className="text-base sm:text-lg text-slate-300 font-normal leading-relaxed">
              {activeDoc.summary}
            </p>
          </div>

          {/* Rich Content Blocks */}
          <div className="space-y-2">
            {activeDoc.blocks?.map((block, index) => (
              <RenderBlock key={index} block={block} index={index} />
            ))}
          </div>

          {/* Legacy Fallback: plain text content if no blocks defined */}
          {!activeDoc.blocks?.length && activeDoc.content && (
            <div className="prose prose-invert max-w-none text-slate-300 space-y-6 text-sm sm:text-base leading-relaxed whitespace-pre-line">
              {activeDoc.content}
            </div>
          )}

          {/* Legacy: code example block */}
          {activeDoc.codeExample && !activeDoc.blocks?.length && (
            <div className="mt-8 bg-slate-950 border border-slate-800 rounded-2xl overflow-hidden shadow-2xl">
              <div className="bg-slate-900/90 px-4 py-2.5 border-b border-slate-800">
                <span className="text-xs font-mono text-cyan-400 font-bold flex items-center space-x-2">
                  <Terminal className="w-4 h-4 text-cyan-400" />
                  <span>{activeDoc.codeLanguage || 'bash'} example</span>
                </span>
              </div>
              <pre className="p-5 font-mono text-xs sm:text-sm text-emerald-300 overflow-x-auto text-left">
                <code>{activeDoc.codeExample}</code>
              </pre>
            </div>
          )}

          {/* Quick Help Callout Box */}
          <div className="mt-12 bg-slate-900/60 border border-cyan-500/30 rounded-2xl p-6 flex items-start space-x-4">
            <div className="w-10 h-10 rounded-xl bg-cyan-500/20 border border-cyan-500/40 flex items-center justify-center flex-shrink-0">
              <Zap className="w-5 h-5 text-cyan-400" />
            </div>
            <div>
              <h4 className="text-sm font-bold text-white font-mono">Need Help or Found a Bug?</h4>
              <p className="text-xs text-slate-400 mt-1 leading-relaxed">
                CodeMRI is built with strict conventional commits and zero-config philosophy. Visit our repository at <a href="https://github.com/KangBasrengg/MRI-Code" target="_blank" rel="noopener noreferrer" className="text-cyan-300 underline hover:text-cyan-200 font-bold">github.com/KangBasrengg/MRI-Code</a> to submit issues, discuss architecture, or contribute to the codebase.
              </p>
            </div>
          </div>
        </div>

      </div>
    </div>
  );
};
