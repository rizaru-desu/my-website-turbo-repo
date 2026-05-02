import type { ReactNode } from "react";

type IconProps = {
  color?: string;
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

export const IconDashboard = ({ color = "#bd93f9" }: IconProps) => (
  <Icon>
    <rect x="2" y="2" width="5" height="5" fill={color} />
    <rect x="9" y="2" width="5" height="3" fill="#8be9fd" />
    <rect x="9" y="7" width="5" height="7" fill="#ff79c6" />
    <rect x="2" y="9" width="5" height="5" fill="#50fa7b" />
  </Icon>
);

export const IconFolder = ({ color = "#f1fa8c" }: IconProps) => (
  <Icon>
    <rect x="1" y="4" width="6" height="2" fill={color} />
    <rect x="1" y="5" width="14" height="9" fill={color} />
    <rect x="1" y="5" width="14" height="1" fill="#ffb86c" />
    <rect x="1" y="13" width="14" height="1" fill="#282a36" />
  </Icon>
);

export const IconPost = ({ color = "#8be9fd" }: IconProps) => (
  <Icon>
    <rect x="2" y="2" width="12" height="12" fill={color} />
    <rect x="4" y="4" width="8" height="1" fill="#282a36" />
    <rect x="4" y="6" width="8" height="1" fill="#282a36" />
    <rect x="4" y="8" width="6" height="1" fill="#282a36" />
    <rect x="4" y="10" width="8" height="1" fill="#282a36" />
  </Icon>
);

export const IconStar = ({ color = "#f1fa8c" }: IconProps) => (
  <Icon>
    <rect x="7" y="1" width="2" height="2" fill={color} />
    <rect x="6" y="3" width="4" height="2" fill={color} />
    <rect x="1" y="5" width="14" height="2" fill={color} />
    <rect x="3" y="7" width="10" height="2" fill={color} />
    <rect x="5" y="9" width="6" height="2" fill={color} />
    <rect x="5" y="12" width="2" height="2" fill={color} />
    <rect x="9" y="12" width="2" height="2" fill={color} />
  </Icon>
);

export const IconNavMail = ({ color = "#ff79c6" }: IconProps) => (
  <Icon>
    <rect x="1" y="3" width="14" height="10" fill={color} />
    <rect x="2" y="4" width="2" height="2" fill="#282a36" />
    <rect x="4" y="6" width="2" height="2" fill="#282a36" />
    <rect x="6" y="8" width="4" height="1" fill="#282a36" />
    <rect x="10" y="6" width="2" height="2" fill="#282a36" />
    <rect x="12" y="4" width="2" height="2" fill="#282a36" />
  </Icon>
);

export const IconChart = ({ color = "#50fa7b" }: IconProps) => (
  <Icon>
    <rect x="2" y="10" width="2" height="4" fill={color} />
    <rect x="5" y="7" width="2" height="7" fill="#8be9fd" />
    <rect x="8" y="4" width="2" height="10" fill="#bd93f9" />
    <rect x="11" y="2" width="2" height="12" fill="#ff79c6" />
  </Icon>
);

export const IconCog = ({ color = "#bdc0cc" }: IconProps) => (
  <Icon>
    <rect x="7" y="1" width="2" height="2" fill={color} />
    <rect x="7" y="13" width="2" height="2" fill={color} />
    <rect x="1" y="7" width="2" height="2" fill={color} />
    <rect x="13" y="7" width="2" height="2" fill={color} />
    <rect x="4" y="4" width="8" height="8" fill={color} />
    <rect x="6" y="6" width="4" height="4" fill="#282a36" />
  </Icon>
);

export const IconAbout = ({ color = "#ffb86c" }: IconProps) => (
  <Icon>
    <rect x="6" y="1" width="4" height="4" fill={color} />
    <rect x="5" y="5" width="6" height="3" fill={color} />
    <rect x="3" y="8" width="10" height="5" fill={color} />
    <rect x="7" y="2" width="2" height="2" fill="#282a36" />
  </Icon>
);

