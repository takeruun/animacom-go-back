package usecase

import (
	models "app/models"

	"github.com/jinzhu/gorm"
)

type PostsRepository interface {
	Add(db *gorm.DB, p models.Post) (post models.Post, err error)
	FindByID(db *gorm.DB, id int) (post models.Post, err error)
	FindByUserId(db *gorm.DB, userId int) (posts []models.Post, err error)
	Remove(db *gorm.DB, id int) (err error)
}
