import { useEffect, useMemo, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useForm } from "@tanstack/react-form";
import { Button, Input, Label, Switch } from "@repo/ui";
import { authService, AxiosError, type SetupTOTPResponse } from "@repo/api";
import { AppLayout } from "../components/layout";
import { navItems, pageMeta } from "../constants/navigation";
import { useAuth } from "../hooks/useAuth";
import type { PageKey } from "../types/dashboard";
import { getFieldError } from "../utils";

export default function SettingsPage() {
  const navigate = useNavigate();
  const { logout, user } = useAuth();
  const [saveMessage, setSaveMessage] = useState("");
  const [saveError, setSaveError] = useState("");
  const [twoFactorEnabled, setTwoFactorEnabled] = useState(false);
  const [twoFactorLoading, setTwoFactorLoading] = useState(true);
  const [twoFactorAction, setTwoFactorAction] = useState<
    "setup" | "enable" | "disable" | "regenerate" | ""
  >("");
  const [twoFactorSetup, setTwoFactorSetup] =
    useState<SetupTOTPResponse | null>(null);
  const [twoFactorCode, setTwoFactorCode] = useState("");
  const [twoFactorPassword, setTwoFactorPassword] = useState("");
  const [twoFactorDisablePassword, setTwoFactorDisablePassword] = useState("");
  const [twoFactorBackupPassword, setTwoFactorBackupPassword] = useState("");
  const [twoFactorBackupCodes, setTwoFactorBackupCodes] = useState<string[]>(
    [],
  );
  const [twoFactorMessage, setTwoFactorMessage] = useState("");
  const [twoFactorError, setTwoFactorError] = useState("");

  const settingsForm = useForm({
    defaultValues: {
      email: user?.email ?? "",
      handle: user?.name
        ? `@${user.name.toLowerCase().replace(/\s+/g, "_")}`
        : "",
      emailNotifications: true,
      changePassword: false,
      currentPassword: "",
      newPassword: "",
      confirmPassword: "",
    },
    onSubmit: async ({ value }) => {
      setSaveMessage("");
      setSaveError("");

      if (!value.changePassword) {
        return;
      }

      try {
        await authService.changePassword({
          currentPassword: value.currentPassword,
          newPassword: value.newPassword,
          confirmPassword: value.confirmPassword,
        });

        settingsForm.setFieldValue("changePassword", false);
        settingsForm.setFieldValue("currentPassword", "");
        settingsForm.setFieldValue("newPassword", "");
        settingsForm.setFieldValue("confirmPassword", "");
        setSaveMessage("PASSWORD UPDATED. USE THE NEW PASSWORD NEXT LOGIN.");
      } catch (error) {
        if (error instanceof AxiosError) {
          setSaveError(error.response?.data?.error || "PASSWORD UPDATE FAILED");
        } else {
          setSaveError("NETWORK ERROR");
        }
      }
    },
  });

  useEffect(() => {
    settingsForm.setFieldValue("email", user?.email ?? "");
    settingsForm.setFieldValue(
      "handle",
      user?.name ? `@${user.name.toLowerCase().replace(/\s+/g, "_")}` : "",
    );
  }, [settingsForm, user?.email, user?.name]);

  useEffect(() => {
    let mounted = true;

    const loadTwoFactorStatus = async () => {
      try {
        setTwoFactorLoading(true);
        const response = await authService.getTwoFactorStatus();
        if (mounted) {
          setTwoFactorEnabled(response.enabled);
        }
      } catch {
        if (mounted) {
          setTwoFactorError("FAILED TO LOAD 2FA STATUS");
        }
      } finally {
        if (mounted) {
          setTwoFactorLoading(false);
        }
      }
    };

    void loadTwoFactorStatus();

    return () => {
      mounted = false;
    };
  }, []);

  const passwordRules = useMemo(
    () => [
      "CURRENT PASSWORD REQUIRED",
      "MINIMUM 8 CHARACTERS",
      "USE A MIX OF LETTERS, NUMBERS, OR SYMBOLS",
      "CONFIRMATION MUST MATCH",
    ],
    [],
  );

  const handleLogout = async () => {
    await logout();
    navigate("/login");
  };

  const handleNavigate = (newPage: PageKey) => {
    if (newPage === "settings") return;
    if (newPage === "about") {
      navigate("/about");
      return;
    }
    navigate("/dashboard", { state: { initialPage: newPage } });
  };

  const twoFactorErrorMessage = (error: unknown, fallback: string) => {
    if (error instanceof AxiosError) {
      return error.response?.data?.error || fallback;
    }

    return "NETWORK ERROR";
  };

  const resetTwoFactorFeedback = () => {
    setTwoFactorMessage("");
    setTwoFactorError("");
  };

  const handleSetupTwoFactor = async () => {
    resetTwoFactorFeedback();
    setTwoFactorAction("setup");

    try {
      if (!twoFactorPassword) {
        setTwoFactorError("PASSWORD REQUIRED");
        return;
      }

      const response = await authService.setupTOTP({
        password: twoFactorPassword,
      });
      setTwoFactorSetup(response);
      setTwoFactorBackupCodes(response.backup_codes);
      setTwoFactorPassword("");
      setTwoFactorMessage("SCAN QR CODE, THEN CONFIRM THE 6-DIGIT CODE.");
    } catch (error) {
      setTwoFactorError(twoFactorErrorMessage(error, "FAILED TO START 2FA"));
    } finally {
      setTwoFactorAction("");
    }
  };

  const handleEnableTwoFactor = async () => {
    resetTwoFactorFeedback();

    const code = twoFactorCode.replace(/\D/g, "").slice(0, 6);
    if (!twoFactorSetup) {
      setTwoFactorError("START SETUP FIRST");
      return;
    }
    if (code.length !== 6) {
      setTwoFactorError("6-DIGIT CODE REQUIRED");
      return;
    }

    setTwoFactorAction("enable");
    try {
      await authService.enableTOTP({
        code,
      });
      setTwoFactorEnabled(true);
      setTwoFactorSetup(null);
      setTwoFactorCode("");
      setTwoFactorMessage("2FA ENABLED. SAVE YOUR BACKUP CODES.");
    } catch (error) {
      setTwoFactorError(twoFactorErrorMessage(error, "FAILED TO ENABLE 2FA"));
    } finally {
      setTwoFactorAction("");
    }
  };

  const handleDisableTwoFactor = async () => {
    resetTwoFactorFeedback();

    if (!twoFactorDisablePassword) {
      setTwoFactorError("PASSWORD REQUIRED");
      return;
    }

    setTwoFactorAction("disable");
    try {
      await authService.disableTOTP({
        password: twoFactorDisablePassword,
      });
      setTwoFactorEnabled(false);
      setTwoFactorSetup(null);
      setTwoFactorBackupCodes([]);
      setTwoFactorDisablePassword("");
      setTwoFactorBackupPassword("");
      setTwoFactorMessage("2FA DISABLED.");
    } catch (error) {
      setTwoFactorError(twoFactorErrorMessage(error, "FAILED TO DISABLE 2FA"));
    } finally {
      setTwoFactorAction("");
    }
  };

  const handleRegenerateBackupCodes = async () => {
    resetTwoFactorFeedback();

    if (!twoFactorBackupPassword) {
      setTwoFactorError("PASSWORD REQUIRED");
      return;
    }

    setTwoFactorAction("regenerate");
    try {
      const response = await authService.regenerateBackupCodes({
        password: twoFactorBackupPassword,
      });
      setTwoFactorBackupCodes(response.backup_codes);
      setTwoFactorBackupPassword("");
      setTwoFactorMessage("NEW BACKUP CODES GENERATED.");
    } catch (error) {
      setTwoFactorError(
        twoFactorErrorMessage(error, "FAILED TO REGENERATE BACKUP CODES"),
      );
    } finally {
      setTwoFactorAction("");
    }
  };

  return (
    <AppLayout
      currentPage="settings"
      onNavigate={handleNavigate}
      onLogout={handleLogout}
      navItems={navItems}
      pageTitle={pageMeta.settings}
    >
      <form
        className="grid-2"
        data-testid="settings-form"
        onSubmit={(event) => {
          event.preventDefault();
          event.stopPropagation();
          settingsForm.handleSubmit();
        }}
      >
        <div className="cms-card accent-purple">
          <div className="card-title-row">
            <div className="card-title">&gt; ACCOUNT.CFG</div>
          </div>

          <settingsForm.Field
            name="email"
            validators={{
              onSubmit: ({ value }) => {
                if (!value) {
                  return "EMAIL REQUIRED";
                }

                if (!value.includes("@")) {
                  return "VALID EMAIL REQUIRED";
                }

                return undefined;
              },
            }}
          >
            {(field) => {
              const error = getFieldError(field.state.meta.errors);

              return (
                <div className="cms-form-row">
                  <Label htmlFor={field.name}>EMAIL</Label>
                  <Input
                    disabled
                    id={field.name}
                    name={field.name}
                    type="email"
                    value={field.state.value}
                    autoComplete="email"
                    data-testid="settings-email"
                    onBlur={field.handleBlur}
                    onChange={(event) => field.handleChange(event.target.value)}
                  />
                  {error ? <p className="field-error">! {error}</p> : null}
                </div>
              );
            }}
          </settingsForm.Field>

          <settingsForm.Field
            name="handle"
            validators={{
              onSubmit: ({ value }) =>
                !value.trim() ? "HANDLE REQUIRED" : undefined,
            }}
          >
            {(field) => {
              const error = getFieldError(field.state.meta.errors);

              return (
                <div className="cms-form-row">
                  <Label htmlFor={field.name}>HANDLE</Label>
                  <Input
                    disabled
                    id={field.name}
                    name={field.name}
                    value={field.state.value}
                    autoComplete="username"
                    data-testid="settings-handle"
                    onBlur={field.handleBlur}
                    onChange={(event) => field.handleChange(event.target.value)}
                  />
                  {error ? <p className="field-error">! {error}</p> : null}
                </div>
              );
            }}
          </settingsForm.Field>

          <div className="setting-row">
            <span>EMAIL NOTIFICATIONS</span>
            <settingsForm.Field name="emailNotifications">
              {(field) => (
                <Switch
                  checked={field.state.value}
                  aria-label="Email notifications"
                  data-testid="settings-email-notifications"
                  onBlur={field.handleBlur}
                  onChange={(event) => field.handleChange(event.target.checked)}
                />
              )}
            </settingsForm.Field>
          </div>

          <div className="setting-row">
            <span>CHANGE PASSWORD</span>
            <settingsForm.Field name="changePassword">
              {(field) => (
                <Switch
                  checked={field.state.value}
                  aria-label="Change password"
                  data-testid="settings-change-password-toggle"
                  onBlur={field.handleBlur}
                  onChange={(event) => {
                    const enabled = event.target.checked;
                    field.handleChange(enabled);

                    if (!enabled) {
                      settingsForm.setFieldValue("currentPassword", "");
                      settingsForm.setFieldValue("newPassword", "");
                      settingsForm.setFieldValue("confirmPassword", "");
                    }
                  }}
                />
              )}
            </settingsForm.Field>
          </div>

          <settingsForm.Subscribe
            selector={(state) => state.values.changePassword}
          >
            {(changePassword) =>
              changePassword ? (
                <div data-testid="settings-change-password-form">
                  <div className="pix-divider" />

                  <settingsForm.Field
                    name="currentPassword"
                    validators={{
                      onSubmit: ({ value, fieldApi }) => {
                        if (!fieldApi.form.getFieldValue("changePassword")) {
                          return undefined;
                        }

                        return value ? undefined : "CURRENT PASSWORD REQUIRED";
                      },
                    }}
                  >
                    {(field) => {
                      const error = getFieldError(field.state.meta.errors);

                      return (
                        <div className="cms-form-row">
                          <Label htmlFor={field.name}>CURRENT PASSWORD</Label>
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
                            data-testid="settings-current-password"
                            onBlur={field.handleBlur}
                            onChange={(event) =>
                              field.handleChange(event.target.value)
                            }
                          />
                          {error ? (
                            <p className="field-error">! {error}</p>
                          ) : null}
                        </div>
                      );
                    }}
                  </settingsForm.Field>

                  <settingsForm.Field
                    name="newPassword"
                    validators={{
                      onSubmit: ({ value, fieldApi }) => {
                        if (!fieldApi.form.getFieldValue("changePassword")) {
                          return undefined;
                        }

                        if (!value) {
                          return "NEW PASSWORD REQUIRED";
                        }

                        if (value.length < 8) {
                          return "PASSWORD TOO SHORT";
                        }

                        return undefined;
                      },
                    }}
                  >
                    {(field) => {
                      const error = getFieldError(field.state.meta.errors);

                      return (
                        <div className="cms-form-row">
                          <Label htmlFor={field.name}>NEW PASSWORD</Label>
                          <Input
                            id={field.name}
                            name={field.name}
                            type="password"
                            value={field.state.value}
                            className="login-credential-input"
                            placeholder="****************"
                            autoComplete="new-password"
                            autoCapitalize="none"
                            autoCorrect="off"
                            spellCheck={false}
                            data-testid="settings-new-password"
                            onBlur={field.handleBlur}
                            onChange={(event) =>
                              field.handleChange(event.target.value)
                            }
                          />
                          {error ? (
                            <p className="field-error">! {error}</p>
                          ) : null}
                        </div>
                      );
                    }}
                  </settingsForm.Field>

                  <settingsForm.Field
                    name="confirmPassword"
                    validators={{
                      onSubmit: ({ value, fieldApi }) => {
                        if (!fieldApi.form.getFieldValue("changePassword")) {
                          return undefined;
                        }

                        if (!value) {
                          return "CONFIRM PASSWORD REQUIRED";
                        }

                        if (
                          value !== fieldApi.form.getFieldValue("newPassword")
                        ) {
                          return "PASSWORDS DO NOT MATCH";
                        }

                        return undefined;
                      },
                    }}
                  >
                    {(field) => {
                      const error = getFieldError(field.state.meta.errors);

                      return (
                        <div className="cms-form-row">
                          <Label htmlFor={field.name}>CONFIRM PASSWORD</Label>
                          <Input
                            id={field.name}
                            name={field.name}
                            type="password"
                            value={field.state.value}
                            className="login-credential-input"
                            placeholder="****************"
                            autoComplete="new-password"
                            autoCapitalize="none"
                            autoCorrect="off"
                            spellCheck={false}
                            data-testid="settings-confirm-password"
                            onBlur={field.handleBlur}
                            onChange={(event) =>
                              field.handleChange(event.target.value)
                            }
                          />
                          {error ? (
                            <p className="field-error">! {error}</p>
                          ) : null}
                        </div>
                      );
                    }}
                  </settingsForm.Field>

                  <ul
                    className="password-rules"
                    aria-label="Password requirements"
                  >
                    {passwordRules.map((rule) => (
                      <li key={rule}>{rule}</li>
                    ))}
                  </ul>
                </div>
              ) : null
            }
          </settingsForm.Subscribe>

          <div className="pix-divider" />
          {saveMessage ? (
            <p
              className="recovery-copy reset-copy"
              data-testid="settings-save-message"
            >
              {saveMessage}
            </p>
          ) : null}
          {saveError ? (
            <p className="field-error" data-testid="settings-save-error">
              ! {saveError}
            </p>
          ) : null}
          <div className="row-actions end">
            <settingsForm.Subscribe
              selector={(state) => [
                state.canSubmit,
                state.isSubmitting,
                state.values.changePassword,
              ]}
            >
              {([canSubmit, isSubmitting, changePassword]) => (
                <Button
                  type="submit"
                  disabled={!canSubmit || isSubmitting || !changePassword}
                  data-testid="settings-submit"
                >
                  {isSubmitting ? "SAVING..." : "SAVE PASSWORD"}
                </Button>
              )}
            </settingsForm.Subscribe>
          </div>
        </div>

        <div className="cms-card accent-green">
          <div className="card-title-row">
            <div className="card-title">&gt; SECURITY.2FA</div>
            <span
              className={`pixel-badge ${twoFactorEnabled ? "live" : "warn"}`}
              data-testid="settings-2fa-status"
            >
              {twoFactorLoading
                ? "LOADING"
                : twoFactorEnabled
                  ? "ENABLED"
                  : "OFF"}
            </span>
          </div>

          <div className="setting-row">
            <span>AUTHENTICATOR APP</span>
            <Button
              type="button"
              disabled={
                twoFactorLoading ||
                twoFactorEnabled ||
                twoFactorAction !== ""
              }
              data-testid="settings-2fa-setup"
              onClick={handleSetupTwoFactor}
            >
              {twoFactorAction === "setup" ? "STARTING..." : "SETUP"}
            </Button>
          </div>

          {!twoFactorEnabled && !twoFactorSetup ? (
            <div
              className="cms-form-row twofa-password-row"
              data-testid="settings-2fa-password-form"
            >
              <Label htmlFor="settings-2fa-password">PASSWORD</Label>
              <Input
                id="settings-2fa-password"
                type="password"
                value={twoFactorPassword}
                autoComplete="current-password"
                className="login-credential-input"
                data-testid="settings-2fa-password"
                onChange={(event) => setTwoFactorPassword(event.target.value)}
              />
            </div>
          ) : null}

          {twoFactorSetup ? (
            <div className="twofa-setup" data-testid="settings-2fa-setup-form">
              <div className="pix-divider" />
              {twoFactorSetup.qr_code ? (
                <img
                  className="twofa-qr"
                  src={twoFactorSetup.qr_code}
                  alt="Authenticator QR code"
                  data-testid="settings-2fa-qr"
                />
              ) : null}
              <div className="cms-form-row">
                <Label htmlFor="settings-2fa-secret">SECRET</Label>
                <Input
                  readOnly
                  id="settings-2fa-secret"
                  value={twoFactorSetup.secret}
                  className="login-credential-input"
                  data-testid="settings-2fa-secret"
                />
              </div>
              <div className="cms-form-row">
                <Label htmlFor="settings-2fa-code">CODE</Label>
                <Input
                  id="settings-2fa-code"
                  value={twoFactorCode}
                  inputMode="numeric"
                  maxLength={6}
                  autoComplete="one-time-code"
                  className="login-credential-input"
                  data-testid="settings-2fa-code"
                  onChange={(event) =>
                    setTwoFactorCode(
                      event.target.value.replace(/\D/g, "").slice(0, 6),
                    )
                  }
                />
              </div>
              <div className="row-actions end">
                <Button
                  type="button"
                  disabled={twoFactorAction !== ""}
                  data-testid="settings-2fa-enable"
                  onClick={handleEnableTwoFactor}
                >
                  {twoFactorAction === "enable" ? "ENABLING..." : "ENABLE 2FA"}
                </Button>
              </div>
            </div>
          ) : null}

          {twoFactorEnabled ? (
            <div data-testid="settings-2fa-enabled-actions">
              <div className="pix-divider" />
              <div className="cms-form-row">
                <Label htmlFor="settings-2fa-backup-password">
                  PASSWORD FOR NEW CODES
                </Label>
                <Input
                  id="settings-2fa-backup-password"
                  type="password"
                  value={twoFactorBackupPassword}
                  autoComplete="current-password"
                  className="login-credential-input"
                  data-testid="settings-2fa-backup-password"
                  onChange={(event) =>
                    setTwoFactorBackupPassword(event.target.value)
                  }
                />
              </div>
              <div className="row-actions end">
                <Button
                  type="button"
                  disabled={twoFactorAction !== ""}
                  data-testid="settings-2fa-regenerate"
                  onClick={handleRegenerateBackupCodes}
                >
                  {twoFactorAction === "regenerate"
                    ? "GENERATING..."
                    : "NEW BACKUP CODES"}
                </Button>
              </div>

              <div className="cms-form-row twofa-disable-row">
                <Label htmlFor="settings-2fa-disable-password">
                  PASSWORD TO DISABLE
                </Label>
                <Input
                  id="settings-2fa-disable-password"
                  type="password"
                  value={twoFactorDisablePassword}
                  autoComplete="current-password"
                  className="login-credential-input"
                  data-testid="settings-2fa-disable-password"
                  onChange={(event) =>
                    setTwoFactorDisablePassword(event.target.value)
                  }
                />
              </div>
              <div className="row-actions end">
                <Button
                  type="button"
                  disabled={twoFactorAction !== ""}
                  data-testid="settings-2fa-disable"
                  onClick={handleDisableTwoFactor}
                >
                  {twoFactorAction === "disable" ? "DISABLING..." : "DISABLE"}
                </Button>
              </div>
            </div>
          ) : null}

          {twoFactorBackupCodes.length > 0 ? (
            <div
              className="backup-code-grid"
              data-testid="settings-2fa-backup-codes"
            >
              {twoFactorBackupCodes.map((code) => (
                <code key={code}>{code}</code>
              ))}
            </div>
          ) : null}

          {twoFactorMessage ? (
            <p
              className="recovery-copy twofa-copy"
              data-testid="settings-2fa-message"
            >
              {twoFactorMessage}
            </p>
          ) : null}
          {twoFactorError ? (
            <p className="field-error" data-testid="settings-2fa-error">
              ! {twoFactorError}
            </p>
          ) : null}
        </div>
      </form>
    </AppLayout>
  );
}
