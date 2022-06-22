package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/repository"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/service"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	// Repository
	userRepository := repository.NewRepositoryUser(db)
	// Service
	userService := service.NewServiceUser(userRepository)
	// handler
	authHandler := NewHandlerAuth(userService)

	router := gin.Default()
	router.Use(cors.Default())

	auth := router.Group("/api/v1/auth")

	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)

	return router
}
