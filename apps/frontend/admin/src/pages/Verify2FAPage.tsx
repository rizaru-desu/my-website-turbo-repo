import { useLocation, useNavigate } from "react-router-dom";
import { Card, CardContent, CardFooter, CardHeader, Separator } from "@repo/ui";
import { PixelCrosshair, PixelFloppy, PixelSprite } from "../components/icons";
import { Verify2faForm } from "../components/forms";

export default function Verify2FAPage() {
  const navigate = useNavigate();
  const location = useLocation();
  const twoFactorToken = location.state?.twoFactorToken || "";

  if (!twoFactorToken) {
    navigate("/login");
    return null;
  }

  return (
    <main className="login-screen" data-testid="login-screen">
      <div className="login-sprite login-sprite--floppy">
        <PixelFloppy />
      </div>
      <div className="login-sprite login-sprite--sprite">
        <PixelSprite />
      </div>
      <div className="login-sprite login-sprite--crosshair">
        <PixelCrosshair />
      </div>

      <Card className="login-card">
        <CardHeader className="login-header">
          <PixelFloppy />
          <div>
            <h1 className="login-title">PIXEL.CMS</h1>
            <p className="login-subtitle">~ ADMIN TERMINAL v4.20</p>
          </div>
        </CardHeader>

        <Separator className="pixel-divider" />

        <CardContent>
          <div className="terminal-copy" aria-hidden="true">
            &gt; <span>2fa_challenge</span>: true
            <br />
            &gt; <span>enter_code</span>: from authenticator
            <span className="blink" />
          </div>

          <Verify2faForm
            twoFactorToken={twoFactorToken}
            onVerified={() => navigate("/dashboard")}
            onBackToLogin={() => navigate("/login")}
          />
        </CardContent>

        <CardFooter>
          <p>DEMO: ANY 6-DIGIT CODE</p>
        </CardFooter>
      </Card>

      <footer className="login-footer">
        (c) 2026 KAI MORIKAWA - MADE WITH &lt;3 &amp; PIXELS
      </footer>
    </main>
  );
}
