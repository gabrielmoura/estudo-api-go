package entity

import (
	"github.com/google/uuid"
)

type Cart struct {
	ID      string    `json:"id"`
	Product []Product `json:"product"`
}

func NewCart(products *[]Product) (*Cart, error) {
	return &Cart{
		ID:      uuid.New().String(),
		Product: *products,
	}, nil
}
