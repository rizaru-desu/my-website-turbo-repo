import { Navigate } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";

interface GuestRouteProps {
  children: React.ReactNode;
}

export default function GuestRoute({ children }: GuestRouteProps) {
  const { isAuthenticated, isLoading } = useAuth();

  if (isLoading) {
    return (
      <main className="login-screen">
        <div style={{ color: "white", textAlign: "center", padding: "2rem" }}>
          LOADING...
        </div>
      </main>
    );
  }

  if (isAuthenticated) {
    return <Navigate to="/dashboard" replace />;
  }

  return <>{children}</>;
}
