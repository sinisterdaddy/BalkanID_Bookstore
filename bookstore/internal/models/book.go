// internal/models/book.go
package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string
	Author string
	// Add more fields as needed
}
