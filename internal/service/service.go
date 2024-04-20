package service

import (
	"currency/internal/integration"
	"currency/internal/repository"
	"currency/internal/service/currency"
)

type Service interface {
	CurrencyServiceInit() currency.CurrencyService
}

type service struct {
	repo repository.Repository
	itg  integration.Integration

	curr currency.CurrencyService
}

func NewService(repo repository.Repository, itg integration.Integration) Service {
	return &service{
		repo: repo,
		itg:  itg,
	}
}

func (s *service) CurrencyServiceInit() currency.CurrencyService {
	if s.curr == nil {
		s.curr = currency.NewCurrencyService(s.repo, s.itg)
	}

	return s.curr
}
