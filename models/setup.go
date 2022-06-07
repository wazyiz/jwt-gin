package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Unable to load .env file")
	}

	DbDriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	DB, err = gorm.Open(DbDriver, DBURL)

	if err != nil {
		fmt.Println("Cannot connect to database ", DbDriver)
		log.Fatal("connection error ", err)
	} else {
		fmt.Println("We are connected to the database ", DbDriver)
	}

	DB.AutoMigrate(&User{})
}
