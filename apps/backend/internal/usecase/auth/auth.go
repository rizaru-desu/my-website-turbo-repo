package auth

import (
	"context"
	"crypto/subtle"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"
)

var (
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrInactiveUser        = errors.New("inactive user")
	ErrInvalidToken        = errors.New("invalid token")
	ErrEmailNotFound       = errors.New("email not found")
	ErrTwoFactorRequired   = errors.New("two factor authentication required")
	ErrInvalidTOTPCode     = errors.New("invalid totp code")
	ErrTwoFactorNotSetup   = errors.New("two factor not configured")
	ErrTwoFactorExists     = errors.New("two factor already enabled")
	ErrInvalidCallbackURL  = errors.New("invalid callback url")
	ErrEmailDeliveryFailed = errors.New("email delivery failed")
	ErrWeakPassword        = errors.New("password does not meet requirements")
)

type CredentialRepository interface {
	FindByEmail(ctx context.Context, email string) (CredentialAccount, error)
	FindByUserID(ctx context.Context, userID string) (CredentialAccount, error)
	MarkEmailVerified(ctx context.Context, userID string) error
	UpdatePassword(ctx context.Context, userID string, passwordHash string) error
}

type SessionRepository interface {
	Create(ctx context.Context, session Session) error
	FindByToken(ctx context.Context, token string) (StoredSession, error)
	RevokeByToken(ctx context.Context, token string) error
}

type TokenManager interface {
	Sign(ctx context.Context, claims TokenClaims) (string, error)
	Verify(ctx context.Context, token string) (TokenClaims, error)
}

type PasswordVerifier interface {
	Verify(hash string, password string) bool
	Hash(password string) (string, error)
}

type VerificationRepository interface {
	Create(ctx context.Context, id string, identifier string, value string, expiresAt time.Time) error
	FindByIdentifier(ctx context.Context, identifier string) (string, error)
	FindIdentifierByValue(ctx context.Context, value string) (string, error)
	DeleteByIdentifier(ctx context.Context, identifier string) error
}

type EmailVerificationLimiter interface {
	Allow(ctx context.Context, email string, now time.Time, cooldown time.Duration, window time.Duration, max int) (bool, error)
}

type MailSender interface {
	Send(to string, subject string, htmlBody string) error
}

type IDGenerator interface {
	NewID() string
}

type Clock interface {
	Now() time.Time
}

type CredentialAccount struct {
	UserID           string
	Name             string
	Email            string
	EmailVerified    bool
	Role             string
	PasswordHash     string
	Active           bool
	TwoFactorEnabled bool
}

type Session struct {
	ID         string
	UserID     string
	Token      string
	ExpiresAt  time.Time
	IPAddress  string
	UserAgent  string
	RememberMe bool
}

type StoredSession struct {
	ID        string
	UserID    string
	Token     string
	ExpiresAt time.Time
	User      AuthenticatedUser
	Active    bool
}

type TokenClaims struct {
	SessionID  string
	UserID     string
	Name       string
	Email      string
	Role       string
	Issuer     string
	IssuedAt   time.Time
	ExpiresAt  time.Time
	RememberMe bool
}

type LoginCommand struct {
	Email      string
	Password   string
	RememberMe bool
	IPAddress  string
	UserAgent  string
}

type LoginResult struct {
	Token             string
	User              AuthenticatedUser
	ExpiresAt         time.Time
	ExpiresIn         time.Duration
	RememberMe        bool
	RequiresTwoFactor bool
	TwoFactorToken    string
}

type AuthenticatedUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role,omitempty"`
}

type ForgotPasswordCommand struct {
	Email string
}

type ResetPasswordCommand struct {
	Email       string
	Token       string
	NewPassword string
}

type ChangePasswordCommand struct {
	UserID          string
	CurrentPassword string
	NewPassword     string
}

type SendVerificationEmailCommand struct {
	Email       string
	CallbackURL string
}

type VerifyEmailCommand struct {
	Token       string
	CallbackURL string
}

