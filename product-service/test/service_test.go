package test

import (
	"log"
	"testing"

	"github.com/jabutech/ecommerce-warung-pintar/product-service/models/domain"
	"github.com/jabutech/ecommerce-warung-pintar/product-service/models/web"
	"github.com/jabutech/ecommerce-warung-pintar/product-service/repository"
	"github.com/jabutech/ecommerce-warung-pintar/product-service/service"
	"github.com/jabutech/ecommerce-warung-pintar/product-service/util"
	"github.com/stretchr/testify/assert"
)

func createServiceRandomProduct(t *testing.T) domain.Product {
	db := util.SetupTestDB()

	productRepository := repository.NewRepositoryProduct(db)

	productService := service.NewServiceProduct(productRepository)

	product := web.CreateProductRequest{
		Name:        util.RandomNameProduct(),
		Description: util.RandomString(20),
		Price:       util.RandomInt(0, 3000),
		Quantity:    util.RandomInt(1, 500),
	}

	newProduct, err := productService.Create(product)
	if err != nil {
		log.Panic(err)
	}

	assert.NotEmpty(t, newProduct.ID)
	assert.NotEmpty(t, newProduct.Sku)

	assert.Equal(t, product.Name, newProduct.Name)
	assert.Equal(t, product.Description, newProduct.Description)
	assert.Equal(t, product.Price, newProduct.Price)
	assert.Equal(t, product.Quantity, newProduct.Quantity)

	return newProduct
}

func TestServiceCreateProduct(t *testing.T) {
	createServiceRandomProduct(t)
}

func TestServiceGetListProducts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRepositoryRandomProduct(t)
	}
	db := util.SetupTestDB()

	productRepository := repository.NewRepositoryProduct(db)
	productService := service.NewServiceProduct(productRepository)

	products, err := productService.GetProducts()
	if err != nil {
		log.Panic(err)
	}

	assert.NotEqual(t, 0, len(products))

	for _, product := range products {
		assert.NotEmpty(t, product.ID)
		assert.NotEmpty(t, product.Name)
		assert.NotEmpty(t, product.Sku)
		assert.NotEmpty(t, product.Description)
		assert.NotEmpty(t, product.Price)
		assert.NotEmpty(t, product.Quantity)
		assert.NotEmpty(t, product.Quantity)
		assert.NotEmpty(t, product.CreatedAt)
		assert.NotEmpty(t, product.UpdatedAt)
	}
}
