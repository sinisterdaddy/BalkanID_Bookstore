// internal/auth/auth.go
package auth

import (
    "net/http"
    "strings"
)

func AuthenticateUser(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        token := strings.Split(authHeader, " ")[1]
        userID, err := GetUserIDFromToken(token)
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        r = r.WithContext(context.WithValue(r.Context(), "userID", userID))
        next.ServeHTTP(w, r)
    }
}
