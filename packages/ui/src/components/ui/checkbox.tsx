import type { ComponentProps } from "react";

import { cn } from "../../lib";

type CheckboxProps = Omit<ComponentProps<"input">, "type">;

export function Checkbox({ className, ...props }: CheckboxProps) {
  return (
    <input
      type="checkbox"
      className={cn("ui-checkbox", className)}
      {...props}
    />
  );
}
