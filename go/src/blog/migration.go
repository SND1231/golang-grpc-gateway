package main

import (
	"app/user/db"
	"app/user/model"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := db.Connection()
	defer db.Close()
	fmt.Printf("migration")

	db.AutoMigrate(&model.User{})
}
