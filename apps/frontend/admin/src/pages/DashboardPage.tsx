import { useNavigate } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";
import AdminDashboard from "../components/dashboard/AdminDashboard";

export default function DashboardPage() {
  const navigate = useNavigate();
  const { logout } = useAuth();

  const handleLogout = async () => {
    await logout();
    navigate("/login");
  };

  return <AdminDashboard onLogout={handleLogout} />;
}
