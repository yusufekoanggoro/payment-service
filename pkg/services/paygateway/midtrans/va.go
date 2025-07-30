package midtrans

import (
	"context"

	"github.com/yusufekoanggoro/payment-service/pkg/services/paygateway"
)

type vaStrategy struct{}

func NewVAStrategy() paygateway.Strategy {
	return &vaStrategy{}
}

func (s *vaStrategy) ProcessPayment(ctx context.Context, req *paygateway.ProcessPaymentRequest) (*paygateway.ProcessPaymentResponse, error) {
	return &paygateway.ProcessPaymentResponse{}, nil
}

func (s *vaStrategy) GetStatus(ctx context.Context, orderID string) (*paygateway.PaymentStatusResponse, error) {
	return &paygateway.PaymentStatusResponse{}, nil
}
