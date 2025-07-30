package paygateway

import (
	"context"
)

type Strategy interface {
	ProcessPayment(ctx context.Context, req *ProcessPaymentRequest) (*ProcessPaymentResponse, error)
	GetStatus(ctx context.Context, orderID string) (*PaymentStatusResponse, error)
}
