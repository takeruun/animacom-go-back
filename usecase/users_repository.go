package usecase

import (
	models "app/models"

	"github.com/jinzhu/gorm"
)

type UsersRepository interface {
	FindByID(db *gorm.DB, id int) (user models.User, err error)
	Add(db *gorm.DB, u models.User) (user models.User, err error)
	FindByEmail(db *gorm.DB, email string) (user models.User, err error)
}
