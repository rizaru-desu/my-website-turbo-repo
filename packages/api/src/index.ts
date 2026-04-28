export { apiClient, setApiBaseURL } from "./client.js";
export { AxiosError } from "axios";
export type { AxiosInstance, AxiosResponse, AxiosRequestConfig } from "axios";
export { authService } from "./auth.js";
export type {
  LoginRequest,
  LoginResponse,
  TwoFactorResponse,
  VerifyTOTPRequest,
  SetupTOTPResponse,
  SetupTOTPRequest,
  TwoFactorStatusResponse,
  EnableTOTPRequest,
  DisableTOTPRequest,
  RegenerateBackupCodesRequest,
  RegenerateBackupCodesResponse,
  ResetPasswordRequest,
  ChangePasswordRequest,
  MeResponse,
  ErrorResponse,
} from "./auth.js";
