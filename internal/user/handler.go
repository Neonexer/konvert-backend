package user

import (
	"net/http"

	core_http_server "github.com/neonexer/konvert-backend/internal/core/transport/http/server"
)

type UsersHandler struct {
	usersService UsersService
}

type UsersService interface {
}

func NewUsersHandler(usersService UsersService) *UsersHandler {
	return &UsersHandler{
		usersService: usersService,
	}
}

func (h *UsersHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
	}
}
