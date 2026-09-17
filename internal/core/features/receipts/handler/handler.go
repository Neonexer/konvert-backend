package receipts_handler

import (
	"context"
	"net/http"

	"github.com/neonexer/konvert-backend/internal/core/domain"
	core_http_server "github.com/neonexer/konvert-backend/internal/core/transport/http/server"
)

type ReceiptsHandler struct {
	// receiptsService ReceiptsService
}

type ReceiptsService interface {
	CreateReceipt(
		ctx context.Context,
		receipt domain.Receipt,
	)(domain.Receipt, error)
}

func NewReceiptsHandler() *ReceiptsHandler {
	return &ReceiptsHandler{}
}

// func NewReceiptsHandler(receiptsService ReceiptsService) *ReceiptsHandler {
// 	return &ReceiptsHandler{
// 		receiptsService: receiptsService,
// 	}
// }

func (h *ReceiptsHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/receipts",
			Handler: h.CreateUser,
		},
	}
}
