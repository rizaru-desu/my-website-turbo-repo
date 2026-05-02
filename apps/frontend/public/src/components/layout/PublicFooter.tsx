interface PublicFooterProps {
  ownerName?: string;
  adminHref?: string;
  socials?: Array<{ label: string; href: string; color: string }>;
}

const defaultSocials = [
  { label: "GH", href: "#", color: "text-cyan" },
  { label: "DR", href: "#", color: "text-pink" },
  { label: "FG", href: "#", color: "text-purple" },
  { label: "TW", href: "#", color: "text-green" },
];

export default function PublicFooter({
  ownerName = "RIZARU.DESU",
  adminHref = "/admin",
  socials = defaultSocials,
}: PublicFooterProps) {
  return (
    <footer className="public-footer" data-testid="public-footer">
      <div className="flex justify-between items-center flex-wrap gap-3">
        <div
          className="text-line tracking-wider"
          style={{ fontFamily: "var(--font-pixel)", fontSize: "0.48rem" }}
        >
          © {new Date().getFullYear()} {ownerName} · ALL PIXELS RESERVED
        </div>
        <div className="flex gap-[18px]" style={{ fontFamily: "var(--font-pixel)", fontSize: "0.48rem" }}>
          {socials.map((s) => (
            <a
              key={s.label}
              href={s.href}
              className={`${s.color} no-underline`}
            >
              {s.label}
            </a>
          ))}
          <a
            href={adminHref}
            className="text-line no-underline"
            data-testid="footer-admin"
          >
            ADMIN
          </a>
        </div>
      </div>
      <div
        className="mt-[10px] text-line"
        style={{ fontFamily: "var(--font-terminal)", fontSize: "0.85rem" }}
      >
        PRESS <span className="text-cyan">[START]</span> TO CONTINUE ~ BUILT
        WITH REACT, TAILWIND &amp; LOVE FOR THE GRID
      </div>
    </footer>
  );
}
