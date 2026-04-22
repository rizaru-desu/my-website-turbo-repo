package auth

import (
	"context"
	"testing"
	"time"

	authusecase "api/internal/usecase/auth"
)

func TestJWTManagerSignsAndVerifiesIssuerAndExpiry(t *testing.T) {
	manager := NewJWTManager("test-secret", "portfolio")
	now := time.Now().UTC().Add(-time.Minute)

	token, err := manager.Sign(context.Background(), authusecase.TokenClaims{
		SessionID:  "session-1",
		UserID:     "user-1",
		Name:       "Admin",
		Email:      "admin@example.com",
		Role:       "admin",
		IssuedAt:   now,
		ExpiresAt:  now.Add(time.Hour),
		RememberMe: true,
	})
	if err != nil {
		t.Fatalf("expected sign to succeed, got %v", err)
	}

	claims, err := manager.Verify(context.Background(), token)
	if err != nil {
		t.Fatalf("expected verify to succeed, got %v", err)
	}

	if claims.Issuer != "portfolio" {
		t.Fatalf("expected issuer portfolio, got %q", claims.Issuer)
	}
	if claims.SessionID != "session-1" || claims.UserID != "user-1" || !claims.RememberMe {
		t.Fatalf("unexpected claims: %+v", claims)
	}
}

func TestJWTManagerRejectsWrongIssuer(t *testing.T) {
	signer := NewJWTManager("test-secret", "portfolio")
	verifier := NewJWTManager("test-secret", "other")
	now := time.Now().UTC()

	token, err := signer.Sign(context.Background(), authusecase.TokenClaims{
		SessionID: "session-1",
		UserID:    "user-1",
		Email:     "admin@example.com",
		IssuedAt:  now,
		ExpiresAt: now.Add(time.Hour),
	})
	if err != nil {
		t.Fatalf("expected sign to succeed, got %v", err)
	}

	if _, err := verifier.Verify(context.Background(), token); err == nil {
		t.Fatal("expected wrong issuer to be rejected")
	}
}
