package repository

import (
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/models/domain"
	"gorm.io/gorm"
)

type Repository interface {
	Save(user domain.User) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user domain.User) (domain.User, error) {
	// Create new user on db
	err := r.db.Save(&user).Error
	// If err return object data user, with error
	if err != nil {
		return user, err
	}

	// If success return new data user, with no error
	return user, nil
}

func (r *repository) FindByEmail(email string) (domain.User, error) {
	var user domain.User

	// Find user by email
	err := r.db.Where("email = ?", email).Find(&user).Error
	// If err return object data user, with error
	if err != nil {
		return user, err
	}

	// If success return new data user, with no error
	return user, nil
}
