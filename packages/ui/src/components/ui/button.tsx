import type { ComponentProps } from "react";

import { cn } from "../../lib";

type ButtonProps = ComponentProps<"button"> & {
  variant?: "default" | "ghost";
  size?: "default" | "lg" | "icon";
};

export function Button({
  className,
  variant = "default",
  size = "default",
  ...props
}: ButtonProps) {
  return (
    <button
      className={cn(
        "ui-button",
        `ui-button--${variant}`,
        `ui-button--${size}`,
        className,
      )}
      {...props}
    />
  );
}
