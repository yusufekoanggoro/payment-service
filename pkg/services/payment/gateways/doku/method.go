package doku

import "fmt"

type DokuBankTransfer struct{}

func (d *DokuBankTransfer) Pay(amount float64) error {
	fmt.Println("[Midtrans - Bank Transfer] Bayar", amount)
	return nil
}

type DokuQRIS struct{}

func (d *DokuQRIS) Pay(amount float64) error {
	fmt.Println("[Midtrans - QRIS] Bayar", amount)
	return nil
}
