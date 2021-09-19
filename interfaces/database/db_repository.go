package database

import "github.com/jinzhu/gorm"

type DbRepository struct {
	DB DB
}

func (db *DbRepository) Begin() *gorm.DB {
	return db.DB.Begin()
}

func (db *DbRepository) Connect() *gorm.DB {
	return db.DB.Connect()
}
