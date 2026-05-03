import { useState, useEffect, useCallback } from "react";
import { Link } from "react-router-dom";
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
  TestimonialCard,
  CertificateCard,
  ProjectCard,
  BlogCard,
} from "../components";
import { IconMedal } from "../components/icons";

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
  phone: "+81 75-XXX-XXXX",
  availability: "AVAILABLE FOR CONTRACT · Q2 2026",
  primaryCta: "HIRE ME",
  profilePhotoUrl: null,
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
    period: "2025 — PRESENT",
    role: "PRINCIPAL DESIGNER",
    company: "FREELANCE",
    location: "TOKYO, JP",
    summary: "Leading end-to-end product design & frontend for early-stage SaaS and indie studios. Specialty: retro-flavored modern UIs.",
    achievements: [
      "Designed & built a pixel-art design system used across 3 active projects.",
      "Reduced time-to-market for MVP prototypes by 40% using custom React components.",
    ],
    sortOrder: 0,
  },
  {
    period: "2022 — 2025",
    role: "SENIOR UI/UX DESIGNER",
    company: "NEBULA LABS",
    location: "REMOTE",
    summary: "Owned the design system powering 5 B2B products. Shipped 120+ Figma components, reduced design-dev handoff time by 60%.",
    achievements: [
      "Led the redesign of the core dashboard, increasing user retention by 25%.",
      "Mentored 4 junior designers in atomic design principles and accessibility.",
    ],
    sortOrder: 1,
  },
  {
    period: "2020 — 2022",
    role: "FRONTEND DEVELOPER",
    company: "PLUMA STUDIO",
    location: "KYOTO, JP",
    summary: "Built production React + Tailwind interfaces for fintech and creator-economy clients. Obsessed over micro-interactions.",
    achievements: [
      "Optimized frontend performance, achieving a 95+ Lighthouse score across all client sites.",
      "Implemented complex data visualizations using D3.js and SVG animations.",
    ],
    sortOrder: 2,
  },
  {
    period: "2018 — 2020",
    role: "JUNIOR DESIGNER",
    company: "BYTECRAFT",
    location: "OSAKA, JP",
    summary: "Cut my teeth on landing pages, brand systems, and the occasional game jam. Where the pixel love started.",
    achievements: [
      "Assisted in the branding of 10+ startup clients.",
      "Winner of the 2019 Internal Game Jam (UI/UX category).",
    ],
    sortOrder: 3,
  },
];

/* ─── Education data (maps to education DB table) ─── */
const education = [
  {
    period: "2014 — 2018",
    degree: "BACHELOR OF DESIGN",
    school: "KYOTO UNIVERSITY OF THE ARTS",
    location: "KYOTO, JP",
    description: "Specialized in Interaction Design and Visual Communication. Explored the synergy between traditional Japanese aesthetics and modern digital interfaces.",
    highlights: [
      "Dean's List for Academic Excellence (2016, 2017).",
      "Lead Designer for the 'Neo-Kyoto' experimental web project.",
      "Thesis on 'Retro-Futurism in Modern UI/UX'."
    ],
    sortOrder: 0,
  },
];

/* ─── Projects data (maps to project DB table) ─── */
const projects = [
  {
    title: "NEO-KYOTO PROTOCOL",
    slug: "neo-kyoto-protocol",
    summary: "A decentralized content platform for independent writers in Kyoto, featuring a retro-future aesthetic and blockchain-backed ownership.",
    category: "WEB APP",
    year: "2026",
    tags: ["REACT", "SOLIDITY", "TAILWIND"],
    projectUrl: "#",
    githubUrl: "#",
    featured: true,
  },
  {
    title: "PIXEL-DASHBOARD",
    slug: "pixel-dashboard",
    summary: "A high-performance SaaS dashboard for developer metrics, using pixel-art primitives and CRT-style rendering for a unique UI experience.",
    category: "DASHBOARD",
    year: "2025",
    tags: ["TYPESCRIPT", "THREE.JS", "D3.JS"],
    projectUrl: "#",
    githubUrl: "#",
    featured: true,
  },
  {
    title: "SYNTH-WAVE MUSIC PLAYER",
    slug: "synth-wave-music-player",
    summary: "A web-based music player with interactive WebGL visualizations that react to frequencies, styled with 80s neon aesthetics.",
    category: "MULTIMEDIA",
    year: "2024",
    tags: ["WEB AUDIO API", "WEBGL", "REACT"],
    projectUrl: "#",
    githubUrl: "#",
    featured: false,
  },
];

