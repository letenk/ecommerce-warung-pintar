package routes

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	auth := router.Group("/api/auth")

	auth.POST("/register", func(ctx *gin.Context) {
		fmt.Println("register")
	})

	return router
}
