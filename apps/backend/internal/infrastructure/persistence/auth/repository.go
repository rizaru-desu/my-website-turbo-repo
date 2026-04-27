package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"api/internal/infrastructure/persistence/ent"
	"api/internal/infrastructure/persistence/ent/account"
	"api/internal/infrastructure/persistence/ent/session"
	"api/internal/infrastructure/persistence/ent/user"
	"api/internal/infrastructure/persistence/ent/verification"
	authusecase "api/internal/usecase/auth"
)

type EntCredentialRepository struct {
	client *ent.Client
}

type EntSessionRepository struct {
	client *ent.Client
	now    func() time.Time
}

func NewEntCredentialRepository(client *ent.Client) *EntCredentialRepository {
	return &EntCredentialRepository{client: client}
}

func NewEntSessionRepository(client *ent.Client, now func() time.Time) *EntSessionRepository {
	if now == nil {
		now = time.Now
	}

	return &EntSessionRepository{client: client, now: now}
}

func (r *EntCredentialRepository) FindByEmail(ctx context.Context, email string) (authusecase.CredentialAccount, error) {
	if r.client == nil {
		return authusecase.CredentialAccount{}, fmt.Errorf("auth database client is not configured")
	}

	record, err := r.client.User.Query().
		Where(user.EmailEqualFold(strings.TrimSpace(email))).
		WithAccounts(func(query *ent.AccountQuery) {
			query.Where(account.PasswordNotNil())
		}).
		Only(ctx)
	if err != nil {
		return authusecase.CredentialAccount{}, err
	}

	accounts, err := record.Edges.AccountsOrErr()
	if err != nil {
		return authusecase.CredentialAccount{}, err
	}

	if len(accounts) == 0 || accounts[0].Password == nil {
		return authusecase.CredentialAccount{}, fmt.Errorf("password account not found")
	}

	return authusecase.CredentialAccount{
		UserID:           record.ID,
		Name:             record.Name,
		Email:            record.Email,
		EmailVerified:    record.EmailVerified,
		Role:             userRole(record),
		PasswordHash:     *accounts[0].Password,
		Active:           userActive(record, time.Now()),
		TwoFactorEnabled: userTwoFactorEnabled(record),
	}, nil
}

func (r *EntCredentialRepository) FindByUserID(ctx context.Context, userID string) (authusecase.CredentialAccount, error) {
	if r.client == nil {
		return authusecase.CredentialAccount{}, fmt.Errorf("auth database client is not configured")
	}

	record, err := r.client.User.Query().
		Where(user.IDEQ(userID)).
		WithAccounts(func(query *ent.AccountQuery) {
			query.Where(account.PasswordNotNil())
		}).
		Only(ctx)
	if err != nil {
		return authusecase.CredentialAccount{}, err
	}

	accounts, err := record.Edges.AccountsOrErr()
	if err != nil {
		return authusecase.CredentialAccount{}, err
	}

	if len(accounts) == 0 || accounts[0].Password == nil {
		return authusecase.CredentialAccount{}, fmt.Errorf("password account not found")
	}

	return authusecase.CredentialAccount{
		UserID:           record.ID,
		Name:             record.Name,
		Email:            record.Email,
		EmailVerified:    record.EmailVerified,
		Role:             userRole(record),
		PasswordHash:     *accounts[0].Password,
		Active:           userActive(record, time.Now()),
		TwoFactorEnabled: userTwoFactorEnabled(record),
	}, nil
}

func (r *EntCredentialRepository) MarkEmailVerified(ctx context.Context, userID string) error {
	if r.client == nil {
		return fmt.Errorf("auth database client is not configured")
	}

	_, err := r.client.User.Update().
		Where(user.IDEQ(userID)).
		SetEmailVerified(true).
		SetUpdatedAt(time.Now().UTC()).
		Save(ctx)
	return err
}

func (r *EntSessionRepository) Create(ctx context.Context, model authusecase.Session) error {
	if r.client == nil {
		return fmt.Errorf("auth database client is not configured")
	}

	now := r.now().UTC()
	builder := r.client.Session.Create().
		SetID(model.ID).
		SetUserID(model.UserID).
		SetToken(model.Token).
		SetExpiresAt(model.ExpiresAt).
		SetCreatedAt(now).
		SetUpdatedAt(now)

	if model.IPAddress != "" {
		builder.SetIPAddress(model.IPAddress)
	}
	if model.UserAgent != "" {
		builder.SetUserAgent(model.UserAgent)
	}

	_, err := builder.Save(ctx)
	return err
}

func (r *EntSessionRepository) FindByToken(ctx context.Context, token string) (authusecase.StoredSession, error) {
	if r.client == nil {
		return authusecase.StoredSession{}, fmt.Errorf("auth database client is not configured")
	}

	record, err := r.client.Session.Query().
		Where(session.TokenEQ(token)).
		WithUser().
		Only(ctx)
	if err != nil {
		return authusecase.StoredSession{}, err
	}

	userRecord, err := record.Edges.UserOrErr()
	if err != nil {
		return authusecase.StoredSession{}, err
	}

	return authusecase.StoredSession{
		ID:        record.ID,
		UserID:    record.UserID,
		Token:     record.Token,
		ExpiresAt: record.ExpiresAt,
		User: authusecase.AuthenticatedUser{
			ID:    userRecord.ID,
			Name:  userRecord.Name,
			Email: userRecord.Email,
			Role:  userRole(userRecord),
		},
		Active: userActive(userRecord, r.now()),
	}, nil
}

func (r *EntSessionRepository) RevokeByToken(ctx context.Context, token string) error {
	if r.client == nil {
		return nil
	}

	_, err := r.client.Session.Delete().Where(session.TokenEQ(token)).Exec(ctx)
	return err
}

