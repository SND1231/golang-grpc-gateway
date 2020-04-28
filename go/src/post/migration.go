package main

import (
	"app/post/db"
	"app/post/model"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := db.Connection()
	defer db.Close()
	fmt.Printf("migration")

	db.AutoMigrate(&model.Post{})
}
