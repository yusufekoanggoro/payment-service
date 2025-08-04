package factory

import (
	"database/sql"

	"github.com/yusufekoanggoro/payment-service/internal/factory/iface"
	"github.com/yusufekoanggoro/payment-service/internal/modules/payment"
	"github.com/yusufekoanggoro/payment-service/pkg/middleware"
	"github.com/yusufekoanggoro/payment-service/pkg/services"
)

type ModuleFactory interface {
	RestHandler() iface.RestHandler
}

func InitAllModule(db *sql.DB, mw middleware.Middleware, services *services.ExternalService) []ModuleFactory {

	modules := []ModuleFactory{
		payment.NewModule(db, mw, services),

		// Add initialization for other modules below
		// "modulename": modulePackage.NewModule(db),
	}

	return modules
}
