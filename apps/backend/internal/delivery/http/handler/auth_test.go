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
	result  authusecase.LoginResult
	options authusecase.Options
}

func (s stubAuthService) Login(context.Context, authusecase.LoginCommand) (authusecase.LoginResult, error) {
	return s.result, nil
}

func (s stubAuthService) Logout(context.Context, string) error {
	return nil
}

func (s stubAuthService) VerifyToken(context.Context, string) (authusecase.TokenClaims, error) {
	return authusecase.TokenClaims{}, nil
}

func (s stubAuthService) Options() authusecase.Options {
	return s.options
}

func TestAuthLoginSetsJWTCookie(t *testing.T) {
	expiresAt := time.Now().UTC().Add(30 * 24 * time.Hour)
	handler := NewAuthHandler(stubAuthService{
		result: authusecase.LoginResult{
			Token:      "jwt-token",
			TokenType:  "Bearer",
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
}

func TestAuthOptionsExposeSettings(t *testing.T) {
	handler := NewAuthHandler(stubAuthService{
		options: authusecase.Options{
			Issuer:         "portfolio",
			AccessTokenTTL: time.Hour,
			RememberMeTTL:  7 * 24 * time.Hour,
		},
	}, AuthCookieConfig{
		Name:     "portfolio_auth",
		Secure:   false,
		HTTPOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	req := httptest.NewRequest(http.MethodGet, authPath+"/options", nil)
	rec := httptest.NewRecorder()

	handler.Options(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
	if !strings.Contains(rec.Body.String(), `"issuer":"portfolio"`) {
		t.Fatalf("expected issuer in options response, got %s", rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), `"cookie_same_site":"lax"`) {
		t.Fatalf("expected same-site option in response, got %s", rec.Body.String())
	}
}
