package util

import (
	"fmt"
	"log"
	"os"

	"github.com/jabutech/ecommerce-warung-pintar/auth-service/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	// Environtment
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open connection to db
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection to database failed!, err: ", err.Error())
	}

	// Auto Migrate
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to auto migration %v", err)
	}

	return db
}
