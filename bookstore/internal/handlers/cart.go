// internal/handlers/cart.go
package handlers

import (
	"bookstore/db"
	"bookstore/internal/models"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddToCartHandler(w http.ResponseWriter, r *http.Request) {
	userID := getLoggedInUserID(r)
	var newItem models.Cart
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.AddToCart(userID, newItem.BookID, newItem.Quantity)
	if err != nil {
		http.Error(w, "Failed to add item to cart", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetCartItemsHandler(w http.ResponseWriter, r *http.Request) {
	userID := getLoggedInUserID(r)

	cartItems, err := db.GetCartItems(userID)
	if err != nil {
		http.Error(w, "Failed to retrieve cart items", http.StatusInternalServerError)
		return
	}

	// Serialize cart items as JSON and send response
}

func UpdateCartItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartItemID, _ := strconv.Atoi(vars["id"])

	var updatedItem models.Cart
	err := json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.UpdateCartItem(uint(cartItemID), updatedItem.Quantity)
	if err != nil {
		http.Error(w, "Failed to update cart item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func RemoveCartItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cartItemID, _ := strconv.Atoi(vars["id"])

	err := db.RemoveCartItem(uint(cartItemID))
	if err != nil {
		http.Error(w, "Failed to remove cart item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
func CartPageHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/cart.html")
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, nil)
}