import type { ComponentProps } from "react";

import { cn } from "../../lib";

type SwitchProps = Omit<ComponentProps<"input">, "type">;

export function Switch({ className, ...props }: SwitchProps) {
  return (
    <input
      type="checkbox"
      role="switch"
      className={cn("ui-switch", className)}
      {...props}
    />
  );
}
