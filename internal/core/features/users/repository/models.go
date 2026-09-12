package users_repository

import (
	"time"

	"github.com/neonexer/konvert-backend/internal/core/domain"
)

type UserModel struct {
	ID        int
	Version   int
	Email     string
	Password  string
	CreatedAt time.Time
}

func userDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))

	for i, user := range users {
		userDomains[i] = domain.NewUser(
			user.ID,
			user.Version,
			user.Email,
			user.Password,
			user.CreatedAt,
		)
	}

	return userDomains
}
