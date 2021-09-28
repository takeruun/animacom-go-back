package database

import (
	models "app/models"
	"errors"

	"github.com/jinzhu/gorm"
)

type PostImagesRepository struct{}

func (repo *PostImagesRepository) Add(db *gorm.DB, p models.PostImage) (postImage models.PostImage, err error) {
	result := db.Create(&p)
	if result.Error != nil {
		return models.PostImage{}, errors.New(result.Error.Error())
	}
	db.First(&postImage, p.ID)

	return postImage, nil
}
