import { useState } from "react";
import { useNavigate } from "react-router-dom";
import "@/App.css";

/* =========================================================
   Pixel SVG icons (all 16x16 chunky sprites, Dracula palette)
   ========================================================= */
const px = (color) => color;

const Icon = ({ children, size = 18 }) => (
  <svg
    width={size}
    height={size}
    viewBox="0 0 16 16"
    shapeRendering="crispEdges"
    xmlns="http://www.w3.org/2000/svg"
    style={{ imageRendering: "pixelated", flexShrink: 0 }}
  >
    {children}
  </svg>
);

// Each icon uses 1px squares — true pixel art
const IconDashboard = ({ c = "#bd93f9" }) => (
  <Icon>
    <rect x="2" y="2" width="5" height="5" fill={c} />
    <rect x="9" y="2" width="5" height="3" fill="#8be9fd" />
    <rect x="9" y="7" width="5" height="7" fill="#ff79c6" />
    <rect x="2" y="9" width="5" height="5" fill="#50fa7b" />
  </Icon>
);
const IconFolder = ({ c = "#f1fa8c" }) => (
  <Icon>
    <rect x="1" y="4" width="6" height="2" fill={c} />
    <rect x="1" y="5" width="14" height="9" fill={c} />
    <rect x="1" y="5" width="14" height="1" fill="#ffb86c" />
    <rect x="1" y="13" width="14" height="1" fill="#282a36" />
  </Icon>
);
const IconPost = ({ c = "#8be9fd" }) => (
  <Icon>
    <rect x="2" y="2" width="12" height="12" fill={c} />
    <rect x="4" y="4" width="8" height="1" fill="#282a36" />
    <rect x="4" y="6" width="8" height="1" fill="#282a36" />
    <rect x="4" y="8" width="6" height="1" fill="#282a36" />
    <rect x="4" y="10" width="8" height="1" fill="#282a36" />
  </Icon>
);
const IconStar = ({ c = "#f1fa8c" }) => (
  <Icon>
    <rect x="7" y="1" width="2" height="2" fill={c} />
    <rect x="6" y="3" width="4" height="2" fill={c} />
    <rect x="1" y="5" width="14" height="2" fill={c} />
    <rect x="3" y="7" width="10" height="2" fill={c} />
    <rect x="2" y="9" width="3" height="3" fill={c} />
    <rect x="11" y="9" width="3" height="3" fill={c} />
    <rect x="5" y="9" width="6" height="2" fill={c} />
    <rect x="5" y="12" width="2" height="2" fill={c} />
    <rect x="9" y="12" width="2" height="2" fill={c} />
  </Icon>
);
const IconMail = ({ c = "#ff79c6" }) => (
  <Icon>
    <rect x="1" y="3" width="14" height="10" fill={c} />
    <rect x="1" y="3" width="14" height="1" fill="#282a36" />
    <rect x="1" y="3" width="1" height="10" fill="#282a36" />
    <rect x="14" y="3" width="1" height="10" fill="#282a36" />
    <rect x="1" y="12" width="14" height="1" fill="#282a36" />
    <rect x="2" y="4" width="2" height="2" fill="#282a36" />
    <rect x="4" y="6" width="2" height="2" fill="#282a36" />
    <rect x="6" y="8" width="4" height="1" fill="#282a36" />
    <rect x="10" y="6" width="2" height="2" fill="#282a36" />
    <rect x="12" y="4" width="2" height="2" fill="#282a36" />
  </Icon>
);
const IconChart = ({ c = "#50fa7b" }) => (
  <Icon>
    <rect x="2" y="10" width="2" height="4" fill={c} />
    <rect x="5" y="7" width="2" height="7" fill="#8be9fd" />
    <rect x="8" y="4" width="2" height="10" fill="#bd93f9" />
    <rect x="11" y="2" width="2" height="12" fill="#ff79c6" />
  </Icon>
);
const IconCog = ({ c = "#bdc0cc" }) => (
  <Icon>
    <rect x="7" y="1" width="2" height="2" fill={c} />
    <rect x="7" y="13" width="2" height="2" fill={c} />
    <rect x="1" y="7" width="2" height="2" fill={c} />
    <rect x="13" y="7" width="2" height="2" fill={c} />
    <rect x="3" y="3" width="2" height="2" fill={c} />
    <rect x="11" y="3" width="2" height="2" fill={c} />
    <rect x="3" y="11" width="2" height="2" fill={c} />
    <rect x="11" y="11" width="2" height="2" fill={c} />
    <rect x="4" y="4" width="8" height="8" fill={c} />
    <rect x="6" y="6" width="4" height="4" fill="#282a36" />
  </Icon>
);
const IconSearch = () => (
  <Icon>
    <rect x="2" y="2" width="8" height="2" fill="#f8f8f2" />
    <rect x="2" y="2" width="2" height="8" fill="#f8f8f2" />
    <rect x="8" y="2" width="2" height="8" fill="#f8f8f2" />
    <rect x="2" y="8" width="8" height="2" fill="#f8f8f2" />
    <rect x="10" y="10" width="2" height="2" fill="#f8f8f2" />
    <rect x="12" y="12" width="2" height="2" fill="#f8f8f2" />
  </Icon>
);
const IconBell = ({ c = "#ffb86c" }) => (
  <Icon>
    <rect x="6" y="2" width="4" height="1" fill={c} />
    <rect x="4" y="3" width="8" height="1" fill={c} />
    <rect x="3" y="4" width="10" height="7" fill={c} />
    <rect x="2" y="11" width="12" height="1" fill={c} />
    <rect x="6" y="12" width="4" height="2" fill={c} />
    <rect x="7" y="14" width="2" height="1" fill="#ff79c6" />
  </Icon>
);
const IconPlus = ({ c = "#282a36" }) => (
  <Icon>
    <rect x="7" y="3" width="2" height="10" fill={c} />
    <rect x="3" y="7" width="10" height="2" fill={c} />
  </Icon>
);
const IconTrash = ({ c = "#ff5555" }) => (
  <Icon>
    <rect x="6" y="2" width="4" height="1" fill={c} />
    <rect x="3" y="3" width="10" height="1" fill={c} />
    <rect x="4" y="4" width="8" height="10" fill={c} />
    <rect x="6" y="6" width="1" height="6" fill="#282a36" />
    <rect x="8" y="6" width="1" height="6" fill="#282a36" />
    <rect x="10" y="6" width="1" height="6" fill="#282a36" />
  </Icon>
);
const IconEdit = ({ c = "#8be9fd" }) => (
  <Icon>
    <rect x="10" y="2" width="3" height="3" fill={c} />
    <rect x="8" y="4" width="3" height="3" fill={c} />
    <rect x="6" y="6" width="3" height="3" fill={c} />
    <rect x="4" y="8" width="3" height="3" fill={c} />
    <rect x="2" y="10" width="3" height="3" fill={c} />
    <rect x="2" y="13" width="3" height="1" fill="#282a36" />
  </Icon>
);
const IconEye = ({ c = "#bd93f9" }) => (
  <Icon>
    <rect x="3" y="6" width="10" height="4" fill={c} />
    <rect x="2" y="7" width="12" height="2" fill={c} />
    <rect x="6" y="6" width="4" height="4" fill="#282a36" />
    <rect x="7" y="7" width="2" height="2" fill="#f8f8f2" />
  </Icon>
);

