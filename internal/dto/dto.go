package dto

import "time"

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	Token      string      `json:"token"`
	PlainToken interface{} `json:"claims"`
}

type ErrorResponse struct {
	Error string `json:"error"`
	Stack string `json:"stack"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductRequest struct {
	Name  string  `json:"name" binding:"required,min=3"`
	Price float64 `json:"price" binding:"required"`
}
