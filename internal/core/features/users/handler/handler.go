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

	GetUser(
		ctx context.Context,
		id int,
	) (domain.User, error)

	PatchUser(
		ctx context.Context,
		id int,
		patch domain.UserPatch,
	) (domain.User, error)

	DeleteUser(
		ctx context.Context,
		id int,
	) error
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
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: h.GetUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{id}",
			Handler: h.PatchUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: h.DeleteUser,
		},
	}
}
