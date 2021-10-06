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

func (repo *PostsRepository) FindByUserId(db *gorm.DB, userId int) (posts []models.Post, err error) {
	result := db.First(&posts, userId)
	if result.Error != nil {
		return []models.Post{}, errors.New(result.Error.Error())
	}

	return posts, nil
}

func (repo *PostsRepository) Remove(db *gorm.DB, id int) (err error) {
	result := db.Delete(id)
	if result.Error != nil {
		return errors.New(result.Error.Error())
	}

	return nil
}
