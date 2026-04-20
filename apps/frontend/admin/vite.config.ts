import { defineConfig } from "vite";
import react, { reactCompilerPreset } from "@vitejs/plugin-react";
import babel from "@rolldown/plugin-babel";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), babel({ presets: [reactCompilerPreset()] })],
  server: {
    host: "127.0.0.1",
    port: 3111, // Port khusus untuk Admin
    strictPort: true, // Akan error jika port 3000 sudah terpakai
  },
});
