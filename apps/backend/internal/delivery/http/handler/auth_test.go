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
	result                   authusecase.LoginResult
	logoutToken              *string
	verifyToken              *string
	forgotPasswordErr        error
	resetPasswordCommand     *authusecase.ResetPasswordCommand
	resetPasswordErr         error
	changePasswordCommand    *authusecase.ChangePasswordCommand
	changePasswordErr        error
	setupTOTPCommand         *authusecase.SetupTOTPCommand
	tokenClaims              authusecase.TokenClaims
	verificationEmail        *authusecase.SendVerificationEmailCommand
	verifyEmailCommand       *authusecase.VerifyEmailCommand
	verifyEmailResult        authusecase.VerifyEmailResult
	sendVerificationEmailErr error
	verifyEmailErr           error
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
	return s.forgotPasswordErr
}

func (s stubAuthService) ResetPassword(_ context.Context, command authusecase.ResetPasswordCommand) error {
	if s.resetPasswordCommand != nil {
		*s.resetPasswordCommand = command
	}

	return s.resetPasswordErr
}

func (s stubAuthService) ChangePassword(_ context.Context, command authusecase.ChangePasswordCommand) error {
	if s.changePasswordCommand != nil {
		*s.changePasswordCommand = command
	}

	return s.changePasswordErr
}

func (s stubAuthService) SendVerificationEmail(_ context.Context, command authusecase.SendVerificationEmailCommand) error {
	if s.verificationEmail != nil {
		*s.verificationEmail = command
	}

	return s.sendVerificationEmailErr
}

func (s stubAuthService) VerifyEmail(_ context.Context, command authusecase.VerifyEmailCommand) (authusecase.VerifyEmailResult, error) {
	if s.verifyEmailCommand != nil {
		*s.verifyEmailCommand = command
	}

	return s.verifyEmailResult, s.verifyEmailErr
}

func (s stubAuthService) VerifyToken(_ context.Context, token string) (authusecase.TokenClaims, error) {
	if s.verifyToken != nil {
		*s.verifyToken = token
	}
	if token == "" {
		return authusecase.TokenClaims{}, authusecase.ErrInvalidToken
	}

	if s.tokenClaims.UserID != "" {
		return s.tokenClaims, nil
	}

	return authusecase.TokenClaims{UserID: "user-1"}, nil
}

func (s stubAuthService) SetupTOTP(_ context.Context, command authusecase.SetupTOTPCommand) (authusecase.SetupTOTPResult, error) {
	if s.setupTOTPCommand != nil {
		*s.setupTOTPCommand = command
	}

	return authusecase.SetupTOTPResult{}, nil
}

func (s stubAuthService) GetTOTPStatus(context.Context, string) (authusecase.TOTPStatusResult, error) {
	return authusecase.TOTPStatusResult{}, nil
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

func TestForgotPasswordReportsDeliveryFailure(t *testing.T) {
	handler := NewAuthHandler(stubAuthService{
		forgotPasswordErr: authusecase.ErrEmailDeliveryFailed,
	}, AuthCookieConfig{})

	req := httptest.NewRequest(http.MethodPost, authPath+"/forgot-password", strings.NewReader(`{"email":"admin@example.com"}`))
	rec := httptest.NewRecorder()

	handler.ForgotPassword(rec, req)

	if rec.Code != http.StatusBadGateway {
		t.Fatalf("expected status %d, got %d: %s", http.StatusBadGateway, rec.Code, rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), "failed to send reset email") {
		t.Fatalf("expected delivery failure message, got %s", rec.Body.String())
	}
}

func TestResetPasswordPassesPayloadToService(t *testing.T) {
	command := authusecase.ResetPasswordCommand{}
	handler := NewAuthHandler(stubAuthService{resetPasswordCommand: &command}, AuthCookieConfig{})

	req := httptest.NewRequest(http.MethodPost, authPath+"/reset-password", strings.NewReader(`{"email":" admin@example.com ","token":"reset-token","password":"new-secret","confirmPassword":"new-secret"}`))
	rec := httptest.NewRecorder()

	handler.ResetPassword(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d: %s", http.StatusOK, rec.Code, rec.Body.String())
	}
	if command.Email != "admin@example.com" || command.Token != "reset-token" || command.NewPassword != "new-secret" {
		t.Fatalf("unexpected reset command: %#v", command)
	}
}

func TestResetPasswordReportsInvalidToken(t *testing.T) {
	handler := NewAuthHandler(stubAuthService{
		resetPasswordErr: authusecase.ErrInvalidToken,
	}, AuthCookieConfig{})

	req := httptest.NewRequest(http.MethodPost, authPath+"/reset-password", strings.NewReader(`{"email":"admin@example.com","token":"bad-token","password":"new-secret"}`))
	rec := httptest.NewRecorder()

	handler.ResetPassword(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d: %s", http.StatusBadRequest, rec.Code, rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), "invalid or expired reset token") {
		t.Fatalf("expected invalid token message, got %s", rec.Body.String())
	}
}

func TestChangePasswordRequiresSessionCookie(t *testing.T) {
	handler := NewAuthHandler(stubAuthService{}, AuthCookieConfig{Name: "portfolio_auth"})

	req := httptest.NewRequest(http.MethodPost, authPath+"/change-password", strings.NewReader(`{"current_password":"secret","new_password":"new-secret"}`))
	rec := httptest.NewRecorder()

	handler.ChangePassword(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d: %s", http.StatusUnauthorized, rec.Code, rec.Body.String())
	}
}

func TestChangePasswordUsesSessionUserAndPasswordPayloadOnly(t *testing.T) {
	command := authusecase.ChangePasswordCommand{}
	handler := NewAuthHandler(stubAuthService{
		changePasswordCommand: &command,
		tokenClaims: authusecase.TokenClaims{
			UserID: "session-user-42",
			Email:  "signed-in@example.com",
		},
	}, AuthCookieConfig{Name: "portfolio_auth"})

	req := httptest.NewRequest(http.MethodPost, authPath+"/change-password", strings.NewReader(`{"currentPassword":"secret","newPassword":"new-secret","confirmPassword":"new-secret"}`))
	req.AddCookie(&http.Cookie{Name: "portfolio_auth", Value: "jwt-token"})
	rec := httptest.NewRecorder()

	handler.ChangePassword(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d: %s", http.StatusOK, rec.Code, rec.Body.String())
	}
	if command.UserID != "session-user-42" || command.CurrentPassword != "secret" || command.NewPassword != "new-secret" {
		t.Fatalf("unexpected change password command: %#v", command)
	}
}

func TestSetupTOTPRequiresSessionCookie(t *testing.T) {
	handler := NewAuthHandler(stubAuthService{}, AuthCookieConfig{Name: "portfolio_auth"})

	req := httptest.NewRequest(http.MethodPost, authPath+"/2fa/setup", strings.NewReader(`{"password":"current-password"}`))
	rec := httptest.NewRecorder()

	handler.SetupTOTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d: %s", http.StatusUnauthorized, rec.Code, rec.Body.String())
	}
}

