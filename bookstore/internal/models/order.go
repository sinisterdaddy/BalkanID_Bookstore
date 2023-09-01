// internal/models/order.go
package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID   uint
	BookID   uint
	Quantity int
	// Add more fields as needed
}
