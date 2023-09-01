// internal/handlers/account.go
package handlers

import (
	"net/http"
	"bookstore/db"
	"encoding/json"
)

func DeactivateAccountHandler(w http.ResponseWriter, r *http.Request) {
	userID := getLoggedInUserID(r)

	err := db.DeactivateUserAccount(userID)
	if err != nil {
		http.Error(w, "Failed to deactivate account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
func ArchiveAccountHandler(w http.ResponseWriter, r *http.Request) {
	userID := getLoggedInUserID(r)

	err := db.ArchiveUserAccount(userID)
	if err != nil {
		http.Error(w, "Failed to archive account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func RestoreAccountHandler(w http.ResponseWriter, r *http.Request) {
	userID := getLoggedInUserID(r)

	err := db.RestoreUserAccount(userID)
	if err != nil {
		http.Error(w, "Failed to restore account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
