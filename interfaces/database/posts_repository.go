package database

import (
	models "app/models"
	"errors"

	"github.com/jinzhu/gorm"
)

type PostsRepository struct{}

func (repo *PostsRepository) Add(db *gorm.DB, p models.Post) (post models.Post, err error) {
	result := db.Create(&p)
	if result.Error != nil {
		return models.Post{}, errors.New(result.Error.Error())
	}
	db.First(&post, p.ID)

	return post, nil
}

func (repo *PostsRepository) FindByID(db *gorm.DB, id int) (post models.Post, err error) {
	result := db.First(&post, id)
	if result.Error != nil {
		return models.Post{}, errors.New(result.Error.Error())
	}

	return post, nil
}
