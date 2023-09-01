// internal/handlers/auth.go
package handlers

import (
	"net/http"
	"bookstore/db"
	"bookstore/internal/auth"
	"bookstore/internal/models"
	"encoding/json"
    "html/template"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword := auth.HashPassword(newUser.Password)
	err = db.CreateUser(newUser.Username, hashedPassword)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/register.html")
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, nil)
}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/login.html")
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, nil)
}
func AuthenticateAdmin(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        token := strings.Split(authHeader, " ")[1]
        userID, err := GetUserIDFromToken(token)
        if err != nil || !IsAdmin(userID) {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        r = r.WithContext(context.WithValue(r.Context(), "userID", userID))
        next.ServeHTTP(w, r)
    }
}

// func IsAdmin(userID string) bool {
//     // Check if the user with the given userID is an admin
//     // Implement your logic here
// }