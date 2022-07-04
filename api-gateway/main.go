package main

import "github.com/jabutech/ecommerce-warung-pintar/api-gateway/routes"

func main() {
	server := routes.SetupRouter()

	server.Run(":8080")
}
