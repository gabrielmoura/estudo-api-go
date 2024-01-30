package dto

type ErrorResponse struct {
	Error string `json:"error"`
	Stack string `json:"stack"`
}

type MessageResponse struct {
	Message string `json:"message"`
}
