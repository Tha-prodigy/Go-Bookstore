package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// create a global variable db pointing to the DB struct in the gorm package
var (
	db *gorm.DB
)

// create connect() for opening a db connection to mysql.
func Connect() {
	d, err := gorm.Open("mysql", "prodigy:Sql_Query@tcp(localhost)/simplerest")
	if err != nil {
		panic(err)
	}
	db = d

}

// this returns the pointer to DB struct
func GetDB() *gorm.DB {
	return db

}