/* Isometric pixel decoration — chest */
const IsoChest = () => (
  <svg viewBox="0 0 64 64" shapeRendering="crispEdges" style={{ width: 80, height: 80 }}>
    {/* top lid */}
    <rect x="16" y="14" width="32" height="6" fill="#bd93f9" />
    <rect x="18" y="12" width="28" height="2" fill="#bd93f9" />
    <rect x="18" y="12" width="28" height="2" fill="#d4b3ff" />
    <rect x="16" y="14" width="32" height="2" fill="#d4b3ff" />
    <rect x="16" y="18" width="32" height="2" fill="#7a5cb8" />
    {/* body */}
    <rect x="14" y="20" width="36" height="28" fill="#44475a" />
    <rect x="14" y="20" width="36" height="3" fill="#5a5d73" />
    <rect x="14" y="44" width="36" height="4" fill="#2e3040" />
    {/* bands */}
    <rect x="14" y="28" width="36" height="3" fill="#ffb86c" />
    <rect x="14" y="38" width="36" height="3" fill="#ffb86c" />
    {/* lock */}
    <rect x="28" y="32" width="8" height="6" fill="#f1fa8c" />
    <rect x="30" y="34" width="4" height="2" fill="#282a36" />
    {/* shadow */}
    <rect x="12" y="50" width="40" height="3" fill="#21222c" />
  </svg>
);

/* =========================================================
   Sample data
   ========================================================= */
const projects = [
  { id: "PRJ-001", name: "Nebula Commerce UI Kit", status: "Published", type: "UI Kit", views: 12482, stars: 324, updated: "2 hrs ago" },
  { id: "PRJ-002", name: "Wayfinder Dashboard", status: "Draft", type: "Dashboard", views: 0, stars: 0, updated: "yesterday" },
  { id: "PRJ-003", name: "Glyph Icon Library", status: "Published", type: "Icons", views: 8741, stars: 512, updated: "3 days ago" },
  { id: "PRJ-004", name: "Runekeeper Mobile App", status: "Review", type: "Mobile", views: 204, stars: 12, updated: "5 days ago" },
  { id: "PRJ-005", name: "Obsidian Landing Page", status: "Published", type: "Landing", type2: "Landing", views: 23910, stars: 891, updated: "1 week ago" },
  { id: "PRJ-006", name: "Grimoire Design System", status: "Published", type: "System", views: 5420, stars: 187, updated: "2 weeks ago" },
];

const posts = [
  { title: "Designing with Constraints: Pixel Grids in 2026", cat: "Design", date: "Jan 08", reads: 4210, status: "Live" },
  { title: "From Figma to Framer: A Motion Toolkit", cat: "Tooling", date: "Jan 02", reads: 2891, status: "Live" },
  { title: "The Case for Retro UI in Modern Dashboards", cat: "Opinion", date: "Dec 24", reads: 9812, status: "Live" },
  { title: "Color Theory for Tired Developers", cat: "Design", date: "Dec 11", reads: 1540, status: "Draft" },
];

const messages = [
  { from: "AUREL @ NEON LABS", subject: "Contract revision — round 3", preview: "Hey, sending over the latest scope doc. Let me know when you're free to...", date: "10:42", unread: true },
  { from: "ZIVA KORR",  subject: "Re: Landing page handoff",  preview: "Assets uploaded to the shared drive. The hero section animation is...",   date: "09:15", unread: true },
  { from: "CMD • DISCORD", subject: "You were mentioned in #design-crit", preview: "@you — love the pixel treatment on the stats cards, very readable.",    date: "Yesterday", unread: false },
  { from: "HIRING · PLUMA CO", subject: "Interview invite — Principal Designer", preview: "We'd love to chat about the senior design role. Availability this week?", date: "Yesterday", unread: false },
  { from: "NEWSLETTER",  subject: "Your weekly analytics digest",   preview: "Traffic up 18%, your top page was /case-studies/nebula...",          date: "Mon", unread: false },
];

const skills = [
  { name: "FIGMA",        lvl: 98, c: "c-cyan" },
  { name: "REACT",        lvl: 92, c: "" },
  { name: "MOTION DESIGN",lvl: 84, c: "c-pink" },
  { name: "TYPOGRAPHY",   lvl: 90, c: "c-yellow" },
  { name: "TAILWIND",     lvl: 88, c: "c-green" },
];

const activity = [
  { ts: "10:42", text: "PUBLISHED ", item: "Nebula Commerce UI Kit v2.1", c: "var(--green)" },
  { ts: "09:18", text: "COMMENTED on ", item: "Wayfinder Dashboard", c: "var(--cyan)" },
  { ts: "08:02", text: "NEW FOLLOWER ", item: "@pixelsmith", c: "var(--pink)" },
  { ts: "Yest.", text: "UPDATED ", item: "Glyph Icon Library (+12 icons)", c: "var(--purple)" },
  { ts: "Yest.", text: "REPLIED to ", item: "hiring@pluma.co", c: "var(--orange)" },
];

/* =========================================================
   Sections
   ========================================================= */
