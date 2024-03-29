package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jabutech/ecommerce-warung-pintar/product-service/models/web"
	"github.com/jabutech/ecommerce-warung-pintar/product-service/routes"
	"github.com/jabutech/ecommerce-warung-pintar/product-service/util"
	"github.com/stretchr/testify/assert"
)

func registerRandomUserWithServiceAuth() {
	user := new(struct {
		Fullname string `json:"fullname"`
		Email    string `json:"email"`
		Address  string `json:"address"`
		City     string `json:"city"`
		Province string `json:"province"`
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
		IsAdmin  bool   `json:"is_admin"`
	})

	client := &http.Client{}
	jsonValue, err := json.Marshal(user)
	if err != nil {
		log.Panic(err.Error())
	}

	baseURL := fmt.Sprintf("http://%s:8801/api/v1/auth/register", "localhost")
	request, err := http.NewRequest(http.MethodPost, baseURL, bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Panic(err.Error())
	}

	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		log.Panic(err.Error())
	}

	defer response.Body.Close()

	var jsonData interface{}
	_ = json.NewDecoder(response.Body).Decode(&jsonData)

	fmt.Println(jsonData)
}

func createRandomProduct(t *testing.T) {
	db := util.SetupTestDB()

	// Use router
	router := routes.SetupRouter(db)

	data := web.CreateProductRequest{
		Name:        util.RandomNameProduct(),
		Description: util.RandomString(20),
		Price:       util.RandomInt(0, 3000),
		Quantity:    util.RandomInt(1, 500),
	}

	dataBody := fmt.Sprintf(`{"name": "%s", "description": "%s", "price": %v, "quantity": %v}`, data.Name, data.Description, data.Price, data.Quantity)
	requestBody := strings.NewReader(dataBody)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8802/api/v1/products", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "success", responseBody["status"])
	assert.Equal(t, "Product successfully created", responseBody["message"])
	// Data is not null
	assert.NotZero(t, responseBody["data"])

	var contextData = responseBody["data"].(map[string]interface{})
	assert.NotEmpty(t, contextData["id"])
	assert.NotEmpty(t, contextData["sku"])
	assert.Equal(t, data.Name, contextData["name"])
	assert.Equal(t, data.Description, contextData["description"])
	assert.Equal(t, data.Price, int64(contextData["price"].(float64)))
	assert.Equal(t, data.Quantity, int64(contextData["quantity"].(float64)))
}

func TestCreateRandomProduct(t *testing.T) {
	registerRandomUserWithServiceAuth()

	createRandomProduct(t)

}

func TestGetListProducts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomProduct(t)
	}
	db := util.SetupTestDB()

	// Use router
	router := routes.SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8802/api/v1/products", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "success", responseBody["status"])
	assert.Equal(t, "List of products", responseBody["message"])

	var listProducts = responseBody["data"].([]interface{})

	assert.NotEqual(t, 0, len(listProducts))
	// Data is not null

	for _, listProduct := range listProducts {
		list := listProduct.(map[string]interface{})
		assert.NotEmpty(t, list["id"])
		assert.NotEmpty(t, list["name"])
		assert.NotEmpty(t, list["sku"])
		assert.NotEmpty(t, list["description"])
		assert.NotEmpty(t, list["price"])
		assert.NotEmpty(t, list["quantity"])
	}
}
