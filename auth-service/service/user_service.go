package service

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/models/domain"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/models/web"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/repository"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input web.RegisterRequest) (domain.User, error)
	Login(req web.LoginRequest) (domain.User, error)
	IsEmailAvailable(email string) (bool, error)
	GenerateToken(user domain.User) (string, error)
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

func (s *service) Login(req web.LoginRequest) (domain.User, error) {
	// Get payload
	email := req.Email
	password := req.Password

	// Find user by email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	// If user not found
	if user.ID == uuid.Nil {
		return user, errors.New("email or password incorrect")
	}

	// If user is available, compare password hash with password from request use bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, errors.New("email or password incorrect")
	}

	return user, nil
}

// EmailIsAvailable for check if email already exists or not
func (s *service) IsEmailAvailable(email string) (bool, error) {

	// Find email on db with repository
	user, err := s.repository.FindByEmail(email)
	// If error
	if err != nil {
		return false, err
	}
	// If user.Id nil
	if user.ID == uuid.Nil {
		return false, nil
	}

	// If is exist
	return true, nil
}

type Claim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func (s *service) GenerateToken(user domain.User) (string, error) {
	// Create 1 day
	expirationTime := time.Now().AddDate(0, 0, 1)

	// Create clain for payload token
	claim := Claim{
		UserID: user.ID.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	SecretJWT := os.Getenv("SECRET_JWT")
	// Signed token with secret key
	signedToken, err := token.SignedString([]byte(SecretJWT))
	if err != nil {
		return signedToken, err
	}

	// If success, return token
	return signedToken, nil
}
