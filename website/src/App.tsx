import { useState } from 'react';
import { Navbar } from './components/Navbar';
import { HallOfFame } from './components/HallOfFame';
import { GetStartedDocs } from './components/GetStartedDocs';
import { BlogViewer } from './components/BlogViewer';
import { Roadmap } from './components/Roadmap';
import { Footer } from './components/Footer';
import type { ActiveTab } from './types';

export function App() {
  const [activeTab, setActiveTab] = useState<ActiveTab>('home');

  const handleScrollToCliInstall = () => {
    setActiveTab('home');
    setTimeout(() => {
      const el = document.getElementById('installation');
      if (el) {
        el.scrollIntoView({ behavior: 'smooth' });
      } else {
        window.scrollTo({ top: 400, behavior: 'smooth' });
      }
    }, 100);
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
          <GetStartedDocs />
        )}

        {activeTab === 'hall-of-fame' && (
          <HallOfFame />
        )}

        {activeTab === 'docs' && (
          <GetStartedDocs />
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
