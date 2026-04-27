import { Link } from "react-router-dom";
import "@/App.css";

/* Pixel avatar — bigger version for hero */
const HeroAvatar = () => (
  <svg viewBox="0 0 16 16" width="160" height="160" shapeRendering="crispEdges" style={{ imageRendering: "pixelated" }}>
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

const IconGrid = () => (
  <svg viewBox="0 0 16 16" width="18" height="18" shapeRendering="crispEdges">
    <rect x="2" y="2" width="5" height="5" fill="#bd93f9"/><rect x="9" y="2" width="5" height="5" fill="#8be9fd"/>
    <rect x="2" y="9" width="5" height="5" fill="#ff79c6"/><rect x="9" y="9" width="5" height="5" fill="#50fa7b"/>
  </svg>
);

const IconMail = () => (
  <svg viewBox="0 0 16 16" width="16" height="16" shapeRendering="crispEdges">
    <rect x="1" y="3" width="14" height="10" fill="#ff79c6"/>
    <rect x="2" y="4" width="12" height="1" fill="#282a36"/>
    <rect x="3" y="5" width="10" height="1" fill="#282a36"/>
  </svg>
);

const IconExternal = () => (
  <svg viewBox="0 0 16 16" width="12" height="12" shapeRendering="crispEdges">
    <rect x="8" y="2" width="6" height="2" fill="#282a36"/>
    <rect x="12" y="2" width="2" height="6" fill="#282a36"/>
    <rect x="7" y="7" width="1" height="1" fill="#282a36"/>
    <rect x="8" y="6" width="1" height="1" fill="#282a36"/>
    <rect x="9" y="5" width="1" height="1" fill="#282a36"/>
    <rect x="10" y="4" width="1" height="1" fill="#282a36"/>
    <rect x="2" y="5" width="6" height="9" fill="none" stroke="#282a36"/>
  </svg>
);

/* Small pixel sword divider */
const Sword = () => (
  <svg viewBox="0 0 32 8" width="64" height="16" shapeRendering="crispEdges" style={{ imageRendering: "pixelated" }}>
    <rect x="0" y="3" width="2" height="2" fill="#ff79c6"/>
    <rect x="2" y="2" width="2" height="4" fill="#ff79c6"/>
    <rect x="4" y="3" width="2" height="2" fill="#bdc0cc"/>
    <rect x="6" y="3" width="18" height="2" fill="#f8f8f2"/>
    <rect x="7" y="2" width="16" height="1" fill="#bdc0cc"/>
    <rect x="7" y="5" width="16" height="1" fill="#6272a4"/>
    <rect x="24" y="2" width="2" height="4" fill="#f1fa8c"/>
    <rect x="26" y="1" width="2" height="6" fill="#f1fa8c"/>
    <rect x="28" y="3" width="2" height="2" fill="#f1fa8c"/>
  </svg>
);

const highlights = [
  { k: "YEARS XP",       v: "8+",    c: "purple", sub: "since 2018" },
  { k: "PROJECTS",       v: "42",    c: "cyan",   sub: "shipped & live" },
  { k: "HAPPY CLIENTS",  v: "28",    c: "pink",   sub: "worldwide" },
  { k: "COFFEES /DAY",   v: "∞",     c: "green",  sub: "house blend" },
];

const timeline = [
  { year: "2025 →", role: "PRINCIPAL DESIGNER", org: "FREELANCE · TOKYO", c: "purple",
    body: "Leading end-to-end product design & frontend for early-stage SaaS and indie studios. Specialty: retro-flavored modern UIs." },
  { year: "2022-25", role: "SENIOR UI/UX DESIGNER", org: "NEBULA LABS", c: "cyan",
    body: "Owned the design system powering 5 B2B products. Shipped 120+ Figma components, reduced design-dev handoff time by 60%." },
  { year: "2020-22", role: "FRONTEND DEVELOPER", org: "PLUMA STUDIO", c: "pink",
    body: "Built production React + Tailwind interfaces for fintech and creator-economy clients. Obsessed over micro-interactions." },
  { year: "2018-20", role: "JUNIOR DESIGNER", org: "BYTECRAFT", c: "green",
    body: "Cut my teeth on landing pages, brand systems, and the occasional game jam. Where the pixel love started." },
];

const stack = [
  { label: "DESIGN",  items: ["FIGMA", "FRAMER", "PROTOPIE", "BLENDER", "ASEPRITE"], c: "purple" },
  { label: "BUILD",   items: ["REACT", "TYPESCRIPT", "TAILWIND", "NEXT.JS", "THREE.JS"], c: "cyan" },
  { label: "CRAFT",   items: ["MOTION", "TYPOGRAPHY", "COLOR", "PIXEL ART", "SOUND"], c: "pink" },
];

const PlaceholderCard = ({ to, icon, title, desc, c = "purple", testid }) => (
  <Link to={to} className={`card accent-${c} placeholder-card`} data-testid={testid}>
    <div style={{ display: "flex", alignItems: "center", gap: 12, marginBottom: 12 }}>
      {icon}
      <div className="card-title" style={{ color: `var(--${c})` }}>{title}</div>
    </div>
    <p style={{ fontFamily: "'VT323',monospace", fontSize: 17, color: "var(--muted)", lineHeight: 1.3, margin: 0 }}>{desc}</p>
    <div style={{
      marginTop: 14,
      display: "flex", justifyContent: "space-between", alignItems: "center",
      fontFamily: "'Press Start 2P',monospace", fontSize: 8,
      color: "var(--line)",
    }}>
      <span>SOON ~ COMING</span>
      <IconExternal />
    </div>
  </Link>
);

export default function LandingPage() {
  return (
    <div className="public-shell" data-testid="public-shell">
      {/* ---------- NAVBAR ---------- */}
      <nav className="public-nav" data-testid="public-nav">
        <Link to="/" className="nav-brand">
          <svg viewBox="0 0 16 16" width="28" height="28" shapeRendering="crispEdges">
            <rect x="11" y="2" width="3" height="3" fill="#bd93f9"/><rect x="10" y="3" width="3" height="3" fill="#bd93f9"/>
            <rect x="9" y="4" width="3" height="3" fill="#bd93f9"/><rect x="8" y="5" width="3" height="3" fill="#bd93f9"/>
            <rect x="7" y="6" width="3" height="3" fill="#bd93f9"/><rect x="6" y="7" width="3" height="3" fill="#bd93f9"/>
            <rect x="5" y="8" width="3" height="3" fill="#bd93f9"/><rect x="4" y="9" width="3" height="3" fill="#bd93f9"/>
            <rect x="2" y="10" width="4" height="3" fill="#ff79c6"/><rect x="3" y="13" width="2" height="2" fill="#8be9fd"/>
          </svg>
          <div>
            <div style={{ fontFamily: "'Press Start 2P',monospace", fontSize: 10, color: "var(--purple)" }}>KAI.MORIKAWA</div>
            <div style={{ fontFamily: "'VT323',monospace", fontSize: 14, color: "var(--cyan)", lineHeight: 1 }}>portfolio ~ v4.20</div>
          </div>
        </Link>

        <div className="nav-links">
          <a href="#about"    data-testid="nav-about">ABOUT</a>
          <a href="#timeline" data-testid="nav-timeline">JOURNEY</a>
          <a href="#stack"    data-testid="nav-stack">STACK</a>
          <a href="#explore"  data-testid="nav-explore">EXPLORE</a>
          <a href="#contact"  data-testid="nav-contact">CONTACT</a>
        </div>

        <div className="nav-actions">
          <Link to="/login" className="pix-btn pix-btn-ghost hover-wiggle" data-testid="nav-login">
            <IconGrid /> ADMIN
          </Link>
        </div>
      </nav>

      {/* ---------- HERO ---------- */}
      <section className="hero" data-testid="hero">
        <div className="hero-grid">
          <div className="hero-left">
            <div className="hero-tag">
              <span className="blinker" /> AVAILABLE FOR CONTRACT · Q2 2026
            </div>
            <h1 className="hero-title">
              <span style={{ color: "var(--text)" }}>HI, I'M</span>{" "}
              <span style={{ color: "var(--purple)" }}>KAI</span>.
              <br/>
              <span style={{ color: "var(--cyan)" }}>PIXEL-PUSHER</span>{" "}
              <span style={{ color: "var(--text)" }}>&amp;</span>{" "}
              <span style={{ color: "var(--pink)" }}>CODE-SMITH</span>.
            </h1>
            <p className="hero-sub">
              UI/UX designer and frontend developer from Kyoto, crafting
              characterful, retro-flavored interfaces for startups, studios
              and the occasional weekend game jam.
            </p>

            <div style={{ display: "flex", gap: 12, flexWrap: "wrap", marginTop: 8 }}>
              <a href="#contact" className="pix-btn pix-btn-pink hover-wiggle" data-testid="hero-cta-hire">
                HIRE ME ~ &gt;
              </a>
              <a href="#timeline" className="pix-btn pix-btn-ghost hover-wiggle" data-testid="hero-cta-resume">
                VIEW RESUME
              </a>
              <Link to="/login" className="pix-btn pix-btn-cyan hover-wiggle" data-testid="hero-cta-admin">
                <IconGrid /> ADMIN CMS
              </Link>
            </div>

            <div className="hero-socials">
              <span>FIND ME ~</span>
              <a href="#" data-testid="social-gh">GITHUB</a>
              <span>/</span>
              <a href="#" data-testid="social-dr">DRIBBBLE</a>
              <span>/</span>
              <a href="#" data-testid="social-fg">FIGMA</a>
              <span>/</span>
              <a href="#" data-testid="social-tw">TWITTER</a>
            </div>
          </div>

          <div className="hero-right">
            <div className="avatar-frame">
              <div className="avatar-inner">
                <HeroAvatar />
              </div>
              <div className="avatar-nameplate">
                <div style={{ fontFamily: "'Press Start 2P',monospace", fontSize: 9, color: "var(--purple)" }}>KAI MORIKAWA</div>
                <div style={{ fontFamily: "'VT323',monospace", fontSize: 15, color: "var(--cyan)", lineHeight: 1.1 }}>
                  LVL 42 · DESIGNER / DEV
                </div>
                <div style={{ marginTop: 6 }}>
                  <div className="xp-bar"><div className="xp-bar-fill" style={{ width: "68%" }} /></div>
                  <div style={{ fontFamily: "'VT323',monospace", fontSize: 13, color: "var(--line)", marginTop: 3 }}>
                    HP 100/100 · MP 86/100
                  </div>
                </div>
              </div>
            </div>

            {/* Floating pixel coins */}
            <div className="coin c1 bob" />
            <div className="coin c2 bob" />
            <div className="coin c3 bob" />
          </div>
        </div>

        {/* Highlights strip */}
        <div className="highlight-strip">
          {highlights.map((h) => (
            <div key={h.k} className={`card accent-${h.c} pixel-corners`} data-testid={`highlight-${h.k.replace(/\s+/g,'-').toLowerCase()}`}>
              <span className="label" style={{ fontFamily: "'Press Start 2P',monospace", fontSize: 8, color: "var(--line)", letterSpacing: 1, display: "block", marginBottom: 10 }}>
                {h.k}
              </span>
              <div style={{ fontFamily: "'Press Start 2P',monospace", fontSize: 22, color: `var(--${h.c})`, marginBottom: 6 }}>
                {h.v}
              </div>
              <div style={{ fontFamily: "'VT323',monospace", fontSize: 15, color: "var(--muted)" }}>
                {h.sub}
              </div>
            </div>
          ))}
        </div>
      </section>

      {/* ---------- ABOUT ---------- */}
      <section className="section" id="about" data-testid="section-about">
        <div className="section-header">
          <div>
            <div className="section-kicker">~ CHAPTER 01</div>
            <h2 className="section-title">ABOUT.TXT</h2>
          </div>
          <Sword />
        </div>
        <div className="about-grid">
          <div className="card accent-cyan">
            <p style={{ fontFamily: "'VT323',monospace", fontSize: 19, lineHeight: 1.35, color: "var(--text)", margin: 0 }}>
              I'm <span style={{ color: "var(--purple)" }}>Kai</span> — half designer, half engineer, fully obsessed with
              the quiet details that make interfaces feel{" "}
              <span style={{ color: "var(--pink)" }}>alive</span>.
              <br/><br/>
              I've spent the last{" "}
              <span style={{ color: "var(--green)" }}>8 years</span> shipping product
              UIs for SaaS tools, dashboards, and content-heavy apps — always
              chasing the balance between{" "}
              <span style={{ color: "var(--cyan)" }}>legibility</span> and{" "}
              <span style={{ color: "var(--yellow)" }}>personality</span>.
              <br/><br/>
              Lately I've been exploring how retro aesthetics (pixel grids,
              CRT textures, arcade typography) can make modern dashboards
              genuinely delightful to return to, without sacrificing
              usability.
            </p>
          </div>
          <div className="stack" style={{ gap: 20 }}>
            <div className="card accent-purple">
              <div className="card-title" style={{ marginBottom: 10 }}>&gt; CURRENT QUEST</div>
              <p style={{ fontFamily: "'VT323',monospace", fontSize: 17, color: "var(--muted)", margin: 0, lineHeight: 1.3 }}>
                Designing a retro-flavored content platform for indie writers.
                Shipping in Q2.
              </p>
            </div>
            <div className="card accent-pink">
              <div className="card-title" style={{ marginBottom: 10 }}>&gt; OBSESSIONS</div>
              <div style={{ display: "flex", flexWrap: "wrap", gap: 6 }}>
                <span className="tag tag-purple">PIXEL ART</span>
                <span className="tag tag-cyan">SYNTHWAVE</span>
                <span className="tag tag-pink">KEYBOARDS</span>
                <span className="tag tag-green">GAME UI</span>
                <span className="tag tag-yellow">TYPE</span>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* ---------- TIMELINE / RESUME SUMMARY ---------- */}
      <section className="section" id="timeline" data-testid="section-timeline">
        <div className="section-header">
          <div>
            <div className="section-kicker">~ CHAPTER 02</div>
            <h2 className="section-title">JOURNEY.LOG</h2>
          </div>
          <button className="pix-btn pix-btn-ghost hover-wiggle" data-testid="download-cv">
            ↓ DOWNLOAD CV
          </button>
        </div>

        <div className="timeline">
          {timeline.map((t, i) => (
            <div className="timeline-row" key={i}>
              <div className="timeline-year" style={{ color: `var(--${t.c})` }}>{t.year}</div>
              <div className="timeline-dot" style={{ background: `var(--${t.c})` }} />
              <div className={`card accent-${t.c} timeline-card`} data-testid={`timeline-${i}`}>
                <div style={{ display: "flex", justifyContent: "space-between", alignItems: "flex-start", flexWrap: "wrap", gap: 8, marginBottom: 8 }}>
                  <div className="card-title" style={{ color: `var(--${t.c})` }}>{t.role}</div>
                  <span className="pixel-badge" style={{ color: "var(--line)" }}>{t.org}</span>
                </div>
                <p style={{ fontFamily: "'VT323',monospace", fontSize: 17, color: "var(--muted)", lineHeight: 1.3, margin: 0 }}>
                  {t.body}
                </p>
              </div>
            </div>
          ))}
        </div>
      </section>

      {/* ---------- STACK ---------- */}
      <section className="section" id="stack" data-testid="section-stack">
        <div className="section-header">
          <div>
            <div className="section-kicker">~ CHAPTER 03</div>
            <h2 className="section-title">INVENTORY</h2>
          </div>
        </div>
        <div className="grid-3">
          {stack.map((s) => (
            <div key={s.label} className={`card accent-${s.c}`} data-testid={`stack-${s.label.toLowerCase()}`}>
              <div className="card-title" style={{ marginBottom: 14, color: `var(--${s.c})` }}>&gt; {s.label}</div>
              <div style={{ display: "flex", flexDirection: "column", gap: 10 }}>
                {s.items.map((it) => (
                  <div key={it} style={{ display: "flex", alignItems: "center", gap: 10 }}>
                    <span style={{
                      width: 10, height: 10,
                      background: `var(--${s.c})`,
                      boxShadow: "inset 0 2px 0 rgba(255,255,255,0.3), inset 0 -2px 0 rgba(0,0,0,0.3)",
                    }} />
                    <span style={{ fontFamily: "'Press Start 2P',monospace", fontSize: 9, color: "var(--text)" }}>{it}</span>
                  </div>
                ))}
              </div>
            </div>
          ))}
        </div>
      </section>

      {/* ---------- EXPLORE (placeholders for future pages) ---------- */}
      <section className="section" id="explore" data-testid="section-explore">
        <div className="section-header">
          <div>
            <div className="section-kicker">~ CHAPTER 04</div>
            <h2 className="section-title">EXPLORE.~</h2>
            <div style={{ fontFamily: "'VT323',monospace", fontSize: 16, color: "var(--muted)", marginTop: 8 }}>
              Dedicated pages below are still loading in the next update.
            </div>
          </div>
        </div>
        <div className="grid-3">
          <PlaceholderCard
            testid="explore-projects"
            to="#"
            c="purple"
            title="&gt; PROJECTS"
            desc="Case studies, UI kits, and shipped products. Deep-dives with process shots, code snippets & lessons learned."
            icon={<svg viewBox="0 0 16 16" width="24" height="24" shapeRendering="crispEdges"><rect x="1" y="4" width="6" height="2" fill="#f1fa8c"/><rect x="1" y="5" width="14" height="9" fill="#bd93f9"/></svg>}
          />
          <PlaceholderCard
            testid="explore-skills"
            to="#"
            c="cyan"
            title="&gt; SKILLS"
            desc="Detailed breakdown of tools, levels, and the projects that taught me each one. Character sheet, but real."
            icon={<svg viewBox="0 0 16 16" width="24" height="24" shapeRendering="crispEdges"><rect x="7" y="1" width="2" height="2" fill="#8be9fd"/><rect x="1" y="5" width="14" height="2" fill="#8be9fd"/><rect x="3" y="7" width="10" height="2" fill="#8be9fd"/><rect x="5" y="9" width="6" height="2" fill="#8be9fd"/></svg>}
          />
          <PlaceholderCard
            testid="explore-posts"
            to="#"
            c="pink"
            title="&gt; WRITING"
            desc="Devlogs, opinions, and the occasional rant. Design systems, motion, and the state of retro UI in 2026."
            icon={<svg viewBox="0 0 16 16" width="24" height="24" shapeRendering="crispEdges"><rect x="2" y="2" width="12" height="12" fill="#ff79c6"/><rect x="4" y="5" width="8" height="1" fill="#282a36"/><rect x="4" y="7" width="8" height="1" fill="#282a36"/><rect x="4" y="9" width="6" height="1" fill="#282a36"/></svg>}
          />
        </div>
      </section>

      {/* ---------- CONTACT ---------- */}
      <section className="section" id="contact" data-testid="section-contact">
        <div className="card accent-green contact-card">
          <Sword />
          <h2 style={{ fontFamily: "'Press Start 2P',monospace", fontSize: 22, color: "var(--green)", margin: "16px 0 10px" }}>
            READY PLAYER ONE?
          </h2>
          <p style={{ fontFamily: "'VT323',monospace", fontSize: 19, color: "var(--muted)", lineHeight: 1.35, maxWidth: 560, textAlign: "center", margin: "0 auto 20px" }}>
            Got a product, studio, or weird side-quest that needs a design-engineer
            hybrid? Let's talk. I reply within 24 hours — promise.
          </p>
          <div style={{ display: "flex", gap: 12, justifyContent: "center", flexWrap: "wrap" }}>
            <a href="mailto:kai@morikawa.studio" className="pix-btn pix-btn-green hover-wiggle" data-testid="contact-email">
              <IconMail /> KAI@MORIKAWA.STUDIO
            </a>
            <a href="#" className="pix-btn pix-btn-ghost hover-wiggle" data-testid="contact-schedule">
              ▸ BOOK A CALL
            </a>
          </div>
        </div>
      </section>

      {/* ---------- FOOTER ---------- */}
      <footer className="public-footer" data-testid="public-footer">
        <div style={{ display: "flex", justifyContent: "space-between", alignItems: "center", flexWrap: "wrap", gap: 12 }}>
          <div style={{ fontFamily: "'Press Start 2P',monospace", fontSize: 8, color: "var(--line)", letterSpacing: 1 }}>
            © 2026 KAI MORIKAWA · ALL PIXELS RESERVED
          </div>
          <div style={{ display: "flex", gap: 18, fontFamily: "'Press Start 2P',monospace", fontSize: 8 }}>
            <a href="#" style={{ color: "var(--cyan)", textDecoration: "none" }}>GH</a>
            <a href="#" style={{ color: "var(--pink)", textDecoration: "none" }}>DR</a>
            <a href="#" style={{ color: "var(--purple)", textDecoration: "none" }}>FG</a>
            <a href="#" style={{ color: "var(--green)", textDecoration: "none" }}>TW</a>
            <Link to="/login" style={{ color: "var(--line)", textDecoration: "none" }} data-testid="footer-login">ADMIN</Link>
          </div>
        </div>
        <div style={{ marginTop: 10, fontFamily: "'VT323',monospace", fontSize: 14, color: "var(--line)" }}>
          PRESS <span style={{ color: "var(--cyan)" }}>[START]</span> TO CONTINUE ~ BUILT WITH REACT, CSS &amp; LOVE FOR THE GRID
        </div>
      </footer>
    </div>
  );
}
