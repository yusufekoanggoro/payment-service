package doku

import (
	"fmt"

	"github.com/yusufekoanggoro/payment-service/pkg/services/payment/strategy"
)

type DokuGateway struct {
	methods map[string]strategy.PaymentMethodStrategy
}

func NewDokuGateway() *DokuGateway {
	return &DokuGateway{
		methods: map[string]strategy.PaymentMethodStrategy{
			"bank_transfer": &DokuBankTransfer{},
			"qris":          &DokuQRIS{},
		},
	}
}

func (d *DokuGateway) Charge(method string, amount float64) error {
	if handler, ok := d.methods[method]; ok {
		return handler.Pay(amount)
	}
	return fmt.Errorf("metode %s tidak dikenali", method)
}

func (doku *DokuGateway) GetPaymentStatus(orderID string) (*strategy.PaymentStatusResponse, error) {
	return &strategy.PaymentStatusResponse{}, nil
}

func (d *DokuGateway) Refund(transactionID string) error {
	// Simulasi refund
	return nil
}