/* ─── Blog data (maps to blogPost DB table) ─── */
const blogPosts = [
  {
    title: "THE ART OF PIXEL-PERFECT DESIGN SYSTEMS",
    slug: "art-of-pixel-perfect-design-systems",
    excerpt: "Exploring why the 8-bit aesthetic is making a comeback in modern SaaS interfaces and how to build one that scales.",
    category: "DESIGN",
    publishDate: "MAY 01, 2026",
    readingTime: "5 MIN READ",
    featured: true,
  },
  {
    title: "WHY I STILL USE VANILLA CSS IN 2026",
    slug: "why-vanilla-css-in-2026",
    excerpt: "Tailwind is great, but there's a certain magic in writing raw CSS variables and utility classes that you fully own.",
    category: "CODE",
    publishDate: "APR 20, 2026",
    readingTime: "8 MIN READ",
    featured: true,
  },
  {
    title: "RETRO-FUTURISM: A UX PERSPECTIVE",
    slug: "retro-futurism-ux-perspective",
    excerpt: "How nostalgic elements can improve user engagement by creating an emotional connection to digital tools.",
    category: "UX RESEARCH",
    publishDate: "APR 10, 2026",
    readingTime: "6 MIN READ",
    featured: true,
  },
];

/* ─── Skills data (maps to skill DB table) ─── */
const skills = [
  // DESIGN
  { name: "FIGMA", category: "DESIGN", level: "EXPERT", featured: true },
  { name: "FRAMER", category: "DESIGN", level: "ADVANCED", featured: true },
  { name: "PROTOPIE", category: "DESIGN", level: "INTERMEDIATE", featured: false },
  { name: "BLENDER", category: "DESIGN", level: "INTERMEDIATE", featured: false },
  { name: "ASEPRITE", category: "DESIGN", level: "ADVANCED", featured: true },
  // BUILD
  { name: "REACT", category: "BUILD", level: "EXPERT", featured: true },
  { name: "TYPESCRIPT", category: "BUILD", level: "EXPERT", featured: true },
  { name: "TAILWIND", category: "BUILD", level: "EXPERT", featured: true },
  { name: "NEXT.JS", category: "BUILD", level: "ADVANCED", featured: true },
  { name: "THREE.JS", category: "BUILD", level: "INTERMEDIATE", featured: false },
  // CRAFT
  { name: "MOTION", category: "CRAFT", level: "ADVANCED", featured: true },
  { name: "TYPOGRAPHY", category: "CRAFT", level: "EXPERT", featured: true },
  { name: "COLOR", category: "CRAFT", level: "EXPERT", featured: true },
  { name: "PIXEL ART", category: "CRAFT", level: "EXPERT", featured: true },
  { name: "SOUND", category: "CRAFT", level: "INTERMEDIATE", featured: false },
];

/* ─── Certificates data (maps to certificate DB table) ─── */
const certificates = [
  {
    name: "ADVANCED REACT PATTERNS",
    issuer: "FRONTEND MASTERS",
    year: "2024",
    link: "#",
    featured: true,
  },
  {
    name: "UI DESIGN SPECIALIZATION",
    issuer: "COURSERA (CALARTS)",
    year: "2023",
    link: "#",
    featured: true,
  },
  {
    name: "TYPESCRIPT EXPERT",
    issuer: "TOTAL TYPESCRIPT",
    year: "2023",
    link: "#",
    featured: false,
  },
  {
    name: "PIXEL ART MASTERCLASS",
    issuer: "UDEMY",
    year: "2022",
    link: "#",
    featured: false,
  },
];

