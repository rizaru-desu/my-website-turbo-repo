import axios from "axios";
import type { AxiosInstance } from "axios";

const DEFAULT_API_BASE_URL = "http://127.0.0.1:3333/api/v1";

function resolveBaseURL(): string {
  const viteEnv = (
    import.meta as unknown as { env?: Record<string, string | undefined> }
  ).env;
  if (viteEnv?.VITE_API_BASE_URL) {
    return viteEnv.VITE_API_BASE_URL;
  }

  const location = globalThis.location;
  if (location?.hostname === "localhost" || location?.hostname === "127.0.0.1") {
    return `${location.protocol}//${location.hostname}:3333/api/v1`;
  }

  return DEFAULT_API_BASE_URL;
}

const apiClient: AxiosInstance = axios.create({
  baseURL: resolveBaseURL(),
  headers: {
    "Content-Type": "application/json",
  },
  withCredentials: true,
  timeout: 15_000,
});

function setApiBaseURL(baseURL: string): void {
  apiClient.defaults.baseURL = baseURL;
}

export { apiClient, setApiBaseURL };
