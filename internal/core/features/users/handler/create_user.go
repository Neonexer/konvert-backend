package users_handler

import (
	"net/http"

	"github.com/neonexer/konvert-backend/internal/core/domain"
	core_logger "github.com/neonexer/konvert-backend/internal/core/logger"
	core_http_request "github.com/neonexer/konvert-backend/internal/core/transport/http/request"
	core_http_response "github.com/neonexer/konvert-backend/internal/core/transport/http/response"
)

// @CreateUser godoc
// @Summary Создание пользователя
// @Description Создание нового пользователя в системе
// @Tags users
// @Accept json
// @Produce json
// @Param request body domain.RegisterRequest true "Тело запроса для регистрации пользователя"
// @Success 201 {object} domain.UserResponse "Успешно созданный пользователь"
	// @Failure 400 {object} core_http_response.ErrorResponse "Ошибка валидации данных"
	// @Failure 409 {object} core_http_response.ErrorResponse "Конфликт"
// @Router /users [post]
func (h *UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var request domain.RegisterRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")

		return
	}

	userDomain := domainFromDTO(request)

	userDomain, err := h.usersService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")

		return
	}

	response := dtoFromDomain(userDomain)

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto domain.RegisterRequest) domain.User {
	return domain.NewUserUninitialized(dto.Email, dto.Password)
}

func dtoFromDomain(user domain.User) domain.UserResponse {
	return domain.UserResponse{
		Id:      user.Id,
		Version: user.Version,
		Email:   user.Email,
	}
}
