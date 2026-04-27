import axios from "axios";
import type { AxiosInstance } from "axios";

const DEFAULT_API_BASE_URL = "http://127.0.0.1:3333/api/v1";

function resolveBaseURL(): string {
  const viteEnv = (
    import.meta as unknown as { env?: Record<string, string | undefined> }
  ).env;
  return viteEnv?.VITE_API_BASE_URL ?? DEFAULT_API_BASE_URL;
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
