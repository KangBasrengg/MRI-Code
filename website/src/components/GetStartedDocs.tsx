import { useState, useEffect } from 'react';
import { Copy, Check, Cpu, Sparkles } from 'lucide-react';

export const GetStartedDocs = () => {
  const [copiedText, setCopiedText] = useState<string | null>(null);
  const [installTab, setInstallTab] = useState<'go' | 'windows' | 'linux' | 'macos'>('go');
  const [activeId, setActiveId] = useState<string>('introduction');

  useEffect(() => {
    const handleScroll = () => {
      const sectionIds = [
        'introduction', 'installation', 'why-codemri',
        'pulse', 'shield', 'velocity', 'vision', 'cortex',
        'architecture', 'sqlite-engine', 'offline-first',
        'usage-local', 'usage-online', 'usage-ai', 'usage-cli'
      ];
      
      const scrollPosition = window.scrollY + 250;

      for (const id of sectionIds) {
        const element = document.getElementById(id);
        if (element) {
          const top = element.offsetTop;
          const height = element.offsetHeight;
          if (scrollPosition >= top && scrollPosition < top + height) {
            setActiveId(id);
            break;
          }
        }
      }
    };

    window.addEventListener('scroll', handleScroll, { passive: true });
    handleScroll();
    return () => window.removeEventListener('scroll', handleScroll);
  }, []);

  const copyToClipboard = (text: string) => {
    navigator.clipboard.writeText(text);
    setCopiedText(text);
    setTimeout(() => setCopiedText(null), 2000);
  };

  const installCommands = {
    go: 'go install github.com/KangBasrengg/MRI-Code/cmd/codemri@latest',
    windows: '# Download executable directly from dist/ or release\niwr https://raw.githubusercontent.com/KangBasrengg/MRI-Code/main/install-windows.ps1 -UseBasicParsing | iex',
    linux: 'curl -fsSL https://raw.githubusercontent.com/KangBasrengg/MRI-Code/main/install-linux.sh | bash',
    macos: 'curl -fsSL https://raw.githubusercontent.com/KangBasrengg/MRI-Code/main/install-mac.sh | bash',
  };

  const renderNavLinks = (items: { id: string; label: string }[]) => {
    return (
      <ul className="space-y-1.5 font-medium">
        {items.map((item) => {
          const isActive = activeId === item.id;
          return (
            <li key={item.id}>
              <a
                href={`#${item.id}`}
                onClick={() => setActiveId(item.id)}
                className={`group flex items-center justify-between px-3.5 py-2.5 rounded-xl text-xs transition-all duration-200 border ${
                  isActive
                    ? `bg-slate-950/95 border-cyan-400 text-cyan-300 font-black shadow-2xl shadow-black translate-x-2 scale-[1.02]`
                    : `border-transparent text-slate-300 hover:bg-slate-950/90 hover:border-slate-700 hover:text-white hover:translate-x-1.5 hover:shadow-xl hover:shadow-black/80`
                }`}
              >
                <span className="truncate pr-1">{item.label}</span>
                {isActive ? (
                  <span className="w-2 h-2 rounded-full bg-cyan-400 shadow-sm shadow-cyan-400 animate-pulse ml-2 flex-shrink-0" />
                ) : (
                  <span className="opacity-0 group-hover:opacity-100 transition-all text-slate-500 group-hover:text-cyan-400 font-mono text-[10px] ml-2 font-bold">→</span>
                )}
              </a>
            </li>
          );
        })}
      </ul>
    );
  };

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8 flex flex-col md:flex-row gap-10">
      
      {/* 1. Left Sidebar: Tailwind/React.dev Style Docs Navigation with Popup Hover */}
      <aside className="w-full md:w-64 flex-shrink-0">
        <div className="sticky top-24 space-y-7 pr-4 max-h-[calc(100vh-6.5rem)] overflow-y-auto font-sans text-sm pb-10 border-r border-slate-800/80 custom-scrollbar">
          
          <div>
            <h5 className="font-mono text-[11px] font-bold uppercase tracking-wider text-cyan-400 mb-2.5 px-3 flex items-center gap-1.5">
              <span className="w-1.5 h-1.5 rounded-full bg-cyan-400"></span>
              Getting Started
            </h5>
            {renderNavLinks([
              { id: 'introduction', label: 'What is CodeMRI?' },
              { id: 'installation', label: 'Quick Installation' },
              { id: 'why-codemri', label: 'Why CodeMRI vs Others' }
            ])}
          </div>

          <div>
            <h5 className="font-mono text-[11px] font-bold uppercase tracking-wider text-blue-400 mb-2.5 px-3 flex items-center gap-1.5">
              <span className="w-1.5 h-1.5 rounded-full bg-blue-400"></span>
              Core Functions
            </h5>
            {renderNavLinks([
              { id: 'pulse', label: '💓 Pulse (Health Scoring)' },
              { id: 'shield', label: '🛡️ Shield (Security Audit)' },
              { id: 'velocity', label: '🚀 Velocity (Performance)' },
              { id: 'vision', label: '👁️ Vision (Impact Radius)' },
              { id: 'cortex', label: '🤖 Cortex AI Copilot' }
            ])}
          </div>

          <div>
            <h5 className="font-mono text-[11px] font-bold uppercase tracking-wider text-emerald-400 mb-2.5 px-3 flex items-center gap-1.5">
              <span className="w-1.5 h-1.5 rounded-full bg-emerald-400"></span>
              How It Works
            </h5>
            {renderNavLinks([
              { id: 'architecture', label: 'Under the Hood' },
              { id: 'sqlite-engine', label: 'SQLite NRG Engine' },
              { id: 'offline-first', label: 'Offline-First Guarantee' }
            ])}
          </div>

          <div>
            <h5 className="font-mono text-[11px] font-bold uppercase tracking-wider text-amber-400 mb-2.5 px-3 flex items-center gap-1.5">
              <span className="w-1.5 h-1.5 rounded-full bg-amber-400"></span>
              Practical Usage
            </h5>
            {renderNavLinks([
              { id: 'usage-local', label: 'One-Command Local' },
              { id: 'usage-online', label: '🌐 Online GitHub Scan' },
              { id: 'usage-ai', label: 'Interactive AI Chatbot' },
              { id: 'usage-cli', label: 'CLI Command Table' }
            ])}
          </div>

        </div>
      </aside>

      {/* 2. Main Center Content: Clean, High-Readability Prose */}
      <main className="flex-1 max-w-5xl space-y-16 text-slate-300 pb-20">
        
        {/* SECTION: INTRODUCTION */}
        <section id="introduction" className="space-y-6">
          <div className="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-cyan-500/10 border border-cyan-500/30 text-cyan-300 text-xs font-mono font-semibold">
            <Sparkles className="w-3.5 h-3.5 text-cyan-400" />
            <span>Official v1.0.0 Documentation & Quickstart Guide</span>
          </div>
          <h1 className="text-4xl sm:text-5xl font-black text-white tracking-tight leading-tight">
            What is <span className="bg-gradient-to-r from-cyan-400 via-blue-400 to-purple-400 bg-clip-text text-transparent">CodeMRI</span>?
          </h1>
          <p className="text-lg sm:text-xl text-slate-300 leading-relaxed font-normal">
            <strong>CodeMRI</strong> acts like a medical MRI machine for software codebases. While platforms like GitHub show you simple flat lists of files and folders, CodeMRI maps how your software actually lives, operates, and intertwines.
          </p>
          <div className="p-5 rounded-2xl bg-slate-900/90 border border-slate-800 border-l-4 border-l-cyan-400 shadow-xl space-y-2">
            <p className="text-sm font-semibold text-white flex items-center gap-2">
              <Cpu className="w-4 h-4 text-cyan-400" />
              Core Architectural Promise:
            </p>
            <p className="text-sm text-slate-400 leading-relaxed">
              In under 60 seconds, CodeMRI traverses hundreds of files, parses Abstract Syntax Trees (ASTs), and indexes your code into a lightning-fast <strong>Neural Repository Graph (NRG)</strong> stored entirely offline in an embedded SQLite database inside your repository (<code className="text-cyan-300 bg-black/40 px-1.5 py-0.5 rounded font-mono text-xs">.codemri/graph.db</code>).
            </p>
          </div>
        </section>

        {/* SECTION: INSTALLATION */}
        <section id="installation" className="space-y-6 pt-6 border-t border-slate-800/80">
          <div className="flex flex-col space-y-2">
            <h2 className="text-2xl sm:text-3xl font-bold text-white tracking-tight">Quick Installation</h2>
            <p className="text-sm text-slate-400">Install CodeMRI locally as a native compiled binary with zero runtime dependencies.</p>
          </div>

          <div className="bg-slate-950 rounded-2xl border border-slate-800 overflow-hidden shadow-2xl">
            <div className="flex items-center bg-slate-900/80 px-4 pt-3 gap-2 border-b border-slate-800 flex-wrap">
              <button
                onClick={() => setInstallTab('go')}
                className={`px-4 py-2 text-xs font-mono font-bold rounded-t-lg border-b-2 transition-all ${installTab === 'go' ? 'text-cyan-400 border-cyan-400 bg-slate-950/80' : 'text-slate-400 border-transparent hover:text-white'}`}
              >
                Go Install (Recommended)
              </button>
              <button
                onClick={() => setInstallTab('windows')}
                className={`px-4 py-2 text-xs font-mono font-bold rounded-t-lg border-b-2 transition-all ${installTab === 'windows' ? 'text-blue-400 border-blue-400 bg-slate-950/80' : 'text-slate-400 border-transparent hover:text-white'}`}
              >
                🪟 Windows (PowerShell)
              </button>
              <button
                onClick={() => setInstallTab('linux')}
                className={`px-4 py-2 text-xs font-mono font-bold rounded-t-lg border-b-2 transition-all ${installTab === 'linux' ? 'text-emerald-400 border-emerald-400 bg-slate-950/80' : 'text-slate-400 border-transparent hover:text-white'}`}
              >
                🐧 Linux
              </button>
              <button
                onClick={() => setInstallTab('macos')}
                className={`px-4 py-2 text-xs font-mono font-bold rounded-t-lg border-b-2 transition-all ${installTab === 'macos' ? 'text-purple-400 border-purple-400 bg-slate-950/80' : 'text-slate-400 border-transparent hover:text-white'}`}
              >
                🍎 macOS (Apple Silicon / Intel)
              </button>
            </div>

            <div className="p-5 flex items-center justify-between font-mono text-sm bg-slate-950 text-slate-200">
              <pre className="overflow-x-auto whitespace-pre-wrap flex-1 pr-4">
                <code className="text-cyan-300">{installCommands[installTab]}</code>
              </pre>
              <button
                onClick={() => copyToClipboard(installCommands[installTab])}
                className="p-2 rounded-lg bg-slate-900 hover:bg-slate-800 border border-slate-700 text-slate-400 hover:text-white transition-all flex-shrink-0"
                title="Copy to clipboard"
              >
                {copiedText === installCommands[installTab] ? <Check className="w-5 h-5 text-emerald-400" /> : <Copy className="w-5 h-5" />}
              </button>
            </div>
          </div>
          <p className="text-xs text-slate-500 font-mono">⚡ Pre-compiled release archives (.zip & .tar.gz) for v1.0.0 are also available directly inside the project's <code>dist/</code> directory.</p>
        </section>

        {/* SECTION: WHY CODEMRI */}
        <section id="why-codemri" className="space-y-6 pt-6 border-t border-slate-800/80">
          <h2 className="text-2xl sm:text-3xl font-bold text-white tracking-tight">Why CodeMRI vs Traditional Linters?</h2>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div className="p-6 rounded-2xl bg-slate-900/50 border border-slate-800 space-y-3">
              <h3 className="font-bold text-white flex items-center gap-2 text-base">
                <span className="text-red-400">❌</span> Traditional Tools (Sonar / Linters)
              </h3>
              <ul className="text-sm space-y-2 text-slate-400 list-disc list-inside leading-relaxed">
                <li>Focus solely on shallow surface regex syntax rules.</li>
                <li>Suffer from expensive RAM spikes & lengthy cloud uploads.</li>
                <li>Require complex configuration YAMLs and build servers.</li>
                <li>Cannot explain structural cascading impacts before you edit.</li>
              </ul>
            </div>
            <div className="p-6 rounded-2xl bg-gradient-to-br from-cyan-950/20 via-slate-900/90 to-blue-950/20 border border-cyan-500/40 space-y-3 shadow-xl">
              <h3 className="font-bold text-white flex items-center gap-2 text-base">
                <span className="text-emerald-400">✅</span> CodeMRI Neural Intelligence
              </h3>
              <ul className="text-sm space-y-2 text-slate-300 list-disc list-inside leading-relaxed">
                <li>Maps relational symbol dependencies across entire modules.</li>
                <li>Runs 100% locally with zero cloud transmission (ADR-0002).</li>
                <li>Instant 1-second start: zero setup, single executable binary.</li>
                <li>Equipped with an interactive AI Copilot to answer architectural questions.</li>
              </ul>
            </div>
          </div>
        </section>

        {/* SECTION: CORE FUNCTIONS */}
        <section id="pulse" className="space-y-10 pt-6 border-t border-slate-800/80">
          <div className="space-y-2">
            <h2 className="text-2xl sm:text-3xl font-bold text-white tracking-tight">Core Functions & Modules</h2>
            <p className="text-sm text-slate-400">CodeMRI is powered by six synchronized analytical engines that inspect different architectural dimensions.</p>
          </div>

          <div className="grid grid-cols-1 gap-6">
            
            {/* Pulse */}
            <div className="p-6 rounded-2xl bg-slate-900/60 border border-pink-500/30 space-y-4 shadow-lg">
              <div className="flex items-center justify-between flex-wrap gap-2">
                <h3 className="text-xl font-bold text-pink-400 flex items-center gap-2">
                  💓 Pulse — Architectural Health Scoring
                </h3>
                <span className="px-2.5 py-0.5 rounded-full text-[11px] font-mono font-semibold bg-pink-500/10 text-pink-300 border border-pink-500/30">Phase 04</span>
              </div>
              <p className="text-sm text-slate-300 leading-relaxed">
                Pulse evaluates the structural wellness of your repository from <strong>0 to 100</strong>. It automatically flags <strong>Dead Code</strong> (isolated functions receiving zero upstream calls), detects dangerous <strong>Circular Dependency Loops</strong> between modules, and compiles a technical debt grade (Grade A+ to F).
              </p>
              <div className="bg-slate-950 p-3.5 rounded-xl border border-slate-800/80 font-mono text-xs text-slate-400 flex justify-between items-center">
                <span>Output file: <code className="text-pink-300">.codemri/pulse.json</code></span>
                <span className="text-emerald-400 font-bold">✔ Automated in scan</span>
              </div>
            </div>

            {/* Shield */}
            <div id="shield" className="p-6 rounded-2xl bg-slate-900/60 border border-amber-500/30 space-y-4 shadow-lg">
              <div className="flex items-center justify-between flex-wrap gap-2">
                <h3 className="text-xl font-bold text-amber-400 flex items-center gap-2">
                  🛡️ Shield — Security Intelligence & Secrets Scan
                </h3>
                <span className="px-2.5 py-0.5 rounded-full text-[11px] font-mono font-semibold bg-amber-500/10 text-amber-300 border border-amber-500/30">Phase 07</span>
              </div>
              <p className="text-sm text-slate-300 leading-relaxed">
                Shield performs localized static security analysis over the AST nodes. It pinpoints hardcoded credentials, API key leaks, insecure cryptography algorithms (MD5/SHA1), and SQL injection vulnerabilities without needing external vulnerability registries.
              </p>
              <div className="bg-slate-950 p-3.5 rounded-xl border border-slate-800/80 font-mono text-xs text-slate-400 flex justify-between items-center">
                <span>Output file: <code className="text-amber-300">.codemri/security.json</code></span>
                <span className="text-amber-400 font-bold">✔ Realtime security assessment</span>
              </div>
            </div>

            {/* Velocity */}
            <div id="velocity" className="p-6 rounded-2xl bg-slate-900/60 border border-emerald-500/30 space-y-4 shadow-lg">
              <div className="flex items-center justify-between flex-wrap gap-2">
                <h3 className="text-xl font-bold text-emerald-400 flex items-center gap-2">
                  🚀 Velocity — Performance Diagnostic
                </h3>
                <span className="px-2.5 py-0.5 rounded-full text-[11px] font-mono font-semibold bg-emerald-500/10 text-emerald-300 border border-emerald-500/30">Phase 08</span>
              </div>
              <p className="text-sm text-slate-300 leading-relaxed">
                Velocity diagnoses bundling performance and compilation bottlenecks. It identifies bloated monolithic source files (&gt;50KB/file), tightly-coupled modules (&gt;15 imports per file), and package namespace collisions across directory hierarchies.
              </p>
              <div className="bg-slate-950 p-3.5 rounded-xl border border-slate-800/80 font-mono text-xs text-slate-400 flex justify-between items-center">
                <span>Output file: <code className="text-emerald-300">.codemri/performance.json</code></span>
                <span className="text-emerald-400 font-bold">✔ Build optimization ready</span>
              </div>
            </div>

            {/* Vision */}
            <div id="vision" className="p-6 rounded-2xl bg-slate-900/60 border border-cyan-500/30 space-y-4 shadow-lg">
              <div className="flex items-center justify-between flex-wrap gap-2">
                <h3 className="text-xl font-bold text-cyan-400 flex items-center gap-2">
                  👁️ Vision — Instant Impact Radius
                </h3>
                <span className="px-2.5 py-0.5 rounded-full text-[11px] font-mono font-semibold bg-cyan-500/10 text-cyan-300 border border-cyan-500/30">Phase 05</span>
              </div>
              <p className="text-sm text-slate-300 leading-relaxed">
                Before you refactor or modify any function, Vision computes its <strong>Impact Radius (0% - 100%)</strong>. By tracing upstream callers and downstream dependencies in the SQLite database, it warns you if editing a symbol risks triggering cascading bugs across distant modules.
              </p>
              <div className="bg-slate-950 p-3.5 rounded-xl border border-slate-800/80 font-mono text-xs text-slate-400 flex justify-between items-center">
                <span>API route: <code className="text-cyan-300">/api/graph/impact/:id</code></span>
                <span className="text-cyan-400 font-bold">✔ Interactive graph UI</span>
              </div>
            </div>

            {/* Cortex */}
            <div id="cortex" className="p-6 rounded-2xl bg-slate-900/60 border border-blue-500/30 space-y-4 shadow-lg">
              <div className="flex items-center justify-between flex-wrap gap-2">
                <h3 className="text-xl font-bold text-blue-400 flex items-center gap-2">
                  🤖 Cortex — AI Reasoning Engine (with freemodel.dev)
                </h3>
                <span className="px-2.5 py-0.5 rounded-full text-[11px] font-mono font-semibold bg-blue-500/10 text-blue-300 border border-blue-500/30">Phase 06 & v1.0.0</span>
              </div>
              <p className="text-sm text-slate-300 leading-relaxed">
                Cortex integrates deep conversational artificial intelligence right into your analysis workflow. Enabled by default with <code className="text-cyan-300 bg-black/40 px-1.5 py-0.5 rounded font-mono text-xs">freemodel.dev</code> API endpoints and backed by local structural offline fallbacks, Cortex lets you chat with your repository directly in the browser or terminal.
              </p>
              <div className="bg-slate-950 p-3.5 rounded-xl border border-slate-800/80 font-mono text-xs text-slate-400 flex justify-between items-center">
                <span>API route: <code className="text-blue-300">POST /api/chat</code></span>
                <span className="text-blue-400 font-bold">✔ Intelligent Conversational AI</span>
              </div>
            </div>

          </div>
        </section>

        {/* SECTION: HOW IT WORKS */}
        <section id="architecture" className="space-y-6 pt-6 border-t border-slate-800/80">
          <h2 className="text-2xl sm:text-3xl font-bold text-white tracking-tight">How It Works (Under the Hood)</h2>
          <p className="text-sm text-slate-400 leading-relaxed">
            CodeMRI bypasses traditional slow interpretation loops by executing a systematic 3-stage high-speed parsing workflow:
          </p>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-6 pt-2">
            <div className="p-5 rounded-2xl bg-slate-900/40 border border-slate-800 space-y-3 relative">
              <div className="w-8 h-8 rounded-full bg-cyan-500/20 border border-cyan-500/50 text-cyan-300 font-black flex items-center justify-center font-mono text-sm">1</div>
              <h4 className="font-bold text-white text-base">Concurrent AST Indexer</h4>
              <p className="text-xs text-slate-400 leading-relaxed">
                The embedded <strong>Atlas Engine</strong> spawns parallel worker goroutines to parse lexical syntax trees across Go, TypeScript, JavaScript, Python, and C/C++ files simultaneously.
              </p>
            </div>

            <div id="sqlite-engine" className="p-5 rounded-2xl bg-slate-900/40 border border-slate-800 space-y-3 relative">
              <div className="w-8 h-8 rounded-full bg-blue-500/20 border border-blue-500/50 text-blue-300 font-black flex items-center justify-center font-mono text-sm">2</div>
              <h4 className="font-bold text-white text-base">SQLite Relational Index</h4>
              <p className="text-xs text-slate-400 leading-relaxed">
                Instead of massive memory bloat, symbol nodes (classes, functions, interfaces) and relational edges (calls, imports, inheritances) are committed to a compact, indexed SQLite database at <code className="text-blue-300">.codemri/graph.db</code>.
              </p>
            </div>

            <div id="offline-first" className="p-5 rounded-2xl bg-slate-900/40 border border-slate-800 space-y-3 relative">
              <div className="w-8 h-8 rounded-full bg-emerald-500/20 border border-emerald-500/50 text-emerald-300 font-black flex items-center justify-center font-mono text-sm">3</div>
              <h4 className="font-bold text-white text-base">Zero-Cloud Guarantee</h4>
              <p className="text-xs text-slate-400 leading-relaxed">
                In adherence to <strong>ADR-0002</strong>, your proprietary source code never gets transmitted to external third-party cloud servers during diagnostic scanning. Everything runs CPU-bound on your local machine.
              </p>
            </div>
          </div>
        </section>

        {/* SECTION: PRACTICAL USAGE */}
        <section id="usage-local" className="space-y-8 pt-6 border-t border-slate-800/80">
          <h2 className="text-2xl sm:text-3xl font-bold text-white tracking-tight">Practical Usage & Examples</h2>
          
          <div className="space-y-6">
            
            {/* Example 1: One-Command Workflow */}
            <div className="p-6 rounded-2xl bg-slate-900/50 border border-slate-800 space-y-4">
              <h3 className="font-bold text-white text-lg flex items-center gap-2">
                ⚡ 1. One-Command Local Workflow
              </h3>
              <p className="text-sm text-slate-400">
                To inspect any project directory, simply navigate into it and type <code className="text-cyan-300 bg-slate-950 px-2 py-0.5 rounded font-mono text-xs">codemri</code> without any arguments. It automatically scans your code and launches the interactive dashboard at <code className="text-white">http://localhost:4000</code>.
              </p>
              <div className="p-4 rounded-xl bg-slate-950 border border-slate-800 font-mono text-xs text-slate-300 flex justify-between items-center">
                <span>$ cd /path/to/my-project &amp;&amp; <strong className="text-cyan-400">codemri</strong></span>
                <button onClick={() => copyToClipboard('codemri')} className="text-slate-500 hover:text-white"><Copy className="w-4 h-4" /></button>
              </div>
            </div>

            {/* Example 2: Online GitHub Scanning */}
            <div id="usage-online" className="p-6 rounded-2xl bg-gradient-to-br from-blue-950/30 via-slate-900/90 to-cyan-950/30 border border-cyan-500/40 space-y-4 shadow-xl">
              <div className="flex items-center justify-between flex-wrap gap-2">
                <h3 className="font-bold text-cyan-300 text-lg flex items-center gap-2">
                  🌐 2. Online GitHub Repository Auto-Cloning
                </h3>
                <span className="px-2 py-0.5 rounded text-[10px] font-mono uppercase bg-cyan-500/20 text-cyan-300 font-extrabold">NEW in v1.0.0</span>
              </div>
              <p className="text-sm text-slate-300 leading-relaxed">
                You do not need to manually clone repositories before analyzing them! You can pass any public GitHub repository URL directly to CodeMRI. It will execute a high-speed shallow clone (<code className="text-cyan-200">--depth=1</code>) directly into your local cache (<code className="text-slate-400">~/.codemri/cache</code>) and serve analytics instantly.
              </p>
              <div className="p-4 rounded-xl bg-slate-950 border border-cyan-500/40 font-mono text-xs sm:text-sm text-slate-200 flex justify-between items-center">
                <span>$ codemri <strong className="text-cyan-400">https://github.com/torvalds/linux</strong></span>
                <button onClick={() => copyToClipboard('codemri https://github.com/torvalds/linux')} className="text-slate-400 hover:text-white"><Copy className="w-4 h-4" /></button>
              </div>
              <p className="text-xs text-slate-400 italic">
                💡 Tip: You can also paste GitHub URLs directly into the search bar inside the web dashboard at <code>localhost:4000</code> to switch analyzed repos on the fly!
              </p>
            </div>

            {/* Example 3: AI Chatbot Usage */}
            <div id="usage-ai" className="p-6 rounded-2xl bg-slate-900/50 border border-blue-500/30 space-y-4">
              <h3 className="font-bold text-white text-lg flex items-center gap-2">
                🤖 3. Interactive AI Chatbot (Web &amp; CLI)
              </h3>
              <p className="text-sm text-slate-400 leading-relaxed">
                Once a repository is scanned, you can ask architectural questions in plain natural language. In the browser dashboard (<code className="text-blue-300">localhost:4000</code>), use the built-in Cortex AI console at the bottom of the screen. In the terminal, execute the <code className="text-cyan-300">explain</code> subcommand:
              </p>
              <div className="p-4 rounded-xl bg-slate-950 border border-slate-800 font-mono text-xs sm:text-sm text-slate-300 flex justify-between items-center">
                <span>$ codemri explain <strong className="text-amber-300">"What are the main functions in the authentication module?"</strong></span>
                <button onClick={() => copyToClipboard('codemri explain "What are the main functions in the authentication module?"')} className="text-slate-500 hover:text-white"><Copy className="w-4 h-4" /></button>
              </div>
            </div>

            {/* Example 4: CLI Cheat Sheet */}
            <div id="usage-cli" className="space-y-4 pt-4">
              <h3 className="font-bold text-white text-xl tracking-tight">Complete CLI Command Reference</h3>
              <div className="overflow-x-auto rounded-2xl border border-slate-800 bg-slate-950 shadow-xl">
                <table className="w-full text-left border-collapse font-sans text-sm">
                  <thead>
                    <tr className="border-b border-slate-800 bg-slate-900/90 text-slate-300 font-bold font-mono text-xs uppercase">
                      <th className="py-3.5 px-5">Command</th>
                      <th className="py-3.5 px-5">Target Argument</th>
                      <th className="py-3.5 px-5">Description &amp; Behavior</th>
                    </tr>
                  </thead>
                  <tbody className="divide-y divide-slate-800/60 text-slate-300 font-normal">
                    <tr className="hover:bg-slate-900/40 transition-colors">
                      <td className="py-3 px-5 font-mono font-bold text-cyan-400">codemri</td>
                      <td className="py-3 px-5 font-mono text-xs text-slate-400">[dir / git-url]</td>
                      <td className="py-3 px-5">One-command workflow: automatically scans directory or clones Git URL, generates SQLite database, and opens visual dashboard.</td>
                    </tr>
                    <tr className="hover:bg-slate-900/40 transition-colors">
                      <td className="py-3 px-5 font-mono font-bold text-cyan-400">codemri scan</td>
                      <td className="py-3 px-5 font-mono text-xs text-slate-400">[dir / git-url]</td>
                      <td className="py-3 px-5">Run high-speed concurrent AST syntax parsers and assemble the Neural Repository Graph without serving UI.</td>
                    </tr>
                    <tr className="hover:bg-slate-900/40 transition-colors">
                      <td className="py-3 px-5 font-mono font-bold text-blue-400">codemri analyze</td>
                      <td className="py-3 px-5 font-mono text-xs text-slate-400">[dir / git-url]</td>
                      <td className="py-3 px-5">Execute full diagnostic suite (Pulse health score, Shield security vulnerabilities, and Velocity performance profile).</td>
                    </tr>
                    <tr className="hover:bg-slate-900/40 transition-colors">
                      <td className="py-3 px-5 font-mono font-bold text-emerald-400">codemri serve</td>
                      <td className="py-3 px-5 font-mono text-xs text-slate-400">[dir]</td>
                      <td className="py-3 px-5">Launch local Fiber HTTP server and render interactive SPA dashboard on port 4000.</td>
                    </tr>
                    <tr className="hover:bg-slate-900/40 transition-colors">
                      <td className="py-3 px-5 font-mono font-bold text-amber-400">codemri explain</td>
                      <td className="py-3 px-5 font-mono text-xs text-amber-300">"question"</td>
                      <td className="py-3 px-5">Ask conversational AI questions about the codebase architecture using Cortex structural engine.</td>
                    </tr>
                    <tr className="hover:bg-slate-900/40 transition-colors">
                      <td className="py-3 px-5 font-mono font-bold text-purple-400">codemri graph</td>
                      <td className="py-3 px-5 font-mono text-xs text-slate-400">[query]</td>
                      <td className="py-3 px-5">Terminal interactive topology search tool for finding symbol connections and JSON exports.</td>
                    </tr>
                    <tr className="hover:bg-slate-900/40 transition-colors">
                      <td className="py-3 px-5 font-mono font-bold text-pink-400">codemri doctor</td>
                      <td className="py-3 px-5 font-mono text-xs text-slate-500">None</td>
                      <td className="py-3 px-5">Verify local Go environment, filesystem read/write capability, SQLite bindings, and network port availability.</td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>

          </div>
        </section>

      </main>
    </div>
  );
};

export default GetStartedDocs;
