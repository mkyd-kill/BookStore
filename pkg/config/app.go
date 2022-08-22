package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// database variable
var (
	db * gorm.DB
)

// database connection
func Connect() {
	dbConnect, err := gorm.Open("mysql", "romeo:Coolkid@34/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = dbConnect
}

// returning the database
func GetDB() *gorm.DB {
	return db
}