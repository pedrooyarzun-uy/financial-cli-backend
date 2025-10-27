package middlewares

import (
	"net/http"
	"time"

	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/helpers"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		claims, err := helpers.ParseJWT(token)

		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		exp, err := claims.GetExpirationTime()

		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		if exp == nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		expTime := exp.Time

		if time.Now().After(expTime) {
			http.Error(w, "expired token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
