package database

import (
	"app/models"
	"errors"

	"github.com/jinzhu/gorm"
)

type UsersRepository struct{}

func (repo *UsersRepository) FindByID(db *gorm.DB, id int) (user models.User, err error) {
	user = models.User{}
	db.First(&user, id)
	if user.ID <= 0 {
		return models.User{}, errors.New("user is not found")
	}
	return user, nil
}

func (repo *UsersRepository) Add(db *gorm.DB, u models.User) (user models.User, err error) {
	result := db.Create(&u)
	if result.Error != nil {
		return models.User{}, errors.New(result.Error.Error())
	}

	return u, nil
}

func (repo *UsersRepository) FindByEmail(db *gorm.DB, email string) (user models.User, err error) {
	result := db.First(&user, "email = ?", email)
	if result.Error != nil {
		return models.User{}, errors.New(result.Error.Error())
	}

	return user, nil
}
