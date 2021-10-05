package storages

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

func Connect() *gorm.DB {
	dsn := "thai1201:thai1201@tcp(127.0.0.1:6033)/go_training?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("cannot connect database")
	}
	MysqlDB = db
	return MysqlDB
}

func Get() *gorm.DB {
	return MysqlDB
}