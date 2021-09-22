package main

import (
	infrastructure "app/infrastructure"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	time.Local = time.FixedZone("JST", 9*60*60)

	db := infrastructure.NewDB()
	r := infrastructure.NewRouting(db)
	r.SetMiddleware()
	r.Run()
}
