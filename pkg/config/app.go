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
	// mysql username, password and table to be used
	mysqlUsername := ""
	mysqlPassword := ""
	mysqlTable := ""
	
	dbConnect, err := gorm.Open("mysql", mysqlUsername + ":" + mysqlPassword + "/" + mysqlTable + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = dbConnect
}

// returning the database
func GetDB() *gorm.DB {
	return db
}