package entity

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/validator.v2"
)

type Product struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" validate:"min=3"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func (p Product) Validate() error {
	if err := validator.Validate(p); err != nil {
		return err
	}
	return nil
}

func NewProduct(name string, price float64) (*Product, error) {
	newP := &Product{
		ID:        uuid.New().String(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := newP.Validate()
	if err != nil {
		return nil, err
	}
	return newP, nil
}
