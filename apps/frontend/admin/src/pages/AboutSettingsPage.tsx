import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { AppLayout } from "../components/layout";
import { navItems, pageMeta } from "../constants/navigation";
import { useAuth } from "../hooks/useAuth";
import type { PageKey } from "../types/dashboard";

/* ─── Mock profile data (will be replaced by API calls) ─── */
const MOCK_PROFILE = {
  fullName: "KAI MORIKAWA",
  headline: "PIXEL-PUSHER & CODE-SMITH",
  shortIntro:
    "UI/UX designer and frontend developer from Kyoto, crafting characterful, retro-flavored interfaces for startups, studios and the occasional weekend game jam.",
  about:
    "I'm Kai — half designer, half engineer, fully obsessed with the quiet details that make interfaces feel alive.\n\nI've spent the last 8 years shipping product UIs for SaaS tools, dashboards, and content-heavy apps — always chasing the balance between legibility and personality.\n\nLately I've been exploring how retro aesthetics (pixel grids, CRT textures, arcade typography) can make modern dashboards genuinely delightful to return to, without sacrificing usability.",
  location: "KYOTO, JP",
  email: "kai@morikawa.studio",
  phone: "+81 90-XXXX-XXXX",
  availability: "AVAILABLE FOR CONTRACT · Q2 2026",
  primaryCta: "HIRE ME",
  profilePhotoUrl: "",
  socialLinks: {
    github: "https://github.com",
    dribbble: "https://dribbble.com",
    figma: "https://figma.com",
    twitter: "https://twitter.com",
  },
  focus: [
    "Designing a retro-flavored content platform for indie writers. Shipping in Q2.",
  ],
  stats: [
    { key: "YEARS XP", value: "8+", sub: "since 2018" },
    { key: "PROJECTS", value: "42", sub: "shipped & live" },
    { key: "HAPPY CLIENTS", value: "28", sub: "worldwide" },
    { key: "COFFEES /DAY", value: "∞", sub: "house blend" },
  ],
};

const MOCK_TAGS = ["PIXEL ART", "SYNTHWAVE", "KEYBOARDS", "GAME UI", "TYPE"];

/* ─── Stat type for the highlight strip ─── */
interface StatItem {
  key: string;
  value: string;
  sub: string;
}

