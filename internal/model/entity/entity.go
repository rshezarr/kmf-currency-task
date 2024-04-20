package entity

import "time"

type Currency struct {
	Title string    `db:"title"`
	Code  string    `db:"code"`
	Value float64   `db:"value"`
	ADate time.Time `db:"a_date"`
}