/* ─── Testimonials data (maps to testimonial DB table) ─── */
const testimonials = [
  {
    name: "SARAH CHEN",
    role: "PRODUCT LEAD",
    company: "NEBULA LABS",
    message:
      "Kai has a rare eye for detail. He didn't just build our design system; he gave it a soul. The retro-modern aesthetic he proposed became our most loved feature.",
    rating: 5,
    relation: "COLLEAGUE",
  },
  {
    name: "MARCUS JAEGER",
    role: "FOUNDER",
    company: "BYTECRAFT",
    message:
      "Working with Kai on our landing page was seamless. He bridges the gap between design and code perfectly. Highly recommended for any high-stakes UI work.",
    rating: 5,
    relation: "CLIENT",
  },
  {
    name: "YUKI TANAKA",
    role: "CREATIVE DIRECTOR",
    company: "PLUMA STUDIO",
    message:
      "The pixel-perfect implementation and attention to micro-interactions Kai brought to our project were outstanding. A true design-engineer hybrid.",
    rating: 5,
    relation: "CLIENT",
  },
  {
    name: "ALEX RIVERA",
    role: "CTO",
    company: "SYNTH WAVE",
    message:
      "Kai delivered a complex dashboard that managed to be both highly functional and visually stunning. His grasp of frontend architecture is top-tier.",
    rating: 5,
    relation: "CLIENT",
  },
  {
    name: "ELENA ROSSI",
    role: "SENIOR DESIGNER",
    company: "PIXEL PERFECT",
    message:
      "I've worked with many developers, but Kai is one of the few who truly 'gets' design. He never compromises on the vision while keeping the code clean.",
    rating: 5,
    relation: "COLLEAGUE",
  },
  {
    name: "DAVID PARK",
    role: "MARKETING HEAD",
    company: "GLITCH INC",
    message:
      "Our conversion rates jumped by 40% after the redesign Kai implemented. He knows exactly how to make a brand feel modern yet nostalgic.",
    rating: 4,
    relation: "CLIENT",
  },
  {
    name: "MIA WONG",
    role: "INDIE DEVELOPER",
    company: "SELF-EMPLOYED",
    message:
      "Kai mentored me through my first React project. His ability to explain complex concepts through the lens of game design was a game-changer for me.",
    rating: 5,
    relation: "MENTOR",
  },
  {
    name: "JORDAN SMITH",
    role: "PRODUCT MANAGER",
    company: "ARCADE CORE",
    message:
      "The attention to detail in the animations and transitions was beyond what we expected. Kai truly lives and breathes the pixel-art aesthetic.",
    rating: 5,
    relation: "CLIENT",
  },
  {
    name: "LIAM O'CONNOR",
    role: "FRONTEND LEAD",
    company: "TECH FORGE",
    message:
      "One of the most reliable engineers I've had the pleasure of working with. Kai's technical skills are only matched by his creative problem-solving.",
    rating: 5,
    relation: "COLLEAGUE",
  },
  {
    name: "SOPHIE MARTIN",
    role: "UI ARTIST",
    company: "DREAM SCAPE",
    message:
      "Kai turned my static designs into a living, breathing interface that exceeded all expectations. His passion for the craft is evident in every commit.",
    rating: 5,
    relation: "COLLEAGUE",
  },
];

const COLORS = ["purple", "cyan", "pink", "green", "orange", "yellow"];

