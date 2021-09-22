package controllers

import (
	database "app/interfaces/database"
	models "app/models"
	usecase "app/usecase"

	"strconv"
)

type UsersController struct {
	Interactor usecase.UsersInteractor
}

func NewUsersController(db database.DB) *UsersController {
	return &UsersController{
		Interactor: usecase.UsersInteractor{
			DB:    &database.DbRepository{DB: db},
			Users: &database.UsersRepository{},
		},
	}
}

func (controller *UsersController) Get(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := controller.Interactor.Get(id)
	if err != nil {
		c.JSON(500, NewH(err.Error(), nil))
		return
	}

	c.JSON(200, NewH("success", user))
}

func (controller *UsersController) Create(c Context) {
	user := models.User{}
	c.BindJSON(&user)

	user, token, err := controller.Interactor.Create(user)
	if err != nil {
		c.JSON(500, NewH(err.Error(), nil))
		return
	}

	c.Header("access-token", token)
	c.JSON(200, NewH("success", user))
}

func (controller *UsersController) Login(c Context) {
	params := models.UserLogin{}
	c.BindJSON(&params)

	token, err := controller.Interactor.Login(params)
	if err != nil {
		c.JSON(500, NewH(err.Error(), nil))
		return
	}

	c.Header("access-token", token)
	c.JSON(200, NewH("success", nil))
}
