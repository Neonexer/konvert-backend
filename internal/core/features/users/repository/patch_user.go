package users_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/neonexer/konvert-backend/internal/core/domain"
	core_errors "github.com/neonexer/konvert-backend/internal/core/errors"
	core_postgres_pool "github.com/neonexer/konvert-backend/internal/core/repository/postgres/pool"
)

func (r *UsersRepository) PatchUser(
	ctx context.Context,
	id int,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE konvert.users
	SET
		email=$1,
		password=$2,
		version=version+1
	WHERE id=$3 AND version=$4
	RETURNING 
		id,
		version,
		email,
		password,
		created_at;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		user.Email,
		user.Password,
		id,
		user.Version,
	)

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
				"user with id='%d' concurrently accesed: %w",
				id,
				core_errors.ErrConflict,
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
