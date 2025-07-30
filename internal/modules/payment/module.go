package payment

import (
	"database/sql"

	"github.com/yusufekoanggoro/payment-service/internal/factory/iface"
	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/delivery/resthandler"
	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/repository"
	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/usecase"
	"github.com/yusufekoanggoro/payment-service/pkg/middleware"
	"github.com/yusufekoanggoro/payment-service/pkg/services/paygateway"
	"github.com/yusufekoanggoro/payment-service/pkg/services/paygateway/midtrans"
)

type module struct {
	restHandler iface.RestHandler
}

func NewModule(db *sql.DB, mw middleware.Middleware) *module {
	var mdl module

	repo := repository.NewPaymentRepository(db)
	idempoRepo := repository.NewIdempotency(db)
	uc := usecase.NewPaymentUsecase(repo, idempoRepo, map[string]paygateway.Strategy{
		"midtrans_va":   midtrans.NewVAStrategy(),
		"midtrans_qris": midtrans.NewQRISStrategy(),
	})
	restHandler := resthandler.NewRestHandler(uc, mw)

	mdl.restHandler = restHandler
	return &mdl
}

func (m *module) RestHandler() (d iface.RestHandler) {
	return m.restHandler
}
