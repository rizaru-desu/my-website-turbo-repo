import { Suspense, type ReactNode } from "react";
import { useLocation } from "react-router-dom";
import { Loading } from "@repo/ui";
import RouteErrorBoundary from "./RouteErrorBoundary";

interface SuspenseBoundaryProps {
  children: ReactNode;
  fallback?: ReactNode;
}

/**
 * A global wrapper that combines Suspense and ErrorBoundary.
 * Used to handle lazy-loaded routes and unexpected render errors.
 */
export default function SuspenseBoundary({ 
  children, 
  fallback = <Loading /> 
}: SuspenseBoundaryProps) {
  const location = useLocation();

  return (
    <RouteErrorBoundary resetKey={location.pathname}>
      <Suspense fallback={fallback}>
        {children}
      </Suspense>
    </RouteErrorBoundary>
  );
}
