package middleware

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/repository"
)

type Middleware interface {
	Idempotency() gin.HandlerFunc
}

type middleware struct {
	repo repository.IdempotencyRepository
}

func NewMiddleware(db *sql.DB) *middleware {
	return &middleware{repo: repository.NewIdempotency(db)}
}
