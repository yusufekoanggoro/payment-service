package paymentgateway

type PaymentGatewayResponse struct {
	TransactionID string `json:"transactionId"`
	Status        string `json:"status"`      // pending, success, failed
	RedirectURL   string `json:"redirectUrl"` // kalau metode butuh redirect (ex: e-wallet)
	PaymentType   string `json:"paymentType"` // va, qris, credit_card, etc
	RawResponse   any    `json:"rawResponse"` // opsional, untuk debugging atau audit
}
