package database

import (
	models "app/models"
	"errors"

	"github.com/jinzhu/gorm"
)

type PostsRepository struct{}

func (repo *PostsRepository) Add(db *gorm.DB, p models.PostForm) (post models.Post, err error) {
	result := db.Create(&p)
	if result.Error != nil {
		return models.Post{}, errors.New(result.Error.Error())
	}
	db.First(&post, p.ID)

	return post, nil
}
