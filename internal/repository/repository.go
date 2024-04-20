package repository

import (
	"currency/internal/config"
	"currency/internal/repository/currency"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CurrencyRepoInit() currency.CurrencyRepository
}

type repo struct {
	cfg *config.Configs
	db  *sqlx.DB

	currencyRepo currency.CurrencyRepository
}

func NewRepository(cfg *config.Configs, db *sqlx.DB) Repository {
	return &repo{
		cfg: cfg,
		db:  db,
	}
}

func (r *repo) CurrencyRepoInit() currency.CurrencyRepository {
	if r.currencyRepo == nil {
		r.currencyRepo = currency.NewCurrencyRepo(r.cfg, r.db)
	}

	return r.currencyRepo
}