type VerifyEmailResult struct {
	Email       string `json:"email"`
	RedirectURL string `json:"redirect_url,omitempty"`
	Verified    bool   `json:"verified"`
}

type Options struct {
	Issuer                    string
	AccessTokenTTL            time.Duration
	RememberMeTTL             time.Duration
	ResetTokenTTL             time.Duration
	EmailVerificationTTL      time.Duration
	EmailVerificationCooldown time.Duration
	EmailVerificationWindow   time.Duration
	EmailVerificationMax      int
	FrontendURL               string
	BackendURL                string
}

type Service struct {
	credentials   CredentialRepository
	sessions      SessionRepository
	tokens        TokenManager
	passwords     PasswordVerifier
	ids           IDGenerator
	clock         Clock
	options       Options
	verifications VerificationRepository
	emailLimiter  EmailVerificationLimiter
	mail          MailSender
	twoFactors    TwoFactorRepository
	totp          TOTPManager
}

func NewService(credentials CredentialRepository, sessions SessionRepository, tokens TokenManager, passwords PasswordVerifier, ids IDGenerator, clock Clock, options Options) *Service {
	if options.ResetTokenTTL == 0 {
		options.ResetTokenTTL = 15 * time.Minute
	}
	if options.EmailVerificationTTL == 0 {
		options.EmailVerificationTTL = 24 * time.Hour
	}
	if options.EmailVerificationCooldown == 0 {
		options.EmailVerificationCooldown = time.Minute
	}
	if options.EmailVerificationWindow == 0 {
		options.EmailVerificationWindow = time.Hour
	}
	if options.EmailVerificationMax == 0 {
		options.EmailVerificationMax = 5
	}

	return &Service{
		credentials: credentials,
		sessions:    sessions,
		tokens:      tokens,
		passwords:   passwords,
		ids:         ids,
		clock:       clock,
		options:     options,
	}
}

func (s *Service) SetVerifications(repo VerificationRepository) { s.verifications = repo }
func (s *Service) SetEmailVerificationLimiter(limiter EmailVerificationLimiter) {
	s.emailLimiter = limiter
}
func (s *Service) SetMailSender(sender MailSender)        { s.mail = sender }
func (s *Service) SetTwoFactors(repo TwoFactorRepository) { s.twoFactors = repo }
func (s *Service) SetTOTP(manager TOTPManager)            { s.totp = manager }

func (s *Service) Login(ctx context.Context, command LoginCommand) (LoginResult, error) {
	account, err := s.credentials.FindByEmail(ctx, command.Email)
	if err != nil {
		return LoginResult{}, ErrInvalidCredentials
	}

	if !account.Active {
		return LoginResult{}, ErrInactiveUser
	}

	if account.PasswordHash == "" || !s.passwords.Verify(account.PasswordHash, command.Password) {
		return LoginResult{}, ErrInvalidCredentials
	}

	if account.TwoFactorEnabled && s.twoFactors != nil && s.verifications != nil {
		twoFactorToken := s.ids.NewID()
		now := s.clock.Now().UTC()
		expiresAt := now.Add(5 * time.Minute)

		identifier := "2fa_token:" + twoFactorToken

		payload := fmt.Sprintf("%s|%s|%v", account.UserID, account.Email, command.RememberMe)
		if err := s.verifications.Create(ctx, s.ids.NewID(), identifier, payload, expiresAt); err != nil {
			return LoginResult{}, fmt.Errorf("create 2fa token: %w", err)
		}

		return LoginResult{
			RequiresTwoFactor: true,
			TwoFactorToken:    twoFactorToken,
		}, nil
	}

	return s.issueSession(ctx, account, command)
}

