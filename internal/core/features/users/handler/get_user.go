package users_handler

import (
	"net/http"

	_ "github.com/neonexer/konvert-backend/internal/core/domain"
	core_logger "github.com/neonexer/konvert-backend/internal/core/logger"
	core_http_request "github.com/neonexer/konvert-backend/internal/core/transport/http/request"
	core_http_response "github.com/neonexer/konvert-backend/internal/core/transport/http/response"
)

// @GetUser godoc
// @Summary Получение пользователя по ID
// @Description Получение пользователя по ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "127"
// @Success 200 {object} domain.UserResponse "Успешно полученный пользователь"
// @Failure 400 {object} core_http_response.ErrorResponse "Ошибка валидации данных"
// @Failure 404 {object} core_http_response.ErrorResponse "Пользователь не найден"
// @Router /user [get]
func (h *UsersHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID path value",
		)

		return
	}

	user, err := h.usersService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user",
		)

		return
	}

	response := dtoFromDomain(user)

	responseHandler.JSONResponse(response, http.StatusOK)
}