package responses

import (
	"currency/internal/model"
	"currency/internal/model/dto"
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

// NewResponse успешный ответ
func NewResponse(w http.ResponseWriter, respBody interface{}) {
	resp, err := json.Marshal(respBody)
	if err != nil {
		zap.S().Errorf("error while marshaling response: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(resp); err != nil {
		zap.S().Errorf("error while sending response: %v", err)
		return
	}
}

// NewErrorResponse ответ с кодом и сообщением
func NewErrorResponse(w http.ResponseWriter, comErr *model.CustomError) {
	zap.S().Error(comErr.Error)
	errResp := &dto.ErrorResp{
		Code:    comErr.StatusCode,
		Message: comErr.RespMessage,
	}

	b, err := json.Marshal(errResp)
	if err != nil {
		zap.S().Error(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(comErr.StatusCode)
	if _, err := w.Write(b); err != nil {
		zap.S().Error(err.Error())
		return
	}
}
