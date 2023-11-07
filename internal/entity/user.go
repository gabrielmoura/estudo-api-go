package entity

import (
	"github.com/gabrielmoura/estudo-api-go/pkg/argon"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
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
