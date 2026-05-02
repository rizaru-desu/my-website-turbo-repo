import type { ReactNode } from "react";

type PageKey =
  | "dashboard"
  | "projects"
  | "posts"
  | "skills"
  | "messages"
  | "analytics"
  | "about"
  | "settings";

type SidebarProps = {
  currentPage: PageKey;
  onNavigate: (page: PageKey) => void;
  navItems: Array<{ key: PageKey; label: string; icon: ReactNode }>;
};

const PixelSword = () => (
  <svg viewBox="0 0 16 16" width="28" height="28" shapeRendering="crispEdges">
    <rect x="11" y="2" width="3" height="3" fill="#bd93f9" />
    <rect x="10" y="3" width="3" height="3" fill="#bd93f9" />
    <rect x="9" y="4" width="3" height="3" fill="#bd93f9" />
    <rect x="8" y="5" width="3" height="3" fill="#bd93f9" />
    <rect x="7" y="6" width="3" height="3" fill="#bd93f9" />
    <rect x="6" y="7" width="3" height="3" fill="#bd93f9" />
    <rect x="5" y="8" width="3" height="3" fill="#bd93f9" />
    <rect x="4" y="9" width="3" height="3" fill="#bd93f9" />
    <rect x="2" y="10" width="4" height="3" fill="#ff79c6" />
    <rect x="3" y="13" width="2" height="2" fill="#8be9fd" />
  </svg>
);

export default function Sidebar({
  currentPage,
  onNavigate,
  navItems,
}: SidebarProps) {
  return (
    <aside className="sidebar" data-testid="sidebar">
      <div className="sidebar-logo">
        <PixelSword />
        <div>
          <div className="title">PIXEL.CMS</div>
          <div className="subtitle">v4.20 ~ retro edition</div>
        </div>
      </div>

      <div className="sidebar-section-label">~ MAIN</div>
      {navItems.slice(0, 4).map((item) => (
        <button
          key={item.key}
          className={`nav-item ${currentPage === item.key ? "active" : ""}`}
          onClick={() => onNavigate(item.key)}
          data-testid={`nav-${item.key}`}
        >
          {item.icon}
          <span>{item.label}</span>
        </button>
      ))}

      <div className="sidebar-section-label">~ ENGAGE</div>
      {navItems.slice(4).map((item) => (
        <button
          key={item.key}
          className={`nav-item ${currentPage === item.key ? "active" : ""}`}
          onClick={() => onNavigate(item.key)}
          data-testid={`nav-${item.key}`}
        >
          {item.icon}
          <span>{item.label}</span>
        </button>
      ))}

      <div className="xp-card" data-testid="xp-card">
        <div className="xp-name">KAI / LVL 42</div>
        <div className="xp-copy">XP TO LVL 43</div>
        <div className="xp-bar">
          <div className="xp-bar-fill" />
        </div>
        <div className="xp-copy">6,820 / 10,000</div>
      </div>
    </aside>
  );
}
