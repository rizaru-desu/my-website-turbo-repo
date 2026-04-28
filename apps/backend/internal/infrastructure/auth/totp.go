package auth

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"image/png"
	"io"
	"strings"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type TOTPManager struct {
	issuer        string
	encryptionKey [32]byte
}

func NewTOTPManager(issuer string, encryptionSecret string) *TOTPManager {
	key := sha256.Sum256([]byte(encryptionSecret))
	return &TOTPManager{
		issuer:        issuer,
		encryptionKey: key,
	}
}

func (m *TOTPManager) GenerateSecret(email string) (secret string, qrURL string, qrCode string, err error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      m.issuer,
		AccountName: email,
		Period:      30,
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA1,
	})
	if err != nil {
		return "", "", "", fmt.Errorf("totp generate: %w", err)
	}

	image, err := key.Image(220, 220)
	if err != nil {
		return "", "", "", fmt.Errorf("totp qr image: %w", err)
	}

	var buffer bytes.Buffer
	if err := png.Encode(&buffer, image); err != nil {
		return "", "", "", fmt.Errorf("totp qr encode: %w", err)
	}

	qrCode = "data:image/png;base64," + base64.StdEncoding.EncodeToString(buffer.Bytes())
	return key.Secret(), key.URL(), qrCode, nil
}

func (m *TOTPManager) ValidateCode(secret string, code string) bool {
	code = strings.TrimSpace(code)
	if len(code) != 6 {
		return false
	}

	valid, err := totp.ValidateCustom(code, secret, time.Now().UTC(), totp.ValidateOpts{
		Period:    30,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})

	if err != nil {
		return false
	}

	return valid
}

func (m *TOTPManager) EncryptSecret(plaintext string) (string, error) {
	block, err := aes.NewCipher(m.encryptionKey[:])
	if err != nil {
		return "", fmt.Errorf("aes cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("aes gcm: %w", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("nonce: %w", err)
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func (m *TOTPManager) DecryptSecret(ciphertext string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("base64 decode: %w", err)
	}

	block, err := aes.NewCipher(m.encryptionKey[:])
	if err != nil {
		return "", fmt.Errorf("aes cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("aes gcm: %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, encrypted := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		return "", fmt.Errorf("decrypt: %w", err)
	}

	return string(plaintext), nil
}

func GenerateBackupCodes(count int) ([]string, error) {
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
