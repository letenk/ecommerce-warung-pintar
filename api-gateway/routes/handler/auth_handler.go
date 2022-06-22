package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jabutech/ecommerce-warung-pintar/api-gateway/models/web"
)

func Register(c *gin.Context) {
	var req web.RegisterRequest
	client := &http.Client{}
	// Get payload
	c.ShouldBindJSON(&req)

	jsonValue, _ := json.Marshal(req)

	request, _ := http.NewRequest(http.MethodPost, "http://localhost:8801/api/v1/auth/register", bytes.NewBuffer(jsonValue))
	request.Header.Add("Content-Type", "application/json")
	response, _ := client.Do(request)

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

	jsonValue, _ := json.Marshal(req)

	request, _ := http.NewRequest(http.MethodPost, "http://localhost:8801/api/v1/auth/login", bytes.NewBuffer(jsonValue))
	request.Header.Add("Content-Type", "application/json")
	response, _ := client.Do(request)

	defer response.Body.Close()

	var jsonData interface{}
	_ = json.NewDecoder(response.Body).Decode(&jsonData)

	c.JSON(http.StatusOK, jsonData)
}
