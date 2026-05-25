package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "prodigy:Sql_Query@tcp(localhost)/simplerest")
	if err != nil {
		panic(err)
	}
	db = d

}

func GetDB() *gorm.DB {
	return db

}