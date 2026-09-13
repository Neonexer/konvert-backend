package users_handler

import (
	"net/http"

	_ "github.com/neonexer/konvert-backend/internal/core/domain"
	core_logger "github.com/neonexer/konvert-backend/internal/core/logger"
	core_http_response "github.com/neonexer/konvert-backend/internal/core/transport/http/response"
	core_http_utils "github.com/neonexer/konvert-backend/internal/core/transport/http/utils"
)

// @DeleteUser godoc
// @Summary Удаление пользователя по ID
// @Description Удаление пользователя по ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "127"
// @Success 204 "Пользователь успешно удалён"
// @Failure 400 {object} core_http_response.ErrorResponse "Ошибка валидации данных"
// @Failure 404 {object} core_http_response.ErrorResponse "Пользователь не найден"
// @Router /user [delete]
func (h *UsersHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID path value",
		)

		return
	}

	if err := h.usersService.DeleteUser(ctx, userID); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete user",
		)

		return
	}

	responseHandler.NoContentResponse()
}