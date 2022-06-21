package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jabutech/ecommerce-warung-pintar/auth-service/routes"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/util"
	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" || env == "DEVELOPMENT" {
		err := godotenv.Load()

		if err != nil {
			log.Fatalf(".env file couldn't loaded %v", env)
		}
	}
	// Open connection
	db := util.ConnectDb()

	// Router
	server := routes.SetupRouter(db)

	envPort := os.Getenv("APP_PORT")
	port := fmt.Sprintf(":%s", envPort)
	if envPort == "" {
		port = ":8801"
	}

	server.Run(port)
}
