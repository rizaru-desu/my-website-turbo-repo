import { useState } from "react";
import { useForm } from "@tanstack/react-form";
import {
  Button,
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  Checkbox,
  Input,
  Label,
  Separator,
} from "@repo/ui";

import AdminDashboard from "./AdminDashboard";
import ForgotPasswordForm from "./ForgotPasswordForm";
import Verify2faForm from "./Verify2faForm";
import {
  IconLock,
  IconUser,
  PixelCrosshair,
  PixelFloppy,
  PixelSprite,
} from "./icons";
import { getFieldError } from "./utils";
import "./App.css";

function App() {
  const [screen, setScreen] = useState<
    "login" | "forgot" | "verify2fa" | "dashboard"
  >("login");
  const loginForm = useForm({
    defaultValues: {
      username: "admin",
      password: "",
      remember: true,
    },
    onSubmit: async () => {
      await new Promise((resolve) => window.setTimeout(resolve, 450));
      setScreen("verify2fa");
    },
  });
  const isForgot = screen === "forgot";
  const is2FA = screen === "verify2fa";

  if (screen === "dashboard") {
    return <AdminDashboard onLogout={() => setScreen("login")} />;
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
            &gt;{" "}
            <span>
              {is2FA
                ? "2fa_challenge"
                : isForgot
                  ? "recovery_mode"
                  : "login_required"}
            </span>
            : true
            <br />
            &gt;{" "}
            <span>
              {is2FA
                ? "enter_code"
                : isForgot
                  ? "reset_token"
                  : "press_any_key"}
            </span>
            :{" "}
            {is2FA
              ? "from authenticator"
              : isForgot
                ? "queued_for_dispatch"
                : "to continue"}
            <span className="blink" />
          </div>

          {is2FA ? (
            <Verify2faForm
              onVerified={() => setScreen("dashboard")}
              onBackToLogin={() => setScreen("login")}
            />
          ) : isForgot ? (
            <ForgotPasswordForm onBackToLogin={() => setScreen("login")} />
          ) : (
            <form
              className="login-form"
              data-testid="login-form"
              onSubmit={(event) => {
                event.preventDefault();
                event.stopPropagation();
                loginForm.handleSubmit();
              }}
            >
              <loginForm.Field
                name="username"
                validators={{
                  onSubmit: ({ value }) =>
                    !value ? "USERNAME REQUIRED" : undefined,
                }}
              >
                {(field) => {
                  const error = getFieldError(field.state.meta.errors);

                  return (
                    <div className="form-row">
                      <Label htmlFor={field.name}>
                        <IconUser />
                        USERNAME
                      </Label>
                      <Input
                        id={field.name}
                        name={field.name}
                        value={field.state.value}
                        autoComplete="username"
                        data-testid="login-username"
                        autoFocus
                        onBlur={field.handleBlur}
                        onChange={(event) =>
                          field.handleChange(event.target.value)
                        }
                      />
                      {error ? <p className="field-error">! {error}</p> : null}
                    </div>
                  );
                }}
              </loginForm.Field>

              <loginForm.Field
                name="password"
                validators={{
                  onSubmit: ({ value }) =>
                    !value ? "PASSWORD REQUIRED" : undefined,
                }}
              >
                {(field) => {
                  const error = getFieldError(field.state.meta.errors);

                  return (
                    <div className="form-row">
                      <Label htmlFor={field.name}>
                        <IconLock />
                        PASSWORD
                      </Label>
                      <Input
                        id={field.name}
                        name={field.name}
                        type="password"
                        value={field.state.value}
                        placeholder="****************"
                        autoComplete="current-password"
                        data-testid="login-password"
                        onBlur={field.handleBlur}
                        onChange={(event) =>
                          field.handleChange(event.target.value)
                        }
                      />
                      {error ? <p className="field-error">! {error}</p> : null}
                    </div>
                  );
                }}
              </loginForm.Field>

              <div className="login-options">
                <loginForm.Field name="remember">
                  {(field) => (
                    <Label
                      className="remember-toggle"
                      data-testid="remember-toggle"
                    >
                      <Checkbox
                        checked={field.state.value}
                        aria-label="Remember me"
                        onBlur={field.handleBlur}
                        onChange={(event) =>
                          field.handleChange(event.target.checked)
                        }
                      />
                      <span>REMEMBER ME</span>
                    </Label>
                  )}
                </loginForm.Field>
                <button
                  type="button"
                  className="forgot-link"
                  data-testid="forgot-toggle"
                  onClick={() => setScreen("forgot")}
                >
                  FORGOT?
                </button>
              </div>

              <loginForm.Subscribe
                selector={(state) => [state.canSubmit, state.isSubmitting]}
              >
                {([canSubmit, isSubmitting]) => (
                  <Button
                    type="submit"
                    size="lg"
                    className="login-submit"
                    disabled={!canSubmit || isSubmitting}
                    data-testid="login-submit"
                  >
                    {isSubmitting ? "CONNECTING..." : "INSERT COIN > LOGIN"}
                  </Button>
                )}
              </loginForm.Subscribe>

              <div className="login-divider-text">~ OR ~</div>

              <Button
                type="button"
                variant="ghost"
                className="login-submit"
                data-testid="login-back"
              >
                &lt;- BACK TO PORTFOLIO
              </Button>
            </form>
          )}
        </CardContent>

        <CardFooter>
          <p>
            {is2FA
              ? "DEMO: ANY 6-DIGIT CODE"
              : isForgot
                ? "ENTER YOUR EMAIL TO RECEIVE A RESET LINK"
                : "DEMO: ANY USERNAME + ANY PASSWORD"}
          </p>
        </CardFooter>
      </Card>

      <footer className="login-footer">
        (c) 2026 KAI MORIKAWA - MADE WITH &lt;3 &amp; PIXELS
      </footer>
    </main>
  );
}

export default App;
