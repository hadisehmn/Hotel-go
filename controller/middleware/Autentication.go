package middleware

import (
	"context"
	"go-practice/HOTEL/models"
	"go-practice/HOTEL/utils"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		const prefix = "Bearer "
		if authHeader == "" || !strings.HasPrefix(authHeader, prefix) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, prefix)

		token, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid claims", http.StatusUnauthorized)
			return
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			http.Error(w, "user id not found", http.StatusUnauthorized)
			return
		}

		userID := int(userIDFloat)
		ctx := context.WithValue(
			r.Context(),
			"userID",
			userID,
		)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	})
}

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		const prefix = "Bearer "
		if authHeader == "" || !strings.HasPrefix(authHeader, prefix) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, prefix)
		token, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}
		role, ok := claims["role"].(string)
		if !ok {
			http.Error(w, "Role not found", http.StatusUnauthorized)
			return
		}

		if role != string(models.UserRoleAdmin) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)

	})

}
