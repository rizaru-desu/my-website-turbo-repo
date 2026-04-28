import { useNavigate } from "react-router-dom";
import { Button, Input, Switch } from "@repo/ui";
import { AppLayout } from "../components/layout";
import { navItems, pageMeta } from "../constants/navigation";
import { useAuth } from "../hooks/useAuth";
import type { PageKey } from "../types/dashboard";

export default function SettingsPage() {
  const navigate = useNavigate();
  const { logout } = useAuth();

  const handleLogout = async () => {
    await logout();
    navigate("/login");
  };

  const handleNavigate = (newPage: PageKey) => {
    if (newPage === "settings") return;
    navigate("/dashboard", { state: { initialPage: newPage } });
  };

  return (
    <AppLayout
      currentPage="settings"
      onNavigate={handleNavigate}
      onLogout={handleLogout}
      navItems={navItems}
      pageTitle={pageMeta.settings}
    >
      <div className="grid-2">
        <div className="cms-card accent-purple">
          <div className="card-title-row">
            <div className="card-title">&gt; ACCOUNT.CFG</div>
          </div>
          <div className="cms-form-row">
            <label>EMAIL</label>
            <Input defaultValue="kai@morikawa.studio" />
          </div>
          <div className="cms-form-row">
            <label>HANDLE</label>
            <Input defaultValue="@kai_pixels" />
          </div>
          <div className="setting-row">
            <span>EMAIL NOTIFICATIONS</span>
            <Switch defaultChecked aria-label="Email notifications" />
          </div>
          <div className="pix-divider" />
          <div className="row-actions end">
            <Button>SAVE CHANGES</Button>
          </div>
        </div>
      </div>
    </AppLayout>
  );
}
