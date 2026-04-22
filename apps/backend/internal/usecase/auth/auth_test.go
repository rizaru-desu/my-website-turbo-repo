package auth

import (
	"context"
	"testing"
	"time"
)

type fakeCredentialRepository struct {
	account CredentialAccount
	err     error
}

func (r fakeCredentialRepository) FindByEmail(context.Context, string) (CredentialAccount, error) {
	return r.account, r.err
}

type fakeSessionRepository struct {
	session Session
}

func (r *fakeSessionRepository) Create(_ context.Context, session Session) error {
	r.session = session
	return nil
}

func (r *fakeSessionRepository) RevokeByToken(context.Context, string) error {
	return nil
}

type fakeTokenManager struct{}

func (fakeTokenManager) Sign(_ context.Context, claims TokenClaims) (string, error) {
	return "signed:" + claims.SessionID, nil
}

func (fakeTokenManager) Verify(context.Context, string) (TokenClaims, error) {
	return TokenClaims{}, nil
}

type fakePasswordVerifier struct{}

func (fakePasswordVerifier) Verify(hash string, password string) bool {
	return hash == "hash:"+password
}

type fakeIDGenerator struct{}

func (fakeIDGenerator) NewID() string {
	return "session-1"
}

type fixedClock struct {
	now time.Time
}

func (c fixedClock) Now() time.Time {
	return c.now
}

func TestLoginUsesDefaultTTLWithoutRememberMe(t *testing.T) {
	now := time.Date(2026, 4, 21, 8, 0, 0, 0, time.UTC)
	sessions := &fakeSessionRepository{}
	service := testService(sessions, now)

	result, err := service.Login(context.Background(), LoginCommand{
		Email:    "admin@example.com",
		Password: "secret",
	})
	if err != nil {
		t.Fatalf("expected login to succeed, got %v", err)
	}

	if result.ExpiresIn != time.Hour {
		t.Fatalf("expected default ttl %s, got %s", time.Hour, result.ExpiresIn)
	}
	if !result.ExpiresAt.Equal(now.Add(time.Hour)) {
		t.Fatalf("expected expiry %s, got %s", now.Add(time.Hour), result.ExpiresAt)
	}
	if sessions.session.RememberMe {
		t.Fatal("expected session to not be remember-me")
	}
}

func TestLoginUsesRememberMeTTL(t *testing.T) {
	now := time.Date(2026, 4, 21, 8, 0, 0, 0, time.UTC)
	sessions := &fakeSessionRepository{}
	service := testService(sessions, now)

	result, err := service.Login(context.Background(), LoginCommand{
		Email:      "admin@example.com",
		Password:   "secret",
		RememberMe: true,
	})
	if err != nil {
		t.Fatalf("expected login to succeed, got %v", err)
	}

	if result.ExpiresIn != 7*24*time.Hour {
		t.Fatalf("expected remember-me ttl %s, got %s", 7*24*time.Hour, result.ExpiresIn)
	}
	if !sessions.session.ExpiresAt.Equal(now.Add(7 * 24 * time.Hour)) {
		t.Fatalf("expected saved session expiry %s, got %s", now.Add(7*24*time.Hour), sessions.session.ExpiresAt)
	}
	if !sessions.session.RememberMe {
		t.Fatal("expected session to be remember-me")
	}
}

func testService(sessions *fakeSessionRepository, now time.Time) *Service {
	return NewService(
		fakeCredentialRepository{
			account: CredentialAccount{
				UserID:       "user-1",
				Name:         "Admin",
				Email:        "admin@example.com",
				Role:         "admin",
				PasswordHash: "hash:secret",
				Active:       true,
			},
		},
		sessions,
		fakeTokenManager{},
		fakePasswordVerifier{},
		fakeIDGenerator{},
		fixedClock{now: now},
		Options{
			Issuer:         "portfolio",
			AccessTokenTTL: time.Hour,
			RememberMeTTL:  7 * 24 * time.Hour,
		},
	)
}
