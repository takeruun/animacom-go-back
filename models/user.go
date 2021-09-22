package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	Email             string    `json:"email"`
	EncryptedPassword string    `json:"password"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
