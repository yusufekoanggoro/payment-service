package paygateway

import (
	"context"

	"github.com/yusufekoanggoro/payment-service/pkg/services/paygateway/midtrans"
)

type Strategy interface {
	ProcessPayment(ctx context.Context, req *ProcessPaymentRequest) (*ProcessPaymentResponse, error)
	GetStatus(ctx context.Context, orderID string) (*PaymentStatusResponse, error)
}

func GetPaymentStrategyByCode(code string) Strategy {
	switch code {
	case "midtrans_qris":
		return midtrans.NewQRISStrategy()
	case "midtrans_va":
		return midtrans.NewVAStrategy()
	default:
		return nil
	}
}
