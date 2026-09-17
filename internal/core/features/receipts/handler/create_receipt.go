package receipts_handler

import (
	"net/http"
	"time"

	"github.com/neonexer/konvert-backend/internal/core/domain"
	core_logger "github.com/neonexer/konvert-backend/internal/core/logger"
	core_http_response "github.com/neonexer/konvert-backend/internal/core/transport/http/response"
)

type ReceiptCreateRequest struct {
	domain.ReceiptRequest
	Expenses []domain.ExpenseRequest
} // @name ReceiptCreateRequest

type ReceiptResponse struct {
	domain.Receipt
	Expenses []domain.Expense
} // @name ReceiptResponse

// @CreateReceipt godoc
// @Summary Создание чека
// @Description Создание нового чека в системе
// @Tags receipts
// @Accept json
// @Produce json
// @Param request body ReceiptCreateRequest true "Тело запроса для регистрации пользователя"
// @Success 201 {object} ReceiptResponse "Успешно созданный чек"
// @Failure 400 {object} core_http_response.ErrorResponse "Ошибка валидации данных"
// @Failure 409 {object} core_http_response.ErrorResponse "Конфликт"
// @Router /receipts [post]
func (h *ReceiptsHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	creationDate := time.Now()

	response := ReceiptResponse{
		domain.Receipt{
			Id:      1,
			Version: 1,

			UserId: 1,
			CreatedAt: creationDate,
			Name:    "чек за вкусвилл",
			Amount:  2500.00,
		},
		[]domain.Expense{
			{
				Id:        1,
				Version:   1,
				Name:      "Пирожок с капустой",
				Amount:    200.00,
				ReceiptId: 1,
				UserId: 1,
				CategoryId: 1,
				CreatedAt: creationDate,
			},
		},
	}

	responseHandler.JSONResponse(response, http.StatusCreated)
}
