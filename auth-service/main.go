package main

import (
	"fmt"
	"os"

	"github.com/jabutech/ecommerce-warung-pintar/auth-service/routes"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/util"
)

func main() {
	// Open connection
	db := util.SetupDB()

	// Router
	server := routes.SetupRouter(db)

	envPort := os.Getenv("APP_PORT")
	port := fmt.Sprintf(":%s", envPort)
	if envPort == "" {
		port = ":8801"
	}

	server.Run(port)
}
