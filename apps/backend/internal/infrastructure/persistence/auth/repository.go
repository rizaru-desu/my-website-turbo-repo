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

	role := ""
	if record.Role != nil {
		role = *record.Role
	}

	active := true
	if record.Banned != nil && *record.Banned {
		active = false
	}
	if record.BanExpires != nil && record.BanExpires.After(time.Now()) {
		active = false
	}

	return authusecase.CredentialAccount{
		UserID:       record.ID,
		Name:         record.Name,
		Email:        record.Email,
		Role:         role,
		PasswordHash: *accounts[0].Password,
		Active:       active,
	}, nil
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

func (r *EntSessionRepository) RevokeByToken(ctx context.Context, token string) error {
	if r.client == nil {
		return nil
	}

	_, err := r.client.Session.Delete().Where(session.TokenEQ(token)).Exec(ctx)
	return err
}
