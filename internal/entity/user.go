package entity

import (
	"time"

	"github.com/gabrielmoura/estudo-api-go/pkg/argon"
	"github.com/google/uuid"
	"gopkg.in/validator.v2"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" validate:"min=3"`
	Email     string    `json:"email" validate:"regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func (u User) Validate() error {
	if err := validator.Validate(u); err != nil {
		return err
	}
	return nil
}

func NewUser(name, email, password string) (*User, error) {
	encrypted, err := argon.GenerateFromPassword(password)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Password:  encrypted,
		CreatedAt: time.Now(),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	ok, err := argon.ComparePasswordAndHash(password, u.Password)
	if err != nil {
		return false
	}
	return ok
}
