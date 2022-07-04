package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jabutech/ecommerce-warung-pintar/product-service/models/web"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get header with name `Authorization`
		authHeader := c.GetHeader("Authorization")

		// If inside authHeader doesn't have `Bearer`
		if !strings.Contains(authHeader, "Bearer") {
			// Create format response with util ApiResponseWithoutData
			response := web.ApiResponseWithoutData(http.StatusUnauthorized, "error", "Unauthorized")
			// Stop process and return response
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// If there is, create new variable with empty string value
		encodedToken := ""
		// Split authHeader with white space
		arrayToken := strings.Split(authHeader, " ")
		// If length arrayToken is same the 2
		if len(arrayToken) == 2 {
			// Get arrayToken with index 1 / only token jwt
			encodedToken = arrayToken[1]
		}

		// Validation token
		token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)

			if !ok {
				return nil, errors.New("invalid token")
			}

			SecretKey := os.Getenv("SECRET_JWT")

			return []byte(SecretKey), nil
		})

		// If error
		if err != nil {
			errorMessage := gin.H{"errors": err.Error()}
			// Create format response with model web
			response := web.ApiResponseWithData(http.StatusUnauthorized, "error", "Unauthorized", errorMessage)
			// Stop process and return response
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Get payload token
		claim, ok := token.Claims.(jwt.MapClaims)
		// If not `ok` and token invalid
		if !ok || !token.Valid {
			// Create format response with model web
			response := web.ApiResponseWithoutData(http.StatusUnauthorized, "error", "Unauthorized")
			// Stop process and return response
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Get payload `id` and convert to type `float64` and type `int`
		userId := fmt.Sprint(claim["user_id"])

		// Set user to context with name `currentUser`
		c.Set("currentUser", userId)
	}
}
