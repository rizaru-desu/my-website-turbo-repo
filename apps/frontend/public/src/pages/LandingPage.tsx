import {
  PublicNav,
  PublicFooter,
  HeroAvatar,
  IconGrid,
  IconMail,
  IconFolder,
  IconStar,
  IconDoc,
  Sword,
  SectionHeader,
  HighlightCard,
  PlaceholderCard,
  TimelineEntry,
} from "../components";

/* ─── Profile data (maps to profileContent DB table) ─── */
const profile = {
  fullName: "KAI MORIKAWA",
  headline: {
    prefix: "HI, I'M",
    name: "KAI",
    roles: [
      { text: "PIXEL-PUSHER", color: "cyan" },
      { text: "CODE-SMITH", color: "pink" },
    ],
  },
  shortIntro:
    "UI/UX designer and frontend developer from Kyoto, crafting characterful, retro-flavored interfaces for startups, studios and the occasional weekend game jam.",
  about:
    "I'm Kai — half designer, half engineer, fully obsessed with the quiet details that make interfaces feel alive.\n\nI've spent the last 8 years shipping product UIs for SaaS tools, dashboards, and content-heavy apps — always chasing the balance between legibility and personality.\n\nLately I've been exploring how retro aesthetics (pixel grids, CRT textures, arcade typography) can make modern dashboards genuinely delightful to return to, without sacrificing usability.",
  location: "KYOTO, JP",
  email: "kai@morikawa.studio",
  availability: "AVAILABLE FOR CONTRACT · Q2 2026",
  primaryCta: "HIRE ME",
  socialLinks: {
    github: { label: "GITHUB", url: "#" },
    dribbble: { label: "DRIBBBLE", url: "#" },
    figma: { label: "FIGMA", url: "#" },
    twitter: { label: "TWITTER", url: "#" },
  },
  focus: [
    "Designing a retro-flavored content platform for indie writers. Shipping in Q2.",
  ],
  stats: [
    { key: "YEARS XP", value: "8+", color: "purple", sub: "since 2018" },
    { key: "PROJECTS", value: "42", color: "cyan", sub: "shipped & live" },
    { key: "HAPPY CLIENTS", value: "28", color: "pink", sub: "worldwide" },
    { key: "COFFEES /DAY", value: "∞", color: "green", sub: "house blend" },
  ],
  obsessions: ["PIXEL ART", "SYNTHWAVE", "KEYBOARDS", "GAME UI", "TYPE"],
  level: { current: 42, xpPercent: 68, hp: "100/100", mp: "86/100" },
};

/* ─── Timeline data (maps to experience DB table) ─── */
const timeline = [
  {
    year: "2025 →",
    role: "PRINCIPAL DESIGNER",
    org: "FREELANCE · TOKYO",
    c: "purple",
    body: "Leading end-to-end product design & frontend for early-stage SaaS and indie studios. Specialty: retro-flavored modern UIs.",
  },
  {
    year: "2022-25",
    role: "SENIOR UI/UX DESIGNER",
    org: "NEBULA LABS",
    c: "cyan",
    body: "Owned the design system powering 5 B2B products. Shipped 120+ Figma components, reduced design-dev handoff time by 60%.",
  },
  {
    year: "2020-22",
    role: "FRONTEND DEVELOPER",
    org: "PLUMA STUDIO",
    c: "pink",
    body: "Built production React + Tailwind interfaces for fintech and creator-economy clients. Obsessed over micro-interactions.",
  },
  {
    year: "2018-20",
    role: "JUNIOR DESIGNER",
    org: "BYTECRAFT",
    c: "green",
    body: "Cut my teeth on landing pages, brand systems, and the occasional game jam. Where the pixel love started.",
  },
];

const stack = [
  {
    label: "DESIGN",
    items: ["FIGMA", "FRAMER", "PROTOPIE", "BLENDER", "ASEPRITE"],
    c: "purple",
  },
  {
    label: "BUILD",
    items: ["REACT", "TYPESCRIPT", "TAILWIND", "NEXT.JS", "THREE.JS"],
    c: "cyan",
  },
  {
    label: "CRAFT",
    items: ["MOTION", "TYPOGRAPHY", "COLOR", "PIXEL ART", "SOUND"],
    c: "pink",
  },
];

