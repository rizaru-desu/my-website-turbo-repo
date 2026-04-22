package middleware

import (
	"context"
	"net/http"
	"strings"

	authusecase "api/internal/usecase/auth"
)

type contextKey string

const authClaimsContextKey contextKey = "auth_claims"

type Authenticator interface {
	VerifyToken(ctx context.Context, token string) (authusecase.TokenClaims, error)
}

func NewJWTAuthentication(authenticator Authenticator, cookieName string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := bearerToken(r)
			if token == "" {
				if cookie, err := r.Cookie(cookieName); err == nil {
					token = cookie.Value
				}
			}

			claims, err := authenticator.VerifyToken(r.Context(), token)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), authClaimsContextKey, claims)))
		})
	}
}

func AuthClaims(ctx context.Context) (authusecase.TokenClaims, bool) {
	claims, ok := ctx.Value(authClaimsContextKey).(authusecase.TokenClaims)
	return claims, ok
}

func bearerToken(r *http.Request) string {
	header := strings.TrimSpace(r.Header.Get("Authorization"))
	if header == "" {
		return ""
	}

	scheme, token, ok := strings.Cut(header, " ")
	if !ok || !strings.EqualFold(scheme, "Bearer") {
		return ""
	}

	return strings.TrimSpace(token)
}
