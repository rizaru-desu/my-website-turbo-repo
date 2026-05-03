import { Component, type ErrorInfo, type ReactNode } from "react";

type RouteErrorBoundaryProps = {
  children: ReactNode;
  resetKey: string;
};

type RouteErrorBoundaryState = {
  error: Error | null;
};

export default class RouteErrorBoundary extends Component<
  RouteErrorBoundaryProps,
  RouteErrorBoundaryState
> {
  state: RouteErrorBoundaryState = {
    error: null,
  };

  static getDerivedStateFromError(error: Error): RouteErrorBoundaryState {
    return { error };
  }

  componentDidCatch(error: Error, errorInfo: ErrorInfo) {
    console.error("Public route render failed:", error, errorInfo);
  }

  componentDidUpdate(previousProps: RouteErrorBoundaryProps) {
    if (
      this.state.error &&
      previousProps.resetKey !== this.props.resetKey
    ) {
      this.setState({ error: null });
    }
  }

  render() {
    if (this.state.error) {
      return (
        <div className="public-shell min-h-svh flex items-center justify-center p-6">
          <div className="card accent-pink max-w-[600px] w-full text-center">
             <h2 className="text-pink mb-4" style={{ fontFamily: "var(--font-pixel)", fontSize: "1.5rem" }}>
              &gt; RENDER_ERROR
            </h2>
            <div className="bg-panel-deep p-4 text-left border-2 border-dashed border-pink/30 mb-6" style={{ fontFamily: "var(--font-terminal)" }}>
              <p className="text-pink mb-2">CRITICAL_EXCEPTION_DETECTED:</p>
              <code className="text-text block overflow-x-auto">
                {this.state.error.message || "Unknown error"}
              </code>
            </div>
            <button 
              onClick={() => window.location.reload()}
              className="pix-btn pix-btn-pink hover-wiggle"
            >
              RELOAD_SYSTEM
            </button>
          </div>
        </div>
      );
    }

    return this.props.children;
  }
}
