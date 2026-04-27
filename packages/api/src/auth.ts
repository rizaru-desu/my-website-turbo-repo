import { apiClient } from "./client.js";
import type { AxiosResponse } from "axios";

export interface LoginRequest {
  email: string;
  password: string;
  remember_me: boolean;
}

export interface LoginResponse {
  expires_at: string;
  expires_in: number;
  remember_me: boolean;
  user: {
    id: string;
    name: string;
    email: string;
    role: string;
  };
}

export interface TwoFactorResponse {
  requires_two_factor: boolean;
  two_factor_token: string;
}

export interface VerifyTOTPRequest {
  two_factor_token: string;
  code: string;
}

export interface MeResponse {
  user: {
    id: string;
    name: string;
    email: string;
    role: string;
  };
  session_id: string;
  issuer: string;
  issued_at: string;
  expires_at: string;
  remember_me: boolean;
}

export interface ErrorResponse {
  error: string;
}

export const authService = {
  async login(
    data: LoginRequest
  ): Promise<LoginResponse | TwoFactorResponse> {
    const response: AxiosResponse<LoginResponse | TwoFactorResponse> =
      await apiClient.post("/auth/login", data);
    return response.data;
  },

  async verifyTOTP(data: VerifyTOTPRequest): Promise<LoginResponse> {
    const response: AxiosResponse<LoginResponse> = await apiClient.post(
      "/auth/2fa/verify",
      data
    );
    return response.data;
  },

  async me(): Promise<MeResponse> {
    const response: AxiosResponse<MeResponse> =
      await apiClient.get("/auth/me");
    return response.data;
  },

  async logout(): Promise<void> {
    await apiClient.post("/auth/logout");
  },

  async forgotPassword(email: string): Promise<{ status: string; message: string }> {
    const response = await apiClient.post("/auth/forgot-password", { email });
    return response.data;
  },
};
