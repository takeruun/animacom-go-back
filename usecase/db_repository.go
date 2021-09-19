package usecase

import "github.com/jinzhu/gorm"

type DbRepository interface {
	Begin() *gorm.DB
	Connect() *gorm.DB
}
