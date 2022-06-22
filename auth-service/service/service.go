package service

import (
	"github.com/google/uuid"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/models/domain"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/models/web"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/repository"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input web.RegisterRequest) (domain.User, error)
}

type service struct {
	repository repository.Repository
}

func NewServiceUser(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) Register(req web.RegisterRequest) (domain.User, error) {
	// Passing request into object user
	user := domain.User{}
	user.Fullname = req.Fullname
	user.Email = req.Email
	user.Address = req.Address
	user.City = req.City
	user.Province = req.Province
	user.Mobile = req.Mobile

	//  Generate uuid
	id := uuid.New()
	user.ID = id

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	// Passing password with passwordHash
	user.Password = string(passwordHash)

	// Create new variable isAdmin with default value zero (0) / false
	isAdmin := 0
	// If value `input.IsAdmin` is available / true
	if req.IsAdmin {
		// Change value isAdmin to (1) true
		isAdmin = 1
	}
	// Passing IsAdmin
	user.IsAdmin = isAdmin

	// Save new user
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
