// internal/routes/routes.go
package routes

import (
	"github.com/gorilla/mux"
	"bookstore/internal/handlers"
	"bookstore/internal/auth"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// ... (existing routes)

	// Book Management
	router.HandleFunc("/books", handlers.ListBooksHandler).Methods("GET")
	router.HandleFunc("/books", handlers.CreateBookHandler).Methods("POST")
	router.HandleFunc("/books/{id:[0-9]+}", handlers.GetBookHandler).Methods("GET")
	router.HandleFunc("/books/{id:[0-9]+}", handlers.UpdateBookHandler).Methods("PUT")
	router.HandleFunc("/books/{id:[0-9]+}", handlers.DeleteBookHandler).Methods("DELETE")
		// Cart Management
	router.HandleFunc("/cart", handlers.AddToCartHandler).Methods("POST")
	router.HandleFunc("/cart", handlers.GetCartItemsHandler).Methods("GET")
	router.HandleFunc("/cart/{id:[0-9]+}", handlers.UpdateCartItemHandler).Methods("PUT")
	router.HandleFunc("/cart/{id:[0-9]+}", handlers.RemoveCartItemHandler).Methods("DELETE")
	// Order Management
	router.HandleFunc("/orders", handlers.PlaceOrderHandler).Methods("POST")
	router.HandleFunc("/orders", handlers.GetOrderHistoryHandler).Methods("GET")
	router.HandleFunc("/orders/{id:[0-9]+}", handlers.GetOrderHandler).Methods("GET")
	// Review Management
	router.HandleFunc("/reviews", handlers.LeaveReviewHandler).Methods("POST")
	router.HandleFunc("/reviews/{id:[0-9]+}", handlers.UpdateReviewHandler).Methods("PUT")
	router.HandleFunc("/books/{id:[0-9]+}/reviews", handlers.GetBookReviewsHandler).Methods("GET")
	// Admin-specific Endpoints
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(auth.AuthenticateAdmin)  // Require admin authentication for all endpoints in this router

	adminRouter.HandleFunc("/inventory", handlers.ManageInventoryHandler).Methods("POST")
	 // Frontend Pages
    router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
     // User Authentication Pages
    router.HandleFunc("/register", handlers.RegisterPageHandler).Methods("GET")
    router.HandleFunc("/login", handlers.LoginPageHandler).Methods("GET")
	router.HandleFunc("/profile", auth.AuthenticateUser(handlers.ProfilePageHandler)).Methods("GET")
	// Book Catalog Page
    router.HandleFunc("/catalog", handlers.CatalogPageHandler).Methods("GET")
	// Shopping Cart Page
    router.HandleFunc("/cart", auth.AuthenticateUser(handlers.CartPageHandler)).Methods("GET")
	// Checkout and Order History Pages
    router.HandleFunc("/checkout", auth.AuthenticateUser(handlers.CheckoutPageHandler)).Methods("GET")
    router.HandleFunc("/order_history", auth.AuthenticateUser(handlers.OrderHistoryPageHandler)).Methods("GET")
	// Book Detail and Review Pages
    router.HandleFunc("/books/{book_id}", handlers.BookDetailPageHandler).Methods("GET")
    router.HandleFunc("/books/{book_id}/review", auth.AuthenticateUser(handlers.ReviewPageHandler)).Methods("GET")
	// Admin Dashboard Page
    router.HandleFunc("/admin/dashboard", auth.AuthenticateAdmin(handlers.AdminDashboardPageHandler)).Methods("GET")


	return router
}
