// internal/models/user.go
package models

import (
	"gorm.io/gorm"
	//"time"
)

type User struct {
	gorm.Model
	Username     string
	Password     string
	IsDeactivated bool
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	// Add more fields as needed
}
