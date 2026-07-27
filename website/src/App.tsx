import { useState } from 'react';
import { Navbar } from './components/Navbar';
import { Hero } from './components/Hero';
import { CliInstallGuide } from './components/CliInstallGuide';
import { TerminalDemo } from './components/TerminalDemo';
import { HallOfFame } from './components/HallOfFame';
import { DocsViewer } from './components/DocsViewer';
import { BlogViewer } from './components/BlogViewer';
import { Roadmap } from './components/Roadmap';
import { Footer } from './components/Footer';
import type { ActiveTab } from './types';

export function App() {
  const [activeTab, setActiveTab] = useState<ActiveTab>('home');

  const handleOpenHallOfFame = () => {
    setActiveTab('hall-of-fame');
    window.scrollTo({ top: 0, behavior: 'smooth' });
  };

  const handleOpenDocs = () => {
    setActiveTab('docs');
    window.scrollTo({ top: 0, behavior: 'smooth' });
  };

  const handleScrollToCliInstall = () => {
    if (activeTab !== 'home') {
      setActiveTab('home');
      setTimeout(() => {
        const el = document.getElementById('cli-install-guide');
        if (el) el.scrollIntoView({ behavior: 'smooth' });
      }, 100);
    } else {
      const el = document.getElementById('cli-install-guide');
      if (el) {
        el.scrollIntoView({ behavior: 'smooth' });
      } else {
        window.scrollTo({ top: 500, behavior: 'smooth' });
      }
    }
  };

  return (
    <div className="min-h-screen text-slate-100 flex flex-col font-sans selection:bg-cyan-500/30 selection:text-white bg-transparent">
      
      {/* Top Navigation Bar */}
      <Navbar
        activeTab={activeTab}
        setActiveTab={(tab) => {
          setActiveTab(tab);
          window.scrollTo({ top: 0, behavior: 'smooth' });
        }}
        onInstallClick={handleScrollToCliInstall}
      />

      {/* Dynamic Content Views */}
      <main className="flex-1 w-full bg-transparent">
        {activeTab === 'home' && (
          <div className="space-y-6">
            {/* 1. Impactful Hero Section */}
            <Hero
              onInstallClick={handleScrollToCliInstall}
              onOpenHallOfFame={handleOpenHallOfFame}
            />
            
            {/* 2. Official CLI Installation & Quickstart Section */}
            <CliInstallGuide
              onOpenFullDocs={handleOpenDocs}
            />

            {/* 3. 15-Second Interactive Terminal Simulation */}
            <TerminalDemo />
          </div>
        )}

        {activeTab === 'hall-of-fame' && (
          <HallOfFame />
        )}

        {activeTab === 'docs' && (
          <DocsViewer />
        )}

        {activeTab === 'roadmap' && (
          <Roadmap />
        )}

        {activeTab === 'blog' && (
          <BlogViewer />
        )}
      </main>

      {/* Footer */}
      <Footer />

    </div>
  );
}

export default App;
