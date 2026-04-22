package auth

import (
	"context"
	"fmt"
	"time"

	authusecase "api/internal/usecase/auth"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secret []byte
	issuer string
}

type jwtClaims struct {
	UserID     string `json:"uid"`
	Name       string `json:"name,omitempty"`
	Email      string `json:"email"`
	Role       string `json:"role,omitempty"`
	RememberMe bool   `json:"remember_me"`
	jwt.RegisteredClaims
}

func NewJWTManager(secret string, issuer string) *JWTManager {
	return &JWTManager{secret: []byte(secret), issuer: issuer}
}

func (m *JWTManager) Sign(_ context.Context, claims authusecase.TokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims{
		UserID:     claims.UserID,
		Name:       claims.Name,
		Email:      claims.Email,
		Role:       claims.Role,
		RememberMe: claims.RememberMe,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    m.issuer,
			Subject:   claims.UserID,
			ExpiresAt: jwt.NewNumericDate(claims.ExpiresAt),
			IssuedAt:  jwt.NewNumericDate(claims.IssuedAt),
			NotBefore: jwt.NewNumericDate(claims.IssuedAt),
			ID:        claims.SessionID,
		},
	})

	signed, err := token.SignedString(m.secret)
	if err != nil {
		return "", fmt.Errorf("sign jwt: %w", err)
	}

	return signed, nil
}

func (m *JWTManager) Verify(_ context.Context, tokenString string) (authusecase.TokenClaims, error) {
	parsedClaims := &jwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, parsedClaims, func(token *jwt.Token) (interface{}, error) {
		return m.secret, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}), jwt.WithIssuer(m.issuer), jwt.WithExpirationRequired(), jwt.WithLeeway(5*time.Second))
	if err != nil {
		return authusecase.TokenClaims{}, authusecase.ErrInvalidToken
	}

	if !token.Valid {
		return authusecase.TokenClaims{}, authusecase.ErrInvalidToken
	}

	return authusecase.TokenClaims{
		SessionID:  parsedClaims.ID,
		UserID:     parsedClaims.UserID,
		Name:       parsedClaims.Name,
		Email:      parsedClaims.Email,
		Role:       parsedClaims.Role,
		Issuer:     parsedClaims.Issuer,
		IssuedAt:   numericDateTime(parsedClaims.IssuedAt),
		ExpiresAt:  numericDateTime(parsedClaims.ExpiresAt),
		RememberMe: parsedClaims.RememberMe,
	}, nil
}

func numericDateTime(value *jwt.NumericDate) time.Time {
	if value == nil {
		return time.Time{}
	}

	return value.Time
}
