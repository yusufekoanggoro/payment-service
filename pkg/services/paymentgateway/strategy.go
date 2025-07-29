package paymentgateway

import (
	"context"

	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/domain/request"
)

type Strategy interface {
	ProcessPayment(ctx context.Context, req *request.CreatePaymentRequest) (*PaymentGatewayResponse, error)
}
