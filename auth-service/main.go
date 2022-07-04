package main

import (
	"os"

	"github.com/jabutech/ecommerce-warung-pintar/auth-service/routes"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/util"
	"github.com/joho/godotenv"
)

func main() {
	// I
	if os.Getenv("ENV") == "DEVELOPMENT" {
		godotenv.Load()
	}

	// Open connection
	db := util.SetupDB()

	// Router
	server := routes.SetupRouter(db)

	server.Run(":8801")
}