func (s *Service) issueSession(ctx context.Context, account CredentialAccount, command LoginCommand) (LoginResult, error) {
	now := s.clock.Now().UTC()
	ttl := s.options.AccessTokenTTL
	if command.RememberMe {
		ttl = s.options.RememberMeTTL
	}
	expiresAt := now.Add(ttl)
	sessionID := s.ids.NewID()

	claims := TokenClaims{
		SessionID:  sessionID,
		UserID:     account.UserID,
		Name:       account.Name,
		Email:      account.Email,
		Role:       account.Role,
		Issuer:     s.options.Issuer,
		IssuedAt:   now,
		ExpiresAt:  expiresAt,
		RememberMe: command.RememberMe,
	}

	token, err := s.tokens.Sign(ctx, claims)
	if err != nil {
		return LoginResult{}, err
	}

	if err := s.sessions.Create(ctx, Session{
		ID:         sessionID,
		UserID:     account.UserID,
		Token:      token,
		ExpiresAt:  expiresAt,
		IPAddress:  command.IPAddress,
		UserAgent:  command.UserAgent,
		RememberMe: command.RememberMe,
	}); err != nil {
		return LoginResult{}, err
	}

	return LoginResult{
		Token: token,
		User: AuthenticatedUser{
			ID:    account.UserID,
			Name:  account.Name,
			Email: account.Email,
			Role:  account.Role,
		},
		ExpiresAt:  expiresAt,
		ExpiresIn:  ttl,
		RememberMe: command.RememberMe,
	}, nil
}

func (s *Service) VerifyToken(ctx context.Context, token string) (TokenClaims, error) {
	if token == "" {
		return TokenClaims{}, ErrInvalidToken
	}

	claims, err := s.tokens.Verify(ctx, token)
	if err != nil {
		return TokenClaims{}, ErrInvalidToken
	}

	session, err := s.sessions.FindByToken(ctx, token)
	if err != nil {
		return TokenClaims{}, ErrInvalidToken
	}
	if !session.Active {
		return TokenClaims{}, ErrInactiveUser
	}
	if session.ID != claims.SessionID || session.UserID != claims.UserID {
		return TokenClaims{}, ErrInvalidToken
	}
	if !session.ExpiresAt.After(s.clock.Now().UTC()) {
		_ = s.sessions.RevokeByToken(ctx, token)
		return TokenClaims{}, ErrInvalidToken
	}

	claims.UserID = session.User.ID
	claims.Name = session.User.Name
	claims.Email = session.User.Email
	claims.Role = session.User.Role
	claims.ExpiresAt = session.ExpiresAt

	return claims, nil
}

func (s *Service) Logout(ctx context.Context, token string) error {
	if token == "" {
		return nil
	}

	return s.sessions.RevokeByToken(ctx, token)
}

func (s *Service) ForgotPassword(ctx context.Context, command ForgotPasswordCommand) error {
	if s.verifications == nil || s.mail == nil {
		return fmt.Errorf("forgot password is not configured")
	}

	account, err := s.credentials.FindByEmail(ctx, command.Email)
	if err != nil {
		log.Printf("forgot password: email %q not found (silent)", command.Email)
		return nil
	}

	if !account.Active {
		return nil
	}

	_ = s.verifications.DeleteByIdentifier(ctx, "password_reset:"+account.Email)

	now := s.clock.Now().UTC()
	expiresAt := now.Add(s.options.ResetTokenTTL)
	resetToken := s.ids.NewID()

	if err := s.verifications.Create(ctx, s.ids.NewID(), "password_reset:"+account.Email, resetToken, expiresAt); err != nil {
		return fmt.Errorf("create reset token: %w", err)
	}

	resetURL, err := s.resetPasswordURL(resetToken, account.Email)
	if err != nil {
		return fmt.Errorf("build reset password url: %w", err)
	}

	htmlBody := passwordResetEmailHTML(account.Name, resetURL, s.options.ResetTokenTTL)

	if err := s.mail.Send(account.Email, "Password Reset Request", htmlBody); err != nil {
		log.Printf("forgot password: failed to send reset email to %s: %v", account.Email, err)
		return fmt.Errorf("%w: send reset email: %v", ErrEmailDeliveryFailed, err)
	}

	log.Printf("forgot password: reset email sent to %s", account.Email)
	return nil
}

