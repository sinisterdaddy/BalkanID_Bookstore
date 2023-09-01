// db/cart.go
package db

import (
	"bookstore/internal/models"
)

func AddToCart(userID, bookID, quantity uint) error {
	cartItem := models.Cart{
		UserID: userID,
		BookID: bookID,
		Quantity: quantity,
	}
	return DB.Create(&cartItem).Error
}

func GetCartItems(userID uint) ([]models.Cart, error) {
	var cartItems []models.Cart
	err := DB.Where("user_id = ?", userID).Find(&cartItems).Error
	return cartItems, err
}

func UpdateCartItem(cartItemID, quantity uint) error {
	return DB.Model(&models.Cart{}).
		Where("id = ?", cartItemID).
		Update("quantity", quantity).
		Error
}

func RemoveCartItem(cartItemID uint) error {
	return DB.Delete(&models.Cart{}, cartItemID).Error
}
