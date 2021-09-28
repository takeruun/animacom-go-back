package usecase

import (
	models "app/models"

	"github.com/jinzhu/gorm"
)

type PostImagesRepository interface {
	Add(db *gorm.DB, p models.PostImage) (post models.PostImage, err error)
}
