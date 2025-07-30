package domain

type VirtualAccountResponse struct {
	VANumber string `json:"va_number"`
	Bank     string `json:"bank"`
}

type VirtualAccountStatusResponse struct {
	TransactionStatus string `json:"transaction_status"`
	SettlementTime    string `json:"settlement_time"`
	VANumber          string `json:"va_number"`
	Bank              string `json:"bank"`
}

type QRISResponse struct {
	QRCodeURL string `json:"qr_code_url"`
}

type QRISStatusResponse struct {
	OrderID           string `json:"order_id"`
	TransactionStatus string `json:"transaction_status"`
	PaymentType       string `json:"payment_type"`
	GrossAmount       string `json:"gross_amount"`
	Currency          string `json:"currency"`
	TransactionTime   string `json:"transaction_time"`
	TransactionID     string `json:"transaction_id"`
	MerchantID        string `json:"merchant_id"`
	QRCodeURL         string `json:"qr_code_url,omitempty"` // bisa dari "actions"
	ExpiryTime        string `json:"expiry_time,omitempty"` // jika tersedia
}
