import React, { useState } from 'react';
import { Terminal, Copy, Check, Apple, Monitor, ArrowRight, BookOpen, Flame } from 'lucide-react';

interface CliInstallGuideProps {
  onOpenFullDocs: () => void;
}

export const CliInstallGuide: React.FC<CliInstallGuideProps> = ({ onOpenFullDocs }) => {
  const [activeOs, setActiveOs] = useState<'windows' | 'mac' | 'linux' | 'go'>('windows');
  const [copied, setCopied] = useState<string | null>(null);

  const handleCopy = (id: string, text: string) => {
    navigator.clipboard.writeText(text);
    setCopied(id);
    setTimeout(() => setCopied(null), 2000);
  };

  const installCommands = {
    windows: `# Step 1: Install globally via Windows Package Manager
winget install codemri

# OR install via Scoop
scoop install codemri

# Step 2: Navigate to ANY local repository and execute one word
cd C:\\Projects\\unfamiliar-codebase
codemri`,
    mac: `# Step 1: Install globally via Homebrew
brew install codemri

# Step 2: Navigate to ANY software project and execute one word
cd /projects/enterprise-app
codemri`,
    linux: `# Step 1: Install globally via secure shell installer
curl -fsSL https://codemri.dev/install.sh | sh

# OR download standalone binary (codemri_v0.4.0_linux_amd64.tar.gz) from GitHub Releases
cd ~/workspace/backend-service
codemri`,
    go: `# For Go developers (Runtime v1.21+)
go install github.com/KangBasrengg/MRI-Code/cmd/codemri@latest

# Navigate to target repository and run one word
cd my-monorepo
codemri`
  };

  return (
    <section id="cli-install-guide" className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 py-16 scroll-mt-24">
      
      {/* Header Banner */}
      <div className="text-center max-w-4xl mx-auto mb-10">
        <div className="inline-flex items-center space-x-2 px-3.5 py-1 rounded-full bg-cyan-500/20 border border-cyan-500/40 text-cyan-300 font-mono text-xs uppercase font-black mb-4 shadow-md">
          <Terminal className="w-4 h-4 text-cyan-400" />
          <span>ZERO CLONING REQUIRED — STANDALONE NATIVE BINARY</span>
        </div>
        <h2 className="text-3xl sm:text-5xl font-extrabold text-white tracking-tight leading-tight">
          Install Like Bun or Biome. <br className="hidden sm:block" />
          <span className="bg-gradient-to-r from-cyan-400 via-blue-400 to-emerald-400 bg-clip-text text-transparent">Run One Word in Any Repository.</span>
        </h2>
        <p className="text-slate-300 text-sm sm:text-base mt-4 leading-relaxed font-sans max-w-2xl mx-auto">
          We never force developers to clone our GitHub repository just to evaluate code. Install the lightweight cross-platform executable once, navigate into your proprietary project, and type simply: <code className="text-cyan-300 font-mono font-black px-2 py-0.5 bg-black/60 rounded border border-cyan-500/30">codemri</code>.
        </p>
      </div>

      {/* Terminal Guide Box */}
      <div className="bg-[#0c101c]/80 backdrop-blur-xl border-2 border-white/15 rounded-3xl overflow-hidden shadow-2xl">
        
        {/* Top OS Selector Tabs */}
        <div className="bg-black/60 px-6 py-3.5 border-b border-white/10 flex flex-wrap items-center justify-between gap-3">
          <div className="flex flex-wrap items-center gap-2 text-xs font-bold">
            <span className="text-slate-400 mr-1 font-mono hidden sm:inline">Distribution:</span>
            
            <button
              onClick={() => setActiveOs('windows')}
              className={`flex items-center space-x-2 px-3.5 py-2 rounded-xl transition-all ${
                activeOs === 'windows'
                  ? 'bg-gradient-to-r from-cyan-500/30 to-blue-500/30 text-white border border-cyan-400 shadow-lg font-black'
                  : 'text-slate-400 hover:text-white hover:bg-white/10'
              }`}
            >
              <Monitor className="w-4 h-4 text-cyan-400" />
              <span>Windows (WinGet / Scoop)</span>
            </button>

            <button
              onClick={() => setActiveOs('mac')}
              className={`flex items-center space-x-2 px-3.5 py-2 rounded-xl transition-all ${
                activeOs === 'mac'
                  ? 'bg-gradient-to-r from-cyan-500/30 to-blue-500/30 text-white border border-cyan-400 shadow-lg font-black'
                  : 'text-slate-400 hover:text-white hover:bg-white/10'
              }`}
            >
              <Apple className="w-4 h-4 text-slate-200" />
              <span>macOS (Homebrew)</span>
            </button>

            <button
              onClick={() => setActiveOs('linux')}
              className={`flex items-center space-x-2 px-3.5 py-2 rounded-xl transition-all ${
                activeOs === 'linux'
                  ? 'bg-gradient-to-r from-cyan-500/30 to-blue-500/30 text-white border border-cyan-400 shadow-lg font-black'
                  : 'text-slate-400 hover:text-white hover:bg-white/10'
              }`}
            >
              <Terminal className="w-4 h-4 text-amber-400" />
              <span>Linux (curl)</span>
            </button>

            <button
              onClick={() => setActiveOs('go')}
              className={`flex items-center space-x-2 px-3.5 py-2 rounded-xl font-mono transition-all ${
                activeOs === 'go'
                  ? 'bg-gradient-to-r from-cyan-500/30 to-blue-500/30 text-white border border-cyan-400 shadow-lg font-black'
                  : 'text-slate-400 hover:text-white hover:bg-white/10'
              }`}
            >
              <span>🐹 go install</span>
            </button>
          </div>

          <button
            onClick={() => handleCopy(activeOs, installCommands[activeOs])}
            className="flex items-center space-x-1.5 px-3 py-1.5 bg-slate-800 hover:bg-slate-700 text-slate-200 text-xs font-mono rounded-xl border border-slate-600 transition-all shadow-sm"
          >
            {copied === activeOs ? <Check className="w-3.5 h-3.5 text-emerald-400" /> : <Copy className="w-3.5 h-3.5 text-slate-400" />}
            <span className="font-bold">{copied === activeOs ? 'Copied to Clipboard!' : 'Copy Workflow'}</span>
          </button>
        </div>

        {/* Code Content Box */}
        <div className="p-6 sm:p-8 font-mono text-xs sm:text-sm text-slate-100 text-left bg-gradient-to-b from-black/80 to-[#05080f]/90 overflow-x-auto">
          <pre className="leading-relaxed">
            <code className="text-emerald-300 font-bold">{installCommands[activeOs]}</code>
          </pre>
        </div>

        {/* The Killer Feature: One-Command Workflow Explained */}
        <div className="bg-black/40 p-6 sm:p-8 border-t border-white/10 grid grid-cols-1 md:grid-cols-3 gap-6 text-left font-sans">
          
          <div className="space-y-2.5 bg-black/50 p-5 rounded-2xl border border-white/10">
            <div className="w-9 h-9 rounded-xl bg-cyan-500/20 text-cyan-400 border border-cyan-500/40 flex items-center justify-center font-mono font-black text-sm shadow-sm">
              1
            </div>
            <h4 className="font-extrabold text-white text-base">Clone Any Unfamiliar Project</h4>
            <p className="text-xs sm:text-sm text-slate-300 leading-relaxed">
              Whether it is Next.js, Laravel, Go, Rust, or Python. Enter the directory of the software project you want to inspect.
            </p>
            <div className="pt-1 font-mono text-xs text-cyan-300 font-bold">
              $ git clone https://github.com/org/app &amp;&amp; cd app
            </div>
          </div>

          <div className="space-y-2.5 bg-black/50 p-5 rounded-2xl border border-white/10">
            <div className="w-9 h-9 rounded-xl bg-amber-500/20 text-amber-300 border border-amber-500/40 flex items-center justify-center font-mono font-black text-sm shadow-sm">
              2
            </div>
            <h4 className="font-extrabold text-white text-base">Run One Word: <span className="text-amber-400">codemri</span></h4>
            <p className="text-xs sm:text-sm text-slate-300 leading-relaxed">
              No flags or config files needed. CodeMRI detects languages, executes deterministic syntax parsers, and stores the Neural Repository Graph cleanly inside your project at <code className="text-amber-300 font-mono font-bold">.codemri/</code>.
            </p>
            <div className="pt-1 font-mono text-xs text-amber-300 font-bold">
              $ codemri
            </div>
          </div>

          <div className="space-y-2.5 bg-black/50 p-5 rounded-2xl border border-white/10">
            <div className="w-9 h-9 rounded-xl bg-emerald-500/20 text-emerald-400 border border-emerald-500/40 flex items-center justify-center font-mono font-black text-sm shadow-sm">
              3
            </div>
            <h4 className="font-extrabold text-white text-base">Instant UI in System Browser</h4>
            <p className="text-xs sm:text-sm text-slate-300 leading-relaxed">
              Within 60 seconds, your system browser opens automatically at <code className="text-emerald-300 font-mono font-bold">http://localhost:4000</code> displaying an interactive architectural map. Just like Prisma Studio or Laravel Telescope!
            </p>
            <div className="pt-1 font-mono text-xs text-emerald-300 font-bold">
              ✨ Browser opened! "I understand the codebase now."
            </div>
          </div>

        </div>

        {/* Architectural Distinction Footer */}
        <div className="bg-slate-900/80 px-6 py-5 border-t border-white/10 flex flex-col md:flex-row items-center justify-between gap-4 text-left font-sans">
          <div className="flex items-center space-x-3.5">
            <div className="w-10 h-10 rounded-xl bg-purple-500/20 border border-purple-500/40 flex items-center justify-center flex-shrink-0">
              <Flame className="w-5 h-5 text-purple-400" />
            </div>
            <div>
              <h4 className="text-sm font-extrabold text-white">Why Go Binary over NPM package distribution?</h4>
              <p className="text-xs text-slate-300 leading-relaxed">
                CodeMRI targets multi-language ecosystems (Go, Rust, Java, C#, Laravel/PHP, Python, TS). Java or Go developers should never require Node or npm simply to audit code structure. One lightning-fast compiled binary runs seamlessly everywhere.
              </p>
            </div>
          </div>
          
          <button
            onClick={onOpenFullDocs}
            className="px-5 py-2.5 bg-gradient-to-r from-blue-500 to-cyan-400 hover:from-blue-400 hover:to-cyan-300 text-slate-950 font-black text-xs sm:text-sm rounded-xl flex items-center space-x-2 whitespace-nowrap shadow-lg shadow-cyan-500/20 transition-all flex-shrink-0"
          >
            <BookOpen className="w-4 h-4 text-slate-950" />
            <span>Read Architecture Docs</span>
            <ArrowRight className="w-4 h-4 text-slate-950" />
          </button>
        </div>

      </div>

    </section>
  );
};
