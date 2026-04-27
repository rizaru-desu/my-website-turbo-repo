import type { ComponentProps } from "react";

import { cn } from "../../lib";

export function Label({ className, ...props }: ComponentProps<"label">) {
  return <label className={cn("ui-label", className)} {...props} />;
}
