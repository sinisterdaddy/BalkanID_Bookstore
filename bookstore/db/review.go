// db/review.go
package db

import (
	"bookstore/internal/models"
)

func CreateReview(userID, bookID uint, rating int, comment string) error {
	review := models.Review{
		UserID:  userID,
		BookID:  bookID,
		Rating:  rating,
		Comment: comment,
	}
	return DB.Create(&review).Error
}

func GetReviewsForBook(bookID uint) ([]models.Review, error) {
	var reviews []models.Review
	err := DB.Where("book_id = ?", bookID).Find(&reviews).Error
	return reviews, err
}

func UpdateReview(reviewID uint, rating int, comment string) error {
	return DB.Model(&models.Review{}).
		Where("id = ?", reviewID).
		Updates(models.Review{Rating: rating, Comment: comment}).
		Error
}
