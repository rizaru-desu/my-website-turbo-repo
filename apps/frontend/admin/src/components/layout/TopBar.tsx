import type { ReactNode } from "react";
import { useAuth } from "../../hooks/useAuth";

type TopBarProps = {
  onLogout: () => void;
};

const Icon = ({ children }: { children: ReactNode }) => (
  <svg
    width="18"
    height="18"
    viewBox="0 0 16 16"
    shapeRendering="crispEdges"
    aria-hidden="true"
  >
    {children}
  </svg>
);

const IconBell = () => (
  <Icon>
    <rect x="6" y="2" width="4" height="1" fill="#ffb86c" />
    <rect x="4" y="3" width="8" height="1" fill="#ffb86c" />
    <rect x="3" y="4" width="10" height="7" fill="#ffb86c" />
    <rect x="2" y="11" width="12" height="1" fill="#ffb86c" />
    <rect x="6" y="12" width="4" height="2" fill="#ffb86c" />
  </Icon>
);

const PixelAvatar = () => (
  <svg viewBox="0 0 16 16" width="32" height="32" shapeRendering="crispEdges">
    <rect x="4" y="2" width="8" height="6" fill="#ffb86c" />
    <rect x="3" y="3" width="10" height="5" fill="#ffb86c" />
    <rect x="5" y="5" width="1" height="1" fill="#282a36" />
    <rect x="10" y="5" width="1" height="1" fill="#282a36" />
    <rect x="6" y="7" width="4" height="1" fill="#282a36" />
    <rect x="4" y="1" width="8" height="2" fill="#bd93f9" />
    <rect x="2" y="9" width="12" height="6" fill="#50fa7b" />
    <rect x="4" y="11" width="8" height="1" fill="#44475a" />
  </svg>
);

const formatAbbreviatedName = (fullName: string): string => {
  const parts = fullName.trim().split(/\s+/);
  if (parts.length === 1) {
    return parts[0].toUpperCase();
  }
  const firstName = parts[0];
  const lastInitial = parts[parts.length - 1].charAt(0);
  return `${firstName.toUpperCase()} ${lastInitial.toUpperCase()}.`;
};

export default function TopBar({ onLogout }: TopBarProps) {
  const { user } = useAuth();

  const displayName = user?.name ? formatAbbreviatedName(user.name) : "GUEST";
  const displayRole = user?.role ? user.role.toUpperCase() : "USER";

  return (
    <header className="topbar" data-testid="topbar">
      <div className="topbar-actions">
        <span className="pixel-badge live" data-testid="status-online">
          ONLINE
        </span>
        <button className="pix-btn pix-btn-ghost bell-btn">
          <IconBell />
          <span className="bell-count">3</span>
        </button>
        <div className="avatar-block" data-testid="user-block">
          <PixelAvatar />
          <div>
            <div className="name">{displayName}</div>
            <div className="role">{displayRole}</div>
          </div>
        </div>
        <button
          className="pix-btn pix-btn-ghost"
          onClick={onLogout}
          data-testid="logout-btn"
        >
          &lt;- LOGOUT
        </button>
      </div>
    </header>
  );
}
