// internal/models/review.go
package models

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID uint
	BookID uint
	Rating int
	Comment string
	// Add more fields as needed
}
