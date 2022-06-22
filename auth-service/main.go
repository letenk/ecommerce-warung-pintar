package main

import (
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/routes"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/util"
)

func main() {
	// Open connection
	db := util.SetupDB()

	// Router
	server := routes.SetupRouter(db)

	server.Run(":8801")
}
