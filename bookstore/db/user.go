// db/user.go
package db

import (
	"bookstore/internal/models"
	"gorm.io/gorm"
)

func CreateUser(username, password string) error {
	user := models.User{
		Username: username,
		Password: password,
	}
	return DB.Create(&user).Error
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := DB.Where("username = ?", username).First(&user).Error
	return user, err
}
func DeactivateUserAccount(userID uint) error {
	return DB.Model(&models.User{}).
		Where("id = ?", userID).
		Update("is_deactivated", true).
		Error
}
func ArchiveUserAccount(userID uint) error {
	return DB.Model(&models.User{}).
		Where("id = ?", userID).
		Update("deleted_at", time.Now()).
		Error
}

func RestoreUserAccount(userID uint) error {
	return DB.Model(&models.User{}).
		Unscoped().
		Where("id = ?", userID).
		Update("deleted_at", nil).
		Error
}