package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jabutech/ecommerce-warung-pintar/product-service/middleware"
	"github.com/jabutech/ecommerce-warung-pintar/product-service/repository"
	"github.com/jabutech/ecommerce-warung-pintar/product-service/service"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	// Repository
	productRepository := repository.NewRepositoryProduct(db)
	// Service
	productService := service.NewServiceProduct(productRepository)
	// handler
	productHandler := NewHandlerProduct(productService)

	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")

	api.POST("/products", middleware.AuthMiddleware(), productHandler.Create)
	api.GET("/products", productHandler.GetProducts)
	return router
}
