// internal/handlers/book.go
package handlers

import (
	"encoding/json"
	"net/http"
	"bookstore/db"
	"bookstore/internal/models"
	"github.com/gorilla/mux"
	"strconv"
)

func ListBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := db.ListBooks()
	if err != nil {
		http.Error(w, "Failed to retrieve books", http.StatusInternalServerError)
		return
	}

	// Serialize books as JSON and send response
}

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.CreateBook(newBook.Title, newBook.Author)
	if err != nil {
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID, _ := strconv.Atoi(vars["id"])

	book, err := db.GetBookByID(uint(bookID))
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Serialize book as JSON and send response
}

func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID, _ := strconv.Atoi(vars["id"])

	var updatedBook models.Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.UpdateBook(uint(bookID), updatedBook.Title, updatedBook.Author)
	if err != nil {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID, _ := strconv.Atoi(vars["id"])

	err := db.DeleteBook(uint(bookID))
	if err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
