import { useRef, useState } from "react";
import { useForm } from "@tanstack/react-form";
import { Button, Label } from "@repo/ui";
import { authService, AxiosError } from "@repo/api";

import { IconShield } from "../icons";
import { getFieldError } from "../../utils";
import { useAuth } from "../../hooks/useAuth";

interface Verify2faFormProps {
  twoFactorToken: string;
  onVerified: () => void;
  onBackToLogin: () => void;
}

export default function Verify2faForm({
  twoFactorToken,
  onVerified,
  onBackToLogin,
}: Verify2faFormProps) {
  const otpRefs = useRef<(HTMLInputElement | null)[]>([]);
  const [verifyError, setVerifyError] = useState<string>("");
  const { login } = useAuth();

  const twoFaForm = useForm({
    defaultValues: {
      code: ["", "", "", "", "", ""] as string[],
    },
    onSubmit: async ({ value }) => {
      try {
        setVerifyError("");
        const response = await authService.verifyTOTP({
          two_factor_token: twoFactorToken,
          code: value.code.join(""),
        });
        login(response);
        onVerified();
      } catch (error) {
        if (error instanceof AxiosError) {
          setVerifyError(
            error.response?.data?.error ||
              "Failed to verify code. Please try again.",
          );
        } else {
          setVerifyError("An unexpected error occurred.");
        }
      }
    },
  });

  return (
    <form
      className="login-form"
      data-testid="2fa-form"
      onSubmit={(event) => {
        event.preventDefault();
        event.stopPropagation();
        twoFaForm.handleSubmit();
      }}
    >
      <p className="recovery-copy twofa-copy">
        ENTER THE 6-DIGIT CODE FROM YOUR AUTHENTICATOR APP.
      </p>

      {verifyError && (
        <div className="field-error" style={{ marginBottom: "1rem" }}>
          ! {verifyError}
        </div>
      )}

      <twoFaForm.Field
        name="code"
        validators={{
          onSubmit: ({ value }) => {
            const joined = value.join("");
            if (joined.length < 6) {
              return "6-DIGIT CODE REQUIRED";
            }
            if (!/^\d{6}$/.test(joined)) {
              return "DIGITS ONLY";
            }
            return undefined;
          },
        }}
      >
        {(field) => {
          const error = getFieldError(field.state.meta.errors);

          return (
            <div className="form-row">
              <Label>
                <IconShield />
                VERIFICATION CODE
              </Label>
              <div className="otp-group" data-testid="otp-group">
                {field.state.value.map((digit: string, index: number) => (
                  <input
                    key={index}
                    ref={(el) => {
                      otpRefs.current[index] = el;
                    }}
                    className="otp-cell"
                    type="text"
                    inputMode="numeric"
                    maxLength={1}
                    value={digit}
                    autoFocus={index === 0}
                    data-testid={`otp-${index}`}
                    onKeyDown={(event) => {
                      if (event.key === "Backspace" && !digit && index > 0) {
                        otpRefs.current[index - 1]?.focus();
                      }
                    }}
                    onPaste={(event) => {
                      event.preventDefault();
                      const paste = event.clipboardData
                        .getData("text")
                        .replace(/\D/g, "")
                        .slice(0, 6);
                      const next = [...field.state.value];
                      for (let i = 0; i < paste.length && index + i < 6; i++) {
                        next[index + i] = paste[i];
                      }
                      field.handleChange(next);
                      const focusIdx = Math.min(index + paste.length, 5);
                      otpRefs.current[focusIdx]?.focus();
                    }}
                    onChange={(event) => {
                      const char = event.target.value
                        .replace(/\D/g, "")
                        .slice(-1);
                      const next = [...field.state.value];
                      next[index] = char;
                      field.handleChange(next);
                      if (char && index < 5) {
                        otpRefs.current[index + 1]?.focus();
                      }
                    }}
                  />
                ))}
              </div>
              {error ? <p className="field-error">! {error}</p> : null}
            </div>
          );
        }}
      </twoFaForm.Field>

      <twoFaForm.Subscribe
        selector={(state) => [state.canSubmit, state.isSubmitting]}
      >
        {([canSubmit, isSubmitting]) => (
          <Button
            type="submit"
            size="lg"
            className="login-submit"
            disabled={!canSubmit || isSubmitting}
            data-testid="2fa-submit"
          >
            {isSubmitting ? "VERIFYING..." : "VERIFY CODE"}
          </Button>
        )}
      </twoFaForm.Subscribe>

      <div className="login-divider-text">~ OR ~</div>

      <Button
        type="button"
        variant="ghost"
        className="login-submit"
        data-testid="2fa-back-login"
        onClick={onBackToLogin}
      >
        &lt;- BACK TO LOGIN
      </Button>
    </form>
  );
}
