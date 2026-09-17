package users_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/neonexer/konvert-backend/internal/core/domain"
	core_errors "github.com/neonexer/konvert-backend/internal/core/errors"
	core_postgres_pool "github.com/neonexer/konvert-backend/internal/core/repository/postgres/pool"
)

func (r *UsersRepository) GetUser(
	ctx context.Context,
	id int,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, version, email, password, created_at
		FROM konvert.users
		WHERE id=$1;
		`

	row := r.pool.QueryRow(ctx, query, id)

	var userModel UserModel

	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.Email,
		&userModel.Password,
		&userModel.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with id='%d': %w",
				id,
				core_errors.ErrNotFound,
			)
		}

		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.Email,
		userModel.Password,
		userModel.CreatedAt,
	)

	return userDomain, nil
}
