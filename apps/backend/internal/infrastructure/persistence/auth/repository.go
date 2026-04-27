package auth

import (
	"context"
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

func NewEntVerificationRepository(client *ent.Client, now func() time.Time) *EntVerificationRepository {
	if now == nil {
		now = time.Now
	}

	return &EntVerificationRepository{client: client, now: now}
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
