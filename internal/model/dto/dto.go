package dto

import (
	"currency/internal/model/entity"
	"encoding/xml"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type Rates struct {
	XMLName     xml.Name   `xml:"rates"`
	Generator   string     `xml:"generator"`
	Title       string     `xml:"title"`
	Link        string     `xml:"link"`
	Description string     `xml:"description"`
	Copyright   string     `xml:"copyright"`
	Date        string     `xml:"date"`
	Item        []RateItem `xml:"item"`
}

func (r *RateItem) Entity(date time.Time) *entity.Currency {
	f, err := strconv.ParseFloat(r.Description, 64)
	if err != nil {
		zap.S().Errorf("invalid description: %s", r.Description)
		return nil
	}

	return &entity.Currency{
		Title: r.Title,
		Code:  strconv.Itoa(r.Quant),
		Value: f,
		ADate: date,
	}
}

type RateItem struct {
	Fullname    string `xml:"fullname"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Quant       int    `xml:"quant"`
	Index       string `xml:"index"`
	Change      string `xml:"change"`
}

type ErrorResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SuccessResp struct {
	Code    int  `json:"code"`
	Success bool `json:"success"`
}
