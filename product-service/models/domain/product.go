package domain

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"primary_key;"`
	Name        string    `gorm:"name"`
	Sku         string    `gorm:"sku"`
	Description string    `gorm:"description"`
	Price       int64     `gorm:"price"`
	Quantity    int64     `gorm:"quantity"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
