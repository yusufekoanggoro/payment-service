package paygateway

type ProcessPaymentResponse struct {
	OrderID    string      `json:"orderId"`
	Status     string      `json:"status"`
	PaymentURL string      `json:"paymentUrl,omitempty"` // e.g., QR code or redirect URL
	Meta       interface{} `json:"meta,omitempty"`       // untuk detail tambahan sesuai metode
}

type PaymentStatusResponse struct {
	OrderID        string      `json:"orderId"`
	Status         string      `json:"status"` // pending, settlement, expire, etc.
	TransactionID  string      `json:"transactionId,omitempty"`
	PaymentType    string      `json:"paymentType,omitempty"`
	SettlementTime string      `json:"settlementTime,omitempty"` // ISO8601
	Meta           interface{} `json:"meta,omitempty"`           // jika butuh detail tambahan
}
