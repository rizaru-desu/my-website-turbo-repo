/* Pixel avatar — bigger version for hero */
export const HeroAvatar = () => (
  <svg
    viewBox="0 0 16 16"
    width="160"
    height="160"
    shapeRendering="crispEdges"
    style={{ imageRendering: "pixelated" }}
    aria-hidden="true"
  >
    {/* background */}
    <rect x="0" y="0" width="16" height="16" fill="#44475a" />
    {/* hair */}
    <rect x="4" y="1" width="8" height="2" fill="#bd93f9" />
    <rect x="3" y="2" width="10" height="1" fill="#bd93f9" />
    {/* face */}
    <rect x="4" y="2" width="8" height="6" fill="#ffb86c" />
    <rect x="3" y="3" width="10" height="5" fill="#ffb86c" />
    {/* eyes */}
    <rect x="5" y="5" width="1" height="1" fill="#282a36" />
    <rect x="10" y="5" width="1" height="1" fill="#282a36" />
    {/* mouth */}
    <rect x="6" y="7" width="4" height="1" fill="#282a36" />
    <rect x="7" y="6" width="2" height="1" fill="#ff79c6" />
    {/* body / jacket */}
    <rect x="2" y="9" width="12" height="6" fill="#50fa7b" />
    <rect x="2" y="9" width="12" height="1" fill="#50fa7b" />
    <rect x="4" y="11" width="8" height="1" fill="#44475a" />
    <rect x="7" y="12" width="2" height="3" fill="#44475a" />
    {/* neck */}
    <rect x="6" y="8" width="4" height="1" fill="#d99857" />
    {/* shadow */}
    <rect x="0" y="15" width="16" height="1" fill="#282a36" />
  </svg>
);

export const IconGrid = () => (
  <svg
    viewBox="0 0 16 16"
    width="18"
    height="18"
    shapeRendering="crispEdges"
    aria-hidden="true"
  >
    <rect x="2" y="2" width="5" height="5" fill="#bd93f9" />
    <rect x="9" y="2" width="5" height="5" fill="#8be9fd" />
    <rect x="2" y="9" width="5" height="5" fill="#ff79c6" />
    <rect x="9" y="9" width="5" height="5" fill="#50fa7b" />
  </svg>
);

export const IconMail = () => (
  <svg
    viewBox="0 0 16 16"
    width="16"
    height="16"
    shapeRendering="crispEdges"
    aria-hidden="true"
  >
    <rect x="1" y="3" width="14" height="10" fill="#ff79c6" />
    <rect x="2" y="4" width="12" height="1" fill="#282a36" />
    <rect x="3" y="5" width="10" height="1" fill="#282a36" />
  </svg>
);

export const IconExternal = () => (
  <svg
    viewBox="0 0 16 16"
    width="12"
    height="12"
    shapeRendering="crispEdges"
    aria-hidden="true"
  >
    <rect x="8" y="2" width="6" height="2" fill="#282a36" />
    <rect x="12" y="2" width="2" height="6" fill="#282a36" />
    <rect x="7" y="7" width="1" height="1" fill="#282a36" />
    <rect x="8" y="6" width="1" height="1" fill="#282a36" />
    <rect x="9" y="5" width="1" height="1" fill="#282a36" />
    <rect x="10" y="4" width="1" height="1" fill="#282a36" />
    <rect x="2" y="5" width="6" height="9" fill="none" stroke="#282a36" />
  </svg>
);

export const IconFolder = () => (
  <svg
    viewBox="0 0 16 16"
    width="24"
    height="24"
    shapeRendering="crispEdges"
    aria-hidden="true"
  >
    <rect x="1" y="4" width="6" height="2" fill="#f1fa8c" />
    <rect x="1" y="5" width="14" height="9" fill="#bd93f9" />
  </svg>
);

export const IconStar = () => (
  <svg
    viewBox="0 0 16 16"
    width="24"
    height="24"
    shapeRendering="crispEdges"
    aria-hidden="true"
  >
    <rect x="7" y="1" width="2" height="2" fill="#8be9fd" />
    <rect x="1" y="5" width="14" height="2" fill="#8be9fd" />
    <rect x="3" y="7" width="10" height="2" fill="#8be9fd" />
    <rect x="5" y="9" width="6" height="2" fill="#8be9fd" />
  </svg>
);

export const IconDoc = () => (
  <svg
    viewBox="0 0 16 16"
    width="24"
    height="24"
    shapeRendering="crispEdges"
    aria-hidden="true"
  >
    <rect x="2" y="2" width="12" height="12" fill="#ff79c6" />
    <rect x="4" y="5" width="8" height="1" fill="#282a36" />
    <rect x="4" y="7" width="8" height="1" fill="#282a36" />
    <rect x="4" y="9" width="6" height="1" fill="#282a36" />
  </svg>
);

/* Small pixel sword divider */
export const Sword = () => (
  <svg
    viewBox="0 0 32 8"
    width="64"
    height="16"
    shapeRendering="crispEdges"
    style={{ imageRendering: "pixelated" }}
    aria-hidden="true"
  >
    <rect x="0" y="3" width="2" height="2" fill="#ff79c6" />
    <rect x="2" y="2" width="2" height="4" fill="#ff79c6" />
    <rect x="4" y="3" width="2" height="2" fill="#bdc0cc" />
    <rect x="6" y="3" width="18" height="2" fill="#f8f8f2" />
    <rect x="7" y="2" width="16" height="1" fill="#bdc0cc" />
    <rect x="7" y="5" width="16" height="1" fill="#6272a4" />
    <rect x="24" y="2" width="2" height="4" fill="#f1fa8c" />
    <rect x="26" y="1" width="2" height="6" fill="#f1fa8c" />
    <rect x="28" y="3" width="2" height="2" fill="#f1fa8c" />
  </svg>
);

export const NavBrandLogo = () => (
  <svg
    viewBox="0 0 16 16"
    width="28"
    height="28"
    shapeRendering="crispEdges"
    aria-hidden="true"
  >
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
