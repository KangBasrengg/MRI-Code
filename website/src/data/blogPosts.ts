import type { BlogPost } from '../types';

export const BLOG_POSTS: BlogPost[] = [
  {
    id: 'how-react-works',
    title: 'How React Fiber Really Works Behind the Scenes (Scanned by CodeMRI)',
    date: 'July 26, 2026',
    readTime: '6 min read',
    tag: 'Architecture Study',
    author: {
      name: 'Muhammad Nuril',
      handle: '@KangBasrengg',
      avatar: 'https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?auto=format&fit=crop&q=80&w=150'
    },
    excerpt: 'Instead of guessing how React handles concurrent rendering, we ran codemri scan on facebook/react to map out the exact reconciler dependency graph.',
    content: `When engineers first explore the official **facebook/react** monorepo, they are frequently overwhelmed by over 4,000 internal source files spanning custom schedulers, experimental compile targets, and deeply decoupled reconciler packages. 

By executing \`codemri scan .\` directly inside the repository root, our Tree-sitter AST pipeline mapped out the entire dependency graph in just **18 seconds**!

### Key Architectural Discovery: Reconciler Decoupling
The scan reveals that \`react-core\` contains practically zero environment-specific rendering instructions. Instead, it delegates abstract component tree mutations directly to \`react-reconciler\` (the famous Fiber engine). Whether you render to standard web DOM elements via \`react-dom\` or native mobile views via \`react-native-renderer\`, both engines simply subscribe to Reconciler work loops!`,
    nodes: [
      { id: 'r-core', data: { label: '⚛️ react-core' }, position: { x: 220, y: 30 }, style: { background: '#0f1623', border: '1px solid #3b82f6', color: '#60a5fa', fontWeight: 'bold' } },
      { id: 'r-rec', data: { label: '🧠 react-reconciler (Fiber Loop)' }, position: { x: 220, y: 150 }, style: { background: '#0f1623', border: '1px solid #00f2ff', color: '#00f2ff', padding: '10px 18px', fontWeight: 'bold' } },
      { id: 'r-dom', data: { label: '🌐 react-dom/client' }, position: { x: 70, y: 280 }, style: { background: '#0f1623', border: '1px solid #10b981', color: '#34d399' } },
      { id: 'r-native', data: { label: '📱 react-native-renderer' }, position: { x: 370, y: 280 }, style: { background: '#0f1623', border: '1px solid #f59e0b', color: '#fbbf24' } }
    ],
    edges: [
      { id: 'er1', source: 'r-core', target: 'r-rec', animated: true, style: { stroke: '#3b82f6', strokeWidth: 2 }, label: 'dispatches updates' },
      { id: 'er2', source: 'r-rec', target: 'r-dom', style: { stroke: '#10b981' } },
      { id: 'er3', source: 'r-rec', target: 'r-native', style: { stroke: '#f59e0b' } }
    ]
  },
  {
    id: 'why-ai-wrappers-fail',
    title: 'Why Traditional AI Wrappers Fail at System Architecture',
    date: 'July 24, 2026',
    readTime: '8 min read',
    tag: 'Neural Repository Graph',
    author: {
      name: 'Muhammad Nuril',
      handle: '@KangBasrengg',
      avatar: 'https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?auto=format&fit=crop&q=80&w=150'
    },
    excerpt: 'Stuffing 5,000 files into LLM context windows is expensive, slow, and prone to hallucinations. Discover why deterministic parsers and Neural Repository Graphs (NRG) are the undeniable future.',
    content: `We have all experienced the limitation: you paste an enterprise repository path into a typical AI coding chatbot and ask: *"How does our billing verification flow interact with PostgreSQL and Redis?"* The assistant spends 45 seconds crunching tokens only to guess incorrect method names or overlook hidden circular dependency chains.

### The Problem: AI Does Not Read Code Like a Compiler
LLMs act as probabilistic token predictors. They do not naturally build Abstract Syntax Trees (AST) or evaluate symbol scopes across distinct packaging directories. 

### The CodeMRI Solution: "Parser Understands Code. AI Understands Meaning."
In CodeMRI, we forbid raw text stuffing. Our Go scanner combines concurrent filesystem walks with **Tree-sitter** parsers to generate a verified **Neural Repository Graph (NRG)**. When the AI reasoning engine is invoked, it queries structured nodes and relational edges—ensuring accurate answers with near-zero latency and fraction-of-a-cent token consumption.`,
    nodes: [
      { id: 'ai1', data: { label: '📁 Raw Source Repository (10,000 files)' }, position: { x: 220, y: 30 }, style: { background: '#1c1018', border: '1px solid #ef4444', color: '#fca5a5', fontWeight: 'bold' } },
      { id: 'ai2', data: { label: '⚡ Tree-sitter Deterministic Parser' }, position: { x: 220, y: 150 }, style: { background: '#0f1623', border: '1px solid #00f2ff', color: '#00f2ff', fontWeight: 'bold' } },
      { id: 'ai3', data: { label: '🧠 Neural Repository Graph (.codemri/graph.db)' }, position: { x: 220, y: 270 }, style: { background: '#0f1623', border: '1px solid #8a2be2', color: '#c084fc', padding: '10px 22px', fontWeight: 'bold' } },
      { id: 'ai4', data: { label: '🤖 AI Reasoning Chat & Explanations' }, position: { x: 220, y: 390 }, style: { background: '#0f1623', border: '1px solid #10b981', color: '#34d399', fontWeight: 'bold' } }
    ],
    edges: [
      { id: 'eai1-2', source: 'ai1', target: 'ai2', animated: true, style: { stroke: '#ef4444', strokeWidth: 2 }, label: '< 60 seconds' },
      { id: 'eai2-3', source: 'ai2', target: 'ai3', animated: true, style: { stroke: '#00f2ff', strokeWidth: 2 }, label: 'builds AST nodes & edges' },
      { id: 'eai3-4', source: 'ai3', target: 'ai4', animated: true, style: { stroke: '#8a2be2', strokeWidth: 2 }, label: 'instant factual query' }
    ]
  }
];
