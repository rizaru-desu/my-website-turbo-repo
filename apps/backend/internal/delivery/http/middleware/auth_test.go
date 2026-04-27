package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	authusecase "api/internal/usecase/auth"
)

type stubAuthenticator struct {
	token string
}

func (s *stubAuthenticator) VerifyToken(_ context.Context, token string) (authusecase.TokenClaims, error) {
	s.token = token
	if token == "" {
		return authusecase.TokenClaims{}, authusecase.ErrInvalidToken
	}

	return authusecase.TokenClaims{SessionID: "session-1"}, nil
}

func TestJWTAuthenticationUsesSessionCookie(t *testing.T) {
	authenticator := &stubAuthenticator{}
	handler := NewJWTAuthentication(authenticator, "portfolio_auth")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := AuthClaims(r.Context()); !ok {
			t.Fatal("expected auth claims in request context")
		}
		w.WriteHeader(http.StatusNoContent)
	}))

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.AddCookie(&http.Cookie{Name: "portfolio_auth", Value: "jwt-token"})
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, rec.Code)
	}
	if authenticator.token != "jwt-token" {
		t.Fatalf("expected cookie token to be verified, got %q", authenticator.token)
	}
}

func TestJWTAuthenticationIgnoresBearerToken(t *testing.T) {
	authenticator := &stubAuthenticator{}
	handler := NewJWTAuthentication(authenticator, "portfolio_auth")(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("bearer-only request should not reach next handler")
	}))

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer jwt-token")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, rec.Code)
	}
	if authenticator.token != "" {
		t.Fatalf("expected bearer token to be ignored, got %q", authenticator.token)
	}
}
