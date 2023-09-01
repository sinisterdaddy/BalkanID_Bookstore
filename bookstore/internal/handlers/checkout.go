// internal/handlers/checkout.go
package handlers

import (
    "net/http"
    "html/template"
)

func CheckoutPageHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/checkout.html")
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, nil)
}
