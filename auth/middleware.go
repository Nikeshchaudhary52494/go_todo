package auth

import (
	"context"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Token not found", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Invalid cookie", http.StatusBadRequest)
			return
		}

		tokenString := cookie.Value
		claims, err := ValidateToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		r = r.WithContext(context.WithValue(ctx, "username", claims.Username))

		next.ServeHTTP(w, r)
	})
}
