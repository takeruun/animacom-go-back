package usecase

import (
	models "app/models"

	"github.com/jinzhu/gorm"
)

type PostsRepository interface {
	Add(db *gorm.DB, p models.PostForm) (post models.Post, err error)
}
