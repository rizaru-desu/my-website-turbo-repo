package auth

import (
	"context"
	"errors"
	"time"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInactiveUser       = errors.New("inactive user")
	ErrInvalidToken       = errors.New("invalid token")
)

type CredentialRepository interface {
	FindByEmail(ctx context.Context, email string) (CredentialAccount, error)
}

type SessionRepository interface {
	Create(ctx context.Context, session Session) error
	RevokeByToken(ctx context.Context, token string) error
}

type TokenManager interface {
	Sign(ctx context.Context, claims TokenClaims) (string, error)
	Verify(ctx context.Context, token string) (TokenClaims, error)
}

type PasswordVerifier interface {
	Verify(hash string, password string) bool
}

type IDGenerator interface {
	NewID() string
}

type Clock interface {
	Now() time.Time
}

type CredentialAccount struct {
	UserID       string
	Name         string
	Email        string
	Role         string
	PasswordHash string
	Active       bool
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
	Token      string
	TokenType  string
	User       AuthenticatedUser
	ExpiresAt  time.Time
	ExpiresIn  time.Duration
	RememberMe bool
}

type AuthenticatedUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role,omitempty"`
}

type Options struct {
	Issuer         string
	AccessTokenTTL time.Duration
	RememberMeTTL  time.Duration
	Cookie         CookieOptions
}

type CookieOptions struct {
	Name     string
	Secure   bool
	HTTPOnly bool
	SameSite string
}

type Service struct {
	credentials CredentialRepository
	sessions    SessionRepository
	tokens      TokenManager
	passwords   PasswordVerifier
	ids         IDGenerator
	clock       Clock
	options     Options
}

func NewService(credentials CredentialRepository, sessions SessionRepository, tokens TokenManager, passwords PasswordVerifier, ids IDGenerator, clock Clock, options Options) *Service {
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
		Token:     token,
		TokenType: "Bearer",
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

	return s.tokens.Verify(ctx, token)
}

func (s *Service) Logout(ctx context.Context, token string) error {
	if token == "" {
		return nil
	}

	return s.sessions.RevokeByToken(ctx, token)
}

func (s *Service) Options() Options {
	return s.options
}
