package web

import (
	"github.com/google/uuid"
	"github.com/jabutech/ecommerce-warung-pintar/product-service/models/domain"
)

type ProductResponse struct {
	ID          uuid.UUID `json:"id"`
	Sku         string    `json:"sku"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	Quantity    int64     `json:"quantity"`
}

// FormatProductResponse a handle format single product
func FormatProductResponse(product domain.Product) ProductResponse {
	formatter := ProductResponse{
		ID:          product.ID,
		Sku:         product.Sku,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
	}
	return formatter
}

// FormatProductsResponse a handle format multiple products
func FormatProductsResponse(products []domain.Product) []ProductResponse {
	// If products not available, return empty array
	if len(products) == 0 {
		return []ProductResponse{}
	}

	var productFormatter []ProductResponse

	for _, product := range products {
		formatter := FormatProductResponse(product)
		productFormatter = append(productFormatter, formatter)
	}

	return productFormatter
}
