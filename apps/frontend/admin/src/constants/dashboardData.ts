import type { Project, Post, Message, Activity, Skill } from "../types";

export const projects: Project[] = [
  {
    id: "PRJ-001",
    name: "Nebula Commerce UI Kit",
    status: "Published",
    type: "UI Kit",
    views: 12482,
    stars: 324,
    updated: "2 hrs ago",
  },
  {
    id: "PRJ-002",
    name: "Wayfinder Dashboard",
    status: "Draft",
    type: "Dashboard",
    views: 0,
    stars: 0,
    updated: "yesterday",
  },
  {
    id: "PRJ-003",
    name: "Glyph Icon Library",
    status: "Published",
    type: "Icons",
    views: 8741,
    stars: 512,
    updated: "3 days ago",
  },
  {
    id: "PRJ-004",
    name: "Runekeeper Mobile App",
    status: "Review",
    type: "Mobile",
    views: 204,
    stars: 12,
    updated: "5 days ago",
  },
  {
    id: "PRJ-005",
    name: "Obsidian Landing Page",
    status: "Published",
    type: "Landing",
    views: 23910,
    stars: 891,
    updated: "1 week ago",
  },
  {
    id: "PRJ-006",
    name: "Grimoire Design System",
    status: "Published",
    type: "System",
    views: 5420,
    stars: 187,
    updated: "2 weeks ago",
  },
];

export const posts: Post[] = [
  {
    title: "Designing with Constraints: Pixel Grids in 2026",
    cat: "Design",
    date: "Jan 08",
    reads: 4210,
    status: "Live",
  },
  {
    title: "From Figma to Framer: A Motion Toolkit",
    cat: "Tooling",
    date: "Jan 02",
    reads: 2891,
    status: "Live",
  },
  {
    title: "The Case for Retro UI in Modern Dashboards",
    cat: "Opinion",
    date: "Dec 24",
    reads: 9812,
    status: "Live",
  },
  {
    title: "Color Theory for Tired Developers",
    cat: "Design",
    date: "Dec 11",
    reads: 1540,
    status: "Draft",
  },
];

export const messages: Message[] = [
  {
    from: "AUREL @ NEON LABS",
    subject: "Contract revision - round 3",
    preview:
      "Sending over the latest scope doc. Let me know when you are free...",
    date: "10:42",
    unread: true,
  },
  {
    from: "ZIVA KORR",
    subject: "Re: Landing page handoff",
    preview:
      "Assets uploaded to the shared drive. The hero animation is ready...",
    date: "09:15",
    unread: true,
  },
  {
    from: "CMD / DISCORD",
    subject: "You were mentioned in #design-crit",
    preview: "Love the pixel treatment on the stats cards, very readable.",
    date: "Yesterday",
    unread: false,
  },
];

export const activity: Activity[] = [
  {
    ts: "10:42",
    text: "PUBLISHED",
    item: "Nebula Commerce UI Kit v2.1",
    color: "var(--green)",
  },
  {
    ts: "09:18",
    text: "COMMENTED",
    item: "Wayfinder Dashboard",
    color: "var(--cyan)",
  },
  {
    ts: "08:02",
    text: "NEW FOLLOWER",
    item: "@pixelsmith",
    color: "var(--pink)",
  },
  {
    ts: "Yest.",
    text: "UPDATED",
    item: "Glyph Icon Library (+12 icons)",
    color: "var(--purple)",
  },
  {
    ts: "Yest.",
    text: "REPLIED",
    item: "hiring@pluma.co",
    color: "var(--orange)",
  },
];

export const skills: Skill[] = [
  { name: "FIGMA", level: 98, colorClass: "c-cyan" },
  { name: "REACT", level: 92, colorClass: "" },
  { name: "MOTION DESIGN", level: 84, colorClass: "c-pink" },
  { name: "TYPOGRAPHY", level: 90, colorClass: "c-yellow" },
  { name: "TAILWIND", level: 88, colorClass: "c-green" },
];
