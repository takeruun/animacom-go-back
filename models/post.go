package models

import (
	"mime/multipart"
)

type Post struct {
	Model
	UserId     uint   `json:"userId" form:"userId"`
	CategoryId uint   `json:"categoryId" form:"categoryId"`
	Title      string `json:"title" form:"title"`
	SubTitle   string `json:"subTitle" form:"subTitle"`
	Body       string `json:"body" form:"body"`
	PostImage  []PostImage
	User       User
	Category   Category
}

type PostImage struct {
	Model
	PostId uint
	Image  string `json:"image"`
	Post   Post
}

type PostForm struct {
	UserId     uint
	CategoryId uint                  `form:"categoryId"`
	Title      string                `form:"title"`
	SubTitle   string                `form:"subTitle"`
	Body       string                `form:"body"`
	Image      *multipart.FileHeader `form:"image"`
}
