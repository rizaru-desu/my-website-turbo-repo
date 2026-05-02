import type { ReactNode } from "react";
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
