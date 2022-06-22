package util

import (
	"fmt"
	"log"

	"github.com/jabutech/ecommerce-warung-pintar/auth-service/models/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	// Load config
	config, err := LoadConfig(".", "dev")
	if err != nil {
		log.Fatal("error load config: ", err.Error())
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	// Open connection to db
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection to database failed!, err: ", err.Error())
	}

	// Auto Migrate
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatalf("Failed to auto migration %v", err)
	}

	return db
}

func SetupTestDB() *gorm.DB {
	// Load config
	config, err := LoadConfig("../", "dev")
	if err != nil {
		log.Fatal("error load config: ", err.Error())
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBNameTest)
	// Open connection to db
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection to database failed!, err: ", err.Error())
	}

	// Auto Migrate
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatalf("Failed to auto migration %v", err)
	}

	return db
}
