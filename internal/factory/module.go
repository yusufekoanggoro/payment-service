package factory

import (
	"database/sql"

	"github.com/yusufekoanggoro/payment-service/internal/factory/iface"
	"github.com/yusufekoanggoro/payment-service/internal/modules/payment"
)

type ModuleFactory interface {
	RestHandler() iface.RestHandler
}

func InitAllModule(db *sql.DB) []ModuleFactory {

	modules := []ModuleFactory{
		payment.NewModule(db),

		// Add initialization for other modules below
		// "modulename": modulePackage.NewModule(db),
	}

	return modules
}
