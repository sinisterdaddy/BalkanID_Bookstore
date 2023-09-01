// internal/handlers/order_history.go
package handlers

import (
    "net/http"
    "html/template"
)

func OrderHistoryPageHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/order_history.html")
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, nil)
}
