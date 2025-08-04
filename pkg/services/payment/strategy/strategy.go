package strategy

type PaymentMethodStrategy interface {
	Pay(amount float64) error
}

type PaymentStrategy interface {
	Charge(method string, amount float64) error
	Refund(transactionID string) error
	GetPaymentStatus(orderID string) (*PaymentStatusResponse, error)
}
