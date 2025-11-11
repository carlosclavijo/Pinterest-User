package middleware

import (
	"context"
	"github.com/carlosclavijo/Pinterest-Services/internal/infrastructure/services"
	"net/http"
	"strings"
)

func JWTMiddleware(jwtService *services.JWTService, blacklistRepo *services.TokenBlacklist) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "missing authorization header", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "invalid authorization header", http.StatusUnauthorized)
				return
			}

			tokenStr := parts[1]

			if blacklisted, err := blacklistRepo.IsBlacklisted(tokenStr); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			} else if blacklisted {
				http.Error(w, "token revoked", http.StatusUnauthorized)
				return
			}

			userID, err := jwtService.Validate(tokenStr)
			if err != nil {
				http.Error(w, "invalid or expired token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), "user_id", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
