package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type verificationPayload struct {
	userID     string
	email      string
	rememberMe bool
}

func (s *Service) findVerificationByToken(ctx context.Context, token string) (verificationPayload, error) {
	if s.verifications == nil {
		return verificationPayload{}, fmt.Errorf("verifications not configured")
	}

	identifier := "2fa_token:" + token
	value, err := s.verifications.FindByIdentifier(ctx, identifier)
	if err != nil {
		return verificationPayload{}, err
	}

	parts := strings.SplitN(value, "|", 3)
	if len(parts) < 3 {
		return verificationPayload{}, fmt.Errorf("invalid verification payload")
	}

	return verificationPayload{
		userID:     parts[0],
		email:      parts[1],
		rememberMe: parts[2] == "true",
	}, nil
}

func generateBackupCodes(count int) ([]string, error) {
	codes := make([]string, count)
	for i := range codes {
		b := make([]byte, 5)
		if _, err := rand.Read(b); err != nil {
			return nil, fmt.Errorf("generate backup code: %w", err)
		}

		code := fmt.Sprintf("%010d", uint64(b[0])<<32|uint64(b[1])<<24|uint64(b[2])<<16|uint64(b[3])<<8|uint64(b[4]))
		codes[i] = code[:5] + "-" + code[5:10]
	}

	return codes, nil
}

func hashBackupCodes(codes []string) string {
	hashed := make([]string, len(codes))
	for i, code := range codes {
		h := sha256.Sum256([]byte(code))
		hashed[i] = hex.EncodeToString(h[:])
	}

	return strings.Join(hashed, ",")
}

func verifyBackupCode(storedHashes string, code string) bool {
	code = strings.TrimSpace(code)
	h := sha256.Sum256([]byte(code))
	codeHash := hex.EncodeToString(h[:])

	for _, stored := range strings.Split(storedHashes, ",") {
		if stored != "" && stored == codeHash {
			return true
		}
	}

	return false
}

func markBackupCodeUsed(storedHashes string, code string) string {
	code = strings.TrimSpace(code)
	h := sha256.Sum256([]byte(code))
	codeHash := hex.EncodeToString(h[:])

	parts := strings.Split(storedHashes, ",")
	result := make([]string, 0, len(parts))
	for _, stored := range parts {
		if stored != codeHash {
			result = append(result, stored)
		}
	}

	return strings.Join(result, ",")
}
