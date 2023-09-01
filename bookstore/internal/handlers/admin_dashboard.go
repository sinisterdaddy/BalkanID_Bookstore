// internal/handlers/admin_dashboard.go
package handlers

import (
    "net/http"
    "html/template"
)

func AdminDashboardPageHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/admin_dashboard.html")
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, nil)
}
