import { useState } from "react";
import type { ReactNode } from "react";
import { useNavigate, useLocation } from "react-router-dom";
import { AppLayout } from "../layout";
import { navItems, pageMeta } from "../../constants/navigation";

type PageKey =
  | "dashboard"
  | "projects"
  | "posts"
  | "skills"
  | "messages"
  | "analytics"
  | "settings";

type IconProps = {
  color?: string;
};

type AdminDashboardProps = {
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

const IconFolder = ({ color = "#f1fa8c" }: IconProps) => (
  <Icon>
    <rect x="1" y="4" width="6" height="2" fill={color} />
    <rect x="1" y="5" width="14" height="9" fill={color} />
    <rect x="1" y="5" width="14" height="1" fill="#ffb86c" />
    <rect x="1" y="13" width="14" height="1" fill="#282a36" />
  </Icon>
);

const IconStar = ({ color = "#f1fa8c" }: IconProps) => (
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

const IconMail = ({ color = "#ff79c6" }: IconProps) => (
  <Icon>
    <rect x="1" y="3" width="14" height="10" fill={color} />
    <rect x="2" y="4" width="2" height="2" fill="#282a36" />
    <rect x="4" y="6" width="2" height="2" fill="#282a36" />
    <rect x="6" y="8" width="4" height="1" fill="#282a36" />
    <rect x="10" y="6" width="2" height="2" fill="#282a36" />
    <rect x="12" y="4" width="2" height="2" fill="#282a36" />
  </Icon>
);

const IconChart = ({ color = "#50fa7b" }: IconProps) => (
  <Icon>
    <rect x="2" y="10" width="2" height="4" fill={color} />
    <rect x="5" y="7" width="2" height="7" fill="#8be9fd" />
    <rect x="8" y="4" width="2" height="10" fill="#bd93f9" />
    <rect x="11" y="2" width="2" height="12" fill="#ff79c6" />
  </Icon>
);

const IconPlus = () => (
  <Icon>
    <rect x="7" y="3" width="2" height="10" fill="#282a36" />
    <rect x="3" y="7" width="10" height="2" fill="#282a36" />
  </Icon>
);

const IconEdit = () => (
  <Icon>
    <rect x="10" y="2" width="3" height="3" fill="#8be9fd" />
    <rect x="8" y="4" width="3" height="3" fill="#8be9fd" />
    <rect x="6" y="6" width="3" height="3" fill="#8be9fd" />
    <rect x="4" y="8" width="3" height="3" fill="#8be9fd" />
    <rect x="2" y="10" width="3" height="3" fill="#8be9fd" />
  </Icon>
);

const IconTrash = () => (
  <Icon>
    <rect x="6" y="2" width="4" height="1" fill="#ff5555" />
    <rect x="3" y="3" width="10" height="1" fill="#ff5555" />
    <rect x="4" y="4" width="8" height="10" fill="#ff5555" />
    <rect x="6" y="6" width="1" height="6" fill="#282a36" />
    <rect x="8" y="6" width="1" height="6" fill="#282a36" />
    <rect x="10" y="6" width="1" height="6" fill="#282a36" />
  </Icon>
);

const IconEye = ({ color = "#bd93f9" }: IconProps) => (
  <Icon>
    <rect x="3" y="6" width="10" height="4" fill={color} />
    <rect x="2" y="7" width="12" height="2" fill={color} />
    <rect x="6" y="6" width="4" height="4" fill="#282a36" />
    <rect x="7" y="7" width="2" height="2" fill="#f8f8f2" />
  </Icon>
);

const IsoChest = () => (
  <svg viewBox="0 0 64 64" shapeRendering="crispEdges" aria-hidden="true">
    <rect x="16" y="14" width="32" height="6" fill="#bd93f9" />
    <rect x="18" y="12" width="28" height="2" fill="#d4b3ff" />
    <rect x="14" y="20" width="36" height="28" fill="#44475a" />
    <rect x="14" y="20" width="36" height="3" fill="#5a5d73" />
    <rect x="14" y="44" width="36" height="4" fill="#2e3040" />
    <rect x="14" y="28" width="36" height="3" fill="#ffb86c" />
    <rect x="14" y="38" width="36" height="3" fill="#ffb86c" />
    <rect x="28" y="32" width="8" height="6" fill="#f1fa8c" />
    <rect x="30" y="34" width="4" height="2" fill="#282a36" />
    <rect x="12" y="50" width="40" height="3" fill="#21222c" />
  </svg>
);

const projects = [
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

const posts = [
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

const messages = [
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

const activity = [
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

const skills = [
  { name: "FIGMA", level: 98, colorClass: "c-cyan" },
  { name: "REACT", level: 92, colorClass: "" },
  { name: "MOTION DESIGN", level: 84, colorClass: "c-pink" },
  { name: "TYPOGRAPHY", level: 90, colorClass: "c-yellow" },
  { name: "TAILWIND", level: 88, colorClass: "c-green" },
];

const StatTile = ({
  label,
  value,
  delta,
  up,
  accent,
  icon,
}: {
  label: string;
  value: string;
  delta: string;
  up: boolean;
  accent: string;
  icon: ReactNode;
}) => (
  <div
    className={`cms-card cms-stat accent-${accent} pixel-corners`}
    data-testid={`stat-${label.toLowerCase().replace(/\s+/g, "-")}`}
  >
    <div className="cms-stat-top">
      <span className="label">{label}</span>
      <div className="bob">{icon}</div>
    </div>
    <div className="value" style={{ color: `var(--${accent})` }}>
      {value}
    </div>
    <div className={up ? "delta delta-up" : "delta delta-down"}>
      {up ? "UP" : "DOWN"} {delta}
    </div>
  </div>
);

const ProjectsTable = ({ slim = false }: { slim?: boolean }) => {
  const rows = slim ? projects.slice(0, 4) : projects;
  const statusTag = (status: string) =>
    status === "Published"
      ? "tag-green"
      : status === "Draft"
        ? "tag-yellow"
        : status === "Review"
          ? "tag-orange"
          : "tag-purple";

  return (
    <div className="table-scroll">
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
            <th>ACTIONS</th>
          </tr>
        </thead>
        <tbody>
          {rows.map((project) => (
            <tr key={project.id} data-testid={`project-row-${project.id}`}>
              <td className="muted-cell">{project.id}</td>
              <td className="cyan-cell">{project.name}</td>
              <td>{project.type}</td>
              <td>
                <span className={`tag ${statusTag(project.status)}`}>
                  {project.status.toUpperCase()}
                </span>
              </td>
              <td>{project.views.toLocaleString()}</td>
              <td className="yellow-cell">* {project.stars}</td>
              <td>{project.updated}</td>
              <td>
                <div className="table-actions">
                  <button className="pix-btn pix-btn-ghost hover-wiggle">
                    <IconEye />
                  </button>
                  <button className="pix-btn pix-btn-ghost hover-wiggle">
                    <IconEdit />
                  </button>
                  <button className="pix-btn pix-btn-ghost hover-wiggle">
                    <IconTrash />
                  </button>
                </div>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

const Dashboard = () => (
  <>
    <div className="stat-grid">
      <StatTile
        label="TOTAL VIEWS"
        value="58.3K"
        delta="+12.4% THIS WEEK"
        up
        accent="purple"
        icon={<IconEye color="#bd93f9" />}
      />
      <StatTile
        label="PROJECTS LIVE"
        value="24"
        delta="+3 NEW"
        up
        accent="cyan"
        icon={<IconFolder color="#8be9fd" />}
      />
      <StatTile
        label="SUBSCRIBERS"
        value="1,842"
        delta="+5.1%"
        up
        accent="pink"
        icon={<IconMail color="#ff79c6" />}
      />
      <StatTile
        label="AVG. RATING"
        value="4.8"
        delta="-0.1 VS LAST WK"
        up={false}
        accent="green"
        icon={<IconStar color="#50fa7b" />}
      />
    </div>

    <div className="grid-2 dashboard-grid">
      <div className="cms-card accent-purple" data-testid="traffic-card">
        <div className="card-title-row">
          <div className="card-title">
            &gt; TRAFFIC.EXE
            <span className="blink" />
          </div>
          <div className="range-pills">
            <button className="pixel-badge">7D</button>
            <button className="pixel-badge active">30D</button>
            <button className="pixel-badge">90D</button>
          </div>
        </div>
        <div className="bar-chart">
          {[60, 82, 45, 98, 70, 120, 88, 140, 102, 75, 160, 128].map(
            (height, index) => (
              <div
                key={index}
                className={`bar c${(index % 4) + 1}`}
                style={{ height }}
              />
            ),
          )}
        </div>
        <div className="month-labels">
          {[
            "JAN",
            "FEB",
            "MAR",
            "APR",
            "MAY",
            "JUN",
            "JUL",
            "AUG",
            "SEP",
            "OCT",
            "NOV",
            "DEC",
          ].map((month) => (
            <span key={month}>{month}</span>
          ))}
        </div>
      </div>

      <div className="cms-card accent-cyan" data-testid="activity-card">
        <div className="iso-deco">
          <IsoChest />
        </div>
        <div className="card-title-row">
          <div className="card-title">&gt; QUEST LOG</div>
          <span className="pixel-badge live">LIVE</span>
        </div>
        {activity.map((item) => (
          <div className="activity-item" key={`${item.ts}-${item.item}`}>
            <span className="ts">{item.ts}</span>
            <span>
              <span className="activity-action" style={{ color: item.color }}>
                {item.text}
              </span>{" "}
              {item.item}
            </span>
          </div>
        ))}
      </div>
    </div>

    <div className="cms-card accent-pink" data-testid="recent-projects-card">
      <div className="card-title-row">
        <div className="card-title">&gt; RECENT PROJECTS</div>
        <button className="pix-btn pix-btn-pink">
          <IconPlus /> NEW PROJECT
        </button>
      </div>
      <ProjectsTable slim />
    </div>
  </>
);

const Projects = () => (
  <div className="stack">
    <div className="cms-card accent-purple" data-testid="projects-card">
      <div className="card-title-row">
        <div className="card-title">&gt; ALL PROJECTS ({projects.length})</div>
        <div className="row-actions">
          <button className="pix-btn pix-btn-ghost">FILTER</button>
          <button className="pix-btn">
            <IconPlus /> NEW PROJECT
          </button>
        </div>
      </div>
      <ProjectsTable />
    </div>

    <div className="grid-2">
      <div className="cms-card accent-cyan">
        <div className="card-title-row">
          <div className="card-title">&gt; EDIT PROJECT ~ PRJ-001</div>
          <span className="pixel-badge warn">UNSAVED *</span>
        </div>
        <div className="cms-form-row">
          <label>PROJECT NAME</label>
          <input className="pix-input" defaultValue="Nebula Commerce UI Kit" />
        </div>
        <div className="cms-form-grid">
          <div className="cms-form-row">
            <label>TYPE</label>
            <select className="pix-select" defaultValue="UI Kit">
              <option>UI Kit</option>
              <option>Dashboard</option>
              <option>Landing</option>
              <option>Mobile</option>
            </select>
          </div>
          <div className="cms-form-row">
            <label>STATUS</label>
            <select className="pix-select" defaultValue="Published">
              <option>Draft</option>
              <option>Review</option>
              <option>Published</option>
            </select>
          </div>
        </div>
        <div className="cms-form-row">
          <label>SHORT DESCRIPTION</label>
          <textarea
            className="pix-textarea"
            defaultValue="A deep-space inspired commerce UI kit with reusable layouts, dark palette, and motion presets."
          />
        </div>
        <div className="pix-divider" />
        <div className="save-row">
          <span className="toggle-label">FEATURED?</span>
          <div className="pix-switch on">
            <div className="pix-switch-knob" />
          </div>
          <button className="pix-btn pix-btn-ghost push-right">DISCARD</button>
          <button className="pix-btn pix-btn-green">SAVE</button>
        </div>
      </div>

      <div className="cms-card accent-green">
        <div className="card-title-row">
          <div className="card-title">&gt; STATS</div>
        </div>
        <div className="terminal-list">
          <div>
            VIEWS (30D)........... <span>12,482</span>
          </div>
          <div>
            STARS................. <span>* 324</span>
          </div>
          <div>
            DOWNLOADS............. <span>1,891</span>
          </div>
          <div>
            CONVERSION............ <span>4.2%</span>
          </div>
        </div>
      </div>
    </div>
  </div>
);

const Posts = () => (
  <div className="grid-2">
    <div className="cms-card accent-purple" data-testid="post-editor-card">
      <div className="card-title-row">
        <div className="card-title">&gt; WRITE NEW POST</div>
        <span className="pixel-badge">AUTOSAVED 2m AGO</span>
      </div>
      <div className="cms-form-row">
        <label>TITLE</label>
        <input
          className="pix-input"
          defaultValue="Designing with Constraints: Pixel Grids in 2026"
        />
      </div>
      <div className="cms-form-row">
        <label>CONTENT ~ MARKDOWN</label>
        <textarea
          className="pix-textarea tall"
          defaultValue={`# The Return of the Grid

Grids are the quiet heroes of every great interface.

## Why 16-bit again?
- Clearer hierarchy under tight constraints
- Memorable brand signal
- Instant nostalgia without losing legibility`}
        />
      </div>
      <div className="row-actions end">
        <button className="pix-btn pix-btn-ghost">PREVIEW</button>
        <button className="pix-btn pix-btn-pink">PUBLISH</button>
      </div>
    </div>

    <div className="cms-card accent-cyan" data-testid="recent-posts-card">
      <div className="card-title-row">
        <div className="card-title">&gt; RECENT POSTS</div>
      </div>
      {posts.map((post) => (
        <div className="post-row" key={post.title}>
          <div className="post-title">{post.title}</div>
          <div className="post-meta">
            <span className="tag tag-purple">{post.cat.toUpperCase()}</span>
            <span>{post.date}</span>
            <span>{post.reads.toLocaleString()} reads</span>
            <span
              className={`tag ${post.status === "Live" ? "tag-green" : "tag-yellow"}`}
            >
              {post.status.toUpperCase()}
            </span>
          </div>
        </div>
      ))}
    </div>
  </div>
);

const Skills = () => (
  <div className="grid-2">
    <div className="cms-card accent-pink" data-testid="skills-card">
      <div className="card-title-row">
        <div className="card-title">&gt; CHARACTER SHEET</div>
        <span className="pixel-badge live">LVL 42</span>
      </div>
      {skills.map((skill) => (
        <div className="skill-row" key={skill.name}>
          <div className="top">
            <span className="name">{skill.name}</span>
            <span className="lvl">{skill.level}/100</span>
          </div>
          <div className="skill-bar">
            <div
              className={`fill ${skill.colorClass}`}
              style={{ width: `${skill.level}%` }}
            />
          </div>
        </div>
      ))}
      <div className="pix-divider" />
      <div className="cms-form-row">
        <label>ADD NEW SKILL</label>
        <div className="input-with-action">
          <input className="pix-input" placeholder="e.g. RUST, AFTER EFFECTS" />
          <button className="pix-btn pix-btn-pink">
            <IconPlus /> ADD
          </button>
        </div>
      </div>
    </div>

    <div className="cms-card accent-cyan">
      <div className="card-title-row">
        <div className="card-title">&gt; BIO / ABOUT</div>
      </div>
      <div className="cms-form-row">
        <label>DISPLAY NAME</label>
        <input className="pix-input" defaultValue="KAI MORIKAWA" />
      </div>
      <div className="cms-form-row">
        <label>TAGLINE</label>
        <input
          className="pix-input"
          defaultValue="UI/UX designer and frontend dev crafting retro-flavored interfaces."
        />
      </div>
      <div className="cms-form-row">
        <label>BIO</label>
        <textarea
          className="pix-textarea"
          defaultValue="Based in Kyoto, working with studios and founders worldwide. I turn scrappy ideas into tactile, characterful products."
        />
      </div>
    </div>
  </div>
);

const Messages = () => {
  const [selected, setSelected] = useState(0);
  const message = messages[selected];

  return (
    <div className="grid-2">
      <div className="cms-card accent-pink inbox-card" data-testid="inbox-card">
        <div className="card-title-row">
          <div className="card-title">
            &gt; INBOX ({messages.filter((item) => item.unread).length} NEW)
          </div>
        </div>
        {messages.map((item, index) => (
          <button
            key={item.subject}
            className={`msg-row ${selected === index ? "selected" : ""}`}
            onClick={() => setSelected(index)}
            data-testid={`msg-row-${index}`}
          >
            <span className={`dot ${item.unread ? "" : "read"}`} />
            <span>
              <span className="from">{item.from}</span>
              <span className="subject">{item.subject}</span>
              <span className="preview">{item.preview}</span>
            </span>
            <span className="date">{item.date}</span>
          </button>
        ))}
      </div>

      <div className="cms-card accent-cyan" data-testid="msg-detail-card">
        <div className="card-title-row">
          <div className="card-title">&gt; MESSAGE</div>
          <button className="pix-btn pix-btn-cyan">REPLY</button>
        </div>
        <div className="message-from">{message.from}</div>
        <div className="message-subject">{message.subject}</div>
        <div className="message-date">{message.date} / priority: normal</div>
        <div className="pix-divider" />
        <p className="message-body">
          Hope you are doing great. Circling back on the project kickoff we
          discussed last Tuesday. We finalized the scope internally and are
          ready to move forward with the updated prototype.
        </p>
        <div className="cms-form-row">
          <label>QUICK REPLY</label>
          <textarea
            className="pix-textarea small"
            placeholder="Type reply..."
          />
        </div>
      </div>
    </div>
  );
};

const Analytics = () => (
  <>
    <div className="stat-grid">
      <StatTile
        label="SESSIONS"
        value="92.1K"
        delta="+8.2%"
        up
        accent="purple"
        icon={<IconChart />}
      />
      <StatTile
        label="AVG. TIME"
        value="3:42"
        delta="+0:18"
        up
        accent="cyan"
        icon={<IconEye color="#8be9fd" />}
      />
      <StatTile
        label="BOUNCE RATE"
        value="28%"
        delta="-2.1%"
        up
        accent="green"
        icon={<IconStar color="#50fa7b" />}
      />
      <StatTile
        label="TOP COUNTRY"
        value="JP"
        delta="41% TRAFFIC"
        up
        accent="pink"
        icon={<IconMail color="#ff79c6" />}
      />
    </div>
    <div className="cms-card accent-green">
      <div className="card-title-row">
        <div className="card-title">&gt; TOP PAGES</div>
      </div>
      <div className="table-scroll">
        <table className="pixel-table">
          <thead>
            <tr>
              <th>#</th>
              <th>PATH</th>
              <th>VIEWS</th>
              <th>AVG TIME</th>
              <th>BOUNCE</th>
            </tr>
          </thead>
          <tbody>
            {[
              ["/case-studies/nebula", "23,910", "4:12", "19%"],
              ["/projects", "18,441", "2:58", "31%"],
              ["/blog/pixel-grids", "9,812", "5:04", "17%"],
              ["/about", "7,120", "1:44", "44%"],
            ].map(([path, views, time, bounce], index) => (
              <tr key={path}>
                <td className="muted-cell">0{index + 1}</td>
                <td className="cyan-cell">{path}</td>
                <td>{views}</td>
                <td>{time}</td>
                <td>{bounce}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  </>
);

const Settings = () => (
  <div className="grid-2">
    <div className="cms-card accent-purple">
      <div className="card-title-row">
        <div className="card-title">&gt; ACCOUNT.CFG</div>
      </div>
      <div className="cms-form-row">
        <label>EMAIL</label>
        <input className="pix-input" defaultValue="kai@morikawa.studio" />
      </div>
      <div className="cms-form-row">
        <label>HANDLE</label>
        <input className="pix-input" defaultValue="@kai_pixels" />
      </div>
      <div className="setting-row">
        <span>EMAIL NOTIFICATIONS</span>
        <div className="pix-switch on">
          <div className="pix-switch-knob" />
        </div>
      </div>
    </div>
  </div>
);

export default function AdminDashboard({ onLogout }: AdminDashboardProps) {
  const navigate = useNavigate();
  const location = useLocation();
  const initialPage =
    (location.state as { initialPage?: PageKey } | null)?.initialPage ??
    "dashboard";
  const [page, setPage] = useState<PageKey>(initialPage);
  const meta = pageMeta[page];

  const handleNavigate = (newPage: PageKey) => {
    if (newPage === "settings") {
      navigate("/settings");
    } else {
      setPage(newPage);
    }
  };

  const renderPage = () => {
    switch (page) {
      case "projects":
        return <Projects />;
      case "posts":
        return <Posts />;
      case "skills":
        return <Skills />;
      case "messages":
        return <Messages />;
      case "analytics":
        return <Analytics />;
      case "settings":
        return <Settings />;
      default:
        return <Dashboard />;
    }
  };

  return (
    <AppLayout
      currentPage={page}
      onNavigate={handleNavigate}
      onLogout={onLogout}
      navItems={navItems}
      pageTitle={meta}
    >
      {renderPage()}
    </AppLayout>
  );
}
