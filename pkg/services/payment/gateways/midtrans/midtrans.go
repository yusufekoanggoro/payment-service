package midtrans

import (
	"fmt"

	"github.com/yusufekoanggoro/payment-service/pkg/services/payment/strategy"
)

type MidtransGateway struct {
	methods map[string]strategy.PaymentMethodStrategy
}

func NewMidtransGateway() *MidtransGateway {
	return &MidtransGateway{
		methods: map[string]strategy.PaymentMethodStrategy{
			"bank_transfer": MidtransBankTransfer{},
			"qris":          MidtransQRIS{},
		},
	}
}

func (m *MidtransGateway) Charge(method string, amount float64) error {
	if handler, ok := m.methods[method]; ok {
		return handler.Pay(amount)
	}
	return fmt.Errorf("metode %s tidak dikenali", method)
}

func (m *MidtransGateway) GetPaymentStatus(orderID string) (*strategy.PaymentStatusResponse, error) {
	return &strategy.PaymentStatusResponse{}, nil
}

func (m *MidtransGateway) Refund(transactionID string) error {
	// Simulasi refund
	return nil
}
