package main

import (
	_ "currency/docs"
	"currency/internal"
)

// @title			kmf tech task currency
func main() {
	app := internal.Init()
	app.Run()
}
