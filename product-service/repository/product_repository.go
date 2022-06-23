package repository

import (
	"github.com/jabutech/ecommerce-warung-pintar/product-service/models/domain"
	"gorm.io/gorm"
)

type Repository interface {
	Save(product domain.Product) (domain.Product, error)
	FindAll(product domain.Product) ([]domain.Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(product domain.Product) (domain.Product, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return product, err
	}

	return product, err
}

func (r *repository) FindAll() ([]domain.Product, error) {
	var products []domain.Product

	err := r.db.Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, err
}
