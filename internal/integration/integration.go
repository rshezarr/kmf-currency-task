package integration

import (
	"currency/internal/config"
	"currency/internal/integration/national_bank"
)

type Integration interface {
	NationalBankInit() national_bank.NationalBank
}
type integration struct {
	cfg *config.NationalBank

	nb national_bank.NationalBank
}

func NewIntegration(cfg *config.NationalBank) Integration {
	return &integration{
		cfg: cfg,
	}
}

func (i *integration) NationalBankInit() national_bank.NationalBank {
	if i.nb == nil {
		i.nb = national_bank.NewNationalBank(i.cfg)
	}

	return i.nb
}
