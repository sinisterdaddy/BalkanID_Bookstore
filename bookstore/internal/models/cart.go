// internal/models/cart.go
package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID uint
	BookID uint
	Quantity int
	// Add more fields as needed
}
