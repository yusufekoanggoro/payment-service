package repository

import (
	"context"

	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/domain"
)

type PaymentRepository interface {
	Save(ctx context.Context, payment *domain.Payment) error
}

type IdempotencyRepository interface {
	GetByKey(ctx context.Context, key string, requestHash string) (*domain.IdempotencyKey, error)
	Save(ctx context.Context, idempotencyKey *domain.IdempotencyKey) error
	Update(ctx context.Context, idempotency *domain.IdempotencyKey) error
}
