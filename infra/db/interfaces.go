package db

import "github.com/gabrielmoura/estudo-api-go/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
