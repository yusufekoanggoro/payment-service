package domain

type VirtualAccountRequest struct {
	OrderID string  `json:"order_id"`
	Amount  float64 `json:"amount"`
	Bank    string  `json:"bank"` // contoh: "bca", "bni"
}

type QRISRequest struct {
	OrderID string  `json:"order_id"`
	Amount  float64 `json:"amount"`
}
