package main

import (
	"GoHome/task3"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(
		mysql.Open("root:123@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}

func main() {
	db := InitDB()
	task3.RunGorm(db)

}
