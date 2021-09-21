package usecase

import (
	"app/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UsersInteractor struct {
	DB    DbRepository
	Users UsersRepository
}

func (interactor *UsersInteractor) Get(id int) (user models.User, err error) {
	db := interactor.DB.Connect()

	user, err = interactor.Users.FindByID(db, id)
	if err != nil {
		return models.User{}, error(err)
	}

	return
}

func (interactor *UsersInteractor) Create(u models.User) (user models.User, err error) {
	db := interactor.DB.Connect()

	hash, err := bcrypt.GenerateFromPassword([]byte(u.EncryptedPassword), bcrypt.DefaultCost)

	user = models.User{Name: u.Name, Email: u.Email, EncryptedPassword: string(hash), CreatedAt: time.Now(), UpdatedAt: time.Now()}
	user, err = interactor.Users.Add(db, user)
	if err != nil {
		return models.User{}, error(err)
	}

	return
}
