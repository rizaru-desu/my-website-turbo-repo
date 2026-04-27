import { useState } from "react";
import { useNavigate } from "react-router-dom";
import "@/App.css";

/* Pixel decorations */
const PixelFloppy = () => (
  <svg viewBox="0 0 32 32" width="72" height="72" shapeRendering="crispEdges">
    <rect x="4"  y="4"  width="24" height="24" fill="#bd93f9" />
    <rect x="4"  y="4"  width="24" height="2"  fill="#d4b3ff" />
    <rect x="4"  y="26" width="24" height="2"  fill="#7a5cb8" />
    <rect x="4"  y="4"  width="2"  height="24" fill="#d4b3ff" />
    <rect x="26" y="4"  width="2"  height="24" fill="#7a5cb8" />
    {/* metal slider */}
    <rect x="8"  y="6"  width="16" height="8"  fill="#bdc0cc" />
    <rect x="18" y="7"  width="4"  height="6"  fill="#44475a" />
    {/* label */}
    <rect x="10" y="18" width="12" height="8"  fill="#f8f8f2" />
    <rect x="12" y="20" width="8"  height="1"  fill="#282a36" />
    <rect x="12" y="22" width="8"  height="1"  fill="#282a36" />
    <rect x="12" y="24" width="6"  height="1"  fill="#282a36" />
  </svg>
);

const IconLock = () => (
  <svg viewBox="0 0 16 16" width="16" height="16" shapeRendering="crispEdges">
    <rect x="5" y="2" width="6" height="2" fill="#f1fa8c" />
    <rect x="4" y="3" width="1" height="5" fill="#f1fa8c" />
    <rect x="11" y="3" width="1" height="5" fill="#f1fa8c" />
    <rect x="3" y="7" width="10" height="7" fill="#f1fa8c" />
    <rect x="7" y="10" width="2" height="2" fill="#282a36" />
  </svg>
);

const IconUser = () => (
  <svg viewBox="0 0 16 16" width="16" height="16" shapeRendering="crispEdges">
    <rect x="6" y="3" width="4" height="4" fill="#8be9fd" />
    <rect x="4" y="9" width="8" height="5" fill="#8be9fd" />
  </svg>
);

