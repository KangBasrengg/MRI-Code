import React, { useState } from 'react';
import type { ActiveTab } from '../types';
import { BookOpen, Star, Sparkles, Map, Newspaper, Terminal, Menu, X } from 'lucide-react';
import codemriLogo from '../assets/codemri.png';

interface NavbarProps {
  activeTab: ActiveTab;
  setActiveTab: (tab: ActiveTab) => void;
  onInstallClick: () => void;
}

export const Navbar: React.FC<NavbarProps> = ({ activeTab, setActiveTab, onInstallClick }) => {
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false);

  const navItems: { id: ActiveTab; label: string; icon: React.ReactNode; badge?: string }[] = [
    { id: 'home', label: 'Get Started & Guide', icon: <BookOpen className="w-4 h-4 text-cyan-400" /> },
    { id: 'hall-of-fame', label: 'Hall of Fame', icon: <Sparkles className="w-4 h-4 text-amber-400" />, badge: 'Viral' },
    { id: 'roadmap', label: 'Roadmap', icon: <Map className="w-4 h-4 text-emerald-400" /> },
    { id: 'blog', label: 'Blog', icon: <Newspaper className="w-4 h-4 text-purple-400" /> },
  ];

  return (
    <header className="sticky top-0 z-50 bg-black/60 backdrop-blur-xl border-b border-white/10 shadow-2xl transition-all">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 h-16 flex items-center justify-between gap-4">

        {/* Brand Logo & Name */}
        <div
          onClick={() => setActiveTab('home')}
          className="flex items-center space-x-3 cursor-pointer group flex-shrink-0"
        >
          <div className="w-9 h-9 rounded-xl bg-gradient-to-br from-cyan-500/20 via-blue-600/30 to-purple-600/30 border border-cyan-500/40 flex items-center justify-center shadow-md group-hover:border-cyan-400 transition-all p-1">
            <img src={codemriLogo} alt="CodeMRI Logo" className="w-full h-full object-contain filter drop-shadow" />
          </div>
          <div className="flex items-center space-x-2.5">
            <span className="font-sans font-black text-lg sm:text-xl tracking-tight text-white whitespace-nowrap">
              Code<span className="text-cyan-400"></span>MRI
            </span>
            <span className="text-[10px] uppercase font-mono px-2 py-0.5 bg-cyan-500/20 border border-cyan-500/40 text-cyan-300 font-bold rounded-full whitespace-nowrap hidden sm:inline-block shadow-sm">
              v1.0.0
            </span>
          </div>
        </div>

        {/* Desktop Navigation (Sleek, Horizontal, No Multiline Wrapping!) */}
        <nav className="hidden md:flex items-center space-x-1 lg:space-x-2">
          {navItems.map((item) => {
            const isActive = activeTab === item.id;
            return (
              <button
                key={item.id}
                onClick={() => setActiveTab(item.id)}
                className={`relative flex items-center space-x-2 px-3 lg:px-4 py-2 rounded-xl text-xs sm:text-sm font-medium whitespace-nowrap transition-all duration-200 ${isActive
                  ? 'bg-white/15 text-white font-black shadow-md border border-white/20'
                  : 'text-slate-300 hover:text-white hover:bg-white/10'
                  }`}
              >
                {item.icon}
                <span>{item.label}</span>
                {item.badge && (
                  <span className="text-[9px] font-mono font-bold bg-amber-500/20 text-amber-300 border border-amber-500/40 px-1.5 py-0.2 rounded-full">
                    {item.badge}
                  </span>
                )}
              </button>
            );
          })}
        </nav>

        {/* Right Actions: Install CLI & Star on GitHub */}
        <div className="hidden md:flex items-center space-x-3 flex-shrink-0">
          <button
            onClick={onInstallClick}
            className="flex items-center space-x-2 px-3.5 py-2 bg-slate-900/80 hover:bg-slate-800 text-cyan-300 font-mono text-xs rounded-xl border border-cyan-500/40 hover:border-cyan-400 transition-all whitespace-nowrap shadow-md"
          >
            <Terminal className="w-3.5 h-3.5 text-cyan-400" />
            <span>Install CLI</span>
          </button>

          <a
            href="https://github.com/KangBasrengg/MRI-Code"
            target="_blank"
            rel="noopener noreferrer"
            className="flex items-center space-x-2 px-4 py-2 rounded-xl bg-gradient-to-r from-cyan-400 to-blue-500 hover:from-cyan-300 hover:to-blue-400 text-slate-950 font-black text-xs sm:text-sm whitespace-nowrap shadow-lg shadow-cyan-500/30 hover:scale-[1.03] transition-all"
          >
            <Star className="w-4 h-4 fill-current text-slate-950" />
            <span>Star on GitHub</span>
          </a>
        </div>

        {/* Mobile Menu Toggle Button */}
        <div className="flex md:hidden items-center space-x-2">
          <a
            href="https://github.com/KangBasrengg/MRI-Code"
            target="_blank"
            rel="noopener noreferrer"
            className="p-2 rounded-lg bg-slate-800 text-amber-300 border border-slate-700"
            title="GitHub Repository"
          >
            <Star className="w-4 h-4 fill-current" />
          </a>
          <button
            onClick={() => setMobileMenuOpen(!mobileMenuOpen)}
            className="p-2 rounded-lg bg-slate-900 border border-slate-800 text-slate-300 hover:text-white"
          >
            {mobileMenuOpen ? <X className="w-6 h-6" /> : <Menu className="w-6 h-6" />}
          </button>
        </div>

      </div>

      {/* Responsive Mobile Menu Dropdown */}
      {mobileMenuOpen && (
        <div className="md:hidden bg-black/95 backdrop-blur-2xl border-b border-slate-800 px-4 py-4 space-y-2">
          {navItems.map((item) => (
            <button
              key={item.id}
              onClick={() => {
                setActiveTab(item.id);
                setMobileMenuOpen(false);
              }}
              className={`w-full flex items-center space-x-3 px-4 py-3 rounded-xl text-sm font-medium ${activeTab === item.id ? 'bg-white/15 text-white font-bold' : 'text-slate-400 hover:text-white hover:bg-slate-900'
                }`}
            >
              {item.icon}
              <span>{item.label}</span>
            </button>
          ))}
          <div className="pt-3 border-t border-slate-800/80 flex flex-col gap-2">
            <button
              onClick={() => {
                onInstallClick();
                setMobileMenuOpen(false);
              }}
              className="w-full py-3 bg-slate-900 hover:bg-slate-800 text-cyan-300 font-mono text-sm rounded-xl border border-cyan-500/40 flex items-center justify-center space-x-2"
            >
              <Terminal className="w-4 h-4 text-cyan-400" />
              <span>Install CLI Guide</span>
            </button>
          </div>
        </div>
      )}
    </header>
  );
};
