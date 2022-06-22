package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/jabutech/ecommerce-warung-pintar/auth-service/models/web"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/routes"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/util"
	"github.com/stretchr/testify/assert"
)

func createRandomAccount(t *testing.T, withIsAdmin bool) web.RegisterRequest {
	db := util.SetupTestDB()

	// Use router
	router := routes.SetupRouter(db)

	var data web.RegisterRequest
	var dataBody string

	if !withIsAdmin {
		data = web.RegisterRequest{
			Fullname: util.RandomFullname(),
			Email:    util.RandomString(5) + "@mail.com",
			Address:  util.RandomString(20),
			City:     util.RandomCity(),
			Province: "sumatera utara",
			Mobile:   strconv.FormatInt(util.RandomMobile(), 10),
			Password: "password",
		}

		dataBody = fmt.Sprintf(`{"fullname": "%s", "email": "%s", "address": "%s", "city": "%s", "province": "%s", "mobile": "%s", "password": "%s"}`, data.Fullname, data.Email, data.Address, data.City, data.Province, data.Mobile, data.Password)
	} else {
		data = web.RegisterRequest{
			Fullname: util.RandomFullname(),
			Email:    util.RandomString(5) + "@mail.com",
			Address:  util.RandomString(20),
			City:     util.RandomCity(),
			Province: "sumatera utara",
			Mobile:   strconv.FormatInt(util.RandomMobile(), 10),
			Password: "password",
			IsAdmin:  true,
		}

		dataBody = fmt.Sprintf(`{"fullname": "%s", "email": "%s", "address": "%s", "city": "%s", "province": "%s", "mobile": "%s", "password": "%s", "is_admin": %t}`, data.Fullname, data.Email, data.Address, data.City, data.Province, data.Mobile, data.Password, data.IsAdmin)
	}

	// Create payload request
	requestBody := strings.NewReader(dataBody)
	// Create request
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8801/api/v1/auth/register", requestBody)
	// Added header content type
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// Get response
	response := recorder.Result()

	// Read response
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	// Response status code must be 200 (success)
	assert.Equal(t, 200, response.StatusCode)
	// Response body status code must be 200 (success)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	// Response body status must be success
	assert.Equal(t, "success", responseBody["status"])
	// Response body message
	assert.Equal(t, "You have successfully registered", responseBody["message"])

	return data
}

// Test register success with is admin true
func TestRegisterSuccessWithIsAdmin(t *testing.T) {
	withIsAdmin := true
	createRandomAccount(t, withIsAdmin)
}

// Test register success with is admin false
func TestRegisterSuccessWithoutIsAdmin(t *testing.T) {
	withIsAdmin := false
	createRandomAccount(t, withIsAdmin)

}

// Test register email is exist
func TestRegisterEmailIsExist(t *testing.T) {
	// Register user random
	newUser := createRandomAccount(t, true)

	db := util.SetupTestDB()

	// Use router
	router := routes.SetupRouter(db)

	var dataBody string
	password := "password"

	data := web.RegisterRequest{
		Fullname: util.RandomFullname(),
		Email:    newUser.Email,
		Address:  util.RandomString(20),
		City:     util.RandomCity(),
		Province: "sumatera utara",
		Mobile:   strconv.FormatInt(util.RandomMobile(), 10),
		Password: password,
	}

	dataBody = fmt.Sprintf(`{"fullname": "%s", "email": "%s", "address": "%s", "city": "%s", "province": "%s", "mobile": "%s", "password": "%s"}`, data.Fullname, data.Email, data.Address, data.City, data.Province, data.Mobile, password)

	// Create payload request
	requestBody := strings.NewReader(dataBody)
	// Create request
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8801/api/v1/auth/register", requestBody)
	// Added header content type
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// Get response
	response := recorder.Result()

	// Read response
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "error", responseBody["status"])
	assert.Equal(t, "Register failed.", responseBody["message"])
	assert.NotZero(t, responseBody["data"])
	assert.Equal(t, "Email already exist.", responseBody["data"].(map[string]interface{})["errors"])
}

// Test validation error
func TestRegisterValidationError(t *testing.T) {
	db := util.SetupTestDB()

	// Use router
	router := routes.SetupRouter(db)

	dataBody := fmt.Sprintf(`{"fullname": "", "email": "", "address": "", "city": "", "province": "", "mobile": "", "password": ""}`)

	// Create payload request
	requestBody := strings.NewReader(dataBody)
	// Create request
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8801/api/v1/auth/register", requestBody)
	// Added header content type
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// Get response
	response := recorder.Result()

	// Read response
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, response.StatusCode)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "error", responseBody["status"])
	assert.Equal(t, "Register failed.", responseBody["message"])
	assert.NotZero(t, responseBody["data"])
	assert.NotNil(t, responseBody["data"].(map[string]interface{})["errors"])
}