export default function LandingPage() {
  const [currentPage, setCurrentPage] = useState(0);
  const [isPaused, setIsPaused] = useState(false);
  const [mousePos, setMousePos] = useState({ x: 0, y: 0 });
  const [scrollY, setScrollY] = useState(0);

  const itemsPerPage = 3;
  const totalPages = Math.ceil(testimonials.length / itemsPerPage);

  const nextTestimonial = useCallback(() => {
    setCurrentPage((prev) => (prev + 1) % totalPages);
  }, [totalPages]);

  const prevTestimonial = useCallback(() => {
    setCurrentPage((prev) => (prev - 1 + totalPages) % totalPages);
  }, [totalPages]);

  // Group skills by category for display
  const skillCategories = Array.from(new Set(skills.map(s => s.category))).map((cat, i) => ({
    label: cat,
    items: skills.filter(s => s.category === cat),
    color: COLORS[i % COLORS.length]
  }));

  // Parallax listeners
  useEffect(() => {
    const handleMouseMove = (e: MouseEvent) => {
      setMousePos({
        x: (e.clientX / window.innerWidth - 0.5) * 20,
        y: (e.clientY / window.innerHeight - 0.5) * 20,
      });
    };

    const handleScroll = () => {
      setScrollY(window.scrollY);
    };

    window.addEventListener("mousemove", handleMouseMove);
    window.addEventListener("scroll", handleScroll, { passive: true });

    return () => {
      window.removeEventListener("mousemove", handleMouseMove);
      window.removeEventListener("scroll", handleScroll);
    };
  }, []);

  // Auto-scroll logic
  useEffect(() => {
    if (isPaused) return;

    const interval = setInterval(() => {
      nextTestimonial();
    }, 5000); // Change every 5 seconds

    return () => clearInterval(interval);
  }, [isPaused, nextTestimonial]);

  const currentTestimonials = testimonials.slice(
    currentPage * itemsPerPage,
    (currentPage + 1) * itemsPerPage
  );

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
            <div 
              className="avatar-frame transition-transform duration-300 ease-out"
              style={{ transform: `translate(${mousePos.x * -0.5}px, ${mousePos.y * -0.5}px)` }}
            >
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

            {/* Floating pixel coins with Mouse + Scroll Parallax */}
            <div 
              className="coin c1 bob transition-transform duration-200 ease-out" 
              style={{ transform: `translate(${mousePos.x * 1.5}px, ${mousePos.y * 1.5 - scrollY * 0.1}px)` }}
            />
            <div 
              className="coin c2 bob transition-transform duration-300 ease-out" 
              style={{ transform: `translate(${mousePos.x * -2}px, ${mousePos.y * -2 - scrollY * 0.15}px)` }}
            />
            <div 
              className="coin c3 bob transition-transform duration-500 ease-out" 
              style={{ transform: `translate(${mousePos.x * 1}px, ${mousePos.y * 1 - scrollY * 0.05}px)` }}
            />
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

      {/* ========== PROJECTS ========== */}
      <section className="section" id="projects" data-testid="section-projects">
        <SectionHeader kicker="~ CHAPTER 02" title="QUESTS.LOG" />

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {projects.map((p, i) => (
            <ProjectCard
              key={p.slug}
              title={p.title}
              summary={p.summary}
              category={p.category}
              year={p.year}
              tags={p.tags}
              projectUrl={p.projectUrl}
              githubUrl={p.githubUrl}
              color={COLORS[i % COLORS.length]}
              index={i}
            />
          ))}
        </div>
      </section>

      {/* ========== TIMELINE ========== */}
      <section
        className="section"
        id="timeline"
        data-testid="section-timeline"
      >
        <SectionHeader kicker="~ CHAPTER 03" title="JOURNEY.LOG">
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
              period={t.period}
              role={t.role}
              company={t.company}
              location={t.location}
              summary={t.summary}
              achievements={t.achievements}
              color={COLORS[i % COLORS.length]}
              index={i}
            />
          ))}
        </div>
      </section>

      {/* ========== EDUCATION ========== */}
      <section
        className="section"
        id="education"
        data-testid="section-education"
      >
        <SectionHeader kicker="~ CHAPTER 04" title="FORMATION.LOG" />

        <div className="timeline">
          {education.map((e, i) => (
            <TimelineEntry
              key={i}
              period={e.period}
              role={e.degree}
              company={e.school}
              location={e.location}
              summary={e.description}
              achievements={e.highlights}
              color={COLORS[(i + timeline.length) % COLORS.length]}
              index={i + timeline.length}
            />
          ))}
        </div>
      </section>

      {/* ========== STACK ========== */}
      <section className="section" id="stack" data-testid="section-stack">
        <SectionHeader kicker="~ CHAPTER 05" title="INVENTORY" />

        <div className="grid-3">
          {skillCategories.map((cat) => (
            <div
              key={cat.label}
              className={`card accent-${cat.color}`}
              data-testid={`stack-${cat.label.toLowerCase()}`}
            >
              <div className={`card-title mb-[14px] text-${cat.color}`}>
                &gt; {cat.label}
              </div>
              <div className="flex flex-col gap-[10px]">
                {cat.items.map((skill) => (
                  <div key={skill.name} className="flex items-center gap-[10px]">
                    <span
                      className={`w-[10px] h-[10px] bg-${cat.color}`}
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
                      {skill.name}
                    </span>
                    {skill.featured && (
                      <span className="text-yellow" style={{ fontSize: "0.5rem" }}>★</span>
                    )}
                  </div>
                ))}
              </div>
            </div>
          ))}
        </div>
      </section>

      {/* ========== CERTIFICATES ========== */}
      <section className="section" id="certificates" data-testid="section-certificates">
        <SectionHeader kicker="~ CHAPTER 06" title="ACHIEVEMENTS.LOG" />

        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
          {certificates.map((cert, i) => (
            <CertificateCard
              key={cert.name}
              name={cert.name}
              issuer={cert.issuer}
              year={cert.year}
              link={cert.link}
              featured={cert.featured}
              color={COLORS[i % COLORS.length]}
            />
          ))}
        </div>
      </section>

      {/* ========== TESTIMONIALS ========== */}
      <section
        className="section"
        id="testimonials"
        data-testid="section-testimonials"
        onMouseEnter={() => setIsPaused(true)}
        onMouseLeave={() => setIsPaused(false)}
      >
        <SectionHeader kicker="~ CHAPTER 07" title="REVIEWS.LOG">
          <Link
            to="/testimonials/new"
            className="pix-btn pix-btn-ghost hover-wiggle"
            data-testid="leave-review"
          >
            + LEAVE A REVIEW
          </Link>
        </SectionHeader>
        
        <div className="relative">
          {/* Main Cards Grid */}
          <div 
            key={currentPage} 
            className="grid grid-cols-1 md:grid-cols-3 gap-7 min-h-[320px] animate-pixel-slide"
          >
            {currentTestimonials.map((t, i) => {
              const globalIndex = currentPage * itemsPerPage + i;
              return (
                <div key={globalIndex} className="flex flex-col h-full">
                  <TestimonialCard
                    name={t.name}
                    role={t.role}
                    company={t.company}
                    message={t.message}
                    rating={t.rating}
                    relation={t.relation}
                    color={COLORS[globalIndex % COLORS.length]}
                    index={globalIndex}
                  />
                </div>
              );
            })}
            {/* Empty placeholders to maintain grid if last page has fewer than 3 items */}
            {currentTestimonials.length < itemsPerPage && 
              [...Array(itemsPerPage - currentTestimonials.length)].map((_, i) => (
                <div key={`empty-${i}`} className="hidden md:block opacity-0 pointer-events-none">
                  <div className="card h-full" />
                </div>
              ))
            }
          </div>

          {/* Carousel Controls */}
          <div className="flex justify-between items-center mt-10 px-4">
            <button
              onClick={prevTestimonial}
              className="pix-btn pix-btn-ghost hover-wiggle"
              aria-label="Previous page"
            >
              &lt; PREV
            </button>

            <div className="flex gap-3">
              {[...Array(totalPages)].map((_, i) => (
                <button
                  key={i}
                  onClick={() => setCurrentPage(i)}
                  className={`w-3 h-3 ${
                    currentPage === i ? "bg-purple" : "bg-panel"
                  }`}
                  style={{
                    boxShadow: currentPage === i 
                      ? "inset 0 2px 0 rgba(255,255,255,0.3), 0 0 0 2px #282a36" 
                      : "inset 0 -2px 0 rgba(0,0,0,0.3), 0 0 0 2px #282a36",
                    transition: "background-color 0.2s steps(2)"
                  }}
                  aria-label={`Go to page ${i + 1}`}
                />
              ))}
            </div>

            <button
              onClick={nextTestimonial}
              className="pix-btn pix-btn-ghost hover-wiggle"
              aria-label="Next page"
            >
              NEXT &gt;
            </button>
          </div>

          {/* Page Counter */}
          <div 
            className="text-center mt-4 text-line"
            style={{ fontFamily: "var(--font-pixel)", fontSize: "0.48rem" }}
          >
            PAGE {currentPage + 1} OF {totalPages}
          </div>
        </div>
      </section>

      {/* ========== WRITING ========== */}
      <section className="section" id="writing" data-testid="section-writing">
        <SectionHeader kicker="~ CHAPTER 08" title="TRANSMISSIONS.LOG">
          <Link
            to="/blog"
            className="pix-btn pix-btn-ghost hover-wiggle"
            data-testid="view-all-posts"
          >
            VIEW ALL ENTRIES
          </Link>
        </SectionHeader>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          {blogPosts.slice(0, 3).map((post, i) => (
            <BlogCard
              key={post.slug}
              title={post.title}
              excerpt={post.excerpt}
              category={post.category}
              publishDate={post.publishDate}
              readingTime={post.readingTime}
              slug={post.slug}
              color={COLORS[(i + 3) % COLORS.length]}
              index={i}
            />
          ))}
        </div>

        <div className="mt-12 flex justify-center">
          <Link
            to="/blog"
            className="pix-btn hover-wiggle inline-flex items-center gap-2"
            style={{ backgroundColor: "var(--purple)", color: "white" }}
            data-testid="read-more-blog"
          >
            READ MORE TRANSMISSIONS »
          </Link>
        </div>
      </section>

      {/* ========== EXPLORE ========== */}
      <section className="section" id="explore" data-testid="section-explore">
        <SectionHeader
          kicker="~ CHAPTER 09"
          title="EXPLORE.~"
          subtitle="Dedicated pages below are still loading in the next update."
        />

        <div className="grid-3">
          <PlaceholderCard
            testId="explore-projects"
            to="/projects"
            color="purple"
            title="&gt; PROJECTS"
            desc="Case studies, UI kits, and shipped products. Deep-dives with process shots, code snippets & lessons learned."
            icon={<IconFolder />}
          />
          <PlaceholderCard
            testId="explore-skills"
            to="/skills"
            color="cyan"
            title="&gt; SKILLS"
            desc="Detailed breakdown of tools, levels, and the projects that taught me each one. Character sheet, but real."
            icon={<IconStar />}
          />
          <PlaceholderCard
            testId="explore-posts"
            to="/blog"
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
