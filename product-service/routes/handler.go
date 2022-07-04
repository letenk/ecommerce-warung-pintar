package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jabutech/ecommerce-warung-pintar/product-service/models/web"
	"github.com/jabutech/ecommerce-warung-pintar/product-service/service"
)

type productHandler struct {
	productService service.Service
}

func NewHandlerProduct(productService service.Service) *productHandler {
	return &productHandler{productService}
}

func (h *productHandler) Create(c *gin.Context) {
	var req web.CreateProductRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		errors := web.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := web.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Create product failed",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Create product
	product, err := h.productService.Create(req)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := web.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Create product failed",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Create format response
	response := web.ApiResponseWithData(
		http.StatusOK,
		"success",
		"Product successfully created",
		web.FormatProductResponse(product),
	)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) GetProducts(c *gin.Context) {
	// Get all products
	products, err := h.productService.GetProducts()
	if err != nil {
		// Create new map for handle error
		errorMessage := gin.H{"errors": "server error"}
		// Create response with helper
		response := web.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Error to get products",
			errorMessage,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Format
	formatter := web.FormatProductsResponse(products)
	response := web.ApiResponseWithData(
		http.StatusOK,
		"success",
		"List of products",
		formatter,
	)

	c.JSON(http.StatusOK, response)
}
