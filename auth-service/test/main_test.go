package test

import (
	"testing"

	"github.com/jabutech/ecommerce-warung-pintar/auth-service/models/domain"
	"github.com/jabutech/ecommerce-warung-pintar/auth-service/util"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	// Load env
	godotenv.Load("../.env")

	m.Run()

	// Drop table after test
	db := util.SetupTestDB()
	db.Migrator().DropTable(&domain.User{})
}
