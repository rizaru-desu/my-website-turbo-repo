package middleware

import (
	"context"
	"net/http"

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
			token := ""
			if cookie, err := r.Cookie(cookieName); err == nil {
				token = cookie.Value
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
