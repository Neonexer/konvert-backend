package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/neonexer/konvert-backend/internal/core/errors"
)

type User struct {
	Id      int `json:"id"`
	Version int `json:"version"`

	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id      int    `json:"id" example:"1"`
	Version int    `json:"version" example:"1"`
	Email   string `json:"email" example:"john.doe@example.com"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,min=3,max=100,email" example:"john.doe@example.com"`
	Password string `json:"password" validate:"required" example:"securepassword123"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(
	id int,
	version int,
	email string,
	password string,
) User {
	return User{
		Id:       id,
		Version:  version,
		Email:    email,
		Password: password,
	}
}

func NewUserUninitialized(
	email string,
	password string,
) User {
	return NewUser(UninitializedID, UninitializedVersion, email, password)
}

func (u *User) Validate() error {
	emailLength := len([]rune(u.Email))
	if emailLength < 3 || emailLength > 100 {
		return fmt.Errorf(
			"invalid `Email` len: %d: %w",
			emailLength,
			core_errors.ErrInvalidArgument,
		)
	}

	re := regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)

	if !re.MatchString(u.Email) {
		return fmt.Errorf(
			"invalid `Email` format: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if u.Password != "" {
		passwordLength := len([]rune(u.Password))
		if passwordLength < 8 || passwordLength > 255 {
			return fmt.Errorf(
				"invalid `Password` len: %d: %w",
				passwordLength,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}
