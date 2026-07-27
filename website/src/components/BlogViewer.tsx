import React, { useState } from 'react';
import { BLOG_POSTS } from '../data/blogPosts';
import type { BlogPost } from '../types';
import { ReactFlow, Background, Controls } from '@xyflow/react';
import { Calendar, Clock, Layers } from 'lucide-react';

export const BlogViewer: React.FC = () => {
  const [selectedPost, setSelectedPost] = useState<BlogPost>(BLOG_POSTS[0]);

  return (
    <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 py-12 text-left font-sans">
      
      {/* Blog Header */}
      <div className="text-center mb-12">
        <span className="text-xs font-mono uppercase px-3 py-1 bg-purple-500/10 border border-purple-500/30 text-purple-300 rounded-full font-bold">
          📰 Engineering & Architecture Blog
        </span>
        <h1 className="text-3xl sm:text-5xl font-black text-white mt-3 mb-4">
          Deep Technical Case Studies
        </h1>
        <p className="text-slate-400 max-w-2xl mx-auto text-sm sm:text-base">
          We publish comprehensive breakdowns of world-class system architectures using diagram visualizations generated directly by CodeMRI!
        </p>
      </div>

      {/* Post Selector Tabs */}
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-10">
        {BLOG_POSTS.map((post) => {
          const isSelected = selectedPost.id === post.id;
          return (
            <div
              key={post.id}
              onClick={() => setSelectedPost(post)}
              className={`p-6 rounded-3xl border cursor-pointer transition-all ${
                isSelected
                  ? 'bg-slate-900 border-cyan-500 shadow-xl shadow-cyan-500/10'
                  : 'bg-[#0b0f19] border-slate-800 hover:border-slate-700 hover:bg-slate-900/60'
              }`}
            >
              <div className="flex items-center space-x-3 text-xs font-mono text-slate-400 mb-3">
                <span className="px-2.5 py-0.5 bg-cyan-500/10 text-cyan-400 rounded-full border border-cyan-500/20 font-bold">
                  {post.tag}
                </span>
                <span className="flex items-center space-x-1">
                  <Calendar className="w-3.5 h-3.5" />
                  <span>{post.date}</span>
                </span>
                <span className="flex items-center space-x-1">
                  <Clock className="w-3.5 h-3.5" />
                  <span>{post.readTime}</span>
                </span>
              </div>

              <h3 className="text-lg sm:text-xl font-bold text-white mb-2 leading-snug">
                {post.title}
              </h3>
              <p className="text-sm text-slate-400 line-clamp-2 leading-relaxed">
                {post.excerpt}
              </p>
            </div>
          );
        })}
      </div>

      {/* Full Article Reader + Embedded React Flow Diagram */}
      <div className="glass-panel border-2 border-slate-800 rounded-3xl p-6 sm:p-10 shadow-2xl">
        
        {/* Article Meta */}
        <div className="border-b border-slate-800 pb-6 mb-8 flex flex-col sm:flex-row sm:items-center justify-between gap-4">
          <div>
            <span className="text-xs uppercase font-mono px-3 py-1 bg-amber-500/10 text-amber-300 border border-amber-500/30 rounded-full font-bold">
              {selectedPost.tag}
            </span>
            <h2 className="text-2xl sm:text-4xl font-black text-white mt-3 leading-tight">
              {selectedPost.title}
            </h2>
          </div>

          {/* Author Badge */}
          <div className="flex items-center space-x-3 bg-slate-900 p-3 rounded-2xl border border-slate-800 flex-shrink-0">
            <img src={selectedPost.author.avatar} alt="Author avatar" className="w-10 h-10 rounded-xl object-cover border border-slate-700" />
            <div>
              <p className="text-xs font-bold text-white">{selectedPost.author.name}</p>
              <p className="text-[11px] font-mono text-cyan-400">{selectedPost.author.handle}</p>
            </div>
          </div>
        </div>

        {/* Text body */}
        <div className="prose prose-invert max-w-none text-slate-300 space-y-6 text-sm sm:text-base leading-relaxed whitespace-pre-line mb-10">
          {selectedPost.content}
        </div>

        {/* Embedded Neural Graph Visualizer */}
        <div className="mt-8 pt-8 border-t border-slate-800">
          <div className="mb-4 flex items-center justify-between">
            <h4 className="text-base font-bold text-white font-mono flex items-center space-x-2">
              <Layers className="w-5 h-5 text-cyan-400" />
              <span>Embedded CodeMRI Architecture Scan Diagram:</span>
            </h4>
            <span className="text-xs font-mono bg-slate-900 text-emerald-400 px-3 py-1 rounded-lg border border-slate-800 font-bold">
              Interactive Nodes
            </span>
          </div>

          <div className="w-full h-[380px] bg-[#0b0f19] rounded-2xl border border-slate-800 overflow-hidden shadow-inner">
            <ReactFlow
              nodes={selectedPost.nodes}
              edges={selectedPost.edges}
              fitView
              key={selectedPost.id}
              proOptions={{ hideAttribution: true }}
            >
              <Background color="#1e293b" gap={20} size={1} />
              <Controls className="bg-slate-900 border border-slate-700 fill-white rounded-xl" />
            </ReactFlow>
          </div>
        </div>

      </div>

    </div>
  );
};
