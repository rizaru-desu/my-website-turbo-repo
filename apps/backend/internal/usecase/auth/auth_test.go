package auth

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"
)

type fakeCredentialRepository struct {
	account            CredentialAccount
	err                error
	verifiedUserID     string
	markVerifiedErr    error
	findByEmailQueries []string
}

func (r *fakeCredentialRepository) FindByEmail(_ context.Context, email string) (CredentialAccount, error) {
	r.findByEmailQueries = append(r.findByEmailQueries, email)
	return r.account, r.err
}

func (r *fakeCredentialRepository) FindByUserID(context.Context, string) (CredentialAccount, error) {
	return r.account, r.err
}

func (r *fakeCredentialRepository) MarkEmailVerified(_ context.Context, userID string) error {
	r.verifiedUserID = userID
	return r.markVerifiedErr
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

type fakeVerificationRepository struct {
	createdIdentifier string
	createdValue      string
	deletedIdentifier string
	value             string
	err               error
}

func (r *fakeVerificationRepository) Create(_ context.Context, _ string, identifier string, value string, _ time.Time) error {
	r.createdIdentifier = identifier
	r.createdValue = value
	return r.err
}

func (r *fakeVerificationRepository) FindByIdentifier(_ context.Context, identifier string) (string, error) {
	if r.err != nil {
		return "", r.err
	}
	if r.createdIdentifier == identifier {
		return r.createdValue, nil
	}
	if r.value != "" {
		return r.value, nil
	}
	return "", errors.New("verification not found")
}

func (r *fakeVerificationRepository) FindIdentifierByValue(_ context.Context, value string) (string, error) {
	if r.err != nil {
		return "", r.err
	}
	if r.createdValue == value {
		return r.createdIdentifier, nil
	}
	return "", errors.New("verification not found")
}

func (r *fakeVerificationRepository) DeleteByIdentifier(_ context.Context, identifier string) error {
	r.deletedIdentifier = identifier
	return nil
}

type fakeMailSender struct {
	to      string
	subject string
	body    string
	count   int
	err     error
}

func (s *fakeMailSender) Send(to string, subject string, htmlBody string) error {
	s.to = to
	s.subject = subject
	s.body = htmlBody
	s.count++
	return s.err
}

type fakeEmailVerificationLimiter struct {
	allowed []bool
	calls   int
}

func (l *fakeEmailVerificationLimiter) Allow(context.Context, string, time.Time, time.Duration, time.Duration, int) (bool, error) {
	l.calls++
	if len(l.allowed) == 0 {
		return true, nil
	}
	allowed := l.allowed[0]
	l.allowed = l.allowed[1:]
	return allowed, nil
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

func TestForgotPasswordSendsResetEmailViaSMTP(t *testing.T) {
	now := time.Date(2026, 4, 21, 8, 0, 0, 0, time.UTC)
	verifications := &fakeVerificationRepository{}
	mail := &fakeMailSender{}
	service := testService(&fakeSessionRepository{}, now)
	service.SetVerifications(verifications)
	service.SetMailSender(mail)
	service.options.FrontendURL = "https://portfolio.example"

	if err := service.ForgotPassword(context.Background(), ForgotPasswordCommand{Email: "admin@example.com"}); err != nil {
		t.Fatalf("expected forgot password email to send, got %v", err)
	}

	if verifications.createdIdentifier != "password_reset:admin@example.com" {
		t.Fatalf("expected password reset identifier, got %q", verifications.createdIdentifier)
	}
	if mail.count != 1 || mail.to != "admin@example.com" || mail.subject != "Password Reset Request" {
		t.Fatalf("unexpected reset email: count=%d to=%q subject=%q", mail.count, mail.to, mail.subject)
	}
	if !strings.Contains(mail.body, "https://portfolio.example/reset-password") || !strings.Contains(mail.body, "token=session-1") {
		t.Fatalf("expected reset password link in email body, got %s", mail.body)
	}
}

func TestForgotPasswordReturnsDeliveryFailure(t *testing.T) {
	now := time.Date(2026, 4, 21, 8, 0, 0, 0, time.UTC)
	service := testService(&fakeSessionRepository{}, now)
	service.SetVerifications(&fakeVerificationRepository{})
	service.SetMailSender(&fakeMailSender{err: errors.New("smtp unavailable")})

	err := service.ForgotPassword(context.Background(), ForgotPasswordCommand{Email: "admin@example.com"})
	if !errors.Is(err, ErrEmailDeliveryFailed) {
		t.Fatalf("expected email delivery failure, got %v", err)
	}
}

func TestSendVerificationEmailStoresTokenAndSendsBackendLinkWithCallback(t *testing.T) {
	now := time.Date(2026, 4, 21, 8, 0, 0, 0, time.UTC)
	verifications := &fakeVerificationRepository{}
	mail := &fakeMailSender{}
	service := testService(&fakeSessionRepository{}, now)
	service.SetVerifications(verifications)
	service.SetMailSender(mail)
	service.options.FrontendURL = "https://portfolio.example"
	service.options.BackendURL = "https://api.example"

	err := service.SendVerificationEmail(context.Background(), SendVerificationEmailCommand{
		Email:       "admin@example.com",
		CallbackURL: "/dashboard",
	})
	if err != nil {
		t.Fatalf("expected verification email to send, got %v", err)
	}

	if verifications.createdIdentifier != "email_verification:admin@example.com" {
		t.Fatalf("expected email verification identifier, got %q", verifications.createdIdentifier)
	}
	if mail.to != "admin@example.com" || mail.subject != "Verify Your Email" {
		t.Fatalf("unexpected email: to=%q subject=%q", mail.to, mail.subject)
	}
	if !strings.Contains(mail.body, "https://api.example/api/v1/auth/verify-email") {
		t.Fatalf("expected backend verification link in email body, got %s", mail.body)
	}
	if strings.Contains(mail.body, "email=admin%40example.com") || strings.Contains(mail.body, "email=admin@example.com") {
		t.Fatalf("expected verification link to avoid email query parameter, got %s", mail.body)
	}
	if !strings.Contains(mail.body, "token=session-1") {
		t.Fatalf("expected token query parameter in verification link, got %s", mail.body)
	}
	if !strings.Contains(mail.body, "callbackURL=https%3A%2F%2Fportfolio.example%2Fdashboard") {
		t.Fatalf("expected normalized callback in verification link, got %s", mail.body)
	}
}

func TestSendVerificationEmailSuppressesWithinCooldown(t *testing.T) {
	now := time.Date(2026, 4, 21, 8, 0, 0, 0, time.UTC)
	verifications := &fakeVerificationRepository{}
	mail := &fakeMailSender{}
	limiter := &fakeEmailVerificationLimiter{allowed: []bool{true, false}}
	service := testService(&fakeSessionRepository{}, now)
	service.SetVerifications(verifications)
	service.SetEmailVerificationLimiter(limiter)
	service.SetMailSender(mail)

	for i := 0; i < 2; i++ {
		if err := service.SendVerificationEmail(context.Background(), SendVerificationEmailCommand{Email: "admin@example.com"}); err != nil {
			t.Fatalf("expected verification request %d to succeed, got %v", i+1, err)
		}
	}

	if mail.count != 1 {
		t.Fatalf("expected only one email within cooldown, got %d", mail.count)
	}
	if limiter.calls != 2 {
		t.Fatalf("expected limiter to be consulted twice, got %d", limiter.calls)
	}
}

func TestSendVerificationEmailAllowsAfterCooldown(t *testing.T) {
	now := time.Date(2026, 4, 21, 8, 0, 0, 0, time.UTC)
	mail := &fakeMailSender{}
	service := testService(&fakeSessionRepository{}, now)
	service.SetVerifications(&fakeVerificationRepository{})
	service.SetEmailVerificationLimiter(&fakeEmailVerificationLimiter{allowed: []bool{true, true}})
	service.SetMailSender(mail)

	for i := 0; i < 2; i++ {
		if err := service.SendVerificationEmail(context.Background(), SendVerificationEmailCommand{Email: "admin@example.com"}); err != nil {
			t.Fatalf("expected verification request %d to succeed, got %v", i+1, err)
		}
	}

	if mail.count != 2 {
		t.Fatalf("expected two emails after cooldown allowance, got %d", mail.count)
	}
}

func TestSendVerificationEmailReturnsDeliveryFailure(t *testing.T) {
	now := time.Date(2026, 4, 21, 8, 0, 0, 0, time.UTC)
	service := testService(&fakeSessionRepository{}, now)
	service.SetVerifications(&fakeVerificationRepository{})
	service.SetEmailVerificationLimiter(&fakeEmailVerificationLimiter{})
	service.SetMailSender(&fakeMailSender{err: errors.New("smtp unavailable")})

	err := service.SendVerificationEmail(context.Background(), SendVerificationEmailCommand{Email: "admin@example.com"})
	if !errors.Is(err, ErrEmailDeliveryFailed) {
		t.Fatalf("expected email delivery failure, got %v", err)
	}
}

func TestSendVerificationEmailUnknownEmailIsSilent(t *testing.T) {
	now := time.Date(2026, 4, 21, 8, 0, 0, 0, time.UTC)
	credentials := &fakeCredentialRepository{err: errors.New("not found")}
	mail := &fakeMailSender{}
	service := NewService(credentials, &fakeSessionRepository{}, fakeTokenManager{}, fakePasswordVerifier{}, fakeIDGenerator{}, fixedClock{now: now}, Options{
		AccessTokenTTL: time.Hour,
		RememberMeTTL:  7 * 24 * time.Hour,
		FrontendURL:    "http://localhost:3111",
		BackendURL:     "http://localhost:3333",
	})
	service.SetVerifications(&fakeVerificationRepository{})
	service.SetEmailVerificationLimiter(&fakeEmailVerificationLimiter{})
	service.SetMailSender(mail)

	if err := service.SendVerificationEmail(context.Background(), SendVerificationEmailCommand{Email: "missing@example.com"}); err != nil {
		t.Fatalf("expected missing email to be silent, got %v", err)
	}
	if mail.count != 0 {
		t.Fatalf("expected missing email not to send mail, got %d sends", mail.count)
	}
}

func TestSendVerificationEmailVerifiedUserIsSilent(t *testing.T) {
	now := time.Date(2026, 4, 21, 8, 0, 0, 0, time.UTC)
	credentials := &fakeCredentialRepository{account: CredentialAccount{
		UserID:        "user-1",
		Name:          "Admin",
		Email:         "admin@example.com",
		EmailVerified: true,
		PasswordHash:  "hash:secret",
		Active:        true,
	}}
	mail := &fakeMailSender{}
	service := NewService(credentials, &fakeSessionRepository{}, fakeTokenManager{}, fakePasswordVerifier{}, fakeIDGenerator{}, fixedClock{now: now}, Options{
		AccessTokenTTL: time.Hour,
		RememberMeTTL:  7 * 24 * time.Hour,
		FrontendURL:    "http://localhost:3111",
		BackendURL:     "http://localhost:3333",
	})
	service.SetVerifications(&fakeVerificationRepository{})
	service.SetEmailVerificationLimiter(&fakeEmailVerificationLimiter{})
	service.SetMailSender(mail)

	if err := service.SendVerificationEmail(context.Background(), SendVerificationEmailCommand{Email: "admin@example.com"}); err != nil {
		t.Fatalf("expected verified email request to be silent, got %v", err)
	}
	if mail.count != 0 {
		t.Fatalf("expected verified user not to send mail, got %d sends", mail.count)
	}
}

func TestSendVerificationEmailInactiveUserIsSilent(t *testing.T) {
	now := time.Date(2026, 4, 21, 8, 0, 0, 0, time.UTC)
	credentials := &fakeCredentialRepository{account: CredentialAccount{
		UserID:       "user-1",
		Name:         "Admin",
		Email:        "admin@example.com",
		PasswordHash: "hash:secret",
		Active:       false,
	}}
	mail := &fakeMailSender{}
	service := NewService(credentials, &fakeSessionRepository{}, fakeTokenManager{}, fakePasswordVerifier{}, fakeIDGenerator{}, fixedClock{now: now}, Options{
		AccessTokenTTL: time.Hour,
		RememberMeTTL:  7 * 24 * time.Hour,
		FrontendURL:    "http://localhost:3111",
		BackendURL:     "http://localhost:3333",
	})
	service.SetVerifications(&fakeVerificationRepository{})
	service.SetEmailVerificationLimiter(&fakeEmailVerificationLimiter{})
	service.SetMailSender(mail)

	if err := service.SendVerificationEmail(context.Background(), SendVerificationEmailCommand{Email: "admin@example.com"}); err != nil {
		t.Fatalf("expected inactive email request to be silent, got %v", err)
	}
	if mail.count != 0 {
		t.Fatalf("expected inactive user not to send mail, got %d sends", mail.count)
	}
}

func TestVerifyEmailMarksUserVerifiedAndReturnsCallback(t *testing.T) {
	now := time.Date(2026, 4, 21, 8, 0, 0, 0, time.UTC)
	credentials := &fakeCredentialRepository{
		account: CredentialAccount{
			UserID:       "user-1",
			Name:         "Admin",
			Email:        "admin@example.com",
			PasswordHash: "hash:secret",
			Active:       true,
		},
	}
	sessions := &fakeSessionRepository{}
	service := NewService(
		credentials,
		sessions,
		fakeTokenManager{},
		fakePasswordVerifier{},
		fakeIDGenerator{},
		fixedClock{now: now},
		Options{
			Issuer:         "portfolio",
			AccessTokenTTL: time.Hour,
			RememberMeTTL:  7 * 24 * time.Hour,
			FrontendURL:    "https://portfolio.example",
			BackendURL:     "https://api.example",
		},
	)
	verifications := &fakeVerificationRepository{}
	service.SetVerifications(verifications)
	service.SetMailSender(&fakeMailSender{})

	if err := service.SendVerificationEmail(context.Background(), SendVerificationEmailCommand{
		Email:       "admin@example.com",
		CallbackURL: "/verified",
	}); err != nil {
		t.Fatalf("expected verification email to send, got %v", err)
	}

	result, err := service.VerifyEmail(context.Background(), VerifyEmailCommand{
		Token:       "session-1",
		CallbackURL: "/verified",
	})
	if err != nil {
		t.Fatalf("expected email verification to succeed, got %v", err)
	}

	if credentials.verifiedUserID != "user-1" {
		t.Fatalf("expected user to be marked verified, got %q", credentials.verifiedUserID)
	}
	if result.RedirectURL != "https://portfolio.example/verified" {
		t.Fatalf("expected callback redirect, got %q", result.RedirectURL)
	}
	if verifications.deletedIdentifier != "email_verification:admin@example.com" {
		t.Fatalf("expected verification token to be deleted, got %q", verifications.deletedIdentifier)
	}
}

func TestSendVerificationEmailRejectsExternalCallback(t *testing.T) {
	service := testService(&fakeSessionRepository{}, time.Date(2026, 4, 21, 8, 0, 0, 0, time.UTC))
	service.SetVerifications(&fakeVerificationRepository{})
	service.SetMailSender(&fakeMailSender{})
	service.options.FrontendURL = "https://portfolio.example"

	err := service.SendVerificationEmail(context.Background(), SendVerificationEmailCommand{
		Email:       "admin@example.com",
		CallbackURL: "https://evil.example/dashboard",
	})
	if !errors.Is(err, ErrInvalidCallbackURL) {
		t.Fatalf("expected invalid callback url, got %v", err)
	}
}

func testService(sessions *fakeSessionRepository, now time.Time) *Service {
	return NewService(
		&fakeCredentialRepository{
			account: CredentialAccount{
				UserID:        "user-1",
				Name:          "Admin",
				Email:         "admin@example.com",
				EmailVerified: false,
				Role:          "admin",
				PasswordHash:  "hash:secret",
				Active:        true,
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
			FrontendURL:    "http://localhost:3111",
			BackendURL:     "http://localhost:3333",
		},
	)
}
