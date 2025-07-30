package paygateway

type ProcessPaymentRequest struct {
	OrderID     string  `json:"orderId"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	CustomerID  string  `json:"customerId"`
	PaymentType string  `json:"paymentType"` // e.g., qris, va, credit_card
	CallbackURL string  `json:"callbackUrl"`
	Description string  `json:"description"`
}