const StatTile = ({ label, value, delta, up, accent, icon }) => (
  <div className={`card stat accent-${accent} pixel-corners`} data-testid={`stat-${label.toLowerCase().replace(/\s+/g,'-')}`}>
    <div style={{ display: "flex", justifyContent: "space-between", alignItems: "flex-start" }}>
      <span className="label">{label}</span>
      <div className="bob">{icon}</div>
    </div>
    <div className="value" style={{ color: `var(--${accent})` }}>{value}</div>
    <div className={`delta ${up ? "delta-up" : "delta-down"}`}>
      {up ? "▲" : "▼"} {delta}
    </div>
  </div>
);

const Dashboard = () => (
  <>
    <div className="stat-grid">
      <StatTile label="TOTAL VIEWS"   value="58.3K"  delta="+12.4% THIS WEEK" up accent="purple" icon={<IconEye c="#bd93f9" />} />
      <StatTile label="PROJECTS LIVE" value="24"     delta="+3 NEW"            up accent="cyan"   icon={<IconFolder c="#8be9fd" />} />
      <StatTile label="SUBSCRIBERS"   value="1,842"  delta="+5.1%"             up accent="pink"   icon={<IconMail c="#ff79c6" />} />
      <StatTile label="AVG. RATING"   value="4.8"    delta="-0.1 VS LAST WK"   up={false} accent="green" icon={<IconStar c="#50fa7b" />} />
    </div>

    <div className="grid-2" style={{ marginBottom: 28 }}>
      <div className="card accent-purple" data-testid="traffic-card">
        <div className="card-title-row">
          <div className="card-title">&gt; TRAFFIC.EXE<span className="blink" /></div>
          <div style={{ display: "flex", gap: 8 }}>
            <button className="pix-badge" data-testid="range-7d">7D</button>
            <button className="pix-badge" style={{ color: "var(--purple)" }}>30D</button>
            <button className="pix-badge">90D</button>
          </div>
        </div>
        <div className="bar-chart">
          {[60, 82, 45, 98, 70, 120, 88, 140, 102, 75, 160, 128].map((h, i) => (
            <div key={i} className={`bar ${i % 4 === 0 ? "" : i % 4 === 1 ? "c2" : i % 4 === 2 ? "c3" : "c4"}`} style={{ height: h }} />
          ))}
        </div>
        <div className="labels" style={{ display: "grid", gridTemplateColumns: "repeat(12,1fr)", gap: 8, paddingTop: 8 }}>
          {["JAN","FEB","MAR","APR","MAY","JUN","JUL","AUG","SEP","OCT","NOV","DEC"].map(m => (
            <span key={m} style={{ textAlign: "center", fontFamily: "'Press Start 2P',monospace", fontSize: 7, color: "var(--line)" }}>{m}</span>
          ))}
        </div>
      </div>

      <div className="card accent-cyan" data-testid="activity-card" style={{ position: "relative" }}>
        <div className="iso-deco"><IsoChest /></div>
        <div className="card-title-row">
          <div className="card-title">&gt; QUEST LOG</div>
          <span className="pixel-badge" style={{ color: "var(--green)" }}>LIVE</span>
        </div>
        {activity.map((a, i) => (
          <div className="activity-item" key={i}>
            <span className="ts">{a.ts}</span>
            <span>
              <span style={{ color: a.c, fontFamily: "'Press Start 2P',monospace", fontSize: 8 }}>{a.text}</span>
              {a.item}
            </span>
          </div>
        ))}
      </div>
    </div>

    <div className="card accent-pink" data-testid="recent-projects-card">
      <div className="card-title-row">
        <div className="card-title">&gt; RECENT PROJECTS</div>
        <button className="pix-btn pix-btn-pink" data-testid="new-project-btn"><IconPlus /> NEW PROJECT</button>
      </div>
      <ProjectsTable slim />
    </div>
  </>
);

