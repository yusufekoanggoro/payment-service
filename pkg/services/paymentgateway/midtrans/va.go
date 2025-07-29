package midtrans

import (
	"context"

	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/domain/request"
	"github.com/yusufekoanggoro/payment-service/pkg/services/paymentgateway"
)

type midtransVA struct{}

func NewMidtransVAHTTP() paymentgateway.Strategy {
	return &midtransVA{}
}

func (m *midtransVA) ProcessPayment(ctx context.Context, req *request.CreatePaymentRequest) (*paymentgateway.PaymentGatewayResponse, error) {
	resp := struct {
		TransactionID     string
		TransactionStatus string
		VANumber          string
		Bank              string
	}{
		TransactionID:     "va-" + req.OrderID,
		TransactionStatus: "PENDING",
		VANumber:          "1234567890",
		Bank:              "bca",
	}

	return &paymentgateway.PaymentGatewayResponse{
		TransactionID: resp.TransactionID,
		Status:        resp.TransactionStatus,
		RedirectURL:   "", // VA tidak butuh redirect
		PaymentType:   "va",
		RawResponse:   resp,
	}, nil
}
