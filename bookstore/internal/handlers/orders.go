// internal/handlers/order.go
package handlers

import (
	"encoding/json"
	"net/http"
	"bookstore/db"
	"bookstore/internal/models"
	"github.com/gorilla/mux"
	"strconv"
)

func PlaceOrderHandler(w http.ResponseWriter, r *http.Request) {
	userID := getLoggedInUserID(r)
	var newOrder models.Order
	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.CreateOrder(userID, newOrder.BookID, newOrder.Quantity)
	if err != nil {
		http.Error(w, "Failed to place order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getLoggedInUserID(r *http.Request) {
	panic("unimplemented")
}

func GetOrderHistoryHandler(w http.ResponseWriter, r *http.Request) {
	userID := getLoggedInUserID(r)

	orderHistory, err := db.GetOrderHistory(userID)
	if err != nil {
		http.Error(w, "Failed to retrieve order history", http.StatusInternalServerError)
		return
	}

	// Serialize order history as JSON and send response
}

func GetOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, _ := strconv.Atoi(vars["id"])

	order, err := db.GetOrderByID(uint(orderID))
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	// Serialize order as JSON and send response
}
