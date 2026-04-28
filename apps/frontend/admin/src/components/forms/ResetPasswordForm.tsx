import { useMemo, useState } from "react";
import { useForm } from "@tanstack/react-form";
import { Button, Input, Label } from "@repo/ui";

import { IconLock, IconMail } from "../icons";
import { getFieldError } from "../../utils";

interface ResetPasswordFormProps {
  email: string;
  token: string;
  onBackToLogin: () => void;
}

export default function ResetPasswordForm({
  email,
  token,
  onBackToLogin,
}: ResetPasswordFormProps) {
  const [isComplete, setIsComplete] = useState(false);
  const hasValidLink = Boolean(email && token);

  const resetForm = useForm({
    defaultValues: {
      email,
      password: "",
      confirmPassword: "",
    },
    onSubmit: async () => {
      setIsComplete(true);
    },
  });

  const passwordRules = useMemo(
    () => [
      "MINIMUM 8 CHARACTERS",
      "USE A MIX OF LETTERS, NUMBERS, OR SYMBOLS",
      "CONFIRMATION MUST MATCH",
    ],
    [],
  );

  if (!hasValidLink) {
    return (
      <div className="login-form" data-testid="reset-invalid">
        <p className="recovery-copy reset-copy">
          RESET LINK IS MISSING TOKEN DATA. REQUEST A NEW PASSWORD RESET EMAIL.
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

  if (isComplete) {
    return (
      <div className="login-form" data-testid="reset-complete">
        <p className="recovery-copy reset-copy">
          PASSWORD RESET UI COMPLETE. BACKEND SUBMISSION WILL BE WIRED NEXT.
        </p>
        <Button
          type="button"
          size="lg"
          className="login-submit"
          onClick={onBackToLogin}
        >
          RETURN TO LOGIN
        </Button>
      </div>
    );
  }

  return (
    <form
      className="login-form"
      data-testid="reset-password-form"
      onSubmit={(event) => {
        event.preventDefault();
        event.stopPropagation();
        resetForm.handleSubmit();
      }}
    >
      <p className="recovery-copy reset-copy">
        CREATE A NEW ADMIN PASSWORD FOR THIS RESET LINK.
      </p>

      <resetForm.Field name="email">
        {(field) => (
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
              autoComplete="email"
              data-testid="reset-email"
              readOnly
            />
          </div>
        )}
      </resetForm.Field>

      <resetForm.Field
        name="password"
        validators={{
          onSubmit: ({ value }) => {
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
            <div className="form-row">
              <Label htmlFor={field.name}>
                <IconLock />
                NEW PASSWORD
              </Label>
              <Input
                id={field.name}
                name={field.name}
                type="password"
                value={field.state.value}
                placeholder="****************"
                autoComplete="new-password"
                data-testid="reset-password"
                autoFocus
                onBlur={field.handleBlur}
                onChange={(event) => field.handleChange(event.target.value)}
              />
              {error ? <p className="field-error">! {error}</p> : null}
            </div>
          );
        }}
      </resetForm.Field>

      <resetForm.Field
        name="confirmPassword"
        validators={{
          onSubmit: ({ value, fieldApi }) => {
            if (!value) {
              return "CONFIRM PASSWORD REQUIRED";
            }

            if (value !== fieldApi.form.getFieldValue("password")) {
              return "PASSWORDS DO NOT MATCH";
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
                <IconLock />
                CONFIRM PASSWORD
              </Label>
              <Input
                id={field.name}
                name={field.name}
                type="password"
                value={field.state.value}
                placeholder="****************"
                autoComplete="new-password"
                data-testid="reset-confirm-password"
                onBlur={field.handleBlur}
                onChange={(event) => field.handleChange(event.target.value)}
              />
              {error ? <p className="field-error">! {error}</p> : null}
            </div>
          );
        }}
      </resetForm.Field>

      <ul className="password-rules" aria-label="Password requirements">
        {passwordRules.map((rule) => (
          <li key={rule}>{rule}</li>
        ))}
      </ul>

      <resetForm.Subscribe
        selector={(state) => [state.canSubmit, state.isSubmitting]}
      >
        {([canSubmit, isSubmitting]) => (
          <Button
            type="submit"
            size="lg"
            className="login-submit"
            disabled={!canSubmit || isSubmitting}
            data-testid="reset-submit"
          >
            {isSubmitting ? "UPDATING..." : "RESET PASSWORD"}
          </Button>
        )}
      </resetForm.Subscribe>

      <div className="login-divider-text">~ OR ~</div>

      <Button
        type="button"
        variant="ghost"
        className="login-submit"
        data-testid="reset-back-login"
        onClick={onBackToLogin}
      >
        &lt;- BACK TO LOGIN
      </Button>
    </form>
  );
}
