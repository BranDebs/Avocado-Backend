// Package middlewares provides middlewares that are used in avocadoro.
package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/BranDebs/Avocado-Backend/internal/jwt"
)

// Authorisation extracts JWT from the "Authorization" request header .
// Only Bearer Token authorisation is valid.
func Authorisation(settings *jwt.Settings) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			tokenStr := tokenFromHeader(r)
			ctx = context.WithValue(ctx, jwt.ContextKey, tokenStr)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

// tokenFromHeader tries to retreive the token string from the
// "Authorization" reqeust header: "Authorization: BEARER <TOKEN?".
func tokenFromHeader(r *http.Request) string {
	// Get token from authorization header.
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}
