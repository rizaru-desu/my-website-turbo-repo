package auth

import (
	"context"
	"fmt"

	"api/internal/infrastructure/persistence/ent"
	"api/internal/infrastructure/persistence/ent/twofactor"
	"api/internal/infrastructure/persistence/ent/user"
	authusecase "api/internal/usecase/auth"
)

type EntTwoFactorRepository struct {
	client *ent.Client
}

func NewEntTwoFactorRepository(client *ent.Client) *EntTwoFactorRepository {
	return &EntTwoFactorRepository{client: client}
}

func (r *EntTwoFactorRepository) FindByUserID(ctx context.Context, userID string) (authusecase.TwoFactorRecord, error) {
	if r.client == nil {
		return authusecase.TwoFactorRecord{}, fmt.Errorf("two factor database client is not configured")
	}

	record, err := r.client.TwoFactor.Query().
		Where(twofactor.UserIDEQ(userID)).
		Only(ctx)
	if err != nil {
		return authusecase.TwoFactorRecord{}, err
	}

	return authusecase.TwoFactorRecord{
		ID:          record.ID,
		UserID:      record.UserID,
		Secret:      record.Secret,
		BackupCodes: record.BackupCodes,
	}, nil
}

func (r *EntTwoFactorRepository) Create(ctx context.Context, id, userID, encryptedSecret, backupCodes string) error {
	if r.client == nil {
		return fmt.Errorf("two factor database client is not configured")
	}

	_, err := r.client.TwoFactor.Create().
		SetID(id).
		SetUserID(userID).
		SetSecret(encryptedSecret).
		SetBackupCodes(backupCodes).
		Save(ctx)
	return err
}

func (r *EntTwoFactorRepository) Delete(ctx context.Context, userID string) error {
	if r.client == nil {
		return nil
	}

	_, err := r.client.TwoFactor.Delete().Where(twofactor.UserIDEQ(userID)).Exec(ctx)
	return err
}

func (r *EntTwoFactorRepository) UpdateBackupCodes(ctx context.Context, userID, backupCodes string) error {
	if r.client == nil {
		return fmt.Errorf("two factor database client is not configured")
	}

	_, err := r.client.TwoFactor.Update().
		Where(twofactor.UserIDEQ(userID)).
		SetBackupCodes(backupCodes).
		Save(ctx)
	return err
}

func (r *EntTwoFactorRepository) SetUserTwoFactorEnabled(ctx context.Context, userID string, enabled bool) error {
	if r.client == nil {
		return fmt.Errorf("two factor database client is not configured")
	}

	_, err := r.client.User.Update().
		Where(user.IDEQ(userID)).
		SetTwoFactorEnabled(enabled).
		Save(ctx)
	return err
}
