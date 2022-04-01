package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	db *gorm.DB
)

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbUsername := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	conn := dbUsername + ":" + dbPass + "@/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	d, err2 := gorm.Open(dbDriver, conn)
	if err2 != nil {
		panic(err2)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
