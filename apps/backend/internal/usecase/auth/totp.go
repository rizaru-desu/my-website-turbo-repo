package auth

import "context"

type TOTPManager interface {
	GenerateSecret(email string) (secret string, qrURL string, err error)
	ValidateCode(secret string, code string) bool
	EncryptSecret(plaintext string) (string, error)
	DecryptSecret(ciphertext string) (string, error)
}

type TwoFactorRepository interface {
	FindByUserID(ctx context.Context, userID string) (TwoFactorRecord, error)
	Create(ctx context.Context, id, userID, encryptedSecret, backupCodes string) error
	Delete(ctx context.Context, userID string) error
	UpdateBackupCodes(ctx context.Context, userID, backupCodes string) error
	SetUserTwoFactorEnabled(ctx context.Context, userID string, enabled bool) error
}

type TwoFactorRecord struct {
	ID          string
	UserID      string
	Secret      string
	BackupCodes string
}

type SetupTOTPResult struct {
	Secret      string   `json:"secret"`
	QRURL       string   `json:"qr_url"`
	BackupCodes []string `json:"backup_codes"`
}

type EnableTOTPCommand struct {
	UserID   string
	Code     string
	Password string
}

type DisableTOTPCommand struct {
	UserID   string
	Password string
}

type VerifyTOTPCommand struct {
	TwoFactorToken string
	Code           string
}

type RegenerateBackupCodesCommand struct {
	UserID   string
	Password string
}
