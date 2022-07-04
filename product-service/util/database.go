package util

import (
	"fmt"
	"log"
	"os"

	"github.com/jabutech/ecommerce-warung-pintar/product-service/models/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)
	// Open connection to db
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection to database failed!, err: ", err.Error())
	}

	// Auto Migrate
	err = db.AutoMigrate(&domain.Product{})
	if err != nil {
		log.Fatalf("Failed to auto migration %v", err)
	}

	return db
}

func SetupTestDB() *gorm.DB {
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")
	DBName := os.Getenv("DB_NAME_TEST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)
	// Open connection to db
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection to database failed!, err: ", err.Error())
	}

	// Auto Migrate
	err = db.AutoMigrate(&domain.Product{})
	if err != nil {
		log.Fatalf("Failed to auto migration %v", err)
	}

	return db
}
