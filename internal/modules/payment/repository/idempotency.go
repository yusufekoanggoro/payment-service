package repository

import (
	"context"
	"database/sql"

	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/domain"
)

type idempotencyRepository struct {
	db *sql.DB
}

func NewIdempotency(db *sql.DB) IdempotencyRepository {
	return &idempotencyRepository{db: db}
}

func (i *idempotencyRepository) GetByKey(ctx context.Context, key string, requestHash string) (*domain.IdempotencyKey, error) {
	stmt, err := i.db.PrepareContext(ctx, `
		SELECT status_code, response_body, request_hash
		FROM idempotency_keys
		WHERE key = $1 AND request_hash = $2
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var res domain.IdempotencyKey
	err = stmt.QueryRowContext(ctx, key, requestHash).Scan(&res.StatusCode, &res.ResponseBody, &res.RequestHash)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (i *idempotencyRepository) Save(ctx context.Context, idempotency *domain.IdempotencyKey) error {
	stmt, err := i.db.PrepareContext(ctx, `
		INSERT INTO idempotency_keys (key, request_hash, status_code, response_body, created_at)
		VALUES (?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, idempotency.Key, idempotency.RequestHash, idempotency.StatusCode, idempotency.ResponseBody, idempotency.CreatedAt)
	return err
}

func (r *idempotencyRepository) Update(ctx context.Context, idempotency *domain.IdempotencyKey) error {
	query := `
		UPDATE idempotency_keys
		SET status_code = $1,
			response_body = $2
		WHERE key = $3
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, idempotency.StatusCode, idempotency.ResponseBody, idempotency.Key)
	return err
}
