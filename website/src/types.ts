import type { Node, Edge } from '@xyflow/react';

export interface RepositoryData {
  id: string;
  name: string;
  stars: string;
  language: string;
  healthScore: number;
  filesCount: number;
  modulesCount: number;
  circularDeps: number;
  description: string;
  nodes: Node[];
  edges: Edge[];
  insights: {
    architecture: string;
    techDebt: string;
    security: string;
  };
}

export interface DocContentBlock {
  type: 'paragraph' | 'heading' | 'code' | 'list' | 'callout' | 'steps';
  text?: string;
  language?: string;
  code?: string;
  items?: string[] | { label: string; description: string }[];
  variant?: 'info' | 'warning' | 'success';
}

export interface DocSection {
  id: string;
  title: string;
  category: string;
  summary: string;
  content: string;
  codeExample?: string;
  codeLanguage?: string;
  blocks?: DocContentBlock[];
}

export interface BlogPost {
  id: string;
  title: string;
  date: string;
  readTime: string;
  tag: string;
  author: {
    name: string;
    handle: string;
    avatar: string;
  };
  excerpt: string;
  content: string;
  nodes: Node[];
  edges: Edge[];
}

export type ActiveTab = 'home' | 'docs' | 'hall-of-fame' | 'roadmap' | 'blog';
