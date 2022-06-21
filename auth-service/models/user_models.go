package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primary_key;"`
	Fullname  string    `gorm:"not null"`
	Email     string    `gorm:"not null;unique"`
	Address   string    `gorm:"not null"`
	City      string    `gorm:"not null"`
	Province  string    `gorm:"not null"`
	Mobile    string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	IsAdmin   int16     `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
