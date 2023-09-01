// db/models.go
package db

import (
	"bookstore/internal/models"
	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Book{},
		&models.Order{},
	)
	// Add more models to AutoMigrate as you create them
}
