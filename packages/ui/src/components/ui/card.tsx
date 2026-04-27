import type { ComponentProps } from "react";

import { cn } from "../../lib";

export function Card({ className, ...props }: ComponentProps<"div">) {
  return <div className={cn("ui-card", className)} {...props} />;
}

export function CardHeader({ className, ...props }: ComponentProps<"div">) {
  return <div className={cn("ui-card-header", className)} {...props} />;
}

export function CardContent({ className, ...props }: ComponentProps<"div">) {
  return <div className={cn("ui-card-content", className)} {...props} />;
}

export function CardFooter({ className, ...props }: ComponentProps<"div">) {
  return <div className={cn("ui-card-footer", className)} {...props} />;
}
