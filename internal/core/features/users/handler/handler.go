package users_handler

import (
	"context"
	"net/http"

	"github.com/neonexer/konvert-backend/internal/core/domain"
	core_http_server "github.com/neonexer/konvert-backend/internal/core/transport/http/server"
)

type UsersHandler struct {
	usersService UsersService
}

type UsersService interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	)(domain.User, error)

	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	)([]domain.User, error)
}

func NewUsersHandler(usersService UsersService) *UsersHandler {
	return &UsersHandler{
		usersService: usersService,
	}
}

func (h *UsersHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: h.GetUsers,
		},
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
	}
}
