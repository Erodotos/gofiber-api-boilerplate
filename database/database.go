package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/erodotos/gofiber-api-boilerplate/config"
	"github.com/erodotos/gofiber-api-boilerplate/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// connectDb
func init() {

	var err error
	p := config.Config("DB_PORT")
	port, _ := strconv.ParseUint(p, 10, 32)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), port, config.Config("DB_NAME"))
	fmt.Println(dsn)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	autoMigrate()
}

func autoMigrate() {
	DB.AutoMigrate(&models.Book{}, &models.User{})
}