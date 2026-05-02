export type PageKey =
  | "dashboard"
  | "projects"
  | "posts"
  | "skills"
  | "messages"
  | "analytics"
  | "about"
  | "settings";

export interface IconProps {
  color?: string;
}

export interface AdminDashboardProps {
  onLogout: () => void;
}

export interface Project {
  id: string;
  name: string;
  status: string;
  type: string;
  views: number;
  stars: number;
  updated: string;
}

export interface Post {
  title: string;
  cat: string;
  date: string;
  reads: number;
  status: string;
}

export interface Message {
  from: string;
  subject: string;
  preview: string;
  date: string;
  unread: boolean;
}

export interface Activity {
  ts: string;
  text: string;
  item: string;
  color: string;
}

export interface Skill {
  name: string;
  level: number;
  colorClass: string;
}

export interface NavItem {
  key: PageKey;
  label: string;
  icon: React.ReactNode;
}

export interface PageMeta {
  crumb: string;
  title: string;
  sub: string;
}
