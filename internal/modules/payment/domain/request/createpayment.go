package request

type CreatePaymentRequest struct {
	OrderID        string  `json:"orderId" binding:"required"`
	PaymentGateway string  `json:"paymentGateway" binding:"required"`
	PaymentType    string  `json:"paymentType" binding:"required"`
	Bank           string  `json:"bank" binding:"required_if=PaymentType va"`
	Amount         float64 `json:"amount" binding:"required,gt=0"`
}
