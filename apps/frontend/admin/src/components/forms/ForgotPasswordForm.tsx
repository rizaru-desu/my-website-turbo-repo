import { useState } from "react";
import { useForm } from "@tanstack/react-form";
import { Button, Input, Label } from "@repo/ui";
import { apiClient, AxiosError } from "@repo/api";

import { IconMail } from "../icons";
import { getFieldError } from "../../utils";

interface ForgotPasswordFormProps {
  onBackToLogin: () => void;
}

export default function ForgotPasswordForm({
  onBackToLogin,
}: ForgotPasswordFormProps) {
  const [sent, setSent] = useState(false);
  const [apiError, setApiError] = useState<string | null>(null);

  const forgotForm = useForm({
    defaultValues: {
      email: "",
    },
    onSubmit: async ({ value }) => {
      setApiError(null);
      try {
        await apiClient.post("/auth/forgot-password", {
          email: value.email,
        });

        setSent(true);
      } catch (err) {
        if (err instanceof AxiosError) {
          setApiError(err.response?.data?.error || "REQUEST FAILED");
        } else {
          setApiError("NETWORK ERROR");
        }
      }
    },
  });

  if (sent) {
    return (
      <div className="login-form" data-testid="forgot-sent">
        <p className="recovery-copy">
          IF THE EMAIL EXISTS, A RESET LINK HAS BEEN SENT. CHECK YOUR INBOX.
        </p>
        <Button
          type="button"
          variant="ghost"
          className="login-submit"
          onClick={onBackToLogin}
        >
          &lt;- BACK TO LOGIN
        </Button>
      </div>
    );
  }

  return (
    <form
      className="login-form"
      data-testid="forgot-form"
      onSubmit={(event) => {
        event.preventDefault();
        event.stopPropagation();
        forgotForm.handleSubmit();
      }}
    >
      <p className="recovery-copy">
        ENTER ADMIN EMAIL TO RECEIVE A ONE-TIME ACCESS LINK.
      </p>

      {apiError ? <p className="field-error">! {apiError}</p> : null}

      <forgotForm.Field
        name="email"
        validators={{
          onSubmit: ({ value }) => {
            if (!value) {
              return "EMAIL REQUIRED";
            }

            if (!value.includes("@")) {
              return "EMAIL FORMAT INVALID";
            }

            return undefined;
          },
        }}
      >
        {(field) => {
          const error = getFieldError(field.state.meta.errors);

          return (
            <div className="form-row">
              <Label htmlFor={field.name}>
                <IconMail />
                EMAIL
              </Label>
              <Input
                id={field.name}
                name={field.name}
                type="email"
                value={field.state.value}
                placeholder="admin@pixel.cms"
                autoComplete="email"
                data-testid="forgot-email"
                autoFocus
                onBlur={field.handleBlur}
                onChange={(event) => field.handleChange(event.target.value)}
              />
              {error ? <p className="field-error">! {error}</p> : null}
            </div>
          );
        }}
      </forgotForm.Field>

      <forgotForm.Subscribe
        selector={(state) => [state.canSubmit, state.isSubmitting]}
      >
        {([canSubmit, isSubmitting]) => (
          <Button
            type="submit"
            size="lg"
            className="login-submit"
            disabled={!canSubmit || isSubmitting}
            data-testid="forgot-submit"
          >
            {isSubmitting ? "TRANSMITTING..." : "SEND RESET LINK"}
          </Button>
        )}
      </forgotForm.Subscribe>

      <div className="login-divider-text">~ OR ~</div>

      <Button
        type="button"
        variant="ghost"
        className="login-submit"
        data-testid="forgot-back-login"
        onClick={onBackToLogin}
      >
        &lt;- BACK TO LOGIN
      </Button>
    </form>
  );
}