const ProjectsTable = ({ slim = false }) => {
  const rows = slim ? projects.slice(0, 4) : projects;
  const statusTag = (s) =>
    s === "Published" ? "tag-green" :
    s === "Draft"     ? "tag-yellow" :
    s === "Review"    ? "tag-orange" : "tag-purple";

  return (
    <table className="pixel-table" data-testid="projects-table">
      <thead>
        <tr>
          <th>ID</th>
          <th>NAME</th>
          <th>TYPE</th>
          <th>STATUS</th>
          <th>VIEWS</th>
          <th>STARS</th>
          <th>UPDATED</th>
          <th style={{ textAlign: "right" }}>ACTIONS</th>
        </tr>
      </thead>
      <tbody>
        {rows.map((p) => (
          <tr key={p.id} data-testid={`project-row-${p.id}`}>
            <td style={{ color: "var(--line)", fontFamily: "'Press Start 2P',monospace", fontSize: 9 }}>{p.id}</td>
            <td style={{ color: "var(--cyan)" }}>{p.name}</td>
            <td>{p.type}</td>
            <td><span className={`tag ${statusTag(p.status)}`}>{p.status.toUpperCase()}</span></td>
            <td>{p.views.toLocaleString()}</td>
            <td style={{ color: "var(--yellow)" }}>★ {p.stars}</td>
            <td style={{ color: "var(--line)" }}>{p.updated}</td>
            <td style={{ textAlign: "right" }}>
              <div style={{ display: "inline-flex", gap: 6 }}>
                <button className="pix-btn pix-btn-ghost hover-wiggle" style={{ padding: "6px 8px" }} data-testid={`view-${p.id}`}><IconEye /></button>
                <button className="pix-btn pix-btn-ghost hover-wiggle" style={{ padding: "6px 8px" }} data-testid={`edit-${p.id}`}><IconEdit /></button>
                <button className="pix-btn pix-btn-ghost hover-wiggle" style={{ padding: "6px 8px" }} data-testid={`delete-${p.id}`}><IconTrash /></button>
              </div>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

const Projects = () => (
  <>
    <div className="card accent-purple" data-testid="projects-card" style={{ marginBottom: 28 }}>
      <div className="card-title-row">
        <div className="card-title">&gt; ALL PROJECTS ({projects.length})</div>
        <div style={{ display: "flex", gap: 10 }}>
          <button className="pix-btn pix-btn-ghost" data-testid="filter-btn">FILTER</button>
          <button className="pix-btn" data-testid="new-project-btn-2"><IconPlus /> NEW PROJECT</button>
        </div>
      </div>
      <ProjectsTable />
    </div>

    <div className="grid-2">
      <div className="card accent-cyan" data-testid="edit-project-card">
        <div className="card-title-row">
          <div className="card-title">&gt; EDIT PROJECT ~ PRJ-001</div>
          <span className="pixel-badge" style={{ color: "var(--pink)" }}>UNSAVED *</span>
        </div>

        <div className="form-row">
          <label>PROJECT NAME</label>
          <input className="pix-input" defaultValue="Nebula Commerce UI Kit" data-testid="project-name-input" />
        </div>
        <div style={{ display: "grid", gridTemplateColumns: "1fr 1fr", gap: 16 }}>
          <div className="form-row">
            <label>TYPE</label>
            <select className="pix-select" defaultValue="UI Kit">
              <option>UI Kit</option><option>Dashboard</option><option>Landing</option><option>Mobile</option>
            </select>
          </div>
          <div className="form-row">
            <label>STATUS</label>
            <select className="pix-select" defaultValue="Published">
              <option>Draft</option><option>Review</option><option>Published</option>
            </select>
          </div>
        </div>
        <div className="form-row">
          <label>SHORT DESCRIPTION</label>
          <textarea className="pix-textarea" defaultValue="A deep-space inspired commerce UI kit — 120+ components, 40+ page layouts, built for Figma and React. Comes with a complete dark palette and motion presets." />
        </div>
        <div className="form-row">
          <label>TAGS</label>
          <div style={{ display: "flex", gap: 6, flexWrap: "wrap" }}>
            <span className="tag tag-purple">FIGMA</span>
            <span className="tag tag-cyan">REACT</span>
            <span className="tag tag-pink">UI-KIT</span>
            <span className="tag tag-green">E-COMMERCE</span>
            <span className="tag tag-yellow">DARK-MODE</span>
          </div>
        </div>
        <div className="pix-divider" />
        <div style={{ display: "flex", justifyContent: "space-between", alignItems: "center" }}>
          <div style={{ display: "flex", gap: 10, alignItems: "center" }}>
            <span style={{ fontFamily: "'Press Start 2P',monospace", fontSize: 9, color: "var(--cyan)" }}>FEATURED?</span>
            <div className="pix-switch on" data-testid="featured-toggle"><div className="pix-switch-knob" /></div>
          </div>
          <div style={{ display: "flex", gap: 10 }}>
            <button className="pix-btn pix-btn-ghost" data-testid="discard-btn">DISCARD</button>
            <button className="pix-btn pix-btn-green" data-testid="save-btn">SAVE ~</button>
          </div>
        </div>
      </div>

      <div className="stack">
        <div className="card accent-pink" data-testid="cover-card">
          <div className="card-title-row"><div className="card-title">&gt; COVER ART</div></div>
          {/* Fake pixel preview — a little landscape */}
          <div style={{
            height: 160,
            background: "linear-gradient(to bottom, #44475a 0 35%, #6272a4 35% 60%, #282a36 60% 100%)",
            position: "relative",
            boxShadow: "inset 0 0 0 4px var(--bg), inset 0 0 0 8px var(--line)",
            overflow: "hidden",
          }}>
            {/* pixel sun */}
            <div style={{ position: "absolute", top: 16, right: 18, width: 30, height: 30, background: "var(--pink)",
              boxShadow: "inset 0 4px 0 rgba(255,255,255,0.35), inset 0 -4px 0 rgba(0,0,0,0.25)"}} />
            {/* mountains */}
            <div style={{ position: "absolute", bottom: 56, left: 20, width: 0, height: 0,
              borderLeft: "40px solid transparent", borderRight: "40px solid transparent", borderBottom: "50px solid #bd93f9" }} />
            <div style={{ position: "absolute", bottom: 56, left: 100, width: 0, height: 0,
              borderLeft: "60px solid transparent", borderRight: "60px solid transparent", borderBottom: "70px solid #8be9fd" }} />
            {/* stars */}
            {[[40,20],[80,12],[140,28],[220,18],[260,32],[310,14]].map(([l,t],i)=>(
              <div key={i} style={{ position: "absolute", left: l, top: t, width: 4, height: 4, background: "var(--yellow)" }} />
            ))}
            {/* grid floor */}
            <div style={{ position: "absolute", bottom: 0, left: 0, right: 0, height: 60,
              backgroundImage: "linear-gradient(rgba(189,147,249,0.35) 1px, transparent 1px), linear-gradient(90deg, rgba(189,147,249,0.35) 1px, transparent 1px)",
              backgroundSize: "24px 12px",
              backgroundPosition: "center bottom" }} />
          </div>
          <div style={{ marginTop: 12, display: "flex", gap: 8 }}>
            <button className="pix-btn pix-btn-cyan">UPLOAD</button>
            <button className="pix-btn pix-btn-ghost">REPLACE</button>
          </div>
        </div>

        <div className="card accent-green">
          <div className="card-title-row"><div className="card-title">&gt; STATS</div></div>
          <div style={{ fontFamily: "'VT323',monospace", fontSize: 17, lineHeight: 1.6 }}>
            <div>VIEWS (30D)........... <span style={{ color: "var(--cyan)" }}>12,482</span></div>
            <div>STARS................. <span style={{ color: "var(--yellow)" }}>★ 324</span></div>
            <div>DOWNLOADS............. <span style={{ color: "var(--pink)" }}>1,891</span></div>
            <div>CONVERSION............ <span style={{ color: "var(--green)" }}>4.2%</span></div>
          </div>
        </div>
      </div>
    </div>
  </>
);

const Posts = () => (
  <div className="grid-2">
    <div className="card accent-purple" data-testid="post-editor-card">
      <div className="card-title-row">
        <div className="card-title">&gt; WRITE NEW POST</div>
        <span className="pixel-badge" style={{ color: "var(--line)" }}>AUTOSAVED 2m AGO</span>
      </div>
      <div className="form-row">
        <label>TITLE</label>
        <input className="pix-input" defaultValue="Designing with Constraints: Pixel Grids in 2026" />
      </div>
      <div style={{ display: "grid", gridTemplateColumns: "2fr 1fr", gap: 14 }}>
        <div className="form-row">
          <label>SLUG</label>
          <input className="pix-input" defaultValue="designing-with-constraints-pixel-grids-2026" />
        </div>
        <div className="form-row">
          <label>CATEGORY</label>
          <select className="pix-select" defaultValue="Design">
            <option>Design</option><option>Tooling</option><option>Opinion</option>
          </select>
        </div>
      </div>
      <div className="form-row">
        <label>CONTENT ~ MARKDOWN</label>
        <textarea className="pix-textarea" defaultValue={`# The Return of the Grid

Grids are the quiet heroes of every great interface. They don't ask for attention — 
they hold everything in place while your pixels do the talking.

## Why 16-bit again?
- Clearer hierarchy under tight constraints
- Memorable, characterful brand signal
- Instant nostalgia without losing legibility

\`\`\`css
.pixel-border { box-shadow: 0 0 0 4px #282a36, 0 0 0 8px #bd93f9; }
\`\`\`

Try it. Commit to the pixel. Your users will remember.`} />
      </div>
      <div style={{ display: "flex", justifyContent: "space-between" }}>
        <div style={{ display: "flex", gap: 8 }}>
          <button className="pix-btn pix-btn-ghost">B</button>
          <button className="pix-btn pix-btn-ghost" style={{ fontStyle: "italic" }}>I</button>
          <button className="pix-btn pix-btn-ghost">{"<>"}</button>
          <button className="pix-btn pix-btn-ghost">IMG</button>
        </div>
        <div style={{ display: "flex", gap: 10 }}>
          <button className="pix-btn pix-btn-ghost" data-testid="preview-btn">PREVIEW</button>
          <button className="pix-btn pix-btn-pink" data-testid="publish-btn">PUBLISH</button>
        </div>
      </div>
    </div>

    <div className="stack">
      <div className="card accent-cyan" data-testid="recent-posts-card">
        <div className="card-title-row"><div className="card-title">&gt; RECENT POSTS</div></div>
        {posts.map((p, i) => (
          <div key={i} style={{ padding: "10px 0", borderBottom: i < posts.length - 1 ? "2px dashed var(--line)" : 0 }}>
            <div style={{ fontFamily: "'Press Start 2P',monospace", fontSize: 10, color: "var(--text)", marginBottom: 6, lineHeight: 1.4 }}>{p.title}</div>
            <div style={{ display: "flex", gap: 10, alignItems: "center", fontFamily: "'VT323',monospace", fontSize: 15 }}>
              <span className={`tag ${p.cat === "Design" ? "tag-purple" : p.cat === "Tooling" ? "tag-cyan" : "tag-pink"}`}>{p.cat.toUpperCase()}</span>
              <span style={{ color: "var(--line)" }}>{p.date}</span>
              <span style={{ color: "var(--muted)" }}>· {p.reads.toLocaleString()} reads</span>
              <span className={`tag ${p.status === "Live" ? "tag-green" : "tag-yellow"}`} style={{ marginLeft: "auto" }}>{p.status.toUpperCase()}</span>
            </div>
          </div>
        ))}
      </div>

      <div className="card accent-green">
        <div className="card-title-row"><div className="card-title">&gt; SEO CHECK</div></div>
        <div style={{ fontFamily: "'VT323',monospace", fontSize: 17 }}>
          <div>META TITLE........ <span style={{ color: "var(--green)" }}>OK</span></div>
          <div>META DESC......... <span style={{ color: "var(--green)" }}>OK</span></div>
          <div>SLUG LENGTH....... <span style={{ color: "var(--yellow)" }}>LONG</span></div>
          <div>READING TIME...... <span style={{ color: "var(--cyan)" }}>6 min</span></div>
          <div>READABILITY....... <span style={{ color: "var(--green)" }}>A-</span></div>
        </div>
      </div>
    </div>
  </div>
);

const Skills = () => (
  <div className="grid-2">
    <div className="card accent-pink" data-testid="skills-card">
      <div className="card-title-row">
        <div className="card-title">&gt; CHARACTER SHEET</div>
        <span className="pixel-badge" style={{ color: "var(--green)" }}>LVL 42</span>
      </div>
      {skills.map((s) => (
        <div className="skill-row" key={s.name}>
          <div className="top">
            <span className="name">{s.name}</span>
            <span className="lvl">{s.lvl}/100</span>
          </div>
          <div className="skill-bar"><div className={`fill ${s.c}`} style={{ width: `${s.lvl}%` }} /></div>
        </div>
      ))}
      <div className="pix-divider" />
      <div className="form-row">
        <label>ADD NEW SKILL</label>
        <div style={{ display: "flex", gap: 8 }}>
          <input className="pix-input" placeholder="e.g. RUST, AFTER EFFECTS..." data-testid="new-skill-input" />
          <button className="pix-btn pix-btn-pink"><IconPlus /> ADD</button>
        </div>
      </div>
    </div>

    <div className="stack">
      <div className="card accent-cyan">
        <div className="card-title-row"><div className="card-title">&gt; BIO / ABOUT</div></div>
        <div className="form-row">
          <label>DISPLAY NAME</label>
          <input className="pix-input" defaultValue="KAI MORIKAWA" />
        </div>
        <div className="form-row">
          <label>TAGLINE</label>
          <input className="pix-input" defaultValue="UI/UX designer & frontend dev crafting retro-flavored, modern interfaces." />
        </div>
        <div className="form-row">
          <label>BIO</label>
          <textarea className="pix-textarea" defaultValue="Based in Kyoto, working with studios and founders worldwide. I turn scrappy ideas into tactile, characterful products — part engineer, part illustrator, fully caffeinated." />
        </div>
      </div>

      <div className="card accent-purple">
        <div className="card-title-row"><div className="card-title">&gt; ACHIEVEMENTS</div></div>
        <div style={{ display: "grid", gridTemplateColumns: "repeat(3,1fr)", gap: 14 }}>
          {[
            { n: "100K VIEWS", c: "purple", ic: <IconEye c="#bd93f9" /> },
            { n: "50 PROJECTS", c: "cyan", ic: <IconFolder c="#8be9fd" /> },
            { n: "TOP RATED", c: "yellow", ic: <IconStar c="#f1fa8c" /> },
          ].map((a, i) => (
            <div key={i} style={{
              background: "var(--bg-2)",
              padding: "12px 10px",
              textAlign: "center",
              boxShadow: `inset 0 0 0 3px var(--${a.c})`,
            }}>
              <div className="bob" style={{ display: "flex", justifyContent: "center", marginBottom: 6 }}>{a.ic}</div>
              <div style={{ fontFamily: "'Press Start 2P',monospace", fontSize: 8, color: `var(--${a.c})` }}>{a.n}</div>
            </div>
          ))}
        </div>
      </div>
    </div>
  </div>
);

const Messages = () => {
  const [selected, setSelected] = useState(0);
  const m = messages[selected];
  return (
    <div className="grid-2">
      <div className="card accent-pink" data-testid="inbox-card" style={{ padding: 0 }}>
        <div className="card-title-row" style={{ padding: 18, marginBottom: 0 }}>
          <div className="card-title">&gt; INBOX ({messages.filter(x=>x.unread).length} NEW)</div>
          <button className="pix-btn pix-btn-ghost" style={{ padding: "6px 10px", fontSize: 8 }}>MARK ALL READ</button>
        </div>
        <div>
          {messages.map((msg, i) => (
            <div
              key={i}
              className="msg-row"
              onClick={() => setSelected(i)}
              style={{ background: selected === i ? "rgba(189,147,249,0.12)" : undefined }}
              data-testid={`msg-row-${i}`}
            >
              <div className={`dot ${msg.unread ? "" : "read"}`} />
              <div>
                <div className="from">{msg.from}</div>
                <div className="subject">{msg.subject}</div>
                <div className="preview">{msg.preview}</div>
              </div>
              <div className="date">{msg.date}</div>
            </div>
          ))}
        </div>
      </div>

      <div className="card accent-cyan" data-testid="msg-detail-card">
        <div className="card-title-row">
          <div className="card-title">&gt; MESSAGE</div>
          <div style={{ display: "flex", gap: 8 }}>
            <button className="pix-btn pix-btn-ghost" style={{ padding: "6px 8px" }}><IconTrash /></button>
            <button className="pix-btn pix-btn-cyan" style={{ padding: "6px 10px", fontSize: 8 }}>REPLY</button>
          </div>
        </div>
        <div style={{ fontFamily: "'Press Start 2P',monospace", fontSize: 10, color: "var(--pink)", marginBottom: 6 }}>{m.from}</div>
        <div style={{ fontFamily: "'VT323',monospace", fontSize: 22, marginBottom: 4 }}>{m.subject}</div>
        <div style={{ fontFamily: "'VT323',monospace", fontSize: 14, color: "var(--line)", marginBottom: 14 }}>{m.date} · priority: normal</div>
        <div className="pix-divider" />
        <div style={{ fontFamily: "'VT323',monospace", fontSize: 18, lineHeight: 1.4 }}>
          <p>Hey Kai,</p>
          <p>Hope you're doing great. Circling back on the project kickoff we discussed last Tuesday — we've finalized the scope internally and I'm sending over the updated SOW as a follow-up.</p>
          <p>Key points on our side:</p>
          <ul>
            <li>Total pages: 14 (was 10)</li>
            <li>Motion explorations included</li>
            <li>Handoff via Figma + coded prototype</li>
          </ul>
          <p>Let me know if Thursday 15:00 JST works for a kickoff call. We're genuinely excited to get this one moving.</p>
          <p>Cheers,<br/>Aurel</p>
        </div>
        <div className="pix-divider" />
        <div className="form-row">
          <label>QUICK REPLY</label>
          <textarea className="pix-textarea" placeholder="Type your reply..." style={{ minHeight: 90 }} />
        </div>
        <div style={{ display: "flex", justifyContent: "flex-end", gap: 10 }}>
          <button className="pix-btn pix-btn-ghost">SAVE DRAFT</button>
          <button className="pix-btn pix-btn-pink">SEND ~</button>
        </div>
      </div>
    </div>
  );
};

const Analytics = () => (
  <>
    <div className="stat-grid">
      <StatTile label="SESSIONS"     value="92.1K" delta="+8.2%" up accent="purple" icon={<IconChart />} />
      <StatTile label="AVG. TIME"    value="3:42"  delta="+0:18" up accent="cyan"   icon={<IconEye c="#8be9fd" />} />
      <StatTile label="BOUNCE RATE"  value="28%"   delta="-2.1%" up accent="green"  icon={<IconStar c="#50fa7b" />} />
      <StatTile label="TOP COUNTRY"  value="JP"    delta="41% of traffic" up accent="pink" icon={<IconMail c="#ff79c6" />} />
    </div>

    <div className="grid-2" style={{ marginBottom: 28 }}>
      <div className="card accent-cyan" data-testid="analytics-bar-card">
        <div className="card-title-row"><div className="card-title">&gt; VISITS PER DAY</div></div>
        <div className="bar-chart">
          {[45, 68, 52, 90, 120, 82, 140, 98, 76, 130, 150, 108].map((h, i) => (
            <div key={i} className={`bar ${i % 2 === 0 ? "" : "c2"}`} style={{ height: h }} />
          ))}
        </div>
      </div>

      <div className="card accent-purple" data-testid="sources-card">
        <div className="card-title-row"><div className="card-title">&gt; TRAFFIC SOURCES</div></div>
        {/* Pixel block pie substitute */}
        <div style={{ display: "flex", gap: 18, alignItems: "center" }}>
          <div style={{
            width: 120, height: 120,
            display: "grid",
            gridTemplateColumns: "repeat(8, 15px)",
            gridTemplateRows: "repeat(8, 15px)",
          }}>
            {Array.from({ length: 64 }).map((_, i) => {
              const c = i < 26 ? "var(--purple)"
                     : i < 44 ? "var(--cyan)"
                     : i < 56 ? "var(--pink)"
                     : "var(--green)";
              return <div key={i} style={{ background: c, boxShadow: "inset 0 2px 0 rgba(255,255,255,0.2), inset 0 -2px 0 rgba(0,0,0,0.2)" }} />;
            })}
          </div>
          <div className="block-legend">
            <div className="li"><span className="sw" style={{ background: "var(--purple)" }}></span> DIRECT ..... 41%</div>
            <div className="li"><span className="sw" style={{ background: "var(--cyan)" }}></span> SEARCH ..... 28%</div>
            <div className="li"><span className="sw" style={{ background: "var(--pink)" }}></span> SOCIAL ..... 19%</div>
            <div className="li"><span className="sw" style={{ background: "var(--green)" }}></span> REFERRAL .. 12%</div>
          </div>
        </div>
      </div>
    </div>

    <div className="card accent-green" data-testid="top-pages-card">
      <div className="card-title-row"><div className="card-title">&gt; TOP PAGES</div></div>
      <table className="pixel-table">
        <thead>
          <tr><th>#</th><th>PATH</th><th>VIEWS</th><th>AVG TIME</th><th>BOUNCE</th></tr>
        </thead>
        <tbody>
          {[
            { p: "/case-studies/nebula", v: "23,910", t: "4:12", b: "19%" },
            { p: "/projects",            v: "18,441", t: "2:58", b: "31%" },
            { p: "/blog/pixel-grids",    v: "9,812",  t: "5:04", b: "17%" },
            { p: "/about",               v: "7,120",  t: "1:44", b: "44%" },
            { p: "/blog/retro-ui",       v: "4,891",  t: "3:21", b: "22%" },
          ].map((r, i) => (
            <tr key={i}>
              <td style={{ color: "var(--line)", fontFamily: "'Press Start 2P',monospace", fontSize: 9 }}>0{i+1}</td>
              <td style={{ color: "var(--cyan)" }}>{r.p}</td>
              <td>{r.v}</td>
              <td>{r.t}</td>
              <td style={{ color: parseInt(r.b) < 30 ? "var(--green)" : "var(--orange)" }}>{r.b}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  </>
);

const Settings = () => (
  <div className="grid-2">
    <div className="card accent-purple">
      <div className="card-title-row"><div className="card-title">&gt; ACCOUNT.CFG</div></div>
      <div className="form-row">
        <label>EMAIL</label>
        <input className="pix-input" defaultValue="kai@morikawa.studio" />
      </div>
      <div className="form-row">
        <label>HANDLE</label>
        <input className="pix-input" defaultValue="@kai_pixels" />
      </div>
      <div className="form-row">
        <label>TIMEZONE</label>
        <select className="pix-select" defaultValue="Asia/Tokyo">
          <option>Asia/Tokyo</option><option>UTC</option><option>America/New_York</option>
        </select>
      </div>

      {[
        ["PUBLIC PROFILE", true],
        ["EMAIL NOTIFICATIONS", true],
        ["WEEKLY DIGEST", false],
        ["BETA FEATURES", true],
      ].map(([n, on]) => (
        <div key={n} style={{ display: "flex", justifyContent: "space-between", alignItems: "center", padding: "10px 0", borderBottom: "2px dashed var(--line)" }}>
          <span style={{ fontFamily: "'Press Start 2P',monospace", fontSize: 9 }}>{n}</span>
          <div className={`pix-switch ${on ? "on" : ""}`}><div className="pix-switch-knob" /></div>
        </div>
      ))}
    </div>

    <div className="stack">
      <div className="card accent-cyan">
        <div className="card-title-row"><div className="card-title">&gt; THEME</div></div>
        <div style={{ display: "grid", gridTemplateColumns: "repeat(4, 1fr)", gap: 10 }}>
          {[
            { n: "DRACULA", c: "var(--purple)", active: true },
            { n: "NORD", c: "#88c0d0" },
            { n: "GRUVBOX", c: "#fabd2f" },
            { n: "SOLAR", c: "#b58900" },
          ].map((t) => (
            <div key={t.n} style={{
              padding: 14, background: "var(--bg-2)", textAlign: "center",
              fontFamily: "'Press Start 2P',monospace", fontSize: 8, color: t.c,
              boxShadow: `inset 0 0 0 3px ${t.active ? "var(--purple)" : "var(--line)"}`,
              cursor: "pointer",
            }}>
              <div style={{ width: 28, height: 28, background: t.c, margin: "0 auto 8px",
                boxShadow: "inset 0 3px 0 rgba(255,255,255,0.35), inset 0 -3px 0 rgba(0,0,0,0.35)" }} />
              {t.n}
            </div>
          ))}
        </div>
      </div>

      <div className="card accent-pink">
        <div className="card-title-row"><div className="card-title">&gt; DANGER ZONE</div></div>
        <div style={{ fontFamily: "'VT323',monospace", fontSize: 17, color: "var(--muted)", marginBottom: 14 }}>
          These actions are permanent. Your saves cannot be reloaded.
        </div>
        <div style={{ display: "flex", gap: 10 }}>
          <button className="pix-btn pix-btn-ghost">EXPORT DATA</button>
          <button className="pix-btn" style={{ background: "var(--red)", color: "var(--bg)" }}>DELETE ACCOUNT</button>
        </div>
      </div>
    </div>
  </div>
);

/* =========================================================
   Shell
   ========================================================= */
const NAV = [
  { key: "dashboard", label: "DASHBOARD", icon: <IconDashboard /> },
  { key: "projects",  label: "PROJECTS",  icon: <IconFolder /> },
  { key: "posts",     label: "POSTS",     icon: <IconPost /> },
  { key: "skills",    label: "SKILLS",    icon: <IconStar /> },
  { key: "messages",  label: "MESSAGES",  icon: <IconMail /> },
  { key: "analytics", label: "ANALYTICS", icon: <IconChart /> },
  { key: "settings",  label: "SETTINGS",  icon: <IconCog /> },
];

const PAGE_META = {
  dashboard: { crumb: "HOME / DASHBOARD", title: "DASHBOARD", sub: "Overview of your portfolio's quests, stats & live sessions." },
  projects:  { crumb: "HOME / PROJECTS",  title: "PROJECTS",  sub: "Manage your published & drafted projects. Level them up." },
  posts:     { crumb: "HOME / POSTS",     title: "BLOG POSTS",sub: "Write, edit, and publish devlog-style entries." },
  skills:    { crumb: "HOME / SKILLS",    title: "SKILLS & BIO", sub: "Your character sheet — abilities, achievements & story." },
  messages:  { crumb: "HOME / INBOX",     title: "MESSAGES",  sub: "Client pigeon-mail & collaboration requests." },
  analytics: { crumb: "HOME / ANALYTICS", title: "ANALYTICS", sub: "Traffic, sources & campaign performance — plain numbers." },
  settings:  { crumb: "HOME / SETTINGS",  title: "SETTINGS",  sub: "Account configuration, themes, and danger zone." },
};

function CmsApp() {
  const [page, setPage] = useState("dashboard");
  const meta = PAGE_META[page];
  const nav = useNavigate();

  const renderPage = () => {
    switch (page) {
      case "projects":  return <Projects />;
      case "posts":     return <Posts />;
      case "skills":    return <Skills />;
      case "messages":  return <Messages />;
      case "analytics": return <Analytics />;
      case "settings":  return <Settings />;
      default:          return <Dashboard />;
    }
  };

  return (
    <div className="app-shell" data-testid="app-shell">
      {/* ---------- SIDEBAR ---------- */}
      <aside className="sidebar" data-testid="sidebar">
        <div className="sidebar-logo">
          {/* Pixel sword logo */}
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
          <div>
            <div className="title">PIXEL.CMS</div>
            <div className="subtitle">v4.20 ~ retro edition</div>
          </div>
        </div>

        <div className="sidebar-section-label">~ MAIN</div>
        {NAV.slice(0, 4).map((n) => (
          <button
            key={n.key}
            className={`nav-item ${page === n.key ? "active" : ""}`}
            onClick={() => setPage(n.key)}
            data-testid={`nav-${n.key}`}
          >
            {n.icon}
            <span>{n.label}</span>
          </button>
        ))}

        <div className="sidebar-section-label">~ ENGAGE</div>
        {NAV.slice(4).map((n) => (
          <button
            key={n.key}
            className={`nav-item ${page === n.key ? "active" : ""}`}
            onClick={() => setPage(n.key)}
            data-testid={`nav-${n.key}`}
          >
            {n.icon}
            <span>{n.label}</span>
          </button>
        ))}

        <div className="xp-card" data-testid="xp-card">
          <div style={{ fontFamily: "'Press Start 2P',monospace", fontSize: 8, color: "var(--cyan)", marginBottom: 6 }}>KAI · LVL 42</div>
          <div style={{ fontFamily: "'VT323',monospace", fontSize: 15, color: "var(--muted)" }}>XP TO LVL 43</div>
          <div className="xp-bar"><div className="xp-bar-fill" style={{ width: "68%" }} /></div>
          <div style={{ fontFamily: "'VT323',monospace", fontSize: 14, color: "var(--line)", marginTop: 4 }}>6,820 / 10,000</div>
        </div>
      </aside>

      {/* ---------- TOPBAR ---------- */}
      <header className="topbar" data-testid="topbar">
        <div className="search-box">
          <IconSearch />
          <input placeholder="SEARCH PROJECTS, POSTS, COMMANDS..." data-testid="topbar-search" />
          <span style={{ fontFamily: "'Press Start 2P',monospace", fontSize: 8, color: "var(--line)" }}>⌘K</span>
        </div>

        <div className="topbar-actions">
          <span className="pixel-badge hover-wiggle" style={{ color: "var(--green)" }} data-testid="status-online">● ONLINE</span>
          <button className="pix-btn pix-btn-ghost hover-wiggle" style={{ padding: "8px 10px", position: "relative" }} data-testid="bell-btn">
            <IconBell />
            <span style={{ position: "absolute", top: -4, right: -4, width: 12, height: 12, background: "var(--pink)",
              boxShadow: "inset 0 2px 0 rgba(255,255,255,0.35), inset 0 -2px 0 rgba(0,0,0,0.25)",
              fontFamily: "'Press Start 2P',monospace", fontSize: 6, color: "var(--bg)",
              display: "flex", alignItems: "center", justifyContent: "center" }}>3</span>
          </button>
          <div className="avatar-block" data-testid="user-block">
            {/* Pixel avatar */}
            <svg viewBox="0 0 16 16" width="32" height="32" shapeRendering="crispEdges">
              <rect x="4" y="2" width="8" height="6" fill="#ffb86c" />
              <rect x="3" y="3" width="10" height="5" fill="#ffb86c" />
              <rect x="5" y="5" width="1" height="1" fill="#282a36" />
              <rect x="10" y="5" width="1" height="1" fill="#282a36" />
              <rect x="6" y="7" width="4" height="1" fill="#282a36" />
              <rect x="4" y="1" width="8" height="2" fill="#bd93f9" />
              <rect x="3" y="2" width="10" height="1" fill="#bd93f9" />
              <rect x="2" y="9" width="12" height="6" fill="#50fa7b" />
              <rect x="4" y="11" width="8" height="1" fill="#44475a" />
            </svg>
            <div>
              <div className="name">KAI M.</div>
              <div className="role">ADMIN</div>
            </div>
          </div>
          <button
            className="pix-btn pix-btn-ghost hover-wiggle"
            onClick={() => nav("/")}
            data-testid="logout-btn"
            style={{ padding: "8px 10px", fontSize: 8 }}
          >
            ← LOGOUT
          </button>
        </div>
      </header>

      {/* ---------- MAIN ---------- */}
      <main className="main" data-testid="main-content">
        <div className="page-header">
          <div className="page-title">
            <div className="crumbs">&gt; <b>{meta.crumb}</b><span className="blink" /></div>
            <h1 data-testid="page-title">{meta.title}</h1>
            <div className="sub">{meta.sub}</div>
          </div>
          <div style={{ display: "flex", gap: 10 }}>
            <button className="pix-btn pix-btn-ghost hover-wiggle" data-testid="export-btn">EXPORT</button>
            <button className="pix-btn hover-wiggle" data-testid="action-btn">
              <IconPlus /> QUICK ACTION
            </button>
          </div>
        </div>

        {renderPage()}

        <div style={{ marginTop: 48, textAlign: "center", fontFamily: "'Press Start 2P',monospace", fontSize: 8, color: "var(--line)" }}>
          PRESS <span style={{ color: "var(--cyan)" }}>[ESC]</span> TO RETURN ~ PIXEL.CMS v4.20 · MADE WITH &lt;3 &amp; PIXELS
        </div>
      </main>
    </div>
  );
}

export default CmsApp;
