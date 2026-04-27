import type { ReactNode } from "react";
import Sidebar from "./Sidebar";
import TopBar from "./TopBar";

type PageKey =
  | "dashboard"
  | "projects"
  | "posts"
  | "skills"
  | "messages"
  | "analytics"
  | "settings";

type AppLayoutProps = {
  children: ReactNode;
  currentPage: PageKey;
  onNavigate: (page: PageKey) => void;
  onLogout: () => void;
  navItems: Array<{ key: PageKey; label: string; icon: ReactNode }>;
  pageTitle?: {
    crumb: string;
    title: string;
    sub: string;
  };
  showPageHeader?: boolean;
};

export default function AppLayout({
  children,
  currentPage,
  onNavigate,
  onLogout,
  navItems,
  pageTitle,
  showPageHeader = true,
}: AppLayoutProps) {
  return (
    <div className="admin-shell" data-testid="app-shell">
      <Sidebar
        currentPage={currentPage}
        onNavigate={onNavigate}
        navItems={navItems}
      />
      <TopBar onLogout={onLogout} />
      <main className="main" data-testid="main-content">
        {showPageHeader && pageTitle && (
          <div className="page-header">
            <div className="page-title">
              <div className="crumbs">
                &gt; <b>{pageTitle.crumb}</b>
                <span className="blink" />
              </div>
              <h1 data-testid="page-title">{pageTitle.title}</h1>
              <div className="sub">{pageTitle.sub}</div>
            </div>
            <div className="row-actions">
              <button className="pix-btn pix-btn-ghost">EXPORT</button>
              <button className="pix-btn">+ QUICK ACTION</button>
            </div>
          </div>
        )}
        {children}
        <div className="cms-footer">
          PRESS [ESC] TO RETURN ~ PIXEL.CMS v4.20 / MADE WITH &lt;3 &amp; PIXELS
        </div>
      </main>
    </div>
  );
}
