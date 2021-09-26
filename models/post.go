package models

import "mime/multipart"

type Post struct {
	Model
	UserId     uint   `json:"userId" form:"userId"`
	CategoryId uint   `json:"categoryId" form:"categoryId"`
	Title      string `json:"title" form:"title"`
	SubTitle   string `json:"subTitle" form:"subTitle"`
	Body       string `json:"body" form:"body"`
	User       User
	Category   Category
}

type PostImage struct {
	Model
	PostId uint
	File   *multipart.FileHeader
	Post   Post
}
