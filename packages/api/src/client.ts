import axios from "axios";
import type { AxiosInstance } from "axios";

const apiClient: AxiosInstance = axios.create({
  baseURL: "/api/v1",
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
