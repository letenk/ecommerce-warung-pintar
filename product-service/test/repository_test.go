package test

import (
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/jabutech/ecommerce-warung-pintar/product-service/models/domain"
	"github.com/jabutech/ecommerce-warung-pintar/product-service/repository"
	"github.com/jabutech/ecommerce-warung-pintar/product-service/util"
	"github.com/stretchr/testify/assert"
)

func createRandomProduct(t *testing.T) domain.Product {
	db := util.SetupTestDB()

	productRepository := repository.NewRepositoryProduct(db)

	id := uuid.New()
	product := domain.Product{
		ID:          id,
		Name:        util.RandomNameProduct(),
		Sku:         "WP" + util.RandomSKU(),
		Description: util.RandomString(20),
		Price:       util.RandomInt(0, 3000),
		Quantity:    util.RandomInt(1, 500),
	}

	newProduct, err := productRepository.Save(product)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, product.ID, newProduct.ID)
	assert.Equal(t, product.Name, newProduct.Name)
	assert.Equal(t, product.Sku, newProduct.Sku)
	assert.Equal(t, product.Description, newProduct.Description)
	assert.Equal(t, product.Price, newProduct.Price)
	assert.Equal(t, product.Quantity, newProduct.Quantity)

	return newProduct
}
func TestSaveProduct(t *testing.T) {
	createRandomProduct(t)
}

func TestFindAllProducts(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandomProduct(t)
	}

	db := util.SetupTestDB()

	productRepository := repository.NewRepositoryProduct(db)

	products, err := productRepository.FindAll()
	if err != nil {
		log.Fatal(err)
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
