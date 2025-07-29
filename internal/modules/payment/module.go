package payment

import (
	"database/sql"

	"github.com/yusufekoanggoro/payment-service/internal/factory/iface"
	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/delivery/resthandler"
	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/repository"
	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/usecase"
	"github.com/yusufekoanggoro/payment-service/pkg/services/paymentgateway"
	"github.com/yusufekoanggoro/payment-service/pkg/services/paymentgateway/midtrans"
)

type module struct {
	restHandler iface.RestHandler
}

func NewModule(db *sql.DB) *module {
	var mdl module

	repo := repository.NewPaymentRepository(db)
	uc := usecase.NewPaymentUsecase(repo, map[string]paymentgateway.Strategy{
		"midtrans_va":   midtrans.NewMidtransVAHTTP(),
		"midtrans_qris": midtrans.NewMidtransQRISHTTP(),
	})
	restHandler := resthandler.NewRestHandler(uc)

	mdl.restHandler = restHandler
	return &mdl
}

func (m *module) RestHandler() (d iface.RestHandler) {
	return m.restHandler
}
