import type { ComponentProps } from "react";

import { cn } from "../../lib";

export function Separator({ className, ...props }: ComponentProps<"div">) {
  return (
    <div
      className={cn("ui-separator", className)}
      role="separator"
      {...props}
    />
  );
}
