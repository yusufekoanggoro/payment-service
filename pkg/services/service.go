package services

import (
	"github.com/yusufekoanggoro/payment-service/pkg/services/payment"
)

type ExternalService struct {
	PaymentGateways *payment.PaymentService
}

func NewExternalService() *ExternalService {
	return &ExternalService{
		PaymentGateways: payment.NewPaymentService(),
	}
}
