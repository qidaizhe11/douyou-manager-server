package models

import (
	//"database/sql"
	//"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql",
		"daizhe:daizhe11@Mysql@tcp(123.56.223.186:3306)/douyou?charset=utf8mb4&parseTime=True&loc=Local")
	if err == nil {
		DB = db
		return db, err
	}
	return nil, err
}

func GetDB() *gorm.DB {
	return DB
}
