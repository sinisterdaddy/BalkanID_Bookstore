// db/order.go
package db

import (
	"bookstore/internal/models"
)

func CreateOrder(userID, bookID, quantity uint) error {
	order := models.Order{
		UserID:   userID,
		BookID:   bookID,
		Quantity: quantity,
	}
	return DB.Create(&order).Error
}

func GetOrderHistory(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := DB.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func GetOrderByID(orderID uint) (models.Order, error) {
	var order models.Order
	err := DB.First(&order, orderID).Error
	return order, err
}
