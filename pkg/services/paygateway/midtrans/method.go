package midtrans

import "fmt"

type PaymentMethodStrategy interface {
	Pay(amount float64) error
}

type MidtransBankTransfer struct{}

func (m MidtransBankTransfer) Pay(amount float64) error {
	fmt.Println("[Midtrans - Bank Transfer] Bayar", amount)
	return nil
}

type MidtransQRIS struct{}

func (m MidtransQRIS) Pay(amount float64) error {
	fmt.Println("[Midtrans - QRIS] Bayar", amount)
	return nil
}
