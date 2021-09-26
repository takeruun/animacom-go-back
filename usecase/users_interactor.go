package usecase

import (
	models "app/models"
	auth "app/usecase/auth"
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

func (interactor *UsersInteractor) Create(u models.User) (user models.User, token string, err error) {
	db := interactor.DB.Connect()

	hash, err := bcrypt.GenerateFromPassword([]byte(u.EncryptedPassword), bcrypt.DefaultCost)

	user = models.User{Name: u.Name, Email: u.Email, EncryptedPassword: string(hash)}
	user, err = interactor.Users.Add(db, user)
	if err != nil {
		return models.User{}, "", error(err)
	}

	token, err = auth.GenerateToken(int(rune(user.ID)), time.Now())

	return user, token, nil
}

func (interactor *UsersInteractor) Login(params models.UserLogin) (string, error) {
	db := interactor.DB.Connect()

	user, err := interactor.Users.FindByEmail(db, params.Email)
	if err != nil {
		return "", error(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(params.Password))
	if err != nil {
		return "", error(err)
	}

	token, err := auth.GenerateToken(int(rune(user.ID)), time.Now())

	return token, nil
}
