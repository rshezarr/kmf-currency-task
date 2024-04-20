package v1

import (
	"currency/internal/http/v1/responses"
	"currency/internal/model"
	"currency/internal/model/dto"
	"currency/internal/service/currency"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

type CurrencyController interface {
	SaveRates(w http.ResponseWriter, r *http.Request)
	GetRates(w http.ResponseWriter, r *http.Request)
}
type currencyCtrl struct {
	svc currency.CurrencyService
}

func NewCurrencyController(svc currency.CurrencyService) CurrencyController {
	return &currencyCtrl{
		svc: svc,
	}
}

// SaveRates godoc
//
//	@Summary		Сохранить записи
//	@Description	Метод для сохранение записей из НацБанка
//	@Tags			SaveRates
//	@Accept			string
//	@Produce		json
//	@Success		200						{object}	dto.SuccessResp
//	@Failure		400 					{object}	dto.ErrorResp
//	@Failure		404 					{object}	dto.ErrorResp
//	@Failure		500 					{object}	dto.ErrorResp
//	@Router			/currency/save/{date} [get]
func (c *currencyCtrl) SaveRates(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	date, ok := params["date"]
	if !ok {
		responses.NewErrorResponse(w, model.NewError(http.StatusBadRequest, "missing date", errors.New("missing date")))
		return
	}

	if comErr := c.svc.SaveRates(r.Context(), date); comErr != nil {
		responses.NewErrorResponse(w, comErr)
		return
	}

	responses.NewResponse(w, &dto.SuccessResp{http.StatusOK, true})
}

// GetRates godoc
//
//	@Summary		Получить записи
//	@Description	Метод для получение записей из базы
//	@Tags			GetRates
//	@Accept			string
//	@Produce		json
//	@Success		200						{object}	entity.Currency
//	@Failure		400 					{object}	dto.ErrorResp
//	@Failure		404 					{object}	dto.ErrorResp
//	@Failure		500 					{object}	dto.ErrorResp
//	@Router			/currency/save/{date}/{code} [get]
func (c *currencyCtrl) GetRates(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	date, ok := params["date"]
	if !ok {
		responses.NewErrorResponse(w, model.NewError(http.StatusBadRequest, "missing date", errors.New("missing date")))
		return
	}

	code := params["code"]

	rates, comErr := c.svc.GetRates(date, code)
	if comErr != nil {
		responses.NewErrorResponse(w, comErr)
		return
	}

	responses.NewResponse(w, rates)
}
