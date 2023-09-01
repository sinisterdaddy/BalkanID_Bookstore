// internal/handlers/review.go
package handlers

import (
	"encoding/json"
	"net/http"
	"bookstore/db"
	"bookstore/internal/models"
	"github.com/gorilla/mux"
	"strconv"
)

func LeaveReviewHandler(w http.ResponseWriter, r *http.Request) {
	userID := getLoggedInUserID(r)
	var newReview models.Review
	err := json.NewDecoder(r.Body).Decode(&newReview)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.CreateReview(userID, newReview.BookID, newReview.Rating, newReview.Comment)
	if err != nil {
		http.Error(w, "Failed to leave review", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetBookReviewsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID, _ := strconv.Atoi(vars["id"])

	reviews, err := db.GetReviewsForBook(uint(bookID))
	if err != nil {
		http.Error(w, "Failed to retrieve reviews", http.StatusInternalServerError)
		return
	}

	//Serialize reviews as JSON and send response
}

func UpdateReviewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reviewID, _ := strconv.Atoi(vars["id"])

	var updatedReview models.Review
	err := json.NewDecoder(r.Body).Decode(&updatedReview)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.UpdateReview(uint(reviewID), updatedReview.Rating, updatedReview.Comment)
	if err != nil {
		http.Error(w, "Failed to update review", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
