import React from 'react';
import { CheckCircle2, Zap, Cpu, Cloud, ShieldCheck } from 'lucide-react';

export const Roadmap: React.FC = () => {
  const roadmapVersions = [
    {
      ver: "v0.1.0",
      code: "Genesis",
      title: "Project Foundation & Core Skeleton",
      status: "COMPLETED",
      items: ["Monorepo setup in Go", "CLI engine via Cobra (scan, serve, version, doctor)", "Go Fiber HTTP backend skeleton", "Neural Repository Graph (NRG) interface structures", "Apache 2.0 Licensing"],
      icon: <CheckCircle2 className="w-6 h-6 text-emerald-400" />,
      badgeColor: "bg-emerald-500/20 text-emerald-300 border-emerald-500/40"
    },
    {
      ver: "v0.2.0",
      code: "Atlas",
      title: "High-Speed Scanner & Universal AST Parser",
      status: "COMPLETED",
      items: ["Multi-language AST pipelines (Go, TS, Python, SQL, Java, etc.)", "High-speed concurrent worker pool scanning", "Automatic language detection & metrics", "Local caching engine"],
      icon: <CheckCircle2 className="w-6 h-6 text-emerald-400" />,
      badgeColor: "bg-emerald-500/20 text-emerald-300 border-emerald-500/40"
    },
    {
      ver: "v0.3.0",
      code: "Neuron",
      title: "Embedded SQLite Relational Engine",
      status: "COMPLETED",
      items: ["Pure-Go SQLite relational graph engine (.codemri/graph.db)", "Sub-millisecond dependency & call graph resolution", "Reactive dark-mode embedded Web UI on port 4000", "Zero CGO & zero external dependencies"],
      icon: <CheckCircle2 className="w-6 h-6 text-emerald-400" />,
      badgeColor: "bg-emerald-500/20 text-emerald-300 border-emerald-500/40"
    },
    {
      ver: "v0.4.0",
      code: "Pulse",
      title: "Architectural Technical Debt & Health Engine",
      status: "COMPLETED",
      items: ["Authoritative 0-100 repository health scores (A+ to F grades)", "Dead code & zero-incoming call symbol discovery", "Circular dependency import cycle detection via DFS", "Actionable AI architectural recommendations"],
      icon: <CheckCircle2 className="w-6 h-6 text-emerald-400" />,
      badgeColor: "bg-emerald-500/20 text-emerald-300 border-emerald-500/40"
    },
    {
      ver: "v0.5.0",
      code: "Vision",
      title: "Interactive Force-Directed 2D/3D Topology Canvas",
      status: "COMPLETED",
      items: ["Force-directed visual node rendering in browser", "Interactive drill-down on circular loops and complexity hotpots", "Real-time module filtering & sub-graph isolation", "Dynamic smooth zoom & search filtering"],
      icon: <CheckCircle2 className="w-6 h-6 text-emerald-400" />,
      badgeColor: "bg-emerald-500/20 text-emerald-300 border-emerald-500/40"
    },
    {
      ver: "v1.0.0+",
      code: "Cloud & Enterprise Ecosystem",
      title: "Team Workspaces & Automated PR Reviews",
      status: "FUTURE ECOSYSTEM",
      items: ["CodeMRI Cloud online synchronization", "Team Dashboard contribution mapping", "Enterprise Governance (RBAC, LDAP/SSO, Audit)", "GitHub App automated AI Pull Request Reviewer"],
      icon: <Cloud className="w-6 h-6 text-amber-400" />,
      badgeColor: "bg-amber-500/20 text-amber-300 border-amber-500/40"
    }
  ];

  return (
    <div className="max-w-5xl mx-auto px-4 sm:px-6 lg:px-8 py-12 text-left font-sans">
      
      {/* Title Header */}
      <div className="text-center mb-14">
        <span className="text-xs uppercase font-mono px-3 py-1 bg-blue-500/10 border border-blue-500/30 text-blue-400 rounded-full font-bold">
          🗺️ Transparent Release Schedule
        </span>
        <h1 className="text-3xl sm:text-5xl font-black text-white mt-3 mb-4">
          The Evolutionary Roadmap
        </h1>
        <p className="text-slate-400 max-w-2xl mx-auto text-sm sm:text-base leading-relaxed">
          Every release carries a meaningful structural codename. We prioritize community trust, code quality, and adoption above immediate monetization.
        </p>
      </div>

      {/* Timeline List */}
      <div className="relative border-l-2 border-slate-800 ml-4 sm:ml-8 space-y-12">
        {roadmapVersions.map((item, index) => (
          <div key={index} className="relative pl-8 sm:pl-12 group">
            
            {/* Circle Node on Timeline */}
            <div className="absolute -left-[18px] top-1 w-9 h-9 rounded-2xl bg-slate-900 border-2 border-slate-700 group-hover:border-cyan-400 flex items-center justify-center transition-all duration-300 shadow-xl">
              {item.icon}
            </div>

            {/* Content Box */}
            <div className="glass-panel p-6 sm:p-8 rounded-3xl border border-slate-800/90 group-hover:border-slate-700 transition-all shadow-2xl">
              <div className="flex flex-wrap items-center justify-between gap-2 mb-3">
                <div className="flex items-center space-x-3">
                  <span className="font-mono font-black text-lg sm:text-2xl text-white">
                    {item.ver}
                  </span>
                  <span className="font-mono text-sm sm:text-lg font-bold text-cyan-400 bg-slate-950 px-3 py-0.5 rounded-full border border-cyan-500/30">
                    "{item.code}"
                  </span>
                </div>
                <span className={`text-[10px] sm:text-xs font-mono font-bold px-3 py-1 rounded-full border ${item.badgeColor}`}>
                  {item.status}
                </span>
              </div>

              <h3 className="text-lg sm:text-xl font-bold text-slate-200 mb-4">
                {item.title}
              </h3>

              <ul className="grid grid-cols-1 sm:grid-cols-2 gap-3 text-xs sm:text-sm text-slate-400">
                {item.items.map((point, idx) => (
                  <li key={idx} className="flex items-center space-x-2 bg-slate-950/60 p-2.5 rounded-xl border border-slate-900">
                    <span className="w-1.5 h-1.5 rounded-full bg-cyan-400 flex-shrink-0" />
                    <span>{point}</span>
                  </li>
                ))}
              </ul>
            </div>

          </div>
        ))}
      </div>

      {/* Bottom Year 1 Goals Callout */}
      <div className="mt-16 bg-gradient-to-r from-cyan-950/40 via-slate-900 to-purple-950/40 border border-cyan-500/40 rounded-3xl p-8 text-center shadow-2xl">
        <h3 className="text-2xl font-black text-white mb-2">Year One Adoption Mandate</h3>
        <p className="text-slate-300 max-w-xl mx-auto text-sm mb-6">
          Before initiating paywalls or SaaS cloud subscriptions, our engineering obsession is establishing CodeMRI as the unquestioned industry reference:
        </p>
        <div className="grid grid-cols-2 sm:grid-cols-4 gap-4 max-w-3xl mx-auto">
          <div className="p-4 bg-slate-950 rounded-2xl border border-slate-800 font-mono">
            <p className="text-xl sm:text-2xl font-black text-amber-300">10,000+</p>
            <p className="text-xs text-slate-400">GitHub Stars ⭐</p>
          </div>
          <div className="p-4 bg-slate-950 rounded-2xl border border-slate-800 font-mono">
            <p className="text-xl sm:text-2xl font-black text-cyan-400">300+</p>
            <p className="text-xs text-slate-400">Contributors 🧑‍💻</p>
          </div>
          <div className="p-4 bg-slate-950 rounded-2xl border border-slate-800 font-mono">
            <p className="text-xl sm:text-2xl font-black text-emerald-400">100k+</p>
            <p className="text-xs text-slate-400">CLI Installs 📦</p>
          </div>
          <div className="p-4 bg-slate-950 rounded-2xl border border-slate-800 font-mono">
            <p className="text-xl sm:text-2xl font-black text-purple-400">Apache 2.0</p>
            <p className="text-xs text-slate-400">Enterprise Safe 🛡️</p>
          </div>
        </div>
      </div>

    </div>
  );
};
