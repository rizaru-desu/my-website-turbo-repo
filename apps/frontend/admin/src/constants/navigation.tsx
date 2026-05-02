import type { PageKey, NavItem, PageMeta } from "../types/dashboard";
import {
  IconDashboard,
  IconFolder,
  IconPost,
  IconStar,
  IconNavMail,
  IconChart,
  IconAbout,
  IconCog,
} from "../components/icons/navIcons";

export const navItems: NavItem[] = [
  { key: "dashboard", label: "DASHBOARD", icon: <IconDashboard /> },
  { key: "projects", label: "PROJECTS", icon: <IconFolder /> },
  { key: "posts", label: "POSTS", icon: <IconPost /> },
  { key: "skills", label: "SKILLS", icon: <IconStar /> },
  { key: "messages", label: "MESSAGES", icon: <IconNavMail /> },
  { key: "analytics", label: "ANALYTICS", icon: <IconChart /> },
  { key: "about", label: "ABOUT", icon: <IconAbout /> },
  { key: "settings", label: "SETTINGS", icon: <IconCog /> },
];

export const pageMeta: Record<PageKey, PageMeta> = {
  dashboard: {
    crumb: "HOME / DASHBOARD",
    title: "DASHBOARD",
    sub: "Overview of your portfolio quests, stats, and live sessions.",
  },
  projects: {
    crumb: "HOME / PROJECTS",
    title: "PROJECTS",
    sub: "Manage published and drafted projects. Level them up.",
  },
  posts: {
    crumb: "HOME / POSTS",
    title: "BLOG POSTS",
    sub: "Write, edit, and publish devlog-style entries.",
  },
  skills: {
    crumb: "HOME / SKILLS",
    title: "SKILLS & BIO",
    sub: "Your character sheet: abilities, achievements, and story.",
  },
  messages: {
    crumb: "HOME / INBOX",
    title: "MESSAGES",
    sub: "Client mail and collaboration requests.",
  },
  analytics: {
    crumb: "HOME / ANALYTICS",
    title: "ANALYTICS",
    sub: "Traffic, sources, and campaign performance.",
  },
  about: {
    crumb: "HOME / ABOUT",
    title: "ABOUT.TXT",
    sub: "Manage your public profile, bio, social links, and highlight stats.",
  },
  settings: {
    crumb: "HOME / SETTINGS",
    title: "SETTINGS",
    sub: "Account configuration and preferences.",
  },
};
