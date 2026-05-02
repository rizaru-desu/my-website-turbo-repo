import type { ReactNode } from "react";
import { Link } from "react-router-dom";
import { Sword } from "../icons";

interface SectionHeaderProps {
  kicker: string;
  title: string;
  subtitle?: string;
  children?: ReactNode;
}

export function SectionHeader({
  kicker,
  title,
  subtitle,
  children,
}: SectionHeaderProps) {
  return (
    <div className="section-header">
      <div>
        <div className="section-kicker">{kicker}</div>
        <h2 className="section-title">{title}</h2>
        {subtitle && (
          <div
            className="mt-2 text-muted"
            style={{ fontFamily: "var(--font-terminal)", fontSize: "0.95rem" }}
          >
            {subtitle}
          </div>
        )}
      </div>
      {children ?? <Sword />}
    </div>
  );
}

interface HighlightCardProps {
  label: string;
  value: string;
  sub: string;
  color: string;
  testId?: string;
}

export function HighlightCard({
  label,
  value,
  sub,
  color,
  testId,
}: HighlightCardProps) {
  return (
    <div
      className={`card accent-${color} relative`}
      data-testid={testId}
    >
      <span
        className="block mb-[10px] text-line tracking-wider"
        style={{ fontFamily: "var(--font-pixel)", fontSize: "0.48rem" }}
      >
        {label}
      </span>
      <div
        className={`mb-[6px] text-${color}`}
        style={{ fontFamily: "var(--font-pixel)", fontSize: "1.35rem" }}
      >
        {value}
      </div>
      <div
        className="text-muted"
        style={{ fontFamily: "var(--font-terminal)", fontSize: "0.9rem" }}
      >
        {sub}
      </div>
    </div>
  );
}

interface PlaceholderCardProps {
  to: string;
  icon: ReactNode;
  title: string;
  desc: string;
  color?: string;
  testId?: string;
}

export function PlaceholderCard({
  to,
  icon,
  title,
  desc,
  color = "purple",
  testId,
}: PlaceholderCardProps) {
  return (
    <a
      href={to}
      className={`card accent-${color} placeholder-card`}
      data-testid={testId}
    >
      <div className="flex items-center gap-3 mb-3">
        {icon}
        <div className={`card-title text-${color}`}>{title}</div>
      </div>
      <p
        className="m-0 text-muted leading-[1.3]"
        style={{ fontFamily: "var(--font-terminal)", fontSize: "1rem" }}
      >
        {desc}
      </p>
      <div
        className="mt-[14px] flex justify-between items-center text-line"
        style={{ fontFamily: "var(--font-pixel)", fontSize: "0.48rem" }}
      >
        <span>SOON ~ COMING</span>
        <svg
          viewBox="0 0 16 16"
          width="12"
          height="12"
          shapeRendering="crispEdges"
          aria-hidden="true"
        >
          <rect x="8" y="2" width="6" height="2" fill="currentColor" />
          <rect x="12" y="2" width="2" height="6" fill="currentColor" />
          <rect x="7" y="7" width="1" height="1" fill="currentColor" />
          <rect x="8" y="6" width="1" height="1" fill="currentColor" />
          <rect x="9" y="5" width="1" height="1" fill="currentColor" />
          <rect x="10" y="4" width="1" height="1" fill="currentColor" />
        </svg>
      </div>
    </a>
  );
}

interface TimelineEntryProps {
  period: string;
  role: string;
  company: string;
  location: string;
  summary: string;
  achievements?: string[];
  color: string;
  index: number;
}

