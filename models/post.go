package models

import "mime/multipart"

type Post struct {
	Model
	UserId     uint   `json:"userId"`
	CategoryId uint   `json:"categoryId"`
	Image      string `json:"image"`
	Title      string `json:"title"`
	SubTitle   string `json:"subTitle"`
	Body       string `json:"body"`
	User       User
	Category   Category
}

type PostForm struct {
	ID         uint
	UserId     uint                  `form:"userId"`
	CategoryId uint                  `form:"categoryId"`
	File       *multipart.FileHeader `form:"image"`
	Title      string                `form:"title"`
	SubTitle   string                `form:"subTitle"`
	Body       string                `form:"body"`
}
