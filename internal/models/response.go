package models

type ErrorResponse struct {
	Error string `json:"error"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type MessageResponse struct {
	Message string `json:"message"`
}
