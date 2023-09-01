package main

import (
	"net/http"

	"github.com/swaggo/http-swagger"
)

func main() {
	// ... (other setup code)

	// Serve Swagger UI
	http.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The location of your generated Swagger JSON file
	))

	// Start your server
	http.ListenAndServe(":8080", nil)
}
