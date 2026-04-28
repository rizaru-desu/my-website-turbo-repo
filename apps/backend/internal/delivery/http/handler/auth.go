package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"strings"
	"time"

	authusecase "api/internal/usecase/auth"
)

const authPath = "/api/v1/auth"

type AuthService interface {
	Login(ctx context.Context, command authusecase.LoginCommand) (authusecase.LoginResult, error)
	Logout(ctx context.Context, token string) error
	VerifyToken(ctx context.Context, token string) (authusecase.TokenClaims, error)
	ForgotPassword(ctx context.Context, command authusecase.ForgotPasswordCommand) error
	ResetPassword(ctx context.Context, command authusecase.ResetPasswordCommand) error
	ChangePassword(ctx context.Context, command authusecase.ChangePasswordCommand) error
	SendVerificationEmail(ctx context.Context, command authusecase.SendVerificationEmailCommand) error
	VerifyEmail(ctx context.Context, command authusecase.VerifyEmailCommand) (authusecase.VerifyEmailResult, error)
	SetupTOTP(ctx context.Context, command authusecase.SetupTOTPCommand) (authusecase.SetupTOTPResult, error)
	GetTOTPStatus(ctx context.Context, userID string) (authusecase.TOTPStatusResult, error)
	EnableTOTP(ctx context.Context, command authusecase.EnableTOTPCommand) error
	DisableTOTP(ctx context.Context, command authusecase.DisableTOTPCommand) error
	VerifyLoginTOTP(ctx context.Context, command authusecase.VerifyTOTPCommand) (authusecase.LoginResult, error)
	RegenerateBackupCodes(ctx context.Context, command authusecase.RegenerateBackupCodesCommand) ([]string, error)
}

type AuthCookieConfig struct {
	Name     string
	Path     string
	Domain   string
	Secure   bool
	HTTPOnly bool
	SameSite http.SameSite
}

type AuthHandler struct {
	service AuthService
	cookie  AuthCookieConfig
}

type LoginRequest struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type ResetPasswordRequest struct {
	Email                string `json:"email"`
	Token                string `json:"token"`
	Password             string `json:"password"`
	NewPassword          string `json:"new_password"`
	NewPasswordCamel     string `json:"newPassword"`
	ConfirmPassword      string `json:"confirm_password"`
	ConfirmPasswordCamel string `json:"confirmPassword"`
}

type ChangePasswordRequest struct {
	CurrentPassword      string `json:"current_password"`
	CurrentPasswordCamel string `json:"currentPassword"`
	NewPassword          string `json:"new_password"`
	NewPasswordCamel     string `json:"newPassword"`
	ConfirmPassword      string `json:"confirm_password"`
	ConfirmPasswordCamel string `json:"confirmPassword"`
}

type SendVerificationEmailRequest struct {
	Email            string `json:"email"`
	CallbackURL      string `json:"callbackURL"`
	CallbackURLSnake string `json:"callback_url"`
}

type VerifyTOTPRequest struct {
	TwoFactorToken string `json:"two_factor_token"`
	Code           string `json:"code"`
}

type EnableTOTPRequest struct {
	Code string `json:"code"`
}

type SetupTOTPRequest struct {
	Password string `json:"password"`
}

type DisableTOTPRequest struct {
	Password string `json:"password"`
}

type RegenerateBackupCodesRequest struct {
	Password string `json:"password"`
}

type LoginResponse struct {
	ExpiresAt  time.Time                     `json:"expires_at"`
	ExpiresIn  int64                         `json:"expires_in"`
	RememberMe bool                          `json:"remember_me"`
	User       authusecase.AuthenticatedUser `json:"user"`
}

