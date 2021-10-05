package controllers

import (
	aws "app/interfaces/aws"
	database "app/interfaces/database"
	models "app/models"
	usecase "app/usecase"
	"fmt"
	"net/http"
	"strconv"
)

type PostsController struct {
	Interactor usecase.PostsInteractor
}

func NewPostsController(db database.DB, awsS3 aws.AwsS3) *PostsController {
	return &PostsController{
		Interactor: usecase.PostsInteractor{
			DB:         &database.DbRepository{DB: db},
			Posts:      &database.PostsRepository{},
			PostImages: &database.PostImagesRepository{},
			AwsS3:      &aws.AwsS3Repository{AwsS3: awsS3},
		},
	}
}

func (controller *PostsController) Show(c Context, accessToken string) {
	id, _ := strconv.Atoi(c.Param("id"))

	post, err := controller.Interactor.Show(id, accessToken)

	if err != nil {
		c.JSON(500, NewH(err.Error(), nil))
		return
	}

	c.JSON(200, NewH("success", post))
}

func (controller *PostsController) Get(c Context, accessToken string) {
	posts, err := controller.Interactor.Get(accessToken)

	if err != nil {
		c.JSON(500, NewH(err.Error(), nil))
		return
	}

	c.JSON(200, NewH("success", posts))
}

func (controller *PostsController) Create(c Context, accessToken string) {
	postForm := models.PostForm{}
	err := c.Bind(&postForm)

	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusBadRequest)
	}
	post, err := controller.Interactor.Create(postForm, accessToken)

	if err != nil {
		c.JSON(500, NewH(err.Error(), nil))
		return
	}

	c.JSON(200, NewH("success", post))
}
