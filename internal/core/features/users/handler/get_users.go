package users_handler

import (
	"fmt"
	"net/http"

	_ "github.com/neonexer/konvert-backend/internal/core/domain"
	core_logger "github.com/neonexer/konvert-backend/internal/core/logger"
	core_http_request "github.com/neonexer/konvert-backend/internal/core/transport/http/request"
	core_http_response "github.com/neonexer/konvert-backend/internal/core/transport/http/response"
)

// @GetUsers godoc
// @Summary Получение пользователей
// @Description Получение пользователей
// @Tags users
// @Accept json
// @Produce json
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Success 200 {array} domain.UserResponse "Успешно полученные пользователи"
// @Failure 400 {object} core_http_response.ErrorResponse "Ошибка валидации данных"
// @Router /users [get]
func (h *UsersHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get 'limit/offset' query param")

		return
	}

	userDomains, err := h.usersService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get users")

		return
	}

	response := usersDTOFromDomains(userDomains)

	responseHandler.JSONResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	const (
		limitQueryParamKey = "limit"
		offsetQueryParamKey = "offset"
	)
	limit, err := core_http_request.GetIntQueryParam(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}

	offset, err := core_http_request.GetIntQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}

	return limit, offset, nil
}
