import { Link } from "react-router-dom";
import { NavBrandLogo, IconGrid } from "../icons";

interface PublicNavProps {
  brandName?: string;
  brandVersion?: string;
  links?: Array<{ href: string; label: string; testId?: string }>;
}

const defaultLinks = [
  { href: "#about", label: "ABOUT", testId: "nav-about" },
  { href: "#timeline", label: "JOURNEY", testId: "nav-timeline" },
  { href: "#testimonials", label: "REVIEWS", testId: "nav-testimonials" },
  { href: "#stack", label: "STACK", testId: "nav-stack" },
  { href: "#explore", label: "EXPLORE", testId: "nav-explore" },
  { href: "#contact", label: "CONTACT", testId: "nav-contact" },
];

export default function PublicNav({
  brandName = "RIZARU.DESU",
  brandVersion = "portfolio ~ v1.0",
  links = defaultLinks,
}: PublicNavProps) {
  return (
    <nav className="public-nav" data-testid="public-nav">
      <Link to="/" className="nav-brand">
        <NavBrandLogo />
        <div>
          <div
            className="text-purple leading-[1.4]"
            style={{ fontFamily: "var(--font-pixel)", fontSize: "0.56rem" }}
          >
            {brandName}
          </div>
          <div
            className="text-cyan leading-none"
            style={{ fontFamily: "var(--font-terminal)", fontSize: "0.9rem" }}
          >
            {brandVersion}
          </div>
        </div>
      </Link>

      <div className="nav-links">
        {links.map((link) => (
          <a key={link.href} href={link.href} data-testid={link.testId}>
            {link.label}
          </a>
        ))}
      </div>

      <div className="nav-actions">
        <a
          href="/admin"
          className="pix-btn pix-btn-ghost hover-wiggle"
          data-testid="nav-admin"
        >
          <IconGrid /> ADMIN
        </a>
      </div>
    </nav>
  );
}
