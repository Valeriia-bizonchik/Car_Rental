package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "mysql", "root:Aria016rnhod4017@tcp(127.0.0.1:3306)/car_share")
	if err != nil {
		panic(err)
	}
	db = d
	fmt.Println("Succesfully connected to db")
}

func GetDB() *gorm.DB {
	return db
}