export default function LoginPage() {
  const [u, setU] = useState("admin");
  const [p, setP] = useState("");
  const [err, setErr] = useState("");
  const [loading, setLoading] = useState(false);
  const nav = useNavigate();

  const submit = (e) => {
    e.preventDefault();
    setErr("");
    if (!u || !p) {
      setErr("BOTH FIELDS REQUIRED");
      return;
    }
    setLoading(true);
    // Mock auth: any non-empty credentials → /admin
    setTimeout(() => {
      setLoading(false);
      nav("/admin");
    }, 900);
  };

  return (
    <div className="login-screen" data-testid="login-screen">
      {/* Decorative floating pixel sprites */}
      <div className="login-sprite s1"><PixelFloppy /></div>
      <div className="login-sprite s2">
        <svg viewBox="0 0 16 16" width="48" height="48" shapeRendering="crispEdges">
          <rect x="6" y="2" width="4" height="2" fill="#50fa7b"/>
          <rect x="4" y="4" width="8" height="8" fill="#50fa7b"/>
          <rect x="2" y="6" width="12" height="4" fill="#50fa7b"/>
          <rect x="6" y="6" width="1" height="2" fill="#282a36"/>
          <rect x="9" y="6" width="1" height="2" fill="#282a36"/>
          <rect x="4" y="12" width="2" height="2" fill="#50fa7b"/>
          <rect x="10" y="12" width="2" height="2" fill="#50fa7b"/>
        </svg>
      </div>
      <div className="login-sprite s3">
        <svg viewBox="0 0 16 16" width="40" height="40" shapeRendering="crispEdges">
          <rect x="7" y="1" width="2" height="14" fill="#ff79c6"/>
          <rect x="1" y="7" width="14" height="2" fill="#ff79c6"/>
          <rect x="5" y="3" width="6" height="2" fill="#ff79c6"/>
          <rect x="5" y="11" width="6" height="2" fill="#ff79c6"/>
          <rect x="3" y="5" width="2" height="6" fill="#ff79c6"/>
          <rect x="11" y="5" width="2" height="6" fill="#ff79c6"/>
        </svg>
      </div>

      <form className="login-card" onSubmit={submit} data-testid="login-form">
        <div className="login-header">
          <PixelFloppy />
          <div>
            <div className="login-title">PIXEL.CMS</div>
            <div className="login-subtitle">~ ADMIN TERMINAL v4.20</div>
          </div>
        </div>

        <div className="pix-divider" style={{ marginTop: 4 }} />

        <div style={{
          fontFamily: "'VT323',monospace",
          fontSize: 17,
          color: "var(--cyan)",
          marginBottom: 18,
          lineHeight: 1.3,
        }}>
          &gt; <span style={{ color: "var(--text)" }}>login_required</span>: true<br/>
          &gt; <span style={{ color: "var(--text)" }}>press_any_key</span>: to continue<span className="blink" />
        </div>

        <div className="form-row">
          <label><IconUser /> USERNAME</label>
          <input
            className="pix-input"
            value={u}
            onChange={(e) => setU(e.target.value)}
            autoComplete="username"
            data-testid="login-username"
            autoFocus
          />
        </div>

        <div className="form-row">
          <label><IconLock /> PASSWORD</label>
          <input
            type="password"
            className="pix-input"
            value={p}
            onChange={(e) => setP(e.target.value)}
            placeholder="****************"
            autoComplete="current-password"
            data-testid="login-password"
          />
        </div>

        <div style={{ display: "flex", justifyContent: "space-between", alignItems: "center", marginBottom: 14 }}>
          <label style={{ display: "flex", gap: 8, alignItems: "center", cursor: "pointer" }} data-testid="remember-toggle">
            <span
              className="pix-switch on"
              style={{ width: 40, height: 20 }}
            >
              <span className="pix-switch-knob" style={{ top: 2, left: 22, width: 14, height: 14 }} />
            </span>
            <span style={{ fontFamily: "'Press Start 2P',monospace", fontSize: 8, color: "var(--text)" }}>REMEMBER ME</span>
          </label>
          <a href="#" style={{
            fontFamily: "'Press Start 2P',monospace", fontSize: 8,
            color: "var(--pink)", textDecoration: "none",
          }}>FORGOT?</a>
        </div>

        {err && (
          <div data-testid="login-error" style={{
            padding: "8px 10px",
            background: "rgba(255,85,85,0.1)",
            color: "var(--red)",
            fontFamily: "'Press Start 2P',monospace",
            fontSize: 8,
            boxShadow: "inset 0 0 0 2px var(--red)",
            marginBottom: 12,
            letterSpacing: 0.5,
          }}>
            ! ERR_401: {err}
          </div>
        )}

        <button
          type="submit"
          className="pix-btn"
          style={{ width: "100%", justifyContent: "center", padding: "14px" }}
          disabled={loading}
          data-testid="login-submit"
        >
          {loading ? "CONNECTING..." : "INSERT COIN ▶ LOGIN"}
        </button>

        <div style={{
          textAlign: "center",
          marginTop: 18,
          fontFamily: "'VT323',monospace",
          fontSize: 15,
          color: "var(--line)",
        }}>
          ~ OR ~
        </div>

        <button
          type="button"
          onClick={() => nav("/")}
          className="pix-btn pix-btn-ghost"
          style={{ width: "100%", justifyContent: "center", marginTop: 10 }}
          data-testid="login-back"
        >
          ← BACK TO PORTFOLIO
        </button>

        <div style={{
          marginTop: 22,
          textAlign: "center",
          fontFamily: "'Press Start 2P',monospace",
          fontSize: 7,
          color: "var(--line)",
          letterSpacing: 1,
        }}>
          DEMO: ANY USERNAME + ANY PASSWORD
        </div>
      </form>

      <div style={{
        position: "absolute",
        bottom: 20,
        left: "50%",
        transform: "translateX(-50%)",
        fontFamily: "'Press Start 2P',monospace",
        fontSize: 8,
        color: "var(--line)",
        letterSpacing: 1,
      }}>
        © 2026 KAI MORIKAWA · MADE WITH &lt;3 &amp; PIXELS
      </div>
    </div>
  );
}