func (s *Service) ResetPassword(ctx context.Context, command ResetPasswordCommand) error {
	email := strings.TrimSpace(command.Email)
	token := strings.TrimSpace(command.Token)
	if email == "" || token == "" || s.verifications == nil {
		return ErrInvalidToken
	}
	if err := validateNewPassword(command.NewPassword); err != nil {
		return err
	}

	identifier := "password_reset:" + email
	storedToken, err := s.verifications.FindByIdentifier(ctx, identifier)
	if err != nil || !sameToken(storedToken, token) {
		return ErrInvalidToken
	}

	account, err := s.credentials.FindByEmail(ctx, email)
	if err != nil || !account.Active {
		return ErrInvalidToken
	}

	passwordHash, err := s.passwords.Hash(command.NewPassword)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	if err := s.credentials.UpdatePassword(ctx, account.UserID, passwordHash); err != nil {
		return fmt.Errorf("update password: %w", err)
	}

	if err := s.verifications.DeleteByIdentifier(ctx, identifier); err != nil {
		log.Printf("password reset: failed to invalidate reset token for %s: %v", account.Email, err)
		return fmt.Errorf("invalidate reset token: %w", err)
	}

	log.Printf("password reset: password updated for user %s", account.UserID)
	return nil
}

func (s *Service) ChangePassword(ctx context.Context, command ChangePasswordCommand) error {
	if strings.TrimSpace(command.UserID) == "" {
		return ErrInvalidToken
	}
	if err := validateNewPassword(command.NewPassword); err != nil {
		return err
	}

	cred, err := s.credentialsByUserID(ctx, command.UserID)
	if err != nil || !cred.Active {
		return ErrInvalidCredentials
	}

	if !s.passwords.Verify(cred.PasswordHash, command.CurrentPassword) {
		return ErrInvalidCredentials
	}

	passwordHash, err := s.passwords.Hash(command.NewPassword)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	if err := s.credentials.UpdatePassword(ctx, command.UserID, passwordHash); err != nil {
		return fmt.Errorf("update password: %w", err)
	}

	log.Printf("password change: password updated for user %s", command.UserID)
	return nil
}

func (s *Service) SendVerificationEmail(ctx context.Context, command SendVerificationEmailCommand) error {
	if s.verifications == nil || s.mail == nil {
		return fmt.Errorf("email verification is not configured")
	}

	email := strings.TrimSpace(command.Email)
	if email == "" || !strings.Contains(email, "@") {
		return ErrEmailNotFound
	}

	callbackURL, err := s.normalizeCallbackURL(command.CallbackURL)
	if err != nil {
		return err
	}

	account, err := s.credentials.FindByEmail(ctx, email)
	if err != nil {
		log.Printf("email verification: email %q not found (silent)", email)
		return nil
	}

	if !account.Active {
		log.Printf("email verification: skipped send for inactive user %s", account.Email)
		return nil
	}

	if account.EmailVerified {
		log.Printf("email verification: skipped send because %s is already verified", account.Email)
		return nil
	}

	now := s.clock.Now().UTC()
	if s.emailLimiter != nil {
		allowed, err := s.emailLimiter.Allow(
			ctx,
			account.Email,
			now,
			s.options.EmailVerificationCooldown,
			s.options.EmailVerificationWindow,
			s.options.EmailVerificationMax,
		)
		if err != nil {
			return fmt.Errorf("check email verification rate limit: %w", err)
		}
		if !allowed {
			log.Printf("email verification: resend suppressed for %s", account.Email)
			return nil
		}
	}

	identifier := "email_verification:" + account.Email
	_ = s.verifications.DeleteByIdentifier(ctx, identifier)

	expiresAt := now.Add(s.options.EmailVerificationTTL)
	verificationToken := s.ids.NewID()

	if err := s.verifications.Create(ctx, s.ids.NewID(), identifier, verificationToken, expiresAt); err != nil {
		return fmt.Errorf("create email verification token: %w", err)
	}

	verificationURL, err := s.emailVerificationURL(account.Email, verificationToken, callbackURL)
	if err != nil {
		return fmt.Errorf("build verification url: %w", err)
	}

	htmlBody := verificationEmailHTML(account.Name, verificationURL, s.options.EmailVerificationTTL)

	log.Printf("email verification: attempting smtp send to %s", account.Email)
	if err := s.mail.Send(account.Email, "Verify Your Email", htmlBody); err != nil {
		log.Printf("email verification: failed to send verification email to %s: %v", account.Email, err)
		return fmt.Errorf("%w: send verification email: %v", ErrEmailDeliveryFailed, err)
	}

	log.Printf("email verification: verification email sent to %s", account.Email)
	return nil
}

