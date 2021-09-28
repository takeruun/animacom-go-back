package main

import (
	infrastructure "app/infrastructure"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	time.Local = time.FixedZone("JST", 9*60*60)

	db := infrastructure.NewDB()
	awsS3 := infrastructure.NewAwsS3()
	r := infrastructure.NewRouting(db, awsS3)
	r.SetMiddleware()
	r.Run()
}
