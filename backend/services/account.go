package services

import "errors"

// type

var (
	ErrEmailAlreadyExists = errors.New("email already exists")
)

type CreateNewAccountScheme struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccountResponseScheme struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type AccountService interface {
	OpenNewAccount(CreateNewAccountScheme) (*AccountResponseScheme, error)
	GetAccount(int) (*AccountResponseScheme, error)
}