func (s *Service) VerifyEmail(ctx context.Context, command VerifyEmailCommand) (VerifyEmailResult, error) {
	token := strings.TrimSpace(command.Token)
	if token == "" || s.verifications == nil {
		return VerifyEmailResult{}, ErrInvalidToken
	}

	identifier, err := s.verifications.FindIdentifierByValue(ctx, token)
	if err != nil {
		return VerifyEmailResult{}, ErrInvalidToken
	}

	email, ok := strings.CutPrefix(identifier, "email_verification:")
	if !ok || strings.TrimSpace(email) == "" {
		return VerifyEmailResult{}, ErrInvalidToken
	}

	storedToken, err := s.verifications.FindByIdentifier(ctx, identifier)
	if err != nil || !sameToken(storedToken, token) {
		return VerifyEmailResult{}, ErrInvalidToken
	}

	account, err := s.credentials.FindByEmail(ctx, email)
	if err != nil {
		return VerifyEmailResult{}, ErrInvalidToken
	}

	redirectURL, err := s.normalizeCallbackURL(command.CallbackURL)
	if err != nil {
		return VerifyEmailResult{}, err
	}

	if err := s.credentials.MarkEmailVerified(ctx, account.UserID); err != nil {
		return VerifyEmailResult{}, fmt.Errorf("mark email verified: %w", err)
	}

	_ = s.verifications.DeleteByIdentifier(ctx, identifier)

	return VerifyEmailResult{
		Email:       account.Email,
		RedirectURL: redirectURL,
		Verified:    true,
	}, nil
}

func (s *Service) SetupTOTP(ctx context.Context, command SetupTOTPCommand) (SetupTOTPResult, error) {
	if s.twoFactors == nil || s.totp == nil {
		return SetupTOTPResult{}, ErrTwoFactorNotSetup
	}

	cred, err := s.credentialsByUserID(ctx, command.UserID)
	if err != nil {
		return SetupTOTPResult{}, fmt.Errorf("find user: %w", err)
	}
	if !s.passwords.Verify(cred.PasswordHash, command.Password) {
		return SetupTOTPResult{}, ErrInvalidCredentials
	}

	existing, err := s.twoFactors.FindByUserID(ctx, command.UserID)
	if err == nil && existing.ID != "" {
		if cred.TwoFactorEnabled {
			return SetupTOTPResult{}, ErrTwoFactorExists
		}

		if err := s.twoFactors.Delete(ctx, command.UserID); err != nil {
			return SetupTOTPResult{}, fmt.Errorf("replace pending totp setup: %w", err)
		}
	}

	secret, qrURL, qrCode, err := s.totp.GenerateSecret(cred.Email)
	if err != nil {
		return SetupTOTPResult{}, fmt.Errorf("generate totp secret: %w", err)
	}

	encrypted, err := s.totp.EncryptSecret(secret)
	if err != nil {
		return SetupTOTPResult{}, fmt.Errorf("encrypt totp secret: %w", err)
	}

	backupCodes, err := generateBackupCodes(8)
	if err != nil {
		return SetupTOTPResult{}, fmt.Errorf("generate backup codes: %w", err)
	}

	hashedCodes := hashBackupCodes(backupCodes)

	if err := s.twoFactors.Create(ctx, s.ids.NewID(), command.UserID, encrypted, hashedCodes); err != nil {
		return SetupTOTPResult{}, fmt.Errorf("store totp: %w", err)
	}

	return SetupTOTPResult{
		Secret:      secret,
		QRURL:       qrURL,
		QRCode:      qrCode,
		BackupCodes: backupCodes,
	}, nil
}

