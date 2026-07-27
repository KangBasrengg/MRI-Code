import type { DocSection } from '../types';

export const DOCS_DATA: DocSection[] = [
  // ─────────────────────────────────────────────────
  // SECTION 1: Installation
  // ─────────────────────────────────────────────────
  {
    id: 'installation',
    title: '⚡ CLI Installation Guide',
    category: 'Getting Started',
    summary: 'Install CodeMRI as a lightweight native Go binary. No cloning, no npm, no runtime dependencies.',
    content: '',
    blocks: [
      {
        type: 'paragraph',
        text: 'CodeMRI is distributed as a single standalone Go binary under the Apache 2.0 License. It compiles to native machine code—meaning it launches in under 5 milliseconds with zero runtime dependencies. No Node.js, no JVM, no Python environment required.'
      },
      {
        type: 'callout',
        variant: 'info',
        text: 'You never need to clone the CodeMRI repository. Just install the binary globally and run it inside any software project you want to inspect.'
      },
      { type: 'heading', text: 'Windows' },
      {
        type: 'paragraph',
        text: 'Install globally using Windows Package Manager or Scoop:'
      },
      {
        type: 'code',
        language: 'powershell',
        code: '# Via WinGet (recommended)\nwinget install codemri\n\n# Or via Scoop\nscoop install codemri'
      },
      { type: 'heading', text: 'macOS' },
      {
        type: 'paragraph',
        text: 'Install globally via Homebrew:'
      },
      {
        type: 'code',
        language: 'bash',
        code: 'brew install codemri'
      },
      { type: 'heading', text: 'Linux' },
      {
        type: 'paragraph',
        text: 'Execute the automated native binary installer or download the release archive:'
      },
      {
        type: 'code',
        language: 'bash',
        code: 'curl -fsSL https://codemri.dev/install.sh | sh'
      },
      { type: 'heading', text: 'Go Developers (go install)' },
      {
        type: 'paragraph',
        text: 'If you already have Go 1.21+ installed, use the native Go module installer:'
      },
      {
        type: 'code',
        language: 'bash',
        code: 'go install github.com/KangBasrengg/MRI-Code/cmd/codemri@latest'
      },
      { type: 'heading', text: 'Manual Download (GitHub Releases)' },
      {
        type: 'paragraph',
        text: 'Download pre-compiled cross-platform archives directly from GitHub Releases:'
      },
      {
        type: 'list',
        items: [
          'codemri_v0.4.0_windows_amd64.zip',
          'codemri_v0.4.0_linux_amd64.tar.gz',
          'codemri_v0.4.0_darwin_arm64.tar.gz'
        ]
      },
      {
        type: 'paragraph',
        text: 'Extract the executable and add it to your system PATH variable. That\'s it—you\'re ready to analyze any codebase.'
      },
      { type: 'heading', text: 'Verify Installation' },
      {
        type: 'code',
        language: 'bash',
        code: '# Confirm the binary is accessible\ncodemri version\n\n# Expected output:\n# CodeMRI v0.4.0 ("Pulse")\n# 🎯 Core Engine : Neural Repository Graph (NRG) - SQLite Relational & Pulse Analytical Engine\n# 🚀 Go Runtime  : go1.26.5 (windows/amd64)\n# 💎 Philosophy  : We don\'t generate code. We generate understanding.'
      }
    ]
  },

  // ─────────────────────────────────────────────────
  // SECTION 2: One-Word Killer Workflow
  // ─────────────────────────────────────────────────
  {
    id: 'killer-workflow',
    title: '🔥 The "One-Word" Killer Workflow',
    category: 'Core Usage',
    summary: 'Clone a new repository. Type one word. Understand the entire architecture in under 60 seconds.',
    content: '',
    blocks: [
      {
        type: 'paragraph',
        text: 'Imagine you just joined a new team. Your tech lead says: "Here\'s our monorepo with 5,000 files. Good luck." Instead of spending days reading folder structures, you follow a three-step routine:'
      },
      {
        type: 'code',
        language: 'bash',
        code: 'git clone https://github.com/org/enterprise-app\ncd enterprise-app\ncodemri'
      },
      {
        type: 'callout',
        variant: 'success',
        text: '"Whenever you clone a new repository, just run codemri first." — This is the habit we want to build in every developer.'
      },
      { type: 'heading', text: 'What Happens Automatically' },
      {
        type: 'paragraph',
        text: 'When you type codemri with zero flags inside any project folder, the engine performs the following automated sequence:'
      },
      {
        type: 'steps',
        items: [
          { label: 'Root & Stack Detection', description: 'Identifies the project root directory, detects primary languages (Go, TypeScript, Python, Java, PHP, Rust, etc.), and determines the framework ecosystem.' },
          { label: 'Deterministic Syntax Scanning', description: 'If no local .codemri/ cache exists, the Atlas Engine executes multi-language AST parsers concurrently across all source files. Typical scan time: 10-50ms for medium repositories.' },
          { label: 'Non-Invasive Storage', description: 'All analytical results are compiled into an isolated .codemri/ folder inside your project. Your actual source code is never modified, uploaded, or touched in any way.' },
          { label: 'Auto-Launch Dashboard', description: 'Starts a local HTTP server on port 4000 (auto-increments if busy) and opens your system web browser automatically—delivering an experience just like Prisma Studio or Laravel Telescope.' }
        ]
      },
      { type: 'heading', text: 'The Complete Flow Diagram' },
      {
        type: 'code',
        language: 'text',
        code: '  Install CodeMRI\n        │\n        ▼\n  Navigate to repository\n        │\n        ▼\n  $ codemri\n        │\n        ├── Detect root project\n        ├── Identify languages & framework\n        ├── Run concurrent AST parsers\n        ├── Compile Neural Repository Graph\n        ├── Store results in .codemri/\n        │\n        ▼\n  Dashboard opens in browser\n        │\n        ▼\n  ☕ "I understand the codebase now."'
      },
      { type: 'heading', text: 'Separate Commands (Advanced Users)' },
      {
        type: 'paragraph',
        text: 'For granular control, you can run the scan and serve steps independently:'
      },
      {
        type: 'code',
        language: 'bash',
        code: '# Step 1: Only scan (no server)\ncodemri scan .\n\n# Step 2: Only serve dashboard (reads cached .codemri/)\ncodemri serve --port 5050'
      }
    ]
  },

  // ─────────────────────────────────────────────────
  // SECTION 3: How the Scanner Works
  // ─────────────────────────────────────────────────
  {
    id: 'scanner-engine',
    title: '🔍 Atlas Scanner Engine',
    category: 'Architecture',
    summary: 'High-speed concurrent directory traversal, language detection, and AST symbol extraction.',
    content: '',
    blocks: [
      {
        type: 'paragraph',
        text: 'The Atlas Engine is CodeMRI\'s Phase 2 scanner implementation. It performs three critical operations in one concurrent pass: directory traversal, language classification, and deterministic syntax extraction.'
      },
      { type: 'heading', text: 'Directory Traversal' },
      {
        type: 'paragraph',
        text: 'The engine walks the entire repository file tree using Go\'s high-performance filepath.WalkDir API. It automatically excludes noise directories that don\'t contain meaningful source code:'
      },
      {
        type: 'list',
        items: [
          '.git — Version control internals',
          'node_modules — NPM dependency trees',
          'vendor — Go vendored dependencies',
          'dist / build / bin — Compiled output artifacts',
          '__pycache__ — Python bytecode caches',
          '.next / .vercel — Framework build caches',
          '.codemri — Our own analytical workspace'
        ]
      },
      { type: 'heading', text: 'Language Detection' },
      {
        type: 'paragraph',
        text: 'Every file passes through the language classifier which identifies 15+ programming languages by file extension, special filenames (Dockerfile, Makefile), and structural patterns:'
      },
      {
        type: 'code',
        language: 'text',
        code: 'Supported Languages:\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n Go          │ .go\n TypeScript  │ .ts, .tsx\n JavaScript  │ .js, .jsx, .mjs, .cjs\n Python      │ .py, .pyw\n Java        │ .java\n PHP         │ .php, .phtml\n Rust        │ .rs\n SQL         │ .sql\n HTML        │ .html, .htm\n CSS         │ .css, .scss, .sass, .less\n Shell       │ .sh, .bash, .zsh\n Markdown    │ .md, .markdown\n JSON        │ .json\n YAML        │ .yml, .yaml\n Docker      │ Dockerfile, Dockerfile.*\n Make        │ Makefile, GNUmakefile'
      },
      { type: 'heading', text: 'Concurrent AST Extraction' },
      {
        type: 'paragraph',
        text: 'After discovery, all source files are dispatched to a pool of concurrent worker goroutines. Each worker invokes the appropriate language-specific parser to extract structural symbols:'
      },
      {
        type: 'list',
        items: [
          'Go Parser — Uses Go\'s native go/parser and go/ast packages for precise extraction of packages, functions, methods, structs, and interfaces',
          'Universal Parser — Regex-driven tokenizer for TypeScript, JavaScript, Python, PHP, Rust, Java, and SQL',
          'Line Metrics — Every file gets LOC (Lines of Code), comment count, and blank line tallies'
        ]
      },
      {
        type: 'callout',
        variant: 'info',
        text: 'Benchmark: The Atlas Engine scanned 71 files across Go + TypeScript codebases and compiled 87 graph nodes with 182 relational edges in just 13.4 milliseconds.'
      }
    ]
  },

  // ─────────────────────────────────────────────────
  // SECTION 4: Neural Repository Graph (NRG)
  // ─────────────────────────────────────────────────
  {
    id: 'nrg-explained',
    title: '🧠 Neural Repository Graph (NRG)',
    category: 'Architecture',
    summary: 'The Single Source of Truth — how CodeMRI represents your entire codebase as a structured relational map.',
    content: '',
    blocks: [
      {
        type: 'paragraph',
        text: 'The Neural Repository Graph is the central architectural concept that makes CodeMRI fundamentally different from other tools. Instead of reading raw source files every time, all features — AI Chat, Architecture View, Impact Analysis, Security Scanning, and Technical Debt assessment — consume the NRG as their single source of truth.'
      },
      {
        type: 'callout',
        variant: 'warning',
        text: 'Parser understands code. AI understands meaning. The NRG bridges these two worlds by providing structured relational data that both humans and machines can reason about.'
      },
      { type: 'heading', text: 'Graph Nodes (Symbols)' },
      {
        type: 'paragraph',
        text: 'Every architectural symbol discovered during scanning becomes a Node in the graph. Each node carries metadata about its type, location, and properties:'
      },
      {
        type: 'code',
        language: 'text',
        code: 'Node Types:\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n FILE           │ Individual source code files\n FUNCTION       │ Functions, methods, handlers\n CLASS          │ Classes, structs, interfaces\n MODULE         │ Modules, packages, namespaces\n PACKAGE        │ Go packages, npm packages\n ROUTE          │ API routes, HTTP endpoints\n DATABASE_TABLE │ SQL tables, ORM models\n SERVICE        │ Microservices, external integrations'
      },
      { type: 'heading', text: 'Graph Edges (Relations)' },
      {
        type: 'paragraph',
        text: 'Relationships between nodes are modeled as directed edges with typed semantics:'
      },
      {
        type: 'code',
        language: 'text',
        code: 'Edge Types:\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n IMPORTS     │ File A imports module B\n CALLS       │ Function X calls function Y\n INHERITS    │ Class C extends class D\n QUERIES     │ Handler H queries table T\n EXPOSES     │ File F exposes function G\n DEPENDS_ON  │ Module M depends on module N'
      },
      { type: 'heading', text: 'Phase 3 ("Neuron") Relational Storage Engine' },
      {
        type: 'paragraph',
        text: 'In Phase 3 ("Neuron"), the compiled NRG is persisted simultaneously as a zero-dependency embedded SQLite database (for sub-millisecond relational queries and SQL indexing) and as JSON inside your project\'s local workspace. Your source code is never modified:'
      },
      {
        type: 'code',
        language: 'text',
        code: 'your-project/\n  ├── src/                    ← Your code (untouched)\n  ├── package.json\n  └── .codemri/               ← CodeMRI workspace\n       ├── graph.db            ← SQLite Relational Graph Index (Phase 3 Neuron)\n       ├── graph.json          ← Compiled NRG JSON backup (nodes + edges)\n       ├── repository.json     ← Analytical metadata & stats\n       ├── cache/              ← Incremental scan cache\n       └── logs/               ← Diagnostic logs'
      }
    ]
  },

  // ─────────────────────────────────────────────────
  // SECTION 5: Two Modes (Offline vs Online)
  // ─────────────────────────────────────────────────
  {
    id: 'two-modes',
    title: '🛡️ Two Modes: Offline vs Online AI',
    category: 'Architecture & Security',
    summary: '100% private local analytics by default. Optional AI reasoning that only shares graph structures, never raw code.',
    content: '',
    blocks: [
      {
        type: 'paragraph',
        text: 'Enterprise developers work with proprietary codebases that must never be leaked to external AI endpoints. CodeMRI separates analysis into two rigorous operating modes to guarantee intellectual property protection.'
      },
      { type: 'heading', text: 'Mode 1: Offline (Default)' },
      {
        type: 'paragraph',
        text: 'In default operating mode, everything runs entirely on your machine. Zero internet traffic. Zero API tokens. Complete intellectual property privacy.'
      },
      {
        type: 'code',
        language: 'text',
        code: 'Offline Pipeline:\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n Source Code → AST Parser → NRG → Dashboard\n        ↑                            ↓\n    Your machine              Your browser\n\n  ✓ No internet connection needed\n  ✓ No API keys or tokens\n  ✓ No data leaves your computer\n  ✓ Works in air-gapped environments'
      },
      { type: 'heading', text: 'Mode 2: Online AI (Optional & Targeted)' },
      {
        type: 'paragraph',
        text: 'When you explicitly invoke AI reasoning commands, CodeMRI does NOT upload your entire repository. Instead, it extracts only the mathematically verified Neural Repository Graph — a structured relational summary of symbols and connections — and sends that lightweight abstract data to your chosen LLM provider.'
      },
      {
        type: 'code',
        language: 'bash',
        code: '# Ask AI to explain a specific module\ncodemri explain auth_service\n\n# What gets transmitted:\n# ✓ Abstract graph: "auth_service CALLS validateToken"\n# ✓ Metadata: "12 functions, 3 dependencies"\n# ✗ NOT transmitted: Actual source code lines\n# ✗ NOT transmitted: API keys, secrets, .env files'
      },
      {
        type: 'callout',
        variant: 'success',
        text: 'This architectural separation guarantees that AI understands structural meaning without exposing raw proprietary source code. Think of it as sharing an X-ray image instead of performing surgery in public.'
      }
    ]
  },

  // ─────────────────────────────────────────────────
  // SECTION 6: Why Go Binary Over NPM
  // ─────────────────────────────────────────────────
  {
    id: 'why-go-binary',
    title: '💎 Why Go Binary Over NPM?',
    category: 'Project Manifesto',
    summary: 'A code intelligence engine must never force Java or Rust developers to install Node.js.',
    content: '',
    blocks: [
      {
        type: 'paragraph',
        text: 'Many developer tools distribute via npm install -g. We deliberately rejected NPM as our distribution mechanism. Here\'s why this decision is fundamental to CodeMRI\'s mission:'
      },
      {
        type: 'steps',
        items: [
          { label: 'Language-Agnostic Adoption', description: 'CodeMRI scans Go, Java, Rust, C#, PHP/Laravel, Python, and TypeScript repositories. A Java architect or a Rust systems programmer should never need to install Node.js just to analyze code dependencies.' },
          { label: 'Zero Runtime Overhead', description: 'Go compiles to a single native machine-code binary. It launches in under 5 milliseconds without spinning up virtual memory interpreters, JIT compilers, or garbage collection warmup phases.' },
          { label: 'Cross-Platform Resilience', description: 'One executable works instantly across Linux servers, WSL2, macOS M-series chips, and Windows environments. No compatibility matrix. No "works on my machine" syndrome.' },
          { label: 'Enterprise Licensing', description: 'Apache 2.0 provides patent protection and explicit commercial use grants. No licensing traps, no CLA controversies, no dual-license gotchas.' }
        ]
      },
      {
        type: 'callout',
        variant: 'info',
        text: 'Think of it this way: Would you trust a hospital MRI machine that requires you to install a JavaScript runtime before scanning your brain? Neither should your code intelligence engine.'
      }
    ]
  },

  // ─────────────────────────────────────────────────
  // SECTION 7: CLI Commands Reference
  // ─────────────────────────────────────────────────
  {
    id: 'commands',
    title: '🛠️ CLI Commands Reference',
    category: 'Reference',
    summary: 'Complete command instruction set with flags, options, and usage examples.',
    content: '',
    blocks: [
      {
        type: 'paragraph',
        text: 'While simply typing codemri automates the primary workflow, advanced operators can use granular subcommands for specific tasks.'
      },
      { type: 'heading', text: 'codemri' },
      {
        type: 'paragraph',
        text: 'The default command. Auto-detects project root, runs the scanner if no .codemri/ cache exists, starts the dashboard server, and opens your browser.'
      },
      {
        type: 'code',
        language: 'bash',
        code: '# Run in current directory\ncodemri\n\n# Run targeting a specific path\ncodemri /path/to/project'
      },
      { type: 'heading', text: 'codemri scan' },
      {
        type: 'paragraph',
        text: 'Forces a fresh deterministic syntax parse and rewrites .codemri/graph.json with updated analytical metrics. Use this to refresh results after code changes.'
      },
      {
        type: 'code',
        language: 'bash',
        code: '# Scan current directory\ncodemri scan .\n\n# Scan a specific repository path\ncodemri scan /path/to/enterprise-monorepo'
      },
      { type: 'heading', text: 'codemri serve' },
      {
        type: 'paragraph',
        text: 'Starts the HTTP dashboard visualization server. Reads the cached .codemri/ workspace and serves an interactive UI. Port auto-increments if the default is in use.'
      },
      {
        type: 'code',
        language: 'bash',
        code: '# Start on default port 4000\ncodemri serve\n\n# Start on custom port\ncodemri serve --port 5050'
      },
      { type: 'heading', text: 'codemri doctor' },
      {
        type: 'paragraph',
        text: 'Diagnoses your system environment — checks Go runtime availability, Git binary path, filesystem permissions, and platform capabilities.'
      },
      {
        type: 'code',
        language: 'bash',
        code: 'codemri doctor'
      },
      { type: 'heading', text: 'codemri version' },
      {
        type: 'paragraph',
        text: 'Displays the active compilation version, engine codename, and runtime target information.'
      },
      {
        type: 'code',
        language: 'bash',
        code: 'codemri version'
      },
      { type: 'heading', text: 'Global Flags' },
      {
        type: 'code',
        language: 'text',
        code: 'Flags:\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n  -v, --verbose     Enable detailed verbose debugging logs\n  -p, --port        Port for dashboard server (default: 4000)\n  -h, --help        Display help information for any command'
      }
    ]
  },

  // ─────────────────────────────────────────────────
  // SECTION 8: Project Structure
  // ─────────────────────────────────────────────────
  {
    id: 'project-structure',
    title: '📁 Project Structure',
    category: 'Reference',
    summary: 'Understanding the CodeMRI monorepo layout — CLI engine, website, documentation, and configuration.',
    content: '',
    blocks: [
      {
        type: 'paragraph',
        text: 'CodeMRI follows a clean monorepo architecture separating the Go CLI engine from the React landing page and documentation assets.'
      },
      {
        type: 'code',
        language: 'text',
        code: 'MRI-Code/\n├── cmd/\n│   └── codemri/\n│       └── main.go              ← CLI entry point\n│\n├── internal/\n│   ├── cli/\n│   │   ├── root.go              ← One-command workflow logic\n│   │   ├── scan.go              ← Atlas & Neuron scan orchestration\n│   │   ├── serve.go             ← Dashboard HTTP server & SQLite query APIs\n│   │   ├── doctor.go            ← System diagnostics & DB checking\n│   │   ├── dashboard_html.go    ← Reactive embedded glassmorphism UI\n│   │   └── version.go           ← v0.3.0 Neuron metadata\n│   │\n│   ├── parser/\n│   │   ├── parser.go            ← Parser interface & FileAST type\n│   │   ├── language.go          ← 15+ language detection engine\n│   │   ├── go_parser.go         ← Native Go AST extraction\n│   │   ├── universal_parser.go  ← Multi-language regex tokenizer\n│   │   └── registry.go          ← Language → Parser dispatcher\n│   │\n│   ├── scanner/\n│   │   ├── scanner.go           ← Scanner interface & metrics\n│   │   └── engine.go            ← Atlas concurrent traversal\n│   │\n│   └── graph/\n│       ├── nrg.go               ← NRG types, Node, Edge, Storage interface\n│       └── sqlite_storage.go    ← CGO-free embedded SQLite Engine\n│\n├── website/                     ← React + Vite landing page & official documentation\n├── docs/\n│   ├── philosophy.md            ← Core principles & beliefs\n│   ├── phases/                  ← Phase specifications (Genesis, Atlas, Neuron)\n│   └── ADR/                     ← Architecture Decision Records'
      },
      {
        type: 'callout',
        variant: 'info',
        text: 'When users run codemri scan on their own projects, the .codemri/ folder is created inside THEIR project directory — not inside the CodeMRI source code. This is a fundamental design principle.'
      }
    ]
  }
];
