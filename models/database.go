package models

import (
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_training?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Print(err)
	}
	return db
}