func (s *Service) GetTOTPStatus(ctx context.Context, userID string) (TOTPStatusResult, error) {
	cred, err := s.credentialsByUserID(ctx, userID)
	if err != nil {
		return TOTPStatusResult{}, fmt.Errorf("find user: %w", err)
	}

	return TOTPStatusResult{Enabled: cred.TwoFactorEnabled}, nil
}

func (s *Service) EnableTOTP(ctx context.Context, command EnableTOTPCommand) error {
	if s.twoFactors == nil || s.totp == nil {
		return ErrTwoFactorNotSetup
	}

	record, err := s.twoFactors.FindByUserID(ctx, command.UserID)
	if err != nil {
		return ErrTwoFactorNotSetup
	}

	secret, err := s.totp.DecryptSecret(record.Secret)
	if err != nil {
		return fmt.Errorf("decrypt totp secret: %w", err)
	}

	if !s.totp.ValidateCode(secret, command.Code) {
		return ErrInvalidTOTPCode
	}

	if err := s.twoFactors.SetUserTwoFactorEnabled(ctx, command.UserID, true); err != nil {
		return fmt.Errorf("enable 2fa: %w", err)
	}

	log.Printf("2fa: enabled for user %s", command.UserID)
	return nil
}

func (s *Service) DisableTOTP(ctx context.Context, command DisableTOTPCommand) error {
	if s.twoFactors == nil {
		return ErrTwoFactorNotSetup
	}

	cred, err := s.credentialsByUserID(ctx, command.UserID)
	if err != nil {
		return fmt.Errorf("find user: %w", err)
	}

	if !s.passwords.Verify(cred.PasswordHash, command.Password) {
		return ErrInvalidCredentials
	}

	if err := s.twoFactors.Delete(ctx, command.UserID); err != nil {
		return fmt.Errorf("delete 2fa: %w", err)
	}

	if err := s.twoFactors.SetUserTwoFactorEnabled(ctx, command.UserID, false); err != nil {
		return fmt.Errorf("disable 2fa flag: %w", err)
	}

	log.Printf("2fa: disabled for user %s", command.UserID)
	return nil
}

func (s *Service) VerifyLoginTOTP(ctx context.Context, command VerifyTOTPCommand) (LoginResult, error) {
	if s.twoFactors == nil || s.totp == nil || s.verifications == nil {
		return LoginResult{}, ErrTwoFactorNotSetup
	}

	verifications, err := s.findVerificationByToken(ctx, command.TwoFactorToken)
	if err != nil {
		return LoginResult{}, ErrInvalidToken
	}

	userID := verifications.userID
	email := verifications.email
	rememberMe := verifications.rememberMe

	record, err := s.twoFactors.FindByUserID(ctx, userID)
	if err != nil {
		return LoginResult{}, ErrTwoFactorNotSetup
	}

	secret, err := s.totp.DecryptSecret(record.Secret)
	if err != nil {
		return LoginResult{}, fmt.Errorf("decrypt totp secret: %w", err)
	}

	isBackupCode := false
	if !s.totp.ValidateCode(secret, command.Code) {
		if !verifyBackupCode(record.BackupCodes, command.Code) {
			return LoginResult{}, ErrInvalidTOTPCode
		}
		isBackupCode = true
	}

	_ = s.verifications.DeleteByIdentifier(ctx, "2fa_token:"+command.TwoFactorToken)

	if isBackupCode {
		updatedCodes := markBackupCodeUsed(record.BackupCodes, command.Code)
		_ = s.twoFactors.UpdateBackupCodes(ctx, userID, updatedCodes)
	}

	account, err := s.credentials.FindByEmail(ctx, email)
	if err != nil {
		return LoginResult{}, ErrInvalidCredentials
	}

	return s.issueSession(ctx, account, LoginCommand{
		Email:      email,
		RememberMe: rememberMe,
	})
}

