package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	authusecase "api/internal/usecase/auth"
)

type stubAuthService struct {
	result      authusecase.LoginResult
	logoutToken *string
	verifyToken *string
}

func (s stubAuthService) Login(context.Context, authusecase.LoginCommand) (authusecase.LoginResult, error) {
	return s.result, nil
}

func (s stubAuthService) Logout(_ context.Context, token string) error {
	if s.logoutToken != nil {
		*s.logoutToken = token
	}

	return nil
}

func (s stubAuthService) ForgotPassword(context.Context, authusecase.ForgotPasswordCommand) error {
	return nil
}

func (s stubAuthService) VerifyToken(_ context.Context, token string) (authusecase.TokenClaims, error) {
	if s.verifyToken != nil {
		*s.verifyToken = token
	}
	if token == "" {
		return authusecase.TokenClaims{}, authusecase.ErrInvalidToken
	}

	return authusecase.TokenClaims{}, nil
}

func (s stubAuthService) SetupTOTP(context.Context, string) (authusecase.SetupTOTPResult, error) {
	return authusecase.SetupTOTPResult{}, nil
}

func (s stubAuthService) EnableTOTP(context.Context, authusecase.EnableTOTPCommand) error {
	return nil
}

func (s stubAuthService) DisableTOTP(context.Context, authusecase.DisableTOTPCommand) error {
	return nil
}

func (s stubAuthService) VerifyLoginTOTP(context.Context, authusecase.VerifyTOTPCommand) (authusecase.LoginResult, error) {
	return authusecase.LoginResult{}, nil
}

func (s stubAuthService) RegenerateBackupCodes(context.Context, authusecase.RegenerateBackupCodesCommand) ([]string, error) {
	return nil, nil
}

func TestAuthLoginSetsJWTCookie(t *testing.T) {
	expiresAt := time.Now().UTC().Add(30 * 24 * time.Hour)
	handler := NewAuthHandler(stubAuthService{
		result: authusecase.LoginResult{
			Token:      "jwt-token",
			ExpiresAt:  expiresAt,
			ExpiresIn:  30 * 24 * time.Hour,
			RememberMe: true,
			User: authusecase.AuthenticatedUser{
				ID:    "user-1",
				Name:  "Admin",
				Email: "admin@example.com",
			},
		},
	}, AuthCookieConfig{
		Name:     "portfolio_auth",
		Path:     "/",
		Secure:   true,
		HTTPOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	req := httptest.NewRequest(http.MethodPost, authPath+"/login", strings.NewReader(`{"email":"admin@example.com","password":"secret","remember_me":true}`))
	rec := httptest.NewRecorder()

	handler.Login(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	cookies := rec.Result().Cookies()
	if len(cookies) != 1 {
		t.Fatalf("expected one auth cookie, got %d", len(cookies))
	}

	cookie := cookies[0]
	if cookie.Name != "portfolio_auth" || cookie.Value != "jwt-token" {
		t.Fatalf("unexpected cookie: %#v", cookie)
	}
	if !cookie.HttpOnly || !cookie.Secure || cookie.SameSite != http.SameSiteStrictMode {
		t.Fatalf("expected secure HttpOnly strict cookie, got %#v", cookie)
	}
	if cookie.MaxAge != int((30 * 24 * time.Hour).Seconds()) {
		t.Fatalf("expected remember-me max age, got %d", cookie.MaxAge)
	}
	if strings.Contains(rec.Body.String(), "access_token") || strings.Contains(rec.Body.String(), "token_type") {
		t.Fatalf("expected login response not to expose bearer token fields, got %s", rec.Body.String())
	}
}

func TestAuthLogoutRequiresSessionCookie(t *testing.T) {
	logoutToken := ""
	handler := NewAuthHandler(stubAuthService{logoutToken: &logoutToken}, AuthCookieConfig{
		Name:     "portfolio_auth",
		Path:     "/",
		HTTPOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	req := httptest.NewRequest(http.MethodPost, authPath+"/logout", nil)
	req.Header.Set("Authorization", "Bearer jwt-token")
	rec := httptest.NewRecorder()

	handler.Logout(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, rec.Code)
	}
	if logoutToken != "" {
		t.Fatalf("expected bearer-only request not to revoke a session, got %q", logoutToken)
	}
}

func TestAuthLogoutRevokesCookieSessionAndClearsCookie(t *testing.T) {
	logoutToken := ""
	handler := NewAuthHandler(stubAuthService{logoutToken: &logoutToken}, AuthCookieConfig{
		Name:     "portfolio_auth",
		Path:     "/",
		HTTPOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	req := httptest.NewRequest(http.MethodPost, authPath+"/logout", nil)
	req.AddCookie(&http.Cookie{Name: "portfolio_auth", Value: "jwt-token"})
	rec := httptest.NewRecorder()

	handler.Logout(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
	if logoutToken != "jwt-token" {
		t.Fatalf("expected cookie token to be revoked, got %q", logoutToken)
	}

	cookies := rec.Result().Cookies()
	if len(cookies) != 1 || cookies[0].Name != "portfolio_auth" || cookies[0].MaxAge != -1 {
		t.Fatalf("expected auth cookie to be cleared, got %#v", cookies)
	}
}

func TestAuthMeIgnoresBearerToken(t *testing.T) {
	verifyToken := "not-called"
	handler := NewAuthHandler(stubAuthService{verifyToken: &verifyToken}, AuthCookieConfig{
		Name: "portfolio_auth",
	})

	req := httptest.NewRequest(http.MethodGet, authPath+"/me", nil)
	req.Header.Set("Authorization", "Bearer jwt-token")
	rec := httptest.NewRecorder()

	handler.Me(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, rec.Code)
	}
	if verifyToken != "" {
		t.Fatalf("expected bearer token to be ignored, got %q", verifyToken)
	}
}

func TestAuthMeUsesCookieSession(t *testing.T) {
	verifyToken := ""
	handler := NewAuthHandler(stubAuthService{verifyToken: &verifyToken}, AuthCookieConfig{
		Name: "portfolio_auth",
	})

	req := httptest.NewRequest(http.MethodGet, authPath+"/me", nil)
	req.AddCookie(&http.Cookie{Name: "portfolio_auth", Value: "jwt-token"})
	rec := httptest.NewRecorder()

	handler.Me(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d: %s", http.StatusOK, rec.Code, rec.Body.String())
	}
	if verifyToken != "jwt-token" {
		t.Fatalf("expected cookie token to be verified, got %q", verifyToken)
	}
}
