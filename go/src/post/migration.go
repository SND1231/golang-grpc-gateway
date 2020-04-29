package main

import (
	"app/post/db"
	"app/post/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := db.Connection()
	defer db.Close()

	db.AutoMigrate(&model.Post{}, &model.Like{})
}