export default function AboutSettingsPage() {
  const navigate = useNavigate();
  const { logout } = useAuth();

  /* ─── Form state (local for UI-only) ─── */
  const [fullName, setFullName] = useState(MOCK_PROFILE.fullName);
  const [headline, setHeadline] = useState(MOCK_PROFILE.headline);
  const [shortIntro, setShortIntro] = useState(MOCK_PROFILE.shortIntro);
  const [about, setAbout] = useState(MOCK_PROFILE.about);
  const [location, setLocation] = useState(MOCK_PROFILE.location);
  const [contactEmail, setContactEmail] = useState(MOCK_PROFILE.email);
  const [phone, setPhone] = useState(MOCK_PROFILE.phone);
  const [availability, setAvailability] = useState(MOCK_PROFILE.availability);
  const [primaryCta, setPrimaryCta] = useState(MOCK_PROFILE.primaryCta);
  const [profilePhotoUrl, setProfilePhotoUrl] = useState(
    MOCK_PROFILE.profilePhotoUrl,
  );

  /* Social links */
  const [github, setGithub] = useState(MOCK_PROFILE.socialLinks.github);
  const [dribbble, setDribbble] = useState(MOCK_PROFILE.socialLinks.dribbble);
  const [figma, setFigma] = useState(MOCK_PROFILE.socialLinks.figma);
  const [twitter, setTwitter] = useState(MOCK_PROFILE.socialLinks.twitter);

  /* Focus / current quest */
  const [focusItems, setFocusItems] = useState<string[]>(MOCK_PROFILE.focus);
  const [newFocus, setNewFocus] = useState("");

  /* Highlight stats */
  const [stats, setStats] = useState<StatItem[]>(MOCK_PROFILE.stats);
  const [newStatKey, setNewStatKey] = useState("");
  const [newStatValue, setNewStatValue] = useState("");
  const [newStatSub, setNewStatSub] = useState("");

  /* Tags / obsessions */
  const [tags, setTags] = useState<string[]>(MOCK_TAGS);
  const [newTag, setNewTag] = useState("");

  const handleLogout = async () => {
    await logout();
    navigate("/login");
  };

  const handleNavigate = (newPage: PageKey) => {
    if (newPage === "about") return;
    if (newPage === "settings") {
      navigate("/settings");
    } else {
      navigate("/dashboard", { state: { initialPage: newPage } });
    }
  };

  /* Focus helpers */
  const addFocus = () => {
    const trimmed = newFocus.trim();
    if (trimmed) {
      setFocusItems((prev) => [...prev, trimmed]);
      setNewFocus("");
    }
  };

  const removeFocus = (index: number) => {
    setFocusItems((prev) => prev.filter((_, i) => i !== index));
  };

  /* Stats helpers */
  const addStat = () => {
    const k = newStatKey.trim();
    const v = newStatValue.trim();
    if (k && v) {
      setStats((prev) => [...prev, { key: k, value: v, sub: newStatSub.trim() }]);
      setNewStatKey("");
      setNewStatValue("");
      setNewStatSub("");
    }
  };

  const removeStat = (index: number) => {
    setStats((prev) => prev.filter((_, i) => i !== index));
  };

  /* Tags helpers */
  const addTag = () => {
    const trimmed = newTag.trim().toUpperCase();
    if (trimmed && !tags.includes(trimmed)) {
      setTags((prev) => [...prev, trimmed]);
      setNewTag("");
    }
  };

  const removeTag = (index: number) => {
    setTags((prev) => prev.filter((_, i) => i !== index));
  };

  /* Character count helper */
  const charCount = (text: string, max: number) => {
    const len = text.length;
    return (
      <span
        className={len > max ? "text-red" : "text-line"}
        style={{ fontFamily: "var(--font-terminal)", fontSize: "0.85rem" }}
      >
        {len}/{max}
      </span>
    );
  };

  return (
    <AppLayout
      currentPage={"about" as PageKey}
      onNavigate={handleNavigate}
      onLogout={handleLogout}
      navItems={navItems}
      pageTitle={pageMeta["about" as PageKey]}
    >
      <div className="grid-2" data-testid="about-settings">
        {/* ──────────── LEFT COLUMN: Identity & Bio ──────────── */}
        <div className="stack">
          {/* Profile Identity */}
          <div className="cms-card accent-purple" data-testid="about-identity">
            <div className="card-title-row">
              <div className="card-title">&gt; IDENTITY.CFG</div>
              <span className="pixel-badge live">LIVE</span>
            </div>

            <div className="cms-form-row">
              <label>FULL NAME</label>
              <input
                className="pix-input"
                value={fullName}
                data-testid="about-fullname"
                onChange={(e) => setFullName(e.target.value)}
              />
            </div>

            <div className="cms-form-row">
              <label>HEADLINE</label>
              <input
                className="pix-input"
                value={headline}
                placeholder="e.g. PIXEL-PUSHER & CODE-SMITH"
                data-testid="about-headline"
                onChange={(e) => setHeadline(e.target.value)}
              />
            </div>

            <div className="cms-form-row">
              <label>
                SHORT INTRO {charCount(shortIntro, 200)}
              </label>
              <textarea
                className="pix-textarea small"
                value={shortIntro}
                placeholder="One-liner shown below hero title..."
                data-testid="about-shortintro"
                onChange={(e) => setShortIntro(e.target.value)}
              />
            </div>

            <div className="cms-form-row">
              <label>
                ABOUT / BIO {charCount(about, 1000)}
              </label>
              <textarea
                className="pix-textarea tall"
                value={about}
                placeholder="Your full about text..."
                data-testid="about-bio"
                onChange={(e) => setAbout(e.target.value)}
              />
            </div>

            <div className="cms-form-row">
              <label>PROFILE PHOTO URL</label>
              <input
                className="pix-input"
                value={profilePhotoUrl}
                placeholder="https://..."
                data-testid="about-photo"
                onChange={(e) => setProfilePhotoUrl(e.target.value)}
              />
            </div>

            <div className="pix-divider" />

            <div className="row-actions end">
              <button
                className="pix-btn pix-btn-ghost"
                type="button"
                data-testid="about-identity-discard"
              >
                DISCARD
              </button>
              <button
                className="pix-btn pix-btn-green"
                type="button"
                data-testid="about-identity-save"
              >
                SAVE IDENTITY
              </button>
            </div>
          </div>

          {/* Contact & Availability */}
          <div className="cms-card accent-cyan" data-testid="about-contact">
            <div className="card-title-row">
              <div className="card-title">&gt; CONTACT.INI</div>
            </div>

            <div className="cms-form-grid">
              <div className="cms-form-row">
                <label>LOCATION</label>
                <input
                  className="pix-input"
                  value={location}
                  data-testid="about-location"
                  onChange={(e) => setLocation(e.target.value)}
                />
              </div>
              <div className="cms-form-row">
                <label>EMAIL</label>
                <input
                  className="pix-input"
                  value={contactEmail}
                  data-testid="about-email"
                  onChange={(e) => setContactEmail(e.target.value)}
                />
              </div>
            </div>

            <div className="cms-form-grid">
              <div className="cms-form-row">
                <label>PHONE</label>
                <input
                  className="pix-input"
                  value={phone}
                  data-testid="about-phone"
                  onChange={(e) => setPhone(e.target.value)}
                />
              </div>
              <div className="cms-form-row">
                <label>AVAILABILITY</label>
                <input
                  className="pix-input"
                  value={availability}
                  placeholder="e.g. AVAILABLE FOR CONTRACT · Q2 2026"
                  data-testid="about-availability"
                  onChange={(e) => setAvailability(e.target.value)}
                />
              </div>
            </div>

            <div className="cms-form-row">
              <label>PRIMARY CTA LABEL</label>
              <input
                className="pix-input"
                value={primaryCta}
                placeholder="e.g. HIRE ME"
                data-testid="about-cta"
                onChange={(e) => setPrimaryCta(e.target.value)}
              />
            </div>

            <div className="pix-divider" />

            <div className="row-actions end">
              <button
                className="pix-btn pix-btn-ghost"
                type="button"
                data-testid="about-contact-discard"
              >
                DISCARD
              </button>
              <button
                className="pix-btn pix-btn-green"
                type="button"
                data-testid="about-contact-save"
              >
                SAVE CONTACT
              </button>
            </div>
          </div>
        </div>

        {/* ──────────── RIGHT COLUMN: Social, Focus, Stats, Tags ──────────── */}
        <div className="stack">
          {/* Social Links */}
          <div className="cms-card accent-pink" data-testid="about-socials">
            <div className="card-title-row">
              <div className="card-title">&gt; SOCIAL.LINKS</div>
            </div>

            <div className="cms-form-row">
              <label>GITHUB</label>
              <input
                className="pix-input"
                value={github}
                placeholder="https://github.com/..."
                data-testid="about-github"
                onChange={(e) => setGithub(e.target.value)}
              />
            </div>

            <div className="cms-form-row">
              <label>DRIBBBLE</label>
              <input
                className="pix-input"
                value={dribbble}
                placeholder="https://dribbble.com/..."
                data-testid="about-dribbble"
                onChange={(e) => setDribbble(e.target.value)}
              />
            </div>

            <div className="cms-form-row">
              <label>FIGMA</label>
              <input
                className="pix-input"
                value={figma}
                placeholder="https://figma.com/..."
                data-testid="about-figma"
                onChange={(e) => setFigma(e.target.value)}
              />
            </div>

            <div className="cms-form-row">
              <label>TWITTER / X</label>
              <input
                className="pix-input"
                value={twitter}
                placeholder="https://twitter.com/..."
                data-testid="about-twitter"
                onChange={(e) => setTwitter(e.target.value)}
              />
            </div>

            <div className="pix-divider" />

            <div className="row-actions end">
              <button
                className="pix-btn pix-btn-green"
                type="button"
                data-testid="about-socials-save"
              >
                SAVE LINKS
              </button>
            </div>
          </div>

          {/* Current Quest / Focus */}
          <div className="cms-card accent-green" data-testid="about-focus">
            <div className="card-title-row">
              <div className="card-title">&gt; CURRENT QUEST</div>
              <span className="pixel-badge">{focusItems.length} ITEMS</span>
            </div>

            {focusItems.map((item, i) => (
              <div className="about-list-item" key={i}>
                <span
                  className="text-text"
                  style={{
                    fontFamily: "var(--font-terminal)",
                    fontSize: "1rem",
                  }}
                >
                  {item}
                </span>
                <button
                  className="pix-btn pix-btn-ghost about-remove-btn"
                  type="button"
                  aria-label={`Remove quest ${i + 1}`}
                  data-testid={`about-focus-remove-${i}`}
                  onClick={() => removeFocus(i)}
                >
                  ✕
                </button>
              </div>
            ))}

            <div className="cms-form-row" style={{ marginTop: 12 }}>
              <label>NEW QUEST</label>
              <div className="input-with-action">
                <input
                  className="pix-input"
                  value={newFocus}
                  placeholder="e.g. Building a CMS for indie devs"
                  data-testid="about-focus-input"
                  onChange={(e) => setNewFocus(e.target.value)}
                  onKeyDown={(e) => e.key === "Enter" && addFocus()}
                />
                <button
                  className="pix-btn pix-btn-green"
                  type="button"
                  data-testid="about-focus-add"
                  onClick={addFocus}
                >
                  + ADD
                </button>
              </div>
            </div>
          </div>

          {/* Highlight Stats */}
          <div className="cms-card accent-purple" data-testid="about-stats">
            <div className="card-title-row">
              <div className="card-title">&gt; HIGHLIGHT.STATS</div>
              <span className="pixel-badge">{stats.length} / 4</span>
            </div>

            <div className="about-stats-grid">
              {stats.map((stat, i) => (
                <div className="about-stat-chip" key={i}>
                  <div
                    className="text-purple"
                    style={{
                      fontFamily: "var(--font-pixel)",
                      fontSize: "0.48rem",
                    }}
                  >
                    {stat.key}
                  </div>
                  <div
                    className="text-text"
                    style={{
                      fontFamily: "var(--font-pixel)",
                      fontSize: "0.9rem",
                    }}
                  >
                    {stat.value}
                  </div>
                  <div
                    className="text-muted"
                    style={{
                      fontFamily: "var(--font-terminal)",
                      fontSize: "0.85rem",
                    }}
                  >
                    {stat.sub}
                  </div>
                  <button
                    className="about-stat-remove"
                    type="button"
                    aria-label={`Remove stat ${stat.key}`}
                    data-testid={`about-stat-remove-${i}`}
                    onClick={() => removeStat(i)}
                  >
                    ✕
                  </button>
                </div>
              ))}
            </div>

            {stats.length < 4 && (
              <>
                <div className="pix-divider" />
                <div className="cms-form-grid" style={{ marginBottom: 8 }}>
                  <div className="cms-form-row">
                    <label>LABEL</label>
                    <input
                      className="pix-input"
                      value={newStatKey}
                      placeholder="e.g. YEARS XP"
                      data-testid="about-stat-key"
                      onChange={(e) => setNewStatKey(e.target.value)}
                    />
                  </div>
                  <div className="cms-form-row">
                    <label>VALUE</label>
                    <input
                      className="pix-input"
                      value={newStatValue}
                      placeholder="e.g. 8+"
                      data-testid="about-stat-value"
                      onChange={(e) => setNewStatValue(e.target.value)}
                    />
                  </div>
                </div>
                <div className="cms-form-row">
                  <label>SUBTITLE</label>
                  <div className="input-with-action">
                    <input
                      className="pix-input"
                      value={newStatSub}
                      placeholder="e.g. since 2018"
                      data-testid="about-stat-sub"
                      onChange={(e) => setNewStatSub(e.target.value)}
                      onKeyDown={(e) => e.key === "Enter" && addStat()}
                    />
                    <button
                      className="pix-btn pix-btn-purple"
                      type="button"
                      data-testid="about-stat-add"
                      onClick={addStat}
                    >
                      + ADD
                    </button>
                  </div>
                </div>
              </>
            )}
          </div>

          {/* Obsessions / Tags */}
          <div className="cms-card accent-cyan" data-testid="about-tags">
            <div className="card-title-row">
              <div className="card-title">&gt; OBSESSIONS.TAG</div>
            </div>

            <div className="about-tags-wrap">
              {tags.map((tag, i) => (
                <span className="tag tag-purple about-tag-chip" key={tag}>
                  {tag}
                  <button
                    className="about-tag-remove"
                    type="button"
                    aria-label={`Remove tag ${tag}`}
                    data-testid={`about-tag-remove-${i}`}
                    onClick={() => removeTag(i)}
                  >
                    ✕
                  </button>
                </span>
              ))}
            </div>

            <div className="cms-form-row" style={{ marginTop: 14 }}>
              <label>ADD TAG</label>
              <div className="input-with-action">
                <input
                  className="pix-input"
                  value={newTag}
                  placeholder="e.g. MOTION DESIGN"
                  data-testid="about-tag-input"
                  onChange={(e) => setNewTag(e.target.value)}
                  onKeyDown={(e) => e.key === "Enter" && addTag()}
                />
                <button
                  className="pix-btn pix-btn-cyan"
                  type="button"
                  data-testid="about-tag-add"
                  onClick={addTag}
                >
                  + ADD
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </AppLayout>
  );
}