func (s *Service) RegenerateBackupCodes(ctx context.Context, command RegenerateBackupCodesCommand) ([]string, error) {
	if s.twoFactors == nil {
		return nil, ErrTwoFactorNotSetup
	}

	cred, err := s.credentialsByUserID(ctx, command.UserID)
	if err != nil {
		return nil, fmt.Errorf("find user: %w", err)
	}

	if !s.passwords.Verify(cred.PasswordHash, command.Password) {
		return nil, ErrInvalidCredentials
	}

	_, err = s.twoFactors.FindByUserID(ctx, command.UserID)
	if err != nil {
		return nil, ErrTwoFactorNotSetup
	}

	codes, err := generateBackupCodes(8)
	if err != nil {
		return nil, fmt.Errorf("generate backup codes: %w", err)
	}

	hashed := hashBackupCodes(codes)
	if err := s.twoFactors.UpdateBackupCodes(ctx, command.UserID, hashed); err != nil {
		return nil, fmt.Errorf("update backup codes: %w", err)
	}

	log.Printf("2fa: backup codes regenerated for user %s", command.UserID)
	return codes, nil
}

func (s *Service) credentialsByUserID(ctx context.Context, userID string) (CredentialAccount, error) {
	return s.credentials.FindByUserID(ctx, userID)
}

func (s *Service) emailVerificationURL(email string, token string, callbackURL string) (string, error) {
	baseURL := strings.TrimRight(s.options.BackendURL, "/")
	if baseURL == "" {
		baseURL = strings.TrimRight(s.options.FrontendURL, "/")
	}
	if baseURL == "" {
		baseURL = "http://localhost:3333"
	}

	verificationURL, err := url.Parse(baseURL + "/api/v1/auth/verify-email")
	if err != nil {
		return "", err
	}

	query := verificationURL.Query()
	query.Set("token", token)
	if callbackURL != "" {
		query.Set("callbackURL", callbackURL)
	}
	verificationURL.RawQuery = query.Encode()

	return verificationURL.String(), nil
}

func (s *Service) resetPasswordURL(token string, email string) (string, error) {
	baseURL := strings.TrimRight(s.options.FrontendURL, "/")
	if baseURL == "" {
		baseURL = "http://localhost:3111"
	}

	resetURL, err := url.Parse(baseURL + "/reset-password")
	if err != nil {
		return "", err
	}

	query := resetURL.Query()
	query.Set("token", token)
	query.Set("email", email)
	resetURL.RawQuery = query.Encode()

	return resetURL.String(), nil
}

func (s *Service) normalizeCallbackURL(raw string) (string, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", nil
	}

	frontendURL := strings.TrimRight(s.options.FrontendURL, "/")
	if frontendURL == "" {
		return "", ErrInvalidCallbackURL
	}

	frontendBase, err := url.Parse(frontendURL)
	if err != nil || frontendBase.Scheme == "" || frontendBase.Host == "" {
		return "", ErrInvalidCallbackURL
	}

	parsed, err := url.Parse(raw)
	if err != nil {
		return "", ErrInvalidCallbackURL
	}

	if !parsed.IsAbs() {
		if !strings.HasPrefix(raw, "/") || strings.HasPrefix(raw, "//") {
			return "", ErrInvalidCallbackURL
		}

		return frontendBase.ResolveReference(parsed).String(), nil
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return "", ErrInvalidCallbackURL
	}
	if !strings.EqualFold(parsed.Scheme, frontendBase.Scheme) || !strings.EqualFold(parsed.Host, frontendBase.Host) {
		return "", ErrInvalidCallbackURL
	}

	return parsed.String(), nil
}

func sameToken(expected string, actual string) bool {
	if len(expected) != len(actual) {
		return false
	}

	return subtle.ConstantTimeCompare([]byte(expected), []byte(actual)) == 1
}

func validateNewPassword(password string) error {
	if len([]rune(password)) < 8 || strings.TrimSpace(password) == "" {
		return ErrWeakPassword
	}

	return nil
}
