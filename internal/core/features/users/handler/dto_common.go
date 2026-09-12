package users_handler

import "github.com/neonexer/konvert-backend/internal/core/domain"

func domainFromDTO(dto domain.RegisterRequest) domain.User {
	return domain.NewUserUninitialized(dto.Email, dto.Password)
}

func dtoFromDomain(user domain.User) domain.UserResponse {
	return domain.UserResponse{
		Id:        user.Id,
		Version:   user.Version,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

func usersDTOFromDomains(users []domain.User) []domain.UserResponse {
	usersDTO := make([]domain.UserResponse, len(users))

	for i, user := range users {
		usersDTO[i] = dtoFromDomain(user)
	}

	return usersDTO
}
