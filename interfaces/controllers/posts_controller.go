package controllers

import "app/interfaces/database"

type PostsController struct{}

func NewPostsController(db database.DB) *PostsController {
	return &PostsController{}
}

func (controller *PostsController) Get(c Context) {

	c.JSON(200, NewH("success", nil))
}
