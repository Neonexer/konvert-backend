package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_logger "github.com/neonexer/konvert-backend/internal/core/logger"
)


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