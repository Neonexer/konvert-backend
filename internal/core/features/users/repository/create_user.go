package users_repository

import (
	"context"
	"fmt"

	"github.com/neonexer/konvert-backend/internal/core/domain"
)

func (r *UsersRepository) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO konvert.users (email, password) 
	VALUES ($1, $2)
	RETURNING id, version, email, created_At; 
	`

	row := r.pool.QueryRow(ctx, query, user.Email, user.Password)

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.Email,
		&userModel.CreatedAt,
	)
	if err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.Email,
		"",
		userModel.CreatedAt,
	)

	return userDomain, nil
}
