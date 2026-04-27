import { useNavigate } from "react-router-dom";
import { Card, CardContent, CardFooter, CardHeader, Separator } from "@repo/ui";
import { PixelCrosshair, PixelFloppy, PixelSprite } from "../components/icons";
import { ForgotPasswordForm } from "../components/forms";

export default function ForgotPasswordPage() {
  const navigate = useNavigate();

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
            &gt; <span>recovery_mode</span>: true
            <br />
            &gt; <span>reset_token</span>: queued_for_dispatch
            <span className="blink" />
          </div>

          <ForgotPasswordForm onBackToLogin={() => navigate("/login")} />
        </CardContent>

        <CardFooter>
          <p>ENTER YOUR EMAIL TO RECEIVE A RESET LINK</p>
        </CardFooter>
      </Card>

      <footer className="login-footer">
        (c) 2026 KAI MORIKAWA - MADE WITH &lt;3 &amp; PIXELS
      </footer>
    </main>
  );
}
