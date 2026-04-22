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
	Options() authusecase.Options
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

type LoginResponse struct {
	Token      string                        `json:"access_token"`
	TokenType  string                        `json:"token_type"`
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

type AuthOptionsResponse struct {
	Issuer               string `json:"issuer"`
	AccessTokenExpiresIn int64  `json:"access_token_expires_in"`
	RememberMeExpiresIn  int64  `json:"remember_me_expires_in"`
	CookieName           string `json:"cookie_name"`
	CookieSecure         bool   `json:"cookie_secure"`
	CookieHTTPOnly       bool   `json:"cookie_http_only"`
	CookieSameSite       string `json:"cookie_same_site"`
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
	mux.HandleFunc("GET "+authPath+"/options", h.Options)
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

	http.SetCookie(w, h.authCookie(result.Token, result.ExpiresAt, result.ExpiresIn))
	writeJSON(w, http.StatusOK, LoginResponse{
		Token:      result.Token,
		TokenType:  result.TokenType,
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

	token := tokenFromRequest(r, h.cookie.Name)
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

	claims, err := h.service.VerifyToken(r.Context(), tokenFromRequest(r, h.cookie.Name))
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

func (h *AuthHandler) Options(w http.ResponseWriter, _ *http.Request) {
	if h.service == nil {
		writeJSON(w, http.StatusServiceUnavailable, ErrorResponse{Error: "auth service is not configured"})
		return
	}

	options := h.service.Options()
	writeJSON(w, http.StatusOK, AuthOptionsResponse{
		Issuer:               options.Issuer,
		AccessTokenExpiresIn: int64(options.AccessTokenTTL.Seconds()),
		RememberMeExpiresIn:  int64(options.RememberMeTTL.Seconds()),
		CookieName:           h.cookie.Name,
		CookieSecure:         h.cookie.Secure,
		CookieHTTPOnly:       h.cookie.HTTPOnly,
		CookieSameSite:       sameSiteName(h.cookie.SameSite),
	})
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

func tokenFromRequest(r *http.Request, cookieName string) string {
	if token := bearerToken(r); token != "" {
		return token
	}

	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return ""
	}

	return cookie.Value
}

func bearerToken(r *http.Request) string {
	header := strings.TrimSpace(r.Header.Get("Authorization"))
	scheme, token, ok := strings.Cut(header, " ")
	if !ok || !strings.EqualFold(scheme, "Bearer") {
		return ""
	}

	return strings.TrimSpace(token)
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

func authErrorMessage(err error) string {
	switch {
	case errors.Is(err, authusecase.ErrInactiveUser):
		return "user is inactive"
	case errors.Is(err, authusecase.ErrInvalidCredentials):
		return "invalid email or password"
	default:
		return "failed to authenticate"
	}
}

func sameSiteName(value http.SameSite) string {
	switch value {
	case http.SameSiteStrictMode:
		return "strict"
	case http.SameSiteNoneMode:
		return "none"
	case http.SameSiteLaxMode:
		return "lax"
	default:
		return "default"
	}
}
