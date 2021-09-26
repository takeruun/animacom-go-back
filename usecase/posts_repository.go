package usecase

import (
	models "app/models"

	"github.com/jinzhu/gorm"
)

type PostsRepository interface {
	Add(db *gorm.DB, p models.Post) (post models.Post, err error)
}