type MeResponse struct {
	User       authusecase.AuthenticatedUser `json:"user"`
	SessionID  string                        `json:"session_id"`
	Issuer     string                        `json:"issuer"`
	IssuedAt   time.Time                     `json:"issued_at"`
	ExpiresAt  time.Time                     `json:"expires_at"`
	RememberMe bool                          `json:"remember_me"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewAuthHandler(service AuthService, cookie AuthCookieConfig) *AuthHandler {
	return &AuthHandler{service: service, cookie: cookie}
}

func (h *AuthHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST "+authPath+"/login", h.Login)
	mux.HandleFunc("POST "+authPath+"/logout", h.Logout)
	mux.HandleFunc("GET "+authPath+"/me", h.Me)
	mux.HandleFunc("POST "+authPath+"/forgot-password", h.ForgotPassword)
	mux.HandleFunc("POST "+authPath+"/reset-password", h.ResetPassword)
	mux.HandleFunc("POST "+authPath+"/change-password", h.ChangePassword)
	mux.HandleFunc("POST "+authPath+"/email-verification", h.SendVerificationEmail)
	mux.HandleFunc("GET "+authPath+"/verify-email", h.VerifyEmail)
	mux.HandleFunc("GET "+authPath+"/2fa/status", h.GetTOTPStatus)
	mux.HandleFunc("POST "+authPath+"/2fa/setup", h.SetupTOTP)
	mux.HandleFunc("POST "+authPath+"/2fa/enable", h.EnableTOTP)
	mux.HandleFunc("POST "+authPath+"/2fa/disable", h.DisableTOTP)
	mux.HandleFunc("POST "+authPath+"/2fa/verify", h.VerifyLoginTOTP)
	mux.HandleFunc("POST "+authPath+"/2fa/backup-codes/regenerate", h.RegenerateBackupCodes)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if h.service == nil {
		writeJSON(w, http.StatusServiceUnavailable, ErrorResponse{Error: "auth service is not configured"})
		return
	}

	var payload LoginRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&payload); err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid login payload"})
		return
	}

	result, err := h.service.Login(r.Context(), authusecase.LoginCommand{
		Email:      strings.TrimSpace(payload.Email),
		Password:   payload.Password,
		RememberMe: payload.RememberMe,
		IPAddress:  clientIP(r),
		UserAgent:  r.UserAgent(),
	})
	if err != nil {
		status := http.StatusUnauthorized
		if !errors.Is(err, authusecase.ErrInvalidCredentials) && !errors.Is(err, authusecase.ErrInactiveUser) {
			status = http.StatusInternalServerError
		}
		writeJSON(w, status, ErrorResponse{Error: authErrorMessage(err)})
		return
	}

	if result.RequiresTwoFactor {
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"requires_two_factor": true,
			"two_factor_token":    result.TwoFactorToken,
		})
		return
	}

	http.SetCookie(w, h.authCookie(result.Token, result.ExpiresAt, result.ExpiresIn))
	writeJSON(w, http.StatusOK, LoginResponse{
		ExpiresAt:  result.ExpiresAt,
		ExpiresIn:  int64(result.ExpiresIn.Seconds()),
		RememberMe: result.RememberMe,
		User:       result.User,
	})
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	if h.service == nil {
		writeJSON(w, http.StatusServiceUnavailable, ErrorResponse{Error: "auth service is not configured"})
		return
	}

	token := sessionTokenFromCookie(r, h.cookie.Name)
	if token == "" {
		writeJSON(w, http.StatusUnauthorized, ErrorResponse{Error: "invalid or expired token"})
		return
	}

	if err := h.service.Logout(r.Context(), token); err != nil {
		writeJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "failed to logout"})
		return
	}

	http.SetCookie(w, h.clearCookie())
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) {
	if h.service == nil {
		writeJSON(w, http.StatusServiceUnavailable, ErrorResponse{Error: "auth service is not configured"})
		return
	}

	claims, err := h.service.VerifyToken(r.Context(), sessionTokenFromCookie(r, h.cookie.Name))
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, ErrorResponse{Error: "invalid or expired token"})
		return
	}

	writeJSON(w, http.StatusOK, MeResponse{
		User: authusecase.AuthenticatedUser{
			ID:    claims.UserID,
			Name:  claims.Name,
			Email: claims.Email,
			Role:  claims.Role,
		},
		SessionID:  claims.SessionID,
		Issuer:     claims.Issuer,
		IssuedAt:   claims.IssuedAt,
		ExpiresAt:  claims.ExpiresAt,
		RememberMe: claims.RememberMe,
	})
}

func (h *AuthHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	if h.service == nil {
		writeJSON(w, http.StatusServiceUnavailable, ErrorResponse{Error: "auth service is not configured"})
		return
	}

	var payload ForgotPasswordRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&payload); err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid request payload"})
		return
	}

	email := strings.TrimSpace(payload.Email)
	if email == "" || !strings.Contains(email, "@") {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "valid email is required"})
		return
	}

	if err := h.service.ForgotPassword(r.Context(), authusecase.ForgotPasswordCommand{Email: email}); err != nil {
		status := http.StatusInternalServerError
		message := "failed to process request"
		if errors.Is(err, authusecase.ErrEmailDeliveryFailed) {
			status = http.StatusBadGateway
			message = "failed to send reset email"
		}
		writeJSON(w, status, ErrorResponse{Error: message})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "ok", "message": "if the email exists, a reset link has been sent"})
}

func (h *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	if h.service == nil {
		writeJSON(w, http.StatusServiceUnavailable, ErrorResponse{Error: "auth service is not configured"})
		return
	}

	var payload ResetPasswordRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&payload); err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid request payload"})
		return
	}

	email := strings.TrimSpace(payload.Email)
	token := strings.TrimSpace(payload.Token)
	password := firstQueryValue(payload.NewPassword, payload.NewPasswordCamel, payload.Password)
	confirmPassword := firstQueryValue(payload.ConfirmPassword, payload.ConfirmPasswordCamel)

	if email == "" || !strings.Contains(email, "@") || token == "" {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "valid email and token are required"})
		return
	}
	if confirmPassword != "" && password != confirmPassword {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "passwords do not match"})
		return
	}

	err := h.service.ResetPassword(r.Context(), authusecase.ResetPasswordCommand{
		Email:       email,
		Token:       token,
		NewPassword: password,
	})
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, authusecase.ErrInvalidToken) || errors.Is(err, authusecase.ErrWeakPassword) {
			status = http.StatusBadRequest
		}
		writeJSON(w, status, ErrorResponse{Error: passwordErrorMessage(err)})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "ok", "message": "password has been reset"})
}

func (h *AuthHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	if h.service == nil {
		writeJSON(w, http.StatusServiceUnavailable, ErrorResponse{Error: "auth service is not configured"})
		return
	}

	claims, err := h.service.VerifyToken(r.Context(), sessionTokenFromCookie(r, h.cookie.Name))
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, ErrorResponse{Error: "invalid or expired token"})
		return
	}

	var payload ChangePasswordRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&payload); err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid request payload"})
		return
	}

	currentPassword := firstQueryValue(payload.CurrentPassword, payload.CurrentPasswordCamel)
	newPassword := firstQueryValue(payload.NewPassword, payload.NewPasswordCamel)
	confirmPassword := firstQueryValue(payload.ConfirmPassword, payload.ConfirmPasswordCamel)
	if confirmPassword != "" && newPassword != confirmPassword {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "passwords do not match"})
		return
	}

	err = h.service.ChangePassword(r.Context(), authusecase.ChangePasswordCommand{
		UserID:          claims.UserID,
		CurrentPassword: currentPassword,
		NewPassword:     newPassword,
	})
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, authusecase.ErrInvalidCredentials) {
			status = http.StatusUnauthorized
		} else if errors.Is(err, authusecase.ErrWeakPassword) {
			status = http.StatusBadRequest
		}
		writeJSON(w, status, ErrorResponse{Error: passwordErrorMessage(err)})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "ok", "message": "password has been changed"})
}

func (h *AuthHandler) SendVerificationEmail(w http.ResponseWriter, r *http.Request) {
	if h.service == nil {
		writeJSON(w, http.StatusServiceUnavailable, ErrorResponse{Error: "auth service is not configured"})
		return
	}

	var payload SendVerificationEmailRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&payload); err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid request payload"})
		return
	}

	email := strings.TrimSpace(payload.Email)
	if email == "" || !strings.Contains(email, "@") {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "valid email is required"})
		return
	}

	err := h.service.SendVerificationEmail(r.Context(), authusecase.SendVerificationEmailCommand{
		Email:       email,
		CallbackURL: strings.TrimSpace(firstQueryValue(payload.CallbackURL, payload.CallbackURLSnake)),
	})
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, authusecase.ErrInvalidCallbackURL) {
			status = http.StatusBadRequest
		} else if errors.Is(err, authusecase.ErrEmailDeliveryFailed) {
			status = http.StatusBadGateway
		}
		writeJSON(w, status, ErrorResponse{Error: emailVerificationErrorMessage(err)})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "ok", "message": "if the email exists, a verification link has been sent"})
}

func (h *AuthHandler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	if h.service == nil {
		writeJSON(w, http.StatusServiceUnavailable, ErrorResponse{Error: "auth service is not configured"})
		return
	}

	query := r.URL.Query()
	result, err := h.service.VerifyEmail(r.Context(), authusecase.VerifyEmailCommand{
		Token:       strings.TrimSpace(query.Get("token")),
		CallbackURL: strings.TrimSpace(firstQueryValue(query.Get("callbackURL"), query.Get("callback_url"), query.Get("callbackUrl"))),
	})
	if err != nil {
		status := http.StatusBadRequest
		if !errors.Is(err, authusecase.ErrInvalidToken) && !errors.Is(err, authusecase.ErrInvalidCallbackURL) {
			status = http.StatusInternalServerError
		}
		writeJSON(w, status, ErrorResponse{Error: emailVerificationErrorMessage(err)})
		return
	}

	if result.RedirectURL != "" {
		http.Redirect(w, r, result.RedirectURL, http.StatusSeeOther)
		return
	}

	writeJSON(w, http.StatusOK, result)
}

func (h *AuthHandler) authCookie(token string, expiresAt time.Time, ttl time.Duration) *http.Cookie {
	return &http.Cookie{
		Name:     h.cookie.Name,
		Value:    token,
		Path:     h.cookie.Path,
		Domain:   h.cookie.Domain,
		Expires:  expiresAt,
		MaxAge:   int(ttl.Seconds()),
		Secure:   h.cookie.Secure,
		HttpOnly: h.cookie.HTTPOnly,
		SameSite: h.cookie.SameSite,
	}
}

func (h *AuthHandler) clearCookie() *http.Cookie {
	return &http.Cookie{
		Name:     h.cookie.Name,
		Value:    "",
		Path:     h.cookie.Path,
		Domain:   h.cookie.Domain,
		MaxAge:   -1,
		Secure:   h.cookie.Secure,
		HttpOnly: h.cookie.HTTPOnly,
		SameSite: h.cookie.SameSite,
	}
}

func sessionTokenFromCookie(r *http.Request, cookieName string) string {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return ""
	}

	return cookie.Value
}

func clientIP(r *http.Request) string {
	if forwarded := strings.TrimSpace(r.Header.Get("X-Forwarded-For")); forwarded != "" {
		ip, _, _ := strings.Cut(forwarded, ",")
		return strings.TrimSpace(ip)
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		return host
	}

	return r.RemoteAddr
}

func (h *AuthHandler) SetupTOTP(w http.ResponseWriter, r *http.Request) {
	if h.service == nil {
		writeJSON(w, http.StatusServiceUnavailable, ErrorResponse{Error: "auth service is not configured"})
		return
	}

	claims, err := h.service.VerifyToken(r.Context(), sessionTokenFromCookie(r, h.cookie.Name))
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, ErrorResponse{Error: "invalid or expired token"})
		return
	}

	var payload SetupTOTPRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid request payload"})
		return
	}

	result, err := h.service.SetupTOTP(r.Context(), authusecase.SetupTOTPCommand{
		UserID:   claims.UserID,
		Password: payload.Password,
	})
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, authusecase.ErrInvalidCredentials) {
			status = http.StatusUnauthorized
		} else if errors.Is(err, authusecase.ErrTwoFactorExists) {
			status = http.StatusConflict
		}
		writeJSON(w, status, ErrorResponse{Error: err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, result)
}

func (h *AuthHandler) GetTOTPStatus(w http.ResponseWriter, r *http.Request) {
	if h.service == nil {
		writeJSON(w, http.StatusServiceUnavailable, ErrorResponse{Error: "auth service is not configured"})
		return
	}

	claims, err := h.service.VerifyToken(r.Context(), sessionTokenFromCookie(r, h.cookie.Name))
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, ErrorResponse{Error: "invalid or expired token"})
		return
	}

	result, err := h.service.GetTOTPStatus(r.Context(), claims.UserID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "failed to load 2fa status"})
		return
	}

	writeJSON(w, http.StatusOK, result)
}

func (h *AuthHandler) EnableTOTP(w http.ResponseWriter, r *http.Request) {
	if h.service == nil {
		writeJSON(w, http.StatusServiceUnavailable, ErrorResponse{Error: "auth service is not configured"})
		return
	}

	claims, err := h.service.VerifyToken(r.Context(), sessionTokenFromCookie(r, h.cookie.Name))
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, ErrorResponse{Error: "invalid or expired token"})
		return
	}

	var payload EnableTOTPRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid request payload"})
		return
	}

	err = h.service.EnableTOTP(r.Context(), authusecase.EnableTOTPCommand{
		UserID: claims.UserID,
		Code:   strings.TrimSpace(payload.Code),
	})
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, authusecase.ErrInvalidCredentials) {
			status = http.StatusUnauthorized
		} else if errors.Is(err, authusecase.ErrInvalidTOTPCode) {
			status = http.StatusBadRequest
		}
		writeJSON(w, status, ErrorResponse{Error: err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *AuthHandler) DisableTOTP(w http.ResponseWriter, r *http.Request) {
	if h.service == nil {
		writeJSON(w, http.StatusServiceUnavailable, ErrorResponse{Error: "auth service is not configured"})
		return
	}

	claims, err := h.service.VerifyToken(r.Context(), sessionTokenFromCookie(r, h.cookie.Name))
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, ErrorResponse{Error: "invalid or expired token"})
		return
	}

	var payload DisableTOTPRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid request payload"})
		return
	}

	err = h.service.DisableTOTP(r.Context(), authusecase.DisableTOTPCommand{
		UserID:   claims.UserID,
		Password: payload.Password,
	})
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, authusecase.ErrInvalidCredentials) {
			status = http.StatusUnauthorized
		}
		writeJSON(w, status, ErrorResponse{Error: err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *AuthHandler) VerifyLoginTOTP(w http.ResponseWriter, r *http.Request) {
	if h.service == nil {
		writeJSON(w, http.StatusServiceUnavailable, ErrorResponse{Error: "auth service is not configured"})
		return
	}

	var payload VerifyTOTPRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid request payload"})
		return
	}

	result, err := h.service.VerifyLoginTOTP(r.Context(), authusecase.VerifyTOTPCommand{
		TwoFactorToken: strings.TrimSpace(payload.TwoFactorToken),
		Code:           strings.TrimSpace(payload.Code),
	})
	if err != nil {
		status := http.StatusUnauthorized
		if errors.Is(err, authusecase.ErrInvalidTOTPCode) {
			status = http.StatusBadRequest
		}
		writeJSON(w, status, ErrorResponse{Error: authErrorMessage(err)})
		return
	}

	http.SetCookie(w, h.authCookie(result.Token, result.ExpiresAt, result.ExpiresIn))
	writeJSON(w, http.StatusOK, LoginResponse{
		ExpiresAt:  result.ExpiresAt,
		ExpiresIn:  int64(result.ExpiresIn.Seconds()),
		RememberMe: result.RememberMe,
		User:       result.User,
	})
}

func (h *AuthHandler) RegenerateBackupCodes(w http.ResponseWriter, r *http.Request) {
	if h.service == nil {
		writeJSON(w, http.StatusServiceUnavailable, ErrorResponse{Error: "auth service is not configured"})
		return
	}

	claims, err := h.service.VerifyToken(r.Context(), sessionTokenFromCookie(r, h.cookie.Name))
	if err != nil {
		writeJSON(w, http.StatusUnauthorized, ErrorResponse{Error: "invalid or expired token"})
		return
	}

	var payload RegenerateBackupCodesRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid request payload"})
		return
	}

	codes, err := h.service.RegenerateBackupCodes(r.Context(), authusecase.RegenerateBackupCodesCommand{
		UserID:   claims.UserID,
		Password: payload.Password,
	})
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, authusecase.ErrInvalidCredentials) {
			status = http.StatusUnauthorized
		}
		writeJSON(w, status, ErrorResponse{Error: err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{"backup_codes": codes})
}

func authErrorMessage(err error) string {
	switch {
	case errors.Is(err, authusecase.ErrInactiveUser):
		return "user is inactive"
	case errors.Is(err, authusecase.ErrInvalidCredentials):
		return "invalid email or password"
	case errors.Is(err, authusecase.ErrInvalidTOTPCode):
		return "invalid verification code"
	case errors.Is(err, authusecase.ErrInvalidToken):
		return "invalid or expired token"
	default:
		return "failed to authenticate"
	}
}

func emailVerificationErrorMessage(err error) string {
	switch {
	case errors.Is(err, authusecase.ErrInvalidCallbackURL):
		return "invalid callback url"
	case errors.Is(err, authusecase.ErrInvalidToken):
		return "invalid or expired verification token"
	case errors.Is(err, authusecase.ErrEmailDeliveryFailed):
		return "failed to send verification email"
	default:
		return "failed to verify email"
	}
}

func passwordErrorMessage(err error) string {
	switch {
	case errors.Is(err, authusecase.ErrInvalidCredentials):
		return "current password is invalid"
	case errors.Is(err, authusecase.ErrInvalidToken):
		return "invalid or expired reset token"
	case errors.Is(err, authusecase.ErrWeakPassword):
		return "password must be at least 8 characters"
	default:
		return "failed to update password"
	}
}

func firstQueryValue(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}

	return ""
}
