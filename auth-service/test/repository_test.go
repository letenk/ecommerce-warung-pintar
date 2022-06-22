package test

import (
	"log"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/models/domain"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/repository"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/util"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

// Func for create random account
func CreateRandomAccountRepository(t *testing.T, withIsAdmin bool) domain.User {
	// Open connection db
	db := util.SetupTestDB()
	// Use repository
	userRepository := repository.NewRepositoryUser(db)

	user := domain.User{}
	password := "password"

	if withIsAdmin == true {
		// Generate uuid
		id := uuid.New()

		user.ID = id
		user.Fullname = util.RandomFullname()
		user.Email = util.RandomString(5) + "@mail.com"
		user.Address = util.RandomString(20)
		user.City = util.RandomCity()
		user.Province = "sumatera utara"
		user.Mobile = strconv.FormatInt(util.RandomMobile(), 10)
		user.IsAdmin = 1

		// Hash password
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			log.Fatal(err)
		}
		user.Password = string(passwordHash)

	} else {
		// Generate uuid
		id := uuid.New()

		user.ID = id
		user.Fullname = util.RandomFullname()
		user.Email = util.RandomString(5) + "@mail.com"
		user.Address = util.RandomString(20)
		user.City = util.RandomCity()
		user.Province = "sumatera utara"
		user.Mobile = strconv.FormatInt(util.RandomMobile(), 10)

		// Hash password
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			log.Fatal(err)
		}
		user.Password = string(passwordHash)
	}

	// Save to db
	newUser, err := userRepository.Save(user)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, newUser.ID, user.ID)
	assert.Equal(t, newUser.Fullname, user.Fullname)
	assert.Equal(t, newUser.Email, user.Email)
	assert.Equal(t, newUser.Address, user.Address)
	assert.Equal(t, newUser.City, user.City)
	assert.Equal(t, newUser.Province, user.Province)
	assert.Equal(t, newUser.Mobile, user.Mobile)

	// If parameter  withIsAdmin is true
	if withIsAdmin == true {
		// is_admin must be 1
		assert.Equal(t, newUser.IsAdmin, user.IsAdmin)
	} else {
		// is admin must be 0
		assert.Equal(t, newUser.IsAdmin, 0)
	}

	err = bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(password))
	if err != nil {
		log.Fatal("password not exist.")
	}
	assert.Nil(t, err)

	return newUser
}

// Test success save user to db with is admin
func TestUserSaveWithIsAdmin(t *testing.T) {
	// Var withIsAdmin value true
	withIsAdmin := true
	CreateRandomAccountRepository(t, withIsAdmin)
}

// Test success save user to db is admin
func TestUserSaveWithoutIsAdmin(t *testing.T) {
	// Var withIsAdmin value false
	withIsAdmin := false
	CreateRandomAccountRepository(t, withIsAdmin)

}

// Test failed with email unique
func TestUserSaveErrorEmailUnique(t *testing.T) {
	// Create random account
	newUser := CreateRandomAccountRepository(t, true)

	// Open connection db
	db := util.SetupTestDB()
	// Use repository
	userRepository := repository.NewRepositoryUser(db)

	password := "password"
	// Generate uuid
	id := uuid.New()

	user := domain.User{
		ID:       id,
		Fullname: util.RandomFullname(),
		Email:    newUser.Email, // Register with same email from create account random
		Address:  util.RandomString(20),
		City:     util.RandomCity(),
		Province: "sumatera utara",
		Mobile:   strconv.FormatInt(util.RandomMobile(), 10),
		IsAdmin:  1,
	}

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(passwordHash)

	_, err = userRepository.Save(user)
	assert.NotNil(t, err)
}

func TestFindByEmail(t *testing.T) {
	// Var withIsAdmin value true
	withIsAdmin := true
	newUser := CreateRandomAccountRepository(t, withIsAdmin)

	// Open connection db
	db := util.SetupTestDB()
	// Use repository
	userRepository := repository.NewRepositoryUser(db)

	// Find user by email
	user, err := userRepository.FindByEmail(newUser.Email)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, newUser.ID, user.ID)
	assert.Equal(t, newUser.Fullname, user.Fullname)
	assert.Equal(t, newUser.Email, user.Email)
	assert.Equal(t, newUser.Address, user.Address)
	assert.Equal(t, newUser.City, user.City)
	assert.Equal(t, newUser.Province, user.Province)
	assert.Equal(t, newUser.Mobile, user.Mobile)
	assert.Equal(t, newUser.IsAdmin, user.IsAdmin)
}
