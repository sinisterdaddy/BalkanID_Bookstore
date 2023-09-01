// db/init.go
package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "user=<dbuser> password=cricketkrishna dbname=bookstore sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	DB = conn
	MigrateModels(DB)
	// Auto-migrate models
	// Call functions to create tables and perform migrations
	// Example: CreateUserTable()
}
// db/init.go

