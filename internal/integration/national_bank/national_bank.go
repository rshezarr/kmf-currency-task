package national_bank

import (
	"context"
	"currency/internal/config"
	"currency/internal/model"
	"currency/internal/model/dto"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type NationalBank interface {
	GetRates(ctx context.Context, date string) (*dto.Rates, *model.CustomError)
}

type nationalBank struct {
	cfg *config.NationalBank
}

func NewNationalBank(cfg *config.NationalBank) NationalBank {
	return &nationalBank{
		cfg: cfg,
	}
}

func (n *nationalBank) GetRates(ctx context.Context, date string) (*dto.Rates, *model.CustomError) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s?fdate=%s", n.cfg.GetFullURL(), date), nil)
	if err != nil {
		return nil, model.NewError(http.StatusInternalServerError, "Could not create request", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, model.NewError(http.StatusInternalServerError, "Could not make request", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, model.NewError(http.StatusInternalServerError, "Could not read body", err)
	}

	rates := new(dto.Rates)
	rates.Item = make([]dto.RateItem, 0)

	if err := xml.Unmarshal(body, rates); err != nil {
		return nil, model.NewError(http.StatusInternalServerError, "Could not unmarshal body", err)
	}

	return rates, nil
}
