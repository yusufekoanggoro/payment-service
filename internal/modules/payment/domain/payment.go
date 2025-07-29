package domain

import "time"

type Payment struct {
	ID             string    `json:"id" db:"id"`                          // UUID atau string unik
	OrderID        string    `json:"orderId" db:"order_id"`               // Foreign key ke tabel orders
	PaymentGateway string    `json:"paymentGateway" db:"payment_gateway"` // contoh: midtrans, stripe
	PaymentType    string    `json:"paymentType" db:"payment_type"`       // contoh: credit_card, qris, va_bca
	ExternalID     string    `json:"externalId" db:"external_id"`         // ID transaksi dari payment gateway
	Amount         float64   `json:"amount" db:"amount"`                  // Jumlah yang dibayar
	Status         string    `json:"status" db:"status"`                  // pending, success, failed, expired
	PaidAt         time.Time `json:"paidAt,omitempty" db:"paid_at"`       // Nullable jika belum dibayar
	CreatedAt      time.Time `json:"createdAt" db:"created_at"`           // Timestamp pembuatan
}
