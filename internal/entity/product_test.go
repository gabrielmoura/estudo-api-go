package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Product 1", 5.55)
	assert.Nil(t, err)
	assert.NotNil(t, p)

	assert.Equal(t, "Product 1", p.Name)
	assert.Equal(t, 5.55, p.Price)
}
func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 5.55)

	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}