func TestSetupTOTPPassesSessionUserAndPasswordPayload(t *testing.T) {
	command := authusecase.SetupTOTPCommand{}
	handler := NewAuthHandler(stubAuthService{
		setupTOTPCommand: &command,
		tokenClaims: authusecase.TokenClaims{
			UserID: "session-user-42",
		},
	}, AuthCookieConfig{Name: "portfolio_auth"})

	req := httptest.NewRequest(http.MethodPost, authPath+"/2fa/setup", strings.NewReader(`{"password":"current-password"}`))
	req.AddCookie(&http.Cookie{Name: "portfolio_auth", Value: "jwt-token"})
	rec := httptest.NewRecorder()

	handler.SetupTOTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d: %s", http.StatusOK, rec.Code, rec.Body.String())
	}
	if command.UserID != "session-user-42" || command.Password != "current-password" {
		t.Fatalf("unexpected setup totp command: %#v", command)
	}
}

func TestSendVerificationEmailPassesCallbackToService(t *testing.T) {
	command := authusecase.SendVerificationEmailCommand{}
	handler := NewAuthHandler(stubAuthService{verificationEmail: &command}, AuthCookieConfig{})

	req := httptest.NewRequest(http.MethodPost, authPath+"/email-verification", strings.NewReader(`{"email":" admin@example.com ","callbackURL":"/dashboard"}`))
	rec := httptest.NewRecorder()

	handler.SendVerificationEmail(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d: %s", http.StatusOK, rec.Code, rec.Body.String())
	}
	if command.Email != "admin@example.com" || command.CallbackURL != "/dashboard" {
		t.Fatalf("unexpected verification command: %#v", command)
	}
}

func TestSendVerificationEmailReportsDeliveryFailure(t *testing.T) {
	handler := NewAuthHandler(stubAuthService{
		sendVerificationEmailErr: authusecase.ErrEmailDeliveryFailed,
	}, AuthCookieConfig{})

	req := httptest.NewRequest(http.MethodPost, authPath+"/email-verification", strings.NewReader(`{"email":"admin@example.com","callbackURL":"/dashboard"}`))
	rec := httptest.NewRecorder()

	handler.SendVerificationEmail(rec, req)

	if rec.Code != http.StatusBadGateway {
		t.Fatalf("expected status %d, got %d: %s", http.StatusBadGateway, rec.Code, rec.Body.String())
	}
	if !strings.Contains(rec.Body.String(), "failed to send verification email") {
		t.Fatalf("expected delivery failure message, got %s", rec.Body.String())
	}
}

func TestVerifyEmailRedirectsToCallback(t *testing.T) {
	command := authusecase.VerifyEmailCommand{}
	handler := NewAuthHandler(stubAuthService{
		verifyEmailCommand: &command,
		verifyEmailResult: authusecase.VerifyEmailResult{
			Email:       "admin@example.com",
			RedirectURL: "https://portfolio.example/dashboard",
			Verified:    true,
		},
	}, AuthCookieConfig{})

	req := httptest.NewRequest(http.MethodGet, authPath+"/verify-email?token=abc&callbackURL=/dashboard", nil)
	rec := httptest.NewRecorder()

	handler.VerifyEmail(rec, req)

	if rec.Code != http.StatusSeeOther {
		t.Fatalf("expected redirect status, got %d: %s", rec.Code, rec.Body.String())
	}
	if location := rec.Header().Get("Location"); location != "https://portfolio.example/dashboard" {
		t.Fatalf("expected redirect location, got %q", location)
	}
	if command.Token != "abc" || command.CallbackURL != "/dashboard" {
		t.Fatalf("unexpected verify command: %#v", command)
	}
}
