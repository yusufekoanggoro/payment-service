package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/domain"
	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/domain/request"
	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/repository"
	"github.com/yusufekoanggoro/payment-service/pkg/services/paymentgateway"
)

type paymentUsecase struct {
	repo              repository.PaymentRepository
	idempoRepo        repository.IdempotencyRepository
	paymentStrategies map[string]paymentgateway.Strategy
}

func NewPaymentUsecase(repo repository.PaymentRepository, idempoRepo repository.IdempotencyRepository, paymentStrategies map[string]paymentgateway.Strategy) PaymentUsecase {
	return &paymentUsecase{
		repo:              repo,
		idempoRepo:        idempoRepo,
		paymentStrategies: paymentStrategies,
	}
}

func (p *paymentUsecase) CreatePayment(ctx context.Context, req *request.CreatePaymentRequest) (*domain.Payment, error) {
	strategy, ok := p.paymentStrategies[req.PaymentGateway+"_"+req.PaymentType]
	if !ok {
		return nil, fmt.Errorf("unsupported payment gateway: %s", req.PaymentGateway)
	}

	paymentResp, err := strategy.ProcessPayment(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to process payment: %w", err)
	}

	newPayment := &domain.Payment{
		OrderID:        req.OrderID,
		PaymentGateway: req.PaymentGateway,
		PaymentType:    req.PaymentType,
		ExternalID:     paymentResp.TransactionID,
		Amount:         req.Amount,
		Status:         paymentResp.Status, // default status
		CreatedAt:      time.Now(),
		PaidAt:         nil,
	}

	err = p.repo.Save(ctx, newPayment)
	if err != nil {
		return nil, fmt.Errorf("failed to save payment: %w", err)
	}

	return newPayment, nil
}
