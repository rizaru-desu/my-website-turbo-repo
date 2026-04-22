package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/scrypt"
	"golang.org/x/text/unicode/norm"
)

type PasswordVerifier struct {
	secret string
}

type Argon2PasswordVerifier struct{}

type argon2Params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	keyLength   uint32
}

func NewPasswordVerifier(secret string) PasswordVerifier {
	return PasswordVerifier{secret: secret}
}

func (v PasswordVerifier) Verify(hash string, password string) bool {
	if (Argon2PasswordVerifier{}).Verify(hash, password) {
		return true
	}

	if (BetterAuthScryptPasswordVerifier{}).Verify(hash, password) {
		return true
	}

	return SecretScryptPasswordVerifier{secret: v.secret}.Verify(hash, password)
}

func (Argon2PasswordVerifier) Verify(hash string, password string) bool {
	params, salt, expected, err := decodeArgon2IDHash(hash)
	if err != nil {
		return false
	}

	actual := argon2.IDKey([]byte(password), salt, params.iterations, params.memory, params.parallelism, params.keyLength)
	return subtle.ConstantTimeCompare(actual, expected) == 1
}

func decodeArgon2IDHash(encoded string) (argon2Params, []byte, []byte, error) {
	parts := strings.Split(encoded, "$")
	if len(parts) != 6 || parts[1] != "argon2id" {
		return argon2Params{}, nil, nil, fmt.Errorf("invalid argon2id hash")
	}

	versionText := strings.TrimPrefix(parts[2], "v=")
	version, err := strconv.Atoi(versionText)
	if err != nil || version != argon2.Version {
		return argon2Params{}, nil, nil, fmt.Errorf("unsupported argon2 version")
	}

	params, err := decodeArgon2Params(parts[3])
	if err != nil {
		return argon2Params{}, nil, nil, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return argon2Params{}, nil, nil, fmt.Errorf("decode argon2 salt: %w", err)
	}

	key, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return argon2Params{}, nil, nil, fmt.Errorf("decode argon2 key: %w", err)
	}

	params.keyLength = uint32(len(key))
	return params, salt, key, nil
}

func decodeArgon2Params(encoded string) (argon2Params, error) {
	values := map[string]string{}
	for _, entry := range strings.Split(encoded, ",") {
		key, value, ok := strings.Cut(entry, "=")
		if !ok {
			return argon2Params{}, fmt.Errorf("invalid argon2 parameter")
		}
		values[key] = value
	}

	memory, err := strconv.ParseUint(values["m"], 10, 32)
	if err != nil {
		return argon2Params{}, fmt.Errorf("invalid argon2 memory: %w", err)
	}

	iterations, err := strconv.ParseUint(values["t"], 10, 32)
	if err != nil {
		return argon2Params{}, fmt.Errorf("invalid argon2 iterations: %w", err)
	}

	parallelism, err := strconv.ParseUint(values["p"], 10, 8)
	if err != nil {
		return argon2Params{}, fmt.Errorf("invalid argon2 parallelism: %w", err)
	}

	return argon2Params{
		memory:      uint32(memory),
		iterations:  uint32(iterations),
		parallelism: uint8(parallelism),
	}, nil
}

type BetterAuthScryptPasswordVerifier struct{}

func (BetterAuthScryptPasswordVerifier) Verify(hash string, password string) bool {
	return verifyScryptHash(hash, password, "")
}

type SecretScryptPasswordVerifier struct {
	secret string
}

func (v SecretScryptPasswordVerifier) Verify(hash string, password string) bool {
	if v.secret == "" {
		return false
	}

	return verifyScryptHash(hash, password, v.secret)
}

func verifyScryptHash(hash string, password string, secret string) bool {
	saltHex, keyHex, ok := strings.Cut(hash, ":")
	if !ok || saltHex == "" || keyHex == "" {
		return false
	}

	salt, err := hex.DecodeString(saltHex)
	if err != nil {
		return false
	}

	expected, err := hex.DecodeString(keyHex)
	if err != nil {
		return false
	}

	passwordInput := norm.NFKC.String(password)
	if secret != "" {
		passwordInput += secret
	}

	actual, err := scrypt.Key([]byte(passwordInput), salt, 16384, 16, 1, len(expected))
	if err != nil {
		return false
	}

	return subtle.ConstantTimeCompare(actual, expected) == 1
}

type CryptoIDGenerator struct{}

func (CryptoIDGenerator) NewID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return hex.EncodeToString([]byte(time.Now().UTC().Format(time.RFC3339Nano)))
	}

	return hex.EncodeToString(bytes)
}

type SystemClock struct{}

func (SystemClock) Now() time.Time {
	return time.Now()
}
