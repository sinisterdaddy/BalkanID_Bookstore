// internal/auth/admin.go
package auth

import (
	"errors"
	"net/http"
)

func IsAdmin(userID uint) bool {
	// Implement logic to check if the user is an admin
	return false
}

func AuthenticateAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := GetUserIDFromToken(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if !IsAdmin(userID) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	}
}
