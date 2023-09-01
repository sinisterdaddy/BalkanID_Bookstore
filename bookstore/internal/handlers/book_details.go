// internal/handlers/book_detail.go
package handlers

import (
    "net/http"
    "html/template"
)

func BookDetailPageHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/book_detail.html")
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, nil)
}

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
    // Implement logic to fetch books from the database
    books := fetchBooksFromDatabase()

    // Convert the books to JSON and send the response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}