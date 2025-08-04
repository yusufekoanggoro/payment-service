package payment

import (
	"fmt"

	"github.com/yusufekoanggoro/payment-service/pkg/services/payment/gateways/doku"
	"github.com/yusufekoanggoro/payment-service/pkg/services/payment/gateways/midtrans"
	"github.com/yusufekoanggoro/payment-service/pkg/services/payment/strategy"
)

type PaymentService struct {
	Gateways map[string]strategy.PaymentStrategy
}

func NewPaymentService() *PaymentService {
	return &PaymentService{
		Gateways: map[string]strategy.PaymentStrategy{
			"midtrans": midtrans.NewMidtransGateway(),
			"doku":     doku.NewDokuGateway(),
		},
	}
}

func (ps *PaymentService) ProcessPayment(gateway, method string, amount float64) error {
	paymentGateway, ok := ps.Gateways[gateway]
	if !ok {
		return fmt.Errorf("payment gateway %s not supported", gateway)
	}
	return paymentGateway.Charge(method, amount)
}

func (ps *PaymentService) GetPaymentStatus(gateway, orderId string) (*strategy.PaymentStatusResponse, error) {
	paymentGateway, ok := ps.Gateways[gateway]
	if !ok {
		return &strategy.PaymentStatusResponse{}, fmt.Errorf("payment gateway %s not supported", gateway)
	}
	return paymentGateway.GetPaymentStatus(orderId)
}
