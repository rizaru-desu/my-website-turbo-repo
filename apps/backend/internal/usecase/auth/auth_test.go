package auth

import (
	"context"
	"errors"
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

func (r fakeCredentialRepository) FindByUserID(context.Context, string) (CredentialAccount, error) {
	return r.account, r.err
}

type fakeSessionRepository struct {
	session      Session
	stored       StoredSession
	revokedToken string
}

func (r *fakeSessionRepository) Create(_ context.Context, session Session) error {
	r.session = session
	return nil
}

func (r *fakeSessionRepository) FindByToken(_ context.Context, token string) (StoredSession, error) {
	if r.stored.Token == "" || r.stored.Token != token {
		return StoredSession{}, errors.New("session not found")
	}

	return r.stored, nil
}

func (r *fakeSessionRepository) RevokeByToken(_ context.Context, token string) error {
	r.revokedToken = token
	return nil
}

type fakeTokenManager struct{}

func (fakeTokenManager) Sign(_ context.Context, claims TokenClaims) (string, error) {
	return "signed:" + claims.SessionID, nil
}

func (fakeTokenManager) Verify(context.Context, string) (TokenClaims, error) {
	return TokenClaims{
		SessionID: "session-1",
		UserID:    "user-1",
		Name:      "Token Name",
		Email:     "token@example.com",
		Role:      "token-role",
		ExpiresAt: time.Now().Add(time.Hour),
	}, nil
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

func TestVerifyTokenRequiresPersistedSession(t *testing.T) {
	now := time.Date(2026, 4, 21, 8, 0, 0, 0, time.UTC)
	service := testService(&fakeSessionRepository{}, now)

	if _, err := service.VerifyToken(context.Background(), "signed:session-1"); !errors.Is(err, ErrInvalidToken) {
		t.Fatalf("expected invalid token for missing persisted session, got %v", err)
	}
}

func TestVerifyTokenReturnsDatabaseBackedSessionClaims(t *testing.T) {
	now := time.Date(2026, 4, 21, 8, 0, 0, 0, time.UTC)
	expiresAt := now.Add(time.Hour)
	service := testService(&fakeSessionRepository{
		stored: StoredSession{
			ID:        "session-1",
			UserID:    "user-1",
			Token:     "signed:session-1",
			ExpiresAt: expiresAt,
			User: AuthenticatedUser{
				ID:    "user-1",
				Name:  "DB Admin",
				Email: "db-admin@example.com",
				Role:  "admin",
			},
			Active: true,
		},
	}, now)

	claims, err := service.VerifyToken(context.Background(), "signed:session-1")
	if err != nil {
		t.Fatalf("expected token to verify, got %v", err)
	}

	if claims.Name != "DB Admin" || claims.Email != "db-admin@example.com" || claims.Role != "admin" {
		t.Fatalf("expected claims to use database user, got %#v", claims)
	}
	if !claims.ExpiresAt.Equal(expiresAt) {
		t.Fatalf("expected database session expiry %s, got %s", expiresAt, claims.ExpiresAt)
	}
}

func TestVerifyTokenRevokesExpiredPersistedSession(t *testing.T) {
	now := time.Date(2026, 4, 21, 8, 0, 0, 0, time.UTC)
	sessions := &fakeSessionRepository{
		stored: StoredSession{
			ID:        "session-1",
			UserID:    "user-1",
			Token:     "signed:session-1",
			ExpiresAt: now.Add(-time.Minute),
			User: AuthenticatedUser{
				ID:    "user-1",
				Name:  "Admin",
				Email: "admin@example.com",
			},
			Active: true,
		},
	}
	service := testService(sessions, now)

	if _, err := service.VerifyToken(context.Background(), "signed:session-1"); !errors.Is(err, ErrInvalidToken) {
		t.Fatalf("expected invalid token for expired session, got %v", err)
	}
	if sessions.revokedToken != "signed:session-1" {
		t.Fatalf("expected expired token to be revoked, got %q", sessions.revokedToken)
	}
}

func TestLogoutRevokesCurrentSessionToken(t *testing.T) {
	sessions := &fakeSessionRepository{}
	service := testService(sessions, time.Date(2026, 4, 21, 8, 0, 0, 0, time.UTC))

	if err := service.Logout(context.Background(), "signed:session-1"); err != nil {
		t.Fatalf("expected logout to succeed, got %v", err)
	}

	if sessions.revokedToken != "signed:session-1" {
		t.Fatalf("expected logout to revoke current session token, got %q", sessions.revokedToken)
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
