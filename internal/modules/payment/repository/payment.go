package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/domain"
)

type paymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return &paymentRepository{db}
}

func (p *paymentRepository) Save(ctx context.Context, payment *domain.Payment) error {
	stmt, err := p.db.PrepareContext(ctx, `
		INSERT INTO payments (
			order_id, payment_gateway, payment_type, external_id, amount, status, created_at
		) VALUES (?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		return fmt.Errorf("prepare insert payment: %w", err)
	}
	defer stmt.Close() // Terlalu banyak prepared statement terbuka → error: too many open statements

	_, err = stmt.ExecContext(ctx,
		payment.OrderID, payment.PaymentGateway, payment.PaymentType,
		payment.ExternalID, payment.Amount, payment.Status, payment.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("exec insert payment: %w", err)
	}

	return nil
}
