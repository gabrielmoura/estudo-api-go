package entity

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	ErrIdIsRequired    = errors.New("id is required")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
)

func (p *Product) Validate() error {
	if p.ID == "" {
		return ErrIdIsRequired
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price == 0 {
		return ErrPriceIsRequired
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