export function TimelineEntry({
  period,
  role,
  company,
  location,
  summary,
  achievements,
  color,
  index,
}: TimelineEntryProps) {
  return (
    <div className="timeline-row">
      <div className={`timeline-year text-${color}`}>{period}</div>
      <div className="timeline-dot" style={{ background: `var(--color-${color})` }} />
      <div
        className={`card accent-${color} timeline-card`}
        data-testid={`timeline-${index}`}
      >
        <div className="flex justify-between items-start flex-wrap gap-2 mb-2">
          <div>
            <div className={`card-title text-${color}`}>{role}</div>
            <div
              className="text-muted mt-1"
              style={{ fontFamily: "var(--font-terminal)", fontSize: "0.85rem" }}
            >
              {location}
            </div>
          </div>
          <span className="pixel-badge text-line">{company}</span>
        </div>
        <p
          className="m-0 text-muted leading-[1.3] mb-3"
          style={{ fontFamily: "var(--font-terminal)", fontSize: "1rem" }}
        >
          {summary}
        </p>
        {achievements && achievements.length > 0 && (
          <div className="flex flex-col gap-2">
            {achievements.map((achievement, i) => (
              <div key={i} className="flex items-start gap-2">
                <span className={`text-${color} mt-1`} style={{ fontFamily: "var(--font-pixel)", fontSize: "0.5rem" }}>»</span>
                <span
                  className="text-text leading-tight"
                  style={{ fontFamily: "var(--font-terminal)", fontSize: "0.9rem" }}
                >
                  {achievement}
                </span>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}

interface CertificateCardProps {
  name: string;
  issuer: string;
  year: string;
  link: string;
  color: string;
  featured?: boolean;
}

export function CertificateCard({
  name,
  issuer,
  year,
  link,
  color,
  featured,
}: CertificateCardProps) {
  return (
    <div className={`card accent-${color} relative h-full flex flex-col`}>
      <div className="flex justify-between items-start mb-2">
        <span
          className={`text-${color} leading-none`}
          style={{ fontFamily: "var(--font-pixel)", fontSize: "0.48rem" }}
        >
          {year}
        </span>
        {featured && (
          <span className="text-yellow" style={{ fontSize: "0.8rem" }}>★</span>
        )}
      </div>
      <div
        className="card-title mb-1 leading-tight"
        style={{ fontSize: "1.1rem" }}
      >
        {name}
      </div>
      <div
        className="text-muted mb-4"
        style={{ fontFamily: "var(--font-terminal)", fontSize: "0.9rem" }}
      >
        {issuer}
      </div>
      <div className="mt-auto">
        <a
          href={link}
          target="_blank"
          rel="noopener noreferrer"
          className={`pix-btn pix-btn-ghost text-[0.48rem] py-1 px-2 hover-wiggle w-full block text-center`}
          style={{ fontFamily: "var(--font-pixel)" }}
        >
          VERIFY_CREDENTIAL
        </a>
      </div>
    </div>
  );
}

interface TestimonialCardProps {
  name: string;
  role: string;
  company?: string;
  message: string;
  rating: number;
  relation: string;
  color: string;
  index: number;
}

export function TestimonialCard({
  name,
  role,
  company,
  message,
  rating,
  relation,
  color,
  index,
}: TestimonialCardProps) {
  return (
    <div
      className={`card accent-${color} flex flex-col h-full`}
      data-testid={`testimonial-${index}`}
    >
      <div className="flex justify-between items-start mb-3">
        <div className="flex">
          {[...Array(5)].map((_, i) => (
            <span
              key={i}
              className={`text-${i < rating ? "yellow" : "line"} mr-1`}
              style={{ fontFamily: "var(--font-pixel)", fontSize: "0.8rem" }}
            >
              ★
            </span>
          ))}
        </div>
        <span className="pixel-badge text-line" style={{ fontSize: "0.4rem" }}>
          {relation}
        </span>
      </div>
      <p
        className="m-0 text-text leading-[1.4] italic mb-4 flex-grow"
        style={{ fontFamily: "var(--font-terminal)", fontSize: "1rem" }}
      >
        "{message}"
      </p>
      <div className="flex items-center gap-3">
        <div
          className={`w-10 h-10 bg-${color} flex-shrink-0`}
          style={{
            boxShadow:
              "inset 0 3px 0 rgba(255,255,255,0.3), inset 0 -3px 0 rgba(0,0,0,0.3), 0 0 0 3px #282a36",
          }}
        />
        <div>
          <div className={`card-title text-${color}`}>{name}</div>
          <div
            className="text-muted leading-tight"
            style={{ fontFamily: "var(--font-terminal)", fontSize: "0.85rem" }}
          >
            {role} {company && ` @ ${company}`}
          </div>
        </div>
      </div>
    </div>
  );
}

interface BlogCardProps {
  title: string;
  excerpt: string;
  category: string;
  publishDate: string;
  readingTime: string;
  slug: string;
  color: string;
  index: number;
}

export function BlogCard({
  title,
  excerpt,
  category,
  publishDate,
  readingTime,
  slug,
  color,
  index,
}: BlogCardProps) {
  return (
    <div
      className={`card accent-${color} flex flex-col h-full blog-card-hover`}
      data-testid={`blog-post-${index}`}
    >
      <div className="flex justify-between items-start mb-3">
        <span
          className={`text-${color}`}
          style={{ fontFamily: "var(--font-pixel)", fontSize: "0.48rem" }}
        >
          {publishDate} // {category}
        </span>
        <span
          className="text-line"
          style={{ fontFamily: "var(--font-pixel)", fontSize: "0.4rem" }}
        >
          {readingTime}
        </span>
      </div>

      <div
        className="card-title mb-2 text-text"
        style={{ fontSize: "1.2rem" }}
      >
        {title}
      </div>

      <p
        className="m-0 text-muted leading-[1.35] mb-5 flex-grow"
        style={{ fontFamily: "var(--font-terminal)", fontSize: "0.95rem" }}
      >
        {excerpt}
      </p>

      <Link
        to={`/blog/${slug}`}
        className={`pix-btn pix-btn-ghost text-[0.48rem] py-2 px-3 hover-wiggle inline-flex items-center gap-2 self-start`}
        style={{ fontFamily: "var(--font-pixel)" }}
      >
        READ_ENTRY <span style={{ fontSize: "0.6rem" }}>»</span>
      </Link>
    </div>
  );
}

interface ProjectCardProps {
  title: string;
  summary: string;
  category: string;
  year: string;
  tags: string[];
  projectUrl?: string;
  githubUrl?: string;
  color: string;
  index: number;
}

export function ProjectCard({
  title,
  summary,
  category,
  year,
  tags,
  projectUrl,
  githubUrl,
  color,
  index,
}: ProjectCardProps) {
  return (
    <div
      className={`card accent-${color} flex flex-col h-full project-card-hover`}
      data-testid={`project-${index}`}
    >
      <div className="flex justify-between items-start mb-3">
        <span
          className={`text-${color}`}
          style={{ fontFamily: "var(--font-pixel)", fontSize: "0.48rem" }}
        >
          {year} // {category}
        </span>
        <div className="flex gap-2">
          {githubUrl && (
            <a
              href={githubUrl}
              target="_blank"
              rel="noopener noreferrer"
              className="text-muted hover:text-text transition-colors"
              aria-label="GitHub Repository"
            >
              <svg viewBox="0 0 16 16" width="14" height="14" fill="currentColor">
                <path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z" />
              </svg>
            </a>
          )}
          {projectUrl && (
            <a
              href={projectUrl}
              target="_blank"
              rel="noopener noreferrer"
              className="text-muted hover:text-text transition-colors"
              aria-label="Live Demo"
            >
              <svg viewBox="0 0 16 16" width="14" height="14" fill="currentColor">
                <path d="M11 3H5a2 2 0 00-2 2v6a2 2 0 002 2h6a2 2 0 002-2V5a2 2 0 00-2-2zM5 4h6a1 1 0 011 1v6a1 1 0 01-1 1H5a1 1 0 01-1-1V5a1 1 0 011-1z" />
                <path d="M10.5 8.5l-3 2V6.5l3 2z" />
              </svg>
            </a>
          )}
        </div>
      </div>

      <div
        className="card-title mb-2 text-text"
        style={{ fontSize: "1.2rem" }}
      >
        {title}
      </div>

      <p
        className="m-0 text-muted leading-[1.35] mb-4 flex-grow"
        style={{ fontFamily: "var(--font-terminal)", fontSize: "0.95rem" }}
      >
        {summary}
      </p>

      <div className="flex flex-wrap gap-2 mt-auto">
        {tags.slice(0, 3).map((tag) => (
          <span
            key={tag}
            className="pixel-badge text-line"
            style={{ fontSize: "0.4rem", padding: "2px 6px" }}
          >
            {tag.toUpperCase()}
          </span>
        ))}
        {tags.length > 3 && (
          <span
            className="text-line"
            style={{ fontFamily: "var(--font-pixel)", fontSize: "0.48rem" }}
          >
            +{tags.length - 3}
          </span>
        )}
      </div>
    </div>
  );
}
