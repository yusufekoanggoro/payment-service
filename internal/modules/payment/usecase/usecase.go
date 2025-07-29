package usecase

import (
	"context"

	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/domain"
	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/domain/request"
)

type PaymentUsecase interface {
	CreatePayment(ctx context.Context, req *request.CreatePaymentRequest) (*domain.Payment, error)
}
