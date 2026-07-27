import type { RepositoryData } from '../types';

export const HALL_OF_FAME_DATA: RepositoryData[] = [
  {
    id: 'react',
    name: 'facebook/react',
    stars: '228k ⭐',
    language: 'JavaScript / C++',
    healthScore: 98,
    filesCount: 4120,
    modulesCount: 42,
    circularDeps: 0,
    description: 'The JavaScript library for user interfaces powered by standard Fiber concurent render engine.',
    insights: {
      architecture: 'Clean decoupling between Reconciler (Fiber core) and targeted Renderers (DOM, Native, Three).',
      techDebt: 'Minimal debt. Highly modularized packages under /packages/react-reconciler.',
      security: 'Zero detected XSS injection paths in standard jsx-runtime transpilation boundaries.'
    },
    nodes: [
      { id: '1', type: 'default', data: { label: '⚛️ react-core' }, position: { x: 250, y: 50 }, style: { background: '#0f1623', border: '1px solid #3b82f6', color: '#60a5fa', fontWeight: 'bold' } },
      { id: '2', type: 'default', data: { label: '🧠 react-reconciler (Fiber)' }, position: { x: 250, y: 180 }, style: { background: '#0f1623', border: '1px solid #8a2be2', color: '#c084fc', fontWeight: 'bold', padding: '10px 20px' } },
      { id: '3', type: 'default', data: { label: '🌐 react-dom' }, position: { x: 80, y: 320 }, style: { background: '#0f1623', border: '1px solid #10b981', color: '#34d399' } },
      { id: '4', type: 'default', data: { label: '📱 react-native-renderer' }, position: { x: 250, y: 320 }, style: { background: '#0f1623', border: '1px solid #10b981', color: '#34d399' } },
      { id: '5', type: 'default', data: { label: '🧪 scheduler' }, position: { x: 420, y: 320 }, style: { background: '#0f1623', border: '1px solid #f59e0b', color: '#fbbf24' } }
    ],
    edges: [
      { id: 'e1-2', source: '1', target: '2', animated: true, style: { stroke: '#3b82f6', strokeWidth: 2 }, label: 'delegates rendering' },
      { id: 'e2-3', source: '2', target: '3', style: { stroke: '#10b981' } },
      { id: 'e2-4', source: '2', target: '4', style: { stroke: '#10b981' } },
      { id: 'e2-5', source: '2', target: '5', animated: true, style: { stroke: '#f59e0b' }, label: 'prioritizes ticks' }
    ]
  },
  {
    id: 'laravel',
    name: 'laravel/framework',
    stars: '76k ⭐',
    language: 'PHP',
    healthScore: 95,
    filesCount: 6850,
    modulesCount: 64,
    circularDeps: 0,
    description: 'The PHP framework for web artisans with rich dependency injection and eloquent ORM relationships.',
    insights: {
      architecture: 'Service Provider pipeline pattern driving container abstraction & binding evaluation.',
      techDebt: 'Moderate magic method abstraction (Facade runtime resolution), but cleanly mapped in AST.',
      security: 'Strict PDO bindings prevent SQLi; built-in CSRF middleware verified across all public routing groups.'
    },
    nodes: [
      { id: 'l1', type: 'default', data: { label: '🚪 Public Index / Kernel' }, position: { x: 250, y: 40 }, style: { background: '#18121f', border: '1px solid #ef4444', color: '#fca5a5', fontWeight: 'bold' } },
      { id: 'l2', type: 'default', data: { label: '📦 Service Provider Pipeline' }, position: { x: 250, y: 160 }, style: { background: '#0f1623', border: '1px solid #00f2ff', color: '#00f2ff', padding: '8px 16px' } },
      { id: 'l3', type: 'default', data: { label: '🛣️ Routing & Middleware' }, position: { x: 100, y: 280 }, style: { background: '#0f1623', border: '1px solid #8b5cf6', color: '#d8b4fe' } },
      { id: 'l4', type: 'default', data: { label: '🗄️ Eloquent ORM Engine' }, position: { x: 400, y: 280 }, style: { background: '#0f1623', border: '1px solid #f59e0b', color: '#fde68a', fontWeight: 'bold' } },
      { id: 'l5', type: 'default', data: { label: '🎨 Blade View Compiler' }, position: { x: 250, y: 400 }, style: { background: '#0f1623', border: '1px solid #10b981', color: '#6ee7b7' } }
    ],
    edges: [
      { id: 'el1-2', source: 'l1', target: 'l2', animated: true, style: { stroke: '#ef4444', strokeWidth: 2 } },
      { id: 'el2-3', source: 'l2', target: 'l3', style: { stroke: '#8b5cf6' }, label: 'dispatches requests' },
      { id: 'el2-4', source: 'l2', target: 'l4', style: { stroke: '#f59e0b' }, label: 'binds db models' },
      { id: 'el3-5', source: 'l3', target: 'l5', animated: true, style: { stroke: '#10b981' } }
    ]
  },
  {
    id: 'nextjs',
    name: 'vercel/next.js',
    stars: '128k ⭐',
    language: 'Rust (Turbopack) / TypeScript',
    healthScore: 97,
    filesCount: 9400,
    modulesCount: 88,
    circularDeps: 0,
    description: 'The React Framework for the Web integrating Turbopack bundler and App Router RSC engine.',
    insights: {
      architecture: 'Hybrid compiler system branching Server React Server Components (RSC) vs Client Bundle Hydration.',
      techDebt: 'Rapid evolution from Pages to App Router cleanly isolated via internal modular routing bridges.',
      security: 'Server Action boundaries strictly checked; CORS & asset integrity verified by AST parser.'
    },
    nodes: [
      { id: 'n1', type: 'default', data: { label: '🦀 Turbopack Compiler Core (Rust)' }, position: { x: 250, y: 40 }, style: { background: '#1e111a', border: '1px solid #f97316', color: '#fdba74', fontWeight: 'bold', padding: '10px 18px' } },
      { id: 'n2', type: 'default', data: { label: '🚀 App Router & RSC Engine' }, position: { x: 250, y: 170 }, style: { background: '#0f1623', border: '1px solid #3b82f6', color: '#60a5fa', fontWeight: 'bold' } },
      { id: 'n3', type: 'default', data: { label: '🖥️ React Server Components' }, position: { x: 100, y: 290 }, style: { background: '#0f1623', border: '1px solid #8a2be2', color: '#d8b4fe' } },
      { id: 'n4', type: 'default', data: { label: '⚡ Client Bundles (Hydration)' }, position: { x: 400, y: 290 }, style: { background: '#0f1623', border: '1px solid #00f2ff', color: '#67e8f9' } },
      { id: 'n5', type: 'default', data: { label: '🛡️ Server Actions API Layer' }, position: { x: 250, y: 410 }, style: { background: '#0f1623', border: '1px solid #10b981', color: '#34d399' } }
    ],
    edges: [
      { id: 'en1-2', source: 'n1', target: 'n2', animated: true, style: { stroke: '#f97316', strokeWidth: 2 }, label: 'sub-second compilation' },
      { id: 'en2-3', source: 'n2', target: 'n3', style: { stroke: '#8a2be2' }, label: 'zero-bundle execution' },
      { id: 'en2-4', source: 'n2', target: 'n4', style: { stroke: '#00f2ff' } },
      { id: 'en3-5', source: 'n3', target: 'n5', animated: true, style: { stroke: '#10b981' }, label: 'secure rpc mutations' }
    ]
  },
  {
    id: 'hono',
    name: 'honojs/hono',
    stars: '21k ⭐',
    language: 'TypeScript',
    healthScore: 99,
    filesCount: 1450,
    modulesCount: 18,
    circularDeps: 0,
    description: 'Ultrafast web framework for the Edges (Cloudflare Workers, Deno, Bun, Node.js) with RegExpRouter.',
    insights: {
      architecture: 'RegExpRouter combines routing patterns into a single giant regex tree for lightning fast execution.',
      techDebt: 'Practically zero tech debt. Lightweight zero-dependency design with universal Web Standard Request/Response.',
      security: 'Strict TypeScript typing prevents invalid payload bindings; zero prototype pollution paths found.'
    },
    nodes: [
      { id: 'h1', type: 'default', data: { label: '🔥 Universal Hono Instance' }, position: { x: 250, y: 50 }, style: { background: '#1c100b', border: '1px solid #ea580c', color: '#fed7aa', fontWeight: 'bold' } },
      { id: 'h2', type: 'default', data: { label: '⚡ RegExpRouter (Trie Tree Engine)' }, position: { x: 250, y: 180 }, style: { background: '#0f1623', border: '1px solid #00f2ff', color: '#00f2ff', padding: '8px 20px', fontWeight: 'bold' } },
      { id: 'h3', type: 'default', data: { label: '☁️ Cloudflare Worker Edge' }, position: { x: 80, y: 310 }, style: { background: '#0f1623', border: '1px solid #f59e0b', color: '#fde68a' } },
      { id: 'h4', type: 'default', data: { label: '🥟 Bun & Deno Runtime' }, position: { x: 250, y: 310 }, style: { background: '#0f1623', border: '1px solid #10b981', color: '#6ee7b7' } },
      { id: 'h5', type: 'default', data: { label: '🐢 Node.js HTTP Adapter' }, position: { x: 420, y: 310 }, style: { background: '#0f1623', border: '1px solid #3b82f6', color: '#93c5fd' } }
    ],
    edges: [
      { id: 'eh1-2', source: 'h1', target: 'h2', animated: true, style: { stroke: '#ea580c', strokeWidth: 2 }, label: 'resolves in microseconds' },
      { id: 'eh2-3', source: 'h2', target: 'h3', style: { stroke: '#f59e0b' } },
      { id: 'eh2-4', source: 'h2', target: 'h4', style: { stroke: '#10b981' } },
      { id: 'eh2-5', source: 'h2', target: 'h5', style: { stroke: '#3b82f6' } }
    ]
  }
];
