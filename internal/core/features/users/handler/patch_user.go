package users_handler

import (
	"fmt"
	"net/http"

	"github.com/neonexer/konvert-backend/internal/core/domain"
	core_logger "github.com/neonexer/konvert-backend/internal/core/logger"
	core_http_request "github.com/neonexer/konvert-backend/internal/core/transport/http/request"
	core_http_response "github.com/neonexer/konvert-backend/internal/core/transport/http/response"
	core_http_types "github.com/neonexer/konvert-backend/internal/core/transport/http/types"
	core_http_utils "github.com/neonexer/konvert-backend/internal/core/transport/http/utils"
)

type PatchUserRequest struct {
	Email    core_http_types.Nullable[string] `json:"email"`
	Password core_http_types.Nullable[string] `json:"password"`
} // @name PatchUserRequest

func (r *PatchUserRequest) Validate() error {
	if r.Email.Set {
		if r.Email.Value == nil {
			return fmt.Errorf("`Email` can't be null")
		}

		emailLen := len([]rune(*r.Email.Value))
		if emailLen < 3 || emailLen > 100 {
			return fmt.Errorf("`Email` must be between 3 and 100 symbols")
		}
	}

	if r.Password.Set {
		if r.Password.Value != nil {
			passwordLen := len([]rune(*r.Password.Value))
			if passwordLen < 8 {
				return fmt.Errorf("`Password` must be at least 8 symbols")
			}
		}
	}

	return nil
}

// @PatchUser godoc
// @Summary Обновление пользователя
// @Description Обновление пользователя в системе
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "127"
// @Param request body PatchUserRequest true "Тело запроса для регистрации пользователя"
// @Success 200 {object} domain.UserResponse "Успешно обновленный пользователь"
// @Failure 400 {object} core_http_response.ErrorResponse "Ошибка валидации данных"
// @Failure 409 {object} core_http_response.ErrorResponse "Конфликт"
// @Router /users [patch]
func (h *UsersHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userID path value")

		return
	}

	var request PatchUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")

		return
	}

	userPatch := userPatchFromRequest(request)

	userDomain, err := h.usersService.PatchUser(ctx, userID, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch user",
		)

		return
	}

	response := dtoFromDomain(userDomain)
	responseHandler.JSONResponse(response, http.StatusOK)

	w.WriteHeader(http.StatusOK)
}

func userPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.UserPatch{
		Email:    request.Email.ToDomain(),
		Password: request.Password.ToDomain(),
	}
}
