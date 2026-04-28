import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useForm } from "@tanstack/react-form";
import {
  Button,
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  Input,
  Label,
  Separator,
  Switch,
} from "@repo/ui";
import {
  authService,
  AxiosError,
  type TwoFactorResponse,
  type LoginResponse,
} from "@repo/api";
import {
  IconLock,
  IconUser,
  PixelCrosshair,
  PixelFloppy,
  PixelSprite,
} from "../components/icons";
import { getFieldError } from "../utils";
import { useAuth } from "../hooks/useAuth";

export default function LoginPage() {
  const navigate = useNavigate();
  const { login } = useAuth();
  const [loginError, setLoginError] = useState<string>("");

  const loginForm = useForm({
    defaultValues: {
      username: "",
      password: "",
      remember: true,
    },
    onSubmit: async ({ value }) => {
      try {
        setLoginError("");
        const response = await authService.login({
          email: value.username.trim(),
          password: value.password,
          remember_me: value.remember,
        });

        if ("requires_two_factor" in response && response.requires_two_factor) {
          navigate("/verify-2fa", {
            state: {
              twoFactorToken: (response as TwoFactorResponse).two_factor_token,
            },
          });
        } else {
          login(response as LoginResponse);
          navigate("/dashboard");
        }
      } catch (error) {
        if (error instanceof AxiosError) {
          setLoginError(
            error.response?.data?.error || "Failed to login. Please try again.",
          );
        } else {
          setLoginError("An unexpected error occurred.");
        }
      }
    },
  });

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
            &gt; <span>login_required</span>: true
            <br />
            &gt; <span>press_any_key</span>: to continue
            <span className="blink" />
          </div>

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
                  !value ? "USERNAME/EMAIL REQUIRED" : undefined,
              }}
            >
              {(field) => (
                <div className="form-row">
                  <Label htmlFor={field.name}>
                    <IconUser />
                    USERNAME/EMAIL
                  </Label>
                  <Input
                    id={field.name}
                    name={field.name}
                    type="email"
                    value={field.state.value}
                    className="login-credential-input"
                    inputMode="email"
                    autoComplete="email"
                    autoCapitalize="none"
                    autoCorrect="off"
                    spellCheck={false}
                    data-testid="login-username"
                    autoFocus
                    onBlur={field.handleBlur}
                    onChange={(event) => field.handleChange(event.target.value)}
                  />
                </div>
              )}
            </loginForm.Field>

            <loginForm.Field
              name="password"
              validators={{
                onSubmit: ({ value }) =>
                  !value ? "PASSWORD REQUIRED" : undefined,
              }}
            >
              {(field) => (
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
                    className="login-credential-input"
                    placeholder="****************"
                    autoComplete="current-password"
                    autoCapitalize="none"
                    autoCorrect="off"
                    spellCheck={false}
                    data-testid="login-password"
                    onBlur={field.handleBlur}
                    onChange={(event) => field.handleChange(event.target.value)}
                  />
                </div>
              )}
            </loginForm.Field>

            <div className="login-options">
              <loginForm.Field name="remember">
                {(field) => (
                  <Label
                    className="remember-toggle"
                    data-testid="remember-toggle"
                  >
                    <Switch
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
                onClick={() => navigate("/forgot-password")}
              >
                FORGOT?
              </button>
            </div>

            <loginForm.Subscribe
              selector={(state) => ({
                canSubmit: state.canSubmit,
                isSubmitting: state.isSubmitting,
                submissionAttempts: state.submissionAttempts,
                fieldMeta: state.fieldMeta,
              })}
            >
              {({ canSubmit, isSubmitting, submissionAttempts, fieldMeta }) => {
                const fieldErr = Object.values(fieldMeta ?? {})
                  .flatMap((m) => m?.errors ?? [])
                  .map((e) => getFieldError([e]))
                  .find(Boolean);
                const showFieldErr = submissionAttempts > 0 && fieldErr;
                const errorMsg = loginError || (showFieldErr ? fieldErr : "");

                return (
                  <>
                    {errorMsg && (
                      <div
                        className="field-error"
                        style={{ marginBottom: "1rem" }}
                        data-testid="login-error"
                      >
                        ! ERR_401: {errorMsg}
                      </div>
                    )}
                    <Button
                      type="submit"
                      size="lg"
                      className="login-submit"
                      disabled={!canSubmit || isSubmitting}
                      data-testid="login-submit"
                    >
                      {isSubmitting ? "CONNECTING..." : "INSERT COIN > LOGIN"}
                    </Button>
                  </>
                );
              }}
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
        </CardContent>

        <CardFooter>
          <p>DEMO: ANY USERNAME + ANY PASSWORD</p>
        </CardFooter>
      </Card>

      <footer className="login-footer">
        (c) 2026 KAI MORIKAWA - MADE WITH &lt;3 &amp; PIXELS
      </footer>
    </main>
  );
}
