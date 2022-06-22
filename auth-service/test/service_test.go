package test

import (
	"log"
	"strconv"
	"testing"

	"github.com/jabutech/ecommerce-warung-pintar/auth-service/models/domain"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/models/web"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/repository"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/service"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/util"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func createRandomAccountService(t *testing.T, withIsAdmin bool) domain.User {
	// Open connection db
	db := util.SetupTestDB()
	// Use repository
	userRepository := repository.NewRepositoryUser(db)
	// Use Service
	userService := service.NewServiceUser(userRepository)

	req := web.RegisterRequest{}
	password := "password"

	if withIsAdmin == true {
		req.Fullname = util.RandomFullname()
		req.Email = util.RandomString(5) + "@mail.com"
		req.Address = util.RandomString(20)
		req.City = util.RandomCity()
		req.Province = "sumatera utara"
		req.Mobile = strconv.FormatInt(util.RandomMobile(), 10)
		req.Password = password
		req.IsAdmin = true

	} else {
		req.Fullname = util.RandomFullname()
		req.Email = util.RandomString(5) + "@mail.com"
		req.Address = util.RandomString(20)
		req.City = util.RandomCity()
		req.Province = "sumatera utara"
		req.Mobile = strconv.FormatInt(util.RandomMobile(), 10)
		req.Password = password
	}

	// Register
	newUser, err := userService.Register(req)
	if err != nil {
		log.Fatal(err)
	}

	assert.NotEmpty(t, newUser.ID)
	assert.Equal(t, newUser.Fullname, req.Fullname)
	assert.Equal(t, newUser.Email, req.Email)
	assert.Equal(t, newUser.Address, req.Address)
	assert.Equal(t, newUser.City, req.City)
	assert.Equal(t, newUser.Province, req.Province)
	assert.Equal(t, newUser.Mobile, req.Mobile)

	// If parameter  withIsAdmin is true
	if withIsAdmin == true {
		// is_admin must be true
		assert.Equal(t, 1, newUser.IsAdmin)
	} else {
		// is admin must be false
		assert.Equal(t, 0, newUser.IsAdmin)
	}

	err = bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(password))
	if err != nil {
		log.Fatal("password not exist.")
	}
	assert.Nil(t, err)

	return newUser
}

// Test success register with is admin
func TestRegisterWithIsAdmin(t *testing.T) {
	// Var withIsAdmin value true
	withIsAdmin := true
	createRandomAccountService(t, withIsAdmin)
}

// Test success register without is admin
func TestRegisterWithoutIsAdmin(t *testing.T) {
	// Var withIsAdmin value true
	withIsAdmin := false
	createRandomAccountService(t, withIsAdmin)
}

// Test failed register with email unique
func TestRegisterErrorEmailUnique(t *testing.T) {
	// Create random account
	newUser := createRandomAccountService(t, true)

	// Open connection db
	db := util.SetupTestDB()
	// Use repository
	userRepository := repository.NewRepositoryUser(db)
	// Use Service
	userService := service.NewServiceUser(userRepository)

	user := web.RegisterRequest{
		Fullname: util.RandomFullname(),
		Email:    newUser.Email, // Register with same email from create account random
		Address:  util.RandomString(20),
		City:     util.RandomCity(),
		Province: "sumatera utara",
		Mobile:   strconv.FormatInt(util.RandomMobile(), 10),
		Password: "password",
		IsAdmin:  true,
	}

	_, err := userService.Register(user)
	assert.NotNil(t, err)
}

func TestIsEmailAvailable(t *testing.T) {
	// Var withIsAdmin value true
	withIsAdmin := true
	newUser := createRandomAccountService(t, withIsAdmin)

	// Open connection db
	db := util.SetupTestDB()
	// Use repository
	userRepository := repository.NewRepositoryUser(db)
	// Use service
	userService := service.NewServiceUser(userRepository)

	// Find user by email
	user, err := userService.IsEmailAvailable(newUser.Email)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, true, user)
}

func randomLogin(t *testing.T) domain.User {
	// Var withIsAdmin value true
	withIsAdmin := true
	newUser := createRandomAccountService(t, withIsAdmin)

	// Open connection db
	db := util.SetupTestDB()
	// Use repository
	userRepository := repository.NewRepositoryUser(db)
	// Use service
	userService := service.NewServiceUser(userRepository)

	payload := web.LoginRequest{
		Email:    newUser.Email,
		Password: "password",
	}

	// Login
	userLogin, err := userService.Login(payload)
	if err != nil {
		log.Fatal(err)
	}
	assert.NotEmpty(t, userLogin.ID)
	assert.Equal(t, newUser.Fullname, userLogin.Fullname)
	assert.Equal(t, newUser.Email, userLogin.Email)
	assert.Equal(t, newUser.Address, userLogin.Address)
	assert.Equal(t, newUser.City, userLogin.City)
	assert.Equal(t, newUser.Province, userLogin.Province)
	assert.Equal(t, newUser.Mobile, userLogin.Mobile)
	assert.Equal(t, 1, newUser.IsAdmin)

	err = bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(payload.Password))
	if err != nil {
		log.Fatal("password not exist.")
	}
	assert.Nil(t, err)

	return userLogin
}

func TestServiceLoginSuccess(t *testing.T) {
	randomLogin(t)
}

func TestGenerateTokenSuccess(t *testing.T) {
	user := randomLogin(t)

	// Open connection db
	db := util.SetupTestDB()
	// Use repository
	userRepository := repository.NewRepositoryUser(db)
	// Use service
	userService := service.NewServiceUser(userRepository)

	// Generate token
	token, err := userService.GenerateToken(user)
	if err != nil {
		log.Fatal(err)
	}

	assert.NotEmpty(t, token)

}
