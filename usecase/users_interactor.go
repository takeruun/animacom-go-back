package usecase

import "app/models"

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
