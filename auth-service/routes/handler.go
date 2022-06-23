package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/models/web"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/service"
)

type authHandler struct {
	userService service.Service
}

func NewHandlerAuth(userService service.Service) *authHandler {
	return &authHandler{userService}
}

func (h *authHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "healthy")
	return
}

func (h *authHandler) Register(c *gin.Context) {
	var req web.RegisterRequest

	// Get payload
	err := c.ShouldBindJSON(&req)
	if err != nil {
		errors := web.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := web.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Register failed.",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Check email is availablity
	isEmailAvailable, err := h.userService.IsEmailAvailable(req.Email)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := web.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Register failed.",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// If email is availablity
	if isEmailAvailable {
		errorMessage := gin.H{"errors": "Email already exist."}
		response := web.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Register failed.",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Register
	_, err = h.userService.Register(req)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := web.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Register failed.",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Create format response
	response := web.ApiResponseWithoutData(
		http.StatusOK,
		"success",
		"You have successfully registered",
	)
	c.JSON(http.StatusOK, response)
}

func (h *authHandler) Login(c *gin.Context) {
	var req web.LoginRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		errors := web.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := web.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Login failed",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Login
	userLogin, err := h.userService.Login(req)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := web.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Login failed",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Generate token
	token, err := h.userService.GenerateToken(userLogin)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := web.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Register failed.",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	fieldToken := gin.H{"token": token}
	// Create format response
	response := web.ApiResponseWithData(
		http.StatusOK,
		"success",
		"You have successfully Login",
		fieldToken,
	)
	c.JSON(http.StatusOK, response)
}
