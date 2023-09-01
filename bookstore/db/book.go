// db/book.go
package db

import (
	"bookstore/internal/models"
)

func CreateBook(title, author string) error {
	book := models.Book{
		Title:  title,
		Author: author,
	}
	return DB.Create(&book).Error
}

func GetBookByID(bookID uint) (models.Book, error) {
	var book models.Book
	err := DB.First(&book, bookID).Error
	return book, err
}

func UpdateBook(bookID uint, title, author string) error {
	return DB.Model(&models.Book{}).
		Where("id = ?", bookID).
		Updates(models.Book{Title: title, Author: author}).
		Error
}

func DeleteBook(bookID uint) error {
	return DB.Delete(&models.Book{}, bookID).Error
}
