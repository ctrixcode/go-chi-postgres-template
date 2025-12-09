package models

import (
	"time"

	"github.com/google/uuid"
)

type Example struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name" validate:"required,min=3"`
	LuckyNumber float64   `json:"lucky_number" db:"lucky_number" validate:"required"`
	IsPremium   bool      `json:"is_premium" db:"is_premium"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type CreateExampleRequest struct {
	Name        string  `json:"name" validate:"required,min=3"`
	LuckyNumber float64 `json:"lucky_number" validate:"required"`
	IsPremium   bool    `json:"is_premium"`
}

type UpdateExampleRequest struct {
	Name        *string  `json:"name"`
	LuckyNumber *float64 `json:"lucky_number"`
	IsPremium   *bool    `json:"is_premium"`
}
