package models

import (
	"github.com/jinzhu/gorm"
	"github.com/Tha-prodigy/Go-Bookstore/pkg/config"
)

var db *gorm.DB

// creates a struct that will later be used to create a sql table
type Book struct {
	gorm.Model //adds special sql fields to the Book{}  such as id, createdAt, updatedAt ...
	Name        string `gorm:"" json:"name"`
	Autor       string `json:"autor"`
	Publication string `json:"publication"`
}

func init() {
	// open db connection
	config.Connect()
	// get pointer to db connection object
	db = config.GetDB()
	// creates a sql table with schema similar to the fields in  Book{}
	db.AutoMigrate(&Book{})

}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b

}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}
