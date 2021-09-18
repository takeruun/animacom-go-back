package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	Name      string    `form:"name" json:"name"`
	Email     string    `form:"email" json:"email"`
	Password  string    `form:"password" json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	dbConnect()

	fmt.Println("Hello, World!")
}

func dbConnect() (database *gorm.DB) {
	DBMS := "mysql"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	PROTOCOL := "tcp(" + os.Getenv("DB_HOST") + ":3306)"
	DBNAME := os.Getenv("DB_NAME")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true"

	count := 0
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 180 {
				fmt.Println("")
				fmt.Println("DB接続失敗")
				panic(err)
			}
			db, err = gorm.Open(DBMS, CONNECT)
		}
	}
	fmt.Println("DB接続成功")

	return db
}
