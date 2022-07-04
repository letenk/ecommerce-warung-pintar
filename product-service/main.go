package main

import (
	"github.com/jabutech/ecommerce-warung-pintar/product-service/routes"
	"github.com/jabutech/ecommerce-warung-pintar/product-service/util"
)

func main() {
	// Open connection
	db := util.SetupDB()

	// Router
	server := routes.SetupRouter(db)
	server.Run(":8802")
}
