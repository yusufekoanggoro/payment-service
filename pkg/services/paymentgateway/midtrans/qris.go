package midtrans

import (
	"context"

	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/domain/request"
	"github.com/yusufekoanggoro/payment-service/pkg/services/paymentgateway"
)

type midtransQRIS struct{}

func NewMidtransQRISHTTP() paymentgateway.Strategy {
	return &midtransQRIS{}
}

func (m *midtransQRIS) ProcessPayment(ctx context.Context, req *request.CreatePaymentRequest) (*paymentgateway.PaymentGatewayResponse, error) {
	resp := struct {
		TransactionID     string
		TransactionStatus string
		RedirectURL       string
	}{
		TransactionID:     "qris-" + req.OrderID,
		TransactionStatus: "PENDING",
		RedirectURL:       "https://midtrans.com/qris/123",
	}

	return &paymentgateway.PaymentGatewayResponse{
		TransactionID: resp.TransactionID,
		Status:        resp.TransactionStatus,
		RedirectURL:   resp.RedirectURL,
		PaymentType:   "qris",
		RawResponse:   resp,
	}, nil
}
