package currency

import (
	"context"
	"currency/internal/integration"
	"currency/internal/integration/national_bank"
	"currency/internal/model"
	"currency/internal/model/dto"
	"currency/internal/model/entity"
	"currency/internal/repository"
	currencyRepo "currency/internal/repository/currency"
	"go.uber.org/zap"
)

type CurrencyService interface {
	SaveRates(ctx context.Context, date string) *model.CustomError
	GetRates(date, code string) ([]entity.Currency, *model.CustomError)
}

type currencySvc struct {
	currRepo currencyRepo.CurrencyRepository
	nb       national_bank.NationalBank
}

func NewCurrencyService(repo repository.Repository, itg integration.Integration) CurrencyService {
	return &currencySvc{
		currRepo: repo.CurrencyRepoInit(),
		nb:       itg.NationalBankInit(),
	}
}

func (c *currencySvc) SaveRates(ctx context.Context, date string) *model.CustomError {
	rates, comErr := c.nb.GetRates(ctx, date)
	if comErr != nil {
		return comErr
	}

	go func(rates *dto.Rates) {
		if comErr := c.currRepo.SaveRates(rates); comErr != nil {
			zap.S().Error(comErr.Error)
			return
		}
	}(rates)

	return nil
}

func (c *currencySvc) GetRates(date, code string) ([]entity.Currency, *model.CustomError) {
	rates, comErr := c.currRepo.GetRates(date, code)
	if comErr != nil {
		return nil, comErr
	}

	return rates, nil
}
