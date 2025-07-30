package midtrans

import (
	"context"

	"github.com/yusufekoanggoro/payment-service/pkg/services/paygateway"
)

type qrisStrategy struct{}

func NewQRISStrategy() paygateway.Strategy {
	return &qrisStrategy{}
}

func (s *qrisStrategy) ProcessPayment(ctx context.Context, req *paygateway.ProcessPaymentRequest) (*paygateway.ProcessPaymentResponse, error) {
	return &paygateway.ProcessPaymentResponse{}, nil
}

func (s *qrisStrategy) GetStatus(ctx context.Context, orderID string) (*paygateway.PaymentStatusResponse, error) {
	return &paygateway.PaymentStatusResponse{}, nil
}
