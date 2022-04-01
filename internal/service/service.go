package service

import (
	config "payment-service/configs"
	"payment-service/internal/domain"
	"payment-service/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Payment interface {
	CreateTransactions(input domain.PaymentInfo) (domain.Transaction, error)
	MakePayment(input domain.PaymentInfo) (string, error)
}

type Service struct {
	Payment
}

func NewService(repos *repository.Repository, cfg *config.Config) *Service {
	return &Service{
		Payment: NewPaymentService(repos.Payment, cfg),
	}
}
