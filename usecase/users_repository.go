package usecase

import (
	models "app/models"

	"github.com/jinzhu/gorm"
)

type UsersRepository interface {
	FindByID(db *gorm.DB, id int) (user models.User, err error)
}
