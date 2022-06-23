package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jabutech/ecommerce-warung-pintar/api-gateway/models/web"
)

func Register(c *gin.Context) {
	var req web.RegisterRequest
	client := &http.Client{}
	// Get payload
	c.ShouldBindJSON(&req)

	jsonValue, err := json.Marshal(req)
	if err != nil {
		response := web.ApiResponseWithoutData(
			http.StatusInternalServerError,
			"error",
			err.Error(),
		)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	baseURL := fmt.Sprintf("http://%s:8801/api/v1/auth/register", os.Getenv("AUTH_SERVICE_HOST"))
	request, err := http.NewRequest(http.MethodPost, baseURL, bytes.NewBuffer(jsonValue))
	if err != nil {
		response := web.ApiResponseWithoutData(
			http.StatusInternalServerError,
			"error",
			err.Error(),
		)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		response := web.ApiResponseWithoutData(
			http.StatusInternalServerError,
			"error",
			err.Error(),
		)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	defer response.Body.Close()

	var jsonData interface{}
	_ = json.NewDecoder(response.Body).Decode(&jsonData)

	c.JSON(http.StatusOK, jsonData)
}

func Login(c *gin.Context) {
	var req web.LoginRequest
	client := &http.Client{}
	// Get payload
	c.ShouldBindJSON(&req)

	jsonValue, err := json.Marshal(req)
	if err != nil {
		response := web.ApiResponseWithoutData(
			http.StatusInternalServerError,
			"error",
			err.Error(),
		)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	baseURL := fmt.Sprintf("http://%s:8801/api/v1/auth/login", os.Getenv("AUTH_SERVICE_HOST"))
	request, err := http.NewRequest(http.MethodPost, baseURL, bytes.NewBuffer(jsonValue))
	if err != nil {
		response := web.ApiResponseWithoutData(
			http.StatusInternalServerError,
			"error",
			err.Error(),
		)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		response := web.ApiResponseWithoutData(
			http.StatusInternalServerError,
			"error",
			err.Error(),
		)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	defer response.Body.Close()

	var jsonData interface{}
	_ = json.NewDecoder(response.Body).Decode(&jsonData)

	c.JSON(http.StatusOK, jsonData)
}
