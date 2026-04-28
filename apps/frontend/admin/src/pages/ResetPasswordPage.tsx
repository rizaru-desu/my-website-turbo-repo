import { useNavigate, useSearchParams } from "react-router-dom";
import { Card, CardContent, CardFooter, CardHeader, Separator } from "@repo/ui";
import { PixelCrosshair, PixelFloppy, PixelSprite } from "../components/icons";
import { ResetPasswordForm } from "../components/forms";

export default function ResetPasswordPage() {
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const email = searchParams.get("email")?.trim() ?? "";
  const token = searchParams.get("token")?.trim() ?? "";

  return (
    <main className="login-screen" data-testid="reset-password-screen">
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
            &gt; <span>reset_session</span>: active
            <br />
            &gt; <span>credential_write</span>: pending
            <span className="blink" />
          </div>

          <ResetPasswordForm
            email={email}
            token={token}
            onBackToLogin={() => navigate("/login")}
          />
        </CardContent>

        <CardFooter>
          <p>SET A NEW PASSWORD TO RESTORE ADMIN ACCESS</p>
        </CardFooter>
      </Card>

      <footer className="login-footer">
        (c) 2026 KAI MORIKAWA - MADE WITH &lt;3 &amp; PIXELS
      </footer>
    </main>
  );
}