type EntVerificationRepository struct {
	client *ent.Client
	now    func() time.Time
}

type EntEmailVerificationLimiter struct {
	client *ent.Client
}

type emailVerificationLimitState struct {
	Attempts []time.Time `json:"attempts"`
}

func NewEntVerificationRepository(client *ent.Client, now func() time.Time) *EntVerificationRepository {
	if now == nil {
		now = time.Now
	}

	return &EntVerificationRepository{client: client, now: now}
}

func NewEntEmailVerificationLimiter(client *ent.Client) *EntEmailVerificationLimiter {
	return &EntEmailVerificationLimiter{client: client}
}

func (r *EntVerificationRepository) Create(ctx context.Context, id string, identifier string, value string, expiresAt time.Time) error {
	if r.client == nil {
		return fmt.Errorf("verification database client is not configured")
	}

	now := r.now().UTC()
	_, err := r.client.Verification.Create().
		SetID(id).
		SetIdentifier(identifier).
		SetValue(value).
		SetExpiresAt(expiresAt).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)
	return err
}

func (r *EntVerificationRepository) DeleteByIdentifier(ctx context.Context, identifier string) error {
	if r.client == nil {
		return nil
	}

	_, err := r.client.Verification.Delete().Where(verification.IdentifierEQ(identifier)).Exec(ctx)
	return err
}

func (r *EntVerificationRepository) FindByIdentifier(ctx context.Context, identifier string) (string, error) {
	if r.client == nil {
		return "", fmt.Errorf("verification database client is not configured")
	}

	record, err := r.client.Verification.Query().
		Where(verification.IdentifierEQ(identifier)).
		Only(ctx)
	if err != nil {
		return "", err
	}

	if record.ExpiresAt.Before(time.Now().UTC()) {
		_ = r.DeleteByIdentifier(ctx, identifier)
		return "", fmt.Errorf("verification expired")
	}

	return record.Value, nil
}

func (r *EntVerificationRepository) FindIdentifierByValue(ctx context.Context, value string) (string, error) {
	if r.client == nil {
		return "", fmt.Errorf("verification database client is not configured")
	}

	record, err := r.client.Verification.Query().
		Where(verification.ValueEQ(value)).
		Only(ctx)
	if err != nil {
		return "", err
	}

	if record.ExpiresAt.Before(time.Now().UTC()) {
		_ = r.DeleteByIdentifier(ctx, record.Identifier)
		return "", fmt.Errorf("verification expired")
	}

	return record.Identifier, nil
}

func (r *EntEmailVerificationLimiter) Allow(ctx context.Context, email string, now time.Time, cooldown time.Duration, window time.Duration, max int) (bool, error) {
	if r.client == nil {
		return false, fmt.Errorf("email verification limiter database client is not configured")
	}
	if max <= 0 {
		return false, nil
	}
	if window <= 0 {
		window = time.Hour
	}

	identifier := "email_verification_limit:" + strings.ToLower(strings.TrimSpace(email))
	record, err := r.client.Verification.Query().
		Where(verification.IdentifierEQ(identifier)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return r.createEmailVerificationLimit(ctx, identifier, now, window)
		}
		return false, err
	}

	if !record.ExpiresAt.After(now) {
		_ = r.deleteEmailVerificationLimit(ctx, identifier)
		return r.createEmailVerificationLimit(ctx, identifier, now, window)
	}

	state := emailVerificationLimitState{}
	if record.Value != "" {
		if err := json.Unmarshal([]byte(record.Value), &state); err != nil {
			state = emailVerificationLimitState{}
		}
	}

	cutoff := now.Add(-window)
	attempts := make([]time.Time, 0, len(state.Attempts)+1)
	for _, attempt := range state.Attempts {
		if attempt.After(cutoff) {
			attempts = append(attempts, attempt)
		}
	}

	if len(attempts) > 0 && cooldown > 0 && attempts[len(attempts)-1].After(now.Add(-cooldown)) {
		return false, nil
	}
	if len(attempts) >= max {
		return false, nil
	}

	attempts = append(attempts, now)
	value, err := json.Marshal(emailVerificationLimitState{Attempts: attempts})
	if err != nil {
		return false, err
	}

	_, err = r.client.Verification.Update().
		Where(verification.IdentifierEQ(identifier)).
		SetValue(string(value)).
		SetExpiresAt(now.Add(window)).
		SetUpdatedAt(now).
		Save(ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *EntEmailVerificationLimiter) createEmailVerificationLimit(ctx context.Context, identifier string, now time.Time, window time.Duration) (bool, error) {
	value, err := json.Marshal(emailVerificationLimitState{Attempts: []time.Time{now}})
	if err != nil {
		return false, err
	}

	_, err = r.client.Verification.Create().
		SetID(identifier).
		SetIdentifier(identifier).
		SetValue(string(value)).
		SetExpiresAt(now.Add(window)).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *EntEmailVerificationLimiter) deleteEmailVerificationLimit(ctx context.Context, identifier string) error {
	_, err := r.client.Verification.Delete().Where(verification.IdentifierEQ(identifier)).Exec(ctx)
	return err
}

func userTwoFactorEnabled(record *ent.User) bool {
	return record.TwoFactorEnabled != nil && *record.TwoFactorEnabled
}

func userRole(record *ent.User) string {
	if record.Role == nil {
		return ""
	}

	return *record.Role
}

func userActive(record *ent.User, now time.Time) bool {
	if record.Banned != nil && *record.Banned {
		return false
	}
	if record.BanExpires != nil && record.BanExpires.After(now) {
		return false
	}

	return true
}
