package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type UserJSON struct {
	ID        int       `json:"id"`
	Name      string    `form:"name" json:"name"`
	Email     string    `form:"email" json:"email"`
	Password  string    `form:"password" json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
