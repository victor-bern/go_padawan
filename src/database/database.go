package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDatabase() {
	str := "user:1234@tcp(127.0.0.1)/db?charset=utf8&parseTime=True&loc=Local"

	log.Print(str)
	database, err := gorm.Open(mysql.Open(str))
	if err != nil {
		log.Fatal(err.Error())
	}

	db = database
	config, _ := db.DB()
	config.SetConnMaxIdleTime(10)
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)

}

func GetDatabase() *gorm.DB {
	return db
}
