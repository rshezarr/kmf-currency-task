package currency

import (
	"currency/internal/config"
	"currency/internal/model"
	"currency/internal/model/dto"
	"currency/internal/model/entity"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type CurrencyRepository interface {
	SaveRates(currency *dto.Rates) *model.CustomError
	GetRates(date, code string) ([]entity.Currency, *model.CustomError)
}

type currencyRepo struct {
	cfg *config.Configs
	db  *sqlx.DB
}

func NewCurrencyRepo(cfg *config.Configs, db *sqlx.DB) CurrencyRepository {
	return &currencyRepo{
		cfg: cfg,
		db:  db,
	}
}

func (c *currencyRepo) SaveRates(currencies *dto.Rates) *model.CustomError {
	query := fmt.Sprintf(SaveRates, CurrencyTableName)
	zap.S().Info(len(currencies.Item))
	adate, err := time.Parse("02.01.2006", currencies.Date)
	if err != nil {
		return model.NewError(http.StatusInternalServerError, "DataBase error", err)
	}
	zap.S().Info(adate)
	for _, currency := range currencies.Item {
		currEnt := currency.Entity(adate)
		_, err = c.db.Exec(query, sql.Named("Title", currEnt.Title), sql.Named("Code", currEnt.Code), sql.Named("Value", currEnt.Value), sql.Named("A_DATE", currEnt.ADate))
		if err != nil {
			zap.S().Errorf("error in db: %v", err)
			continue
		}
	}
	return nil
}

func (c *currencyRepo) GetRates(date, code string) ([]entity.Currency, *model.CustomError) {
	query := fmt.Sprintf(GetRates, CurrencyTableName)
	if code != "" {
		query += " AND CODE=@Code"
	}

	rows, err := c.db.Query(query, sql.Named("A_DATE", date), sql.Named("Code", code))
	if err != nil {
		return nil, model.NewError(http.StatusInternalServerError, "Error in db", err)
	}
	defer rows.Close()

	currencyData := make([]entity.Currency, 0)

	for rows.Next() {
		var currency entity.Currency
		err := rows.Scan(&currency.Title, &currency.Code, &currency.Value, &currency.ADate)
		if err != nil {
			return nil, model.NewError(http.StatusInternalServerError, "Error in db", err)
		}
		currencyData = append(currencyData, currency)
	}

	return currencyData, nil
}
