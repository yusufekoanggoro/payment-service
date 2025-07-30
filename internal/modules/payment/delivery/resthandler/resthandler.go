package resthandler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/yusufekoanggoro/payment-service/internal/factory/iface"
	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/domain/request"
	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/usecase"
	"github.com/yusufekoanggoro/payment-service/pkg/middleware"
	"github.com/yusufekoanggoro/payment-service/pkg/utils"
)

type restHandler struct {
	uc usecase.PaymentUsecase
	mw middleware.Middleware
}

func NewRestHandler(uc usecase.PaymentUsecase, mw middleware.Middleware) iface.RestHandler {
	return &restHandler{uc: uc, mw: mw}
}

func (h *restHandler) RegisterRoutes(router gin.IRoutes) {
	router.POST("/payments", h.mw.Idempotency(), h.CreatePayment)
	router.POST("/payments/callback", h.PaymentCallback)
}

func (h *restHandler) CreatePayment(c *gin.Context) {
	var req request.CreatePaymentRequest
	req.IdempotencyKey = c.GetHeader("Idempotency-Key")
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request: "+err.Error())
		return
	}

	ctx := c.Request.Context()
	result, err := h.uc.CreatePayment(ctx, &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, result)
}

func (h *restHandler) PaymentCallback(c *gin.Context) {
	utils.SuccessResponse(c, http.StatusCreated, nil)
}
