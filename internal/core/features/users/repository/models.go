package users_repository

import "time"

type UserModel struct {
	ID int
	Version int
	Email string
	Password string 
	CreatedAt time.Time
}