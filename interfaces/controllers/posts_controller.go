package controllers

import (
	"app/interfaces/database"
	"app/models"
	"app/usecase"
	"fmt"
	"net/http"
)

type PostsController struct {
	Interactor usecase.PostsInteractor
}

func NewPostsController(db database.DB) *PostsController {
	return &PostsController{
		Interactor: usecase.PostsInteractor{
			DB:    &database.DbRepository{DB: db},
			Posts: &database.PostsRepository{},
		},
	}
}

func (controller *PostsController) Get(c Context, accessToken string) {

	c.JSON(200, NewH("success", nil))
}

func (controller *PostsController) Create(c Context, accessToken string) {
	post := models.Post{}
	err := c.Bind(&post)

	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
	}

	post, err = controller.Interactor.Create(post, accessToken)

	if err != nil {
		c.JSON(500, NewH(err.Error(), nil))
		return
	}

	c.JSON(200, NewH("success", post))
}
