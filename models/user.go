package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Model
	Name              string `json:"name"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"password"`
	Posts             []Post
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
