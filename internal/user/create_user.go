package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_logger "github.com/neonexer/konvert-backend/internal/core/logger"
)


// @CreateUser godoc
// @Summary Создание пользователя
// @Description Создание нового пользователя в системе
// @Tags users
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "Тело запроса для регистрации пользователя"
// @Success 201 {object} UserResponse "Успешно созданный пользователь"
// @Failure 400 {object} any "Ошибка валидации данных"
// @Router /users [post]
func (h *UsersHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	log.Debug("invoke CreateUser Handler")

	var request RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println("error")
	}

	w.WriteHeader(http.StatusOK)
}