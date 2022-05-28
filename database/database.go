package database

import (
	"errors"
	"log"
	"os"

	"github.com/erodotos/gofiber-api-boilerplate/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// connectDb
func init() {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:1234@tcp(127.0.0.1:3306)/learn-go?charset=utf8mb4&parseTime=True&loc=Local"
	/*
		NOTE:
		To handle time.Time correctly, you need to include parseTime as a parameter. (more parameters)
		To fully support UTF-8 encoding, you need to change charset=utf8 to charset=utf8mb4. See this article for a detailed explanation
	*/

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db = database

	autoMigrate()
}

func autoMigrate() {

	db.AutoMigrate(&models.Book{})
}

func InsertOne(data interface{}) error {
	if result := db.Create(data); result.Error != nil {
		return errors.New("Database Insertion Error")
	}
	return nil
}

func FindOne() {

}

func UpdateOne() {

}

func Delete() {

}
