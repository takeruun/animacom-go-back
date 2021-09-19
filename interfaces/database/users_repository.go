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
