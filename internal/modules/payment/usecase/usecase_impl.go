package usecase

import (
	"context"
	"fmt"

	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/domain"
	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/domain/request"
	"github.com/yusufekoanggoro/payment-service/internal/modules/payment/repository"
	"github.com/yusufekoanggoro/payment-service/pkg/services"
)

type paymentUsecase struct {
	repo       repository.PaymentRepository
	idempoRepo repository.IdempotencyRepository
	services   *services.ExternalService
}

func NewPaymentUsecase(repo repository.PaymentRepository, idempoRepo repository.IdempotencyRepository, services *services.ExternalService) PaymentUsecase {
	return &paymentUsecase{
		repo:       repo,
		idempoRepo: idempoRepo,
		services:   services,
	}
}

func (p *paymentUsecase) CreatePayment(ctx context.Context, req *request.CreatePaymentRequest) (*domain.Payment, error) {
	err := p.services.PaymentGateways.Gateways[req.PaymentGateway].Charge(req.PaymentType, 1000)
	if err != nil {
		return nil, fmt.Errorf("failed to process payment: %w", err)
	}

	// [req.PaymentGateway].Charge(req.PaymentType, 20)

	// payStrategy, ok := p.services.PaymentGateways[req.PaymentGateway+"_"+req.PaymentType]
	// if !ok {
	// 	return nil, fmt.Errorf("unsupported payment gateway: %s", req.PaymentGateway)
	// }

	// paymentResp, err := payStrategy.ProcessPayment(ctx, &paygateway.ProcessPaymentRequest{})
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to process payment: %w", err)
	// }

	// newPayment := &domain.Payment{
	// 	OrderID:        req.OrderID,
	// 	PaymentGateway: req.PaymentGateway,
	// 	PaymentType:    req.PaymentType,
	// 	ExternalID:     paymentResp.OrderID,
	// 	Amount:         req.Amount,
	// 	Status:         paymentResp.Status, // default status
	// 	CreatedAt:      time.Now(),
	// 	PaidAt:         nil,
	// }

	// err = p.repo.Save(ctx, newPayment)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to save payment: %w", err)
	// }

	return newPayment, nil
}
