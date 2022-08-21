package models

import (
	"github.com/jinzhu/gorm"
	"github.com/mkyd-kill/Bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm: ""json: "name" `
	Author string `json: "author" `
	Publication string `json: "publication" `
}

// database initialisation
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}