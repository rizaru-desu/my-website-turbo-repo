package auth

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/scrypt"
	"golang.org/x/text/unicode/norm"
)

func TestArgon2PasswordVerifierVerifiesPHCHash(t *testing.T) {
	hash := testArgon2IDHash("secret")

	if !(Argon2PasswordVerifier{}).Verify(hash, "secret") {
		t.Fatal("expected argon2id password to verify")
	}
	if (Argon2PasswordVerifier{}).Verify(hash, "wrong") {
		t.Fatal("expected wrong password to be rejected")
	}
}

func TestPasswordVerifierHashesWithArgon2(t *testing.T) {
	verifier := NewPasswordVerifier("auth-secret")
	hash, err := verifier.Hash("new-secret")
	if err != nil {
		t.Fatalf("expected password hash, got %v", err)
	}

	if !verifier.Verify(hash, "new-secret") {
		t.Fatal("expected generated password hash to verify")
	}
	if verifier.Verify(hash, "wrong") {
		t.Fatal("expected wrong password to be rejected")
	}
}

func TestBetterAuthScryptPasswordVerifierVerifiesLegacyHash(t *testing.T) {
	hash := testBetterAuthScryptHash("secret")

	if !(BetterAuthScryptPasswordVerifier{}).Verify(hash, "secret") {
		t.Fatal("expected better-auth scrypt password to verify")
	}
	if (BetterAuthScryptPasswordVerifier{}).Verify(hash, "wrong") {
		t.Fatal("expected wrong password to be rejected")
	}
}

func TestPasswordVerifierFallsBackToSecretScryptHash(t *testing.T) {
	hash := testSecretScryptHash("secret", "auth-secret")

	if !NewPasswordVerifier("auth-secret").Verify(hash, "secret") {
		t.Fatal("expected secret-derived scrypt password to verify")
	}
	if NewPasswordVerifier("different").Verify(hash, "secret") {
		t.Fatal("expected different auth secret to be rejected")
	}
}

func testArgon2IDHash(password string) string {
	salt := []byte("1234567890123456")
	memory := uint32(64 * 1024)
	iterations := uint32(1)
	parallelism := uint8(1)
	key := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, 32)

	return fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		memory,
		iterations,
		parallelism,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(key),
	)
}

func testBetterAuthScryptHash(password string) string {
	salt := []byte("1234567890123456")
	key, err := scrypt.Key([]byte(norm.NFKC.String(password)), salt, 16384, 16, 1, 64)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(salt) + ":" + hex.EncodeToString(key)
}

func testSecretScryptHash(password string, secret string) string {
	salt := []byte("abcdefghijklmnop")
	key, err := scrypt.Key([]byte(norm.NFKC.String(password)+secret), salt, 16384, 16, 1, 64)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(salt) + ":" + hex.EncodeToString(key)
}
