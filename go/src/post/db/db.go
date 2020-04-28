package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:a1b2V4d@tcp(localhost:23306)/ramen_app?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database: ")
	}
	db.LogMode(true)
	return db
}
