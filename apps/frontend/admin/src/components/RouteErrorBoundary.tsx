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
    console.error("Admin route render failed:", error, errorInfo);
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
        <main className="login-screen" data-testid="route-error">
          <section className="login-card route-error-card">
            <div className="login-header">
              <div>
                <h1 className="login-title">PIXEL.CMS</h1>
                <p className="login-subtitle">~ RENDER INTERRUPTED</p>
              </div>
            </div>
            <div className="terminal-copy">
              &gt; route_error: true
              <br />
              &gt; message: {this.state.error.message || "Unknown error"}
            </div>
          </section>
        </main>
      );
    }

    return this.props.children;
  }
}