export default function LandingPage() {
  return (
    <div className="public-shell" data-testid="public-shell">
      {/* ========== NAVBAR ========== */}
      <PublicNav />

      {/* ========== HERO ========== */}
      <section className="hero" data-testid="hero">
        <div className="hero-grid">
          <div className="hero-left">
            <div className="hero-tag">
              <span className="blinker" /> {profile.availability}
            </div>

            <h1 className="hero-title">
              <span className="text-text">{profile.headline.prefix}</span>{" "}
              <span className="text-purple">{profile.headline.name}</span>.
              <br />
              {profile.headline.roles.map((role, i) => (
                <span key={role.text}>
                  <span className={`text-${role.color}`}>{role.text}</span>
                  {i < profile.headline.roles.length - 1 && (
                    <>
                      {" "}
                      <span className="text-text">&amp;</span>{" "}
                    </>
                  )}
                </span>
              ))}
              .
            </h1>

            <p className="hero-sub">{profile.shortIntro}</p>

            <div className="flex gap-3 flex-wrap mt-2">
              <a
                href="#contact"
                className="pix-btn pix-btn-pink hover-wiggle"
                data-testid="hero-cta-hire"
              >
                {profile.primaryCta} ~ &gt;
              </a>
              <a
                href="#timeline"
                className="pix-btn pix-btn-ghost hover-wiggle"
                data-testid="hero-cta-resume"
              >
                VIEW RESUME
              </a>
              <a
                href="/admin"
                className="pix-btn pix-btn-cyan hover-wiggle"
                data-testid="hero-cta-admin"
              >
                <IconGrid /> ADMIN CMS
              </a>
            </div>

            <div className="hero-socials">
              <span>FIND ME ~</span>
              {Object.entries(profile.socialLinks).map(
                ([key, link], i, arr) => (
                  <span key={key}>
                    <a href={link.url} data-testid={`social-${key.slice(0, 2)}`}>
                      {link.label}
                    </a>
                    {i < arr.length - 1 && <span> / </span>}
                  </span>
                ),
              )}
            </div>
          </div>

          <div className="hero-right">
            <div className="avatar-frame">
              <div className="avatar-inner">
                <HeroAvatar />
              </div>
              <div className="avatar-nameplate">
                <div
                  className="text-purple"
                  style={{
                    fontFamily: "var(--font-pixel)",
                    fontSize: "0.54rem",
                  }}
                >
                  {profile.fullName}
                </div>
                <div
                  className="text-cyan leading-[1.1]"
                  style={{
                    fontFamily: "var(--font-terminal)",
                    fontSize: "0.9rem",
                  }}
                >
                  LVL {profile.level.current} · DESIGNER / DEV
                </div>
                <div className="mt-[6px]">
                  <div className="xp-bar">
                    <div
                      className="xp-bar-fill"
                      style={{ width: `${profile.level.xpPercent}%` }}
                    />
                  </div>
                  <div
                    className="mt-[3px] text-line"
                    style={{
                      fontFamily: "var(--font-terminal)",
                      fontSize: "0.8rem",
                    }}
                  >
                    HP {profile.level.hp} · MP {profile.level.mp}
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
          {profile.stats.map((s) => (
            <HighlightCard
              key={s.key}
              label={s.key}
              value={s.value}
              sub={s.sub}
              color={s.color}
              testId={`highlight-${s.key.replace(/\s+/g, "-").toLowerCase()}`}
            />
          ))}
        </div>
      </section>

      {/* ========== ABOUT ========== */}
      <section className="section" id="about" data-testid="section-about">
        <SectionHeader kicker="~ CHAPTER 01" title="ABOUT.TXT" />

        <div className="about-grid">
          <div className="card accent-cyan">
            <div
              className="m-0 text-text leading-[1.35]"
              style={{ fontFamily: "var(--font-terminal)", fontSize: "1.1rem" }}
            >
              {profile.about.split("\n\n").map((para, i) => (
                <p key={i} className="m-0" style={{ marginBottom: i < profile.about.split("\n\n").length - 1 ? "1em" : 0 }}>
                  {para}
                </p>
              ))}
            </div>
          </div>

          <div className="stack">
            <div className="card accent-purple">
              <div className="card-title mb-[10px]">&gt; CURRENT QUEST</div>
              {profile.focus.map((item, i) => (
                <p
                  key={i}
                  className="m-0 text-muted leading-[1.3]"
                  style={{
                    fontFamily: "var(--font-terminal)",
                    fontSize: "1rem",
                    marginBottom: i < profile.focus.length - 1 ? "0.5em" : 0,
                  }}
                >
                  {item}
                </p>
              ))}
            </div>

            <div className="card accent-pink">
              <div className="card-title mb-[10px]">&gt; OBSESSIONS</div>
              <div className="flex flex-wrap gap-[6px]">
                {profile.obsessions.map((tag, i) => {
                  const tagColors = ["tag-purple", "", "tag-pink", "tag-green", "tag-yellow"];
                  return (
                    <span key={tag} className={`tag ${tagColors[i % tagColors.length]}`}>
                      {tag}
                    </span>
                  );
                })}
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* ========== TIMELINE ========== */}
      <section
        className="section"
        id="timeline"
        data-testid="section-timeline"
      >
        <SectionHeader kicker="~ CHAPTER 02" title="JOURNEY.LOG">
          <button
            className="pix-btn pix-btn-ghost hover-wiggle"
            data-testid="download-cv"
          >
            ↓ DOWNLOAD CV
          </button>
        </SectionHeader>

        <div className="timeline">
          {timeline.map((t, i) => (
            <TimelineEntry
              key={i}
              year={t.year}
              role={t.role}
              org={t.org}
              body={t.body}
              color={t.c}
              index={i}
            />
          ))}
        </div>
      </section>

      {/* ========== STACK ========== */}
      <section className="section" id="stack" data-testid="section-stack">
        <SectionHeader kicker="~ CHAPTER 03" title="INVENTORY" />

        <div className="grid-3">
          {stack.map((s) => (
            <div
              key={s.label}
              className={`card accent-${s.c}`}
              data-testid={`stack-${s.label.toLowerCase()}`}
            >
              <div className={`card-title mb-[14px] text-${s.c}`}>
                &gt; {s.label}
              </div>
              <div className="flex flex-col gap-[10px]">
                {s.items.map((it) => (
                  <div key={it} className="flex items-center gap-[10px]">
                    <span
                      className={`w-[10px] h-[10px] bg-${s.c}`}
                      style={{
                        boxShadow:
                          "inset 0 2px 0 rgba(255,255,255,0.3), inset 0 -2px 0 rgba(0,0,0,0.3)",
                      }}
                    />
                    <span
                      className="text-text"
                      style={{
                        fontFamily: "var(--font-pixel)",
                        fontSize: "0.54rem",
                      }}
                    >
                      {it}
                    </span>
                  </div>
                ))}
              </div>
            </div>
          ))}
        </div>
      </section>

      {/* ========== EXPLORE ========== */}
      <section className="section" id="explore" data-testid="section-explore">
        <SectionHeader
          kicker="~ CHAPTER 04"
          title="EXPLORE.~"
          subtitle="Dedicated pages below are still loading in the next update."
        />

        <div className="grid-3">
          <PlaceholderCard
            testId="explore-projects"
            to="#"
            color="purple"
            title="&gt; PROJECTS"
            desc="Case studies, UI kits, and shipped products. Deep-dives with process shots, code snippets & lessons learned."
            icon={<IconFolder />}
          />
          <PlaceholderCard
            testId="explore-skills"
            to="#"
            color="cyan"
            title="&gt; SKILLS"
            desc="Detailed breakdown of tools, levels, and the projects that taught me each one. Character sheet, but real."
            icon={<IconStar />}
          />
          <PlaceholderCard
            testId="explore-posts"
            to="#"
            color="pink"
            title="&gt; WRITING"
            desc="Devlogs, opinions, and the occasional rant. Design systems, motion, and the state of retro UI in 2026."
            icon={<IconDoc />}
          />
        </div>
      </section>

      {/* ========== CONTACT ========== */}
      <section className="section" id="contact" data-testid="section-contact">
        <div className="card accent-green contact-card">
          <Sword />
          <h2
            className="mt-4 mb-[10px] text-green"
            style={{ fontFamily: "var(--font-pixel)", fontSize: "1.35rem" }}
          >
            READY PLAYER ONE?
          </h2>
          <p
            className="mx-auto mb-5 text-muted text-center max-w-[560px] leading-[1.35]"
            style={{ fontFamily: "var(--font-terminal)", fontSize: "1.1rem" }}
          >
            Got a product, studio, or weird side-quest that needs a
            design-engineer hybrid? Let's talk. I reply within 24 hours —
            promise.
          </p>
          <div className="flex gap-3 justify-center flex-wrap">
            <a
              href={`mailto:${profile.email}`}
              className="pix-btn pix-btn-green hover-wiggle"
              data-testid="contact-email"
            >
              <IconMail /> {profile.email.toUpperCase()}
            </a>
            <a
              href="#"
              className="pix-btn pix-btn-ghost hover-wiggle"
              data-testid="contact-schedule"
            >
              ▸ BOOK A CALL
            </a>
          </div>
        </div>
      </section>

      {/* ========== FOOTER ========== */}
      <PublicFooter ownerName={profile.fullName} />
    </div>
  );
}
