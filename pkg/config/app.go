package config

import (
	"github.com/jinzhu/gorm"
		"github.com/jinzhu/gorm/dialets/mysql"
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