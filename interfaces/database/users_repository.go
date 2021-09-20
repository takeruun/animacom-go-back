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

func (repo *UsersRepository) Add(db *gorm.DB, u models.UserJSON) (user models.User, err error) {
	result := db.Create(u)
	if result.Error != nil {
		return models.User{}, errors.New(result.Error.Error())
	}

	user = models.User{Name: u.Name, Email: u.Email, EncryptedPassword: u.EncryptedPassword}

	return user, nil
}
