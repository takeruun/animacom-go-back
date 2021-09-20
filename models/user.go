package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	EncryptedPassword string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type UserJSON struct {
	ID                int       `json:"id"`
	Name              string    `form:"name" json:"name"`
	Email             string    `form:"email" json:"email"`
	EncryptedPassword string    `form:"password" json:"encrypted_password"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
