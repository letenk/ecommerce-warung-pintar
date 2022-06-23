package service

import (
	"github.com/google/uuid"
	"github.com/jabutech/ecommerce-warung-pintar/product-service/models/domain"
	"github.com/jabutech/ecommerce-warung-pintar/product-service/models/web"
	"github.com/jabutech/ecommerce-warung-pintar/product-service/repository"
	"github.com/jabutech/ecommerce-warung-pintar/product-service/util"
)

type Service interface {
	Create(req web.CreateProductRequest) (domain.Product, error)
	GetProducts() ([]domain.Product, error)
}

type service struct {
	repository repository.Repository
}

func NewServiceProduct(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) Create(req web.CreateProductRequest) (domain.Product, error) {
	id := uuid.New()
	post := domain.Product{
		ID:          id,
		Name:        req.Name,
		Sku:         "WP" + util.RandomSKU(),
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
	}

	// Save to db
	newProduct, err := s.repository.Save(post)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (s *service) GetProducts() ([]domain.Product, error) {
	products, err := s.repository.FindAll()
	if err != nil {
		return products, err
	}

	return products, nil
}
