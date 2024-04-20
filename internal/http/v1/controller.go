package v1

import (
	v1 "currency/internal/http/v1/currency"
	"currency/internal/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Controller interface {
	StartRoutes() *mux.Router
	InitRoutes()

	CurrencyControllerInit() v1.CurrencyController
}

type controller struct {
	router *mux.Router
	svc    service.Service

	currCtrl v1.CurrencyController
}

func NewController(svc service.Service) Controller {
	return &controller{
		router: mux.NewRouter(),
		svc:    svc,
	}
}

func (c *controller) StartRoutes() *mux.Router {
	c.router.HandleFunc("/currency/save/{date}", c.currCtrl.SaveRates).Methods(http.MethodGet)
	c.router.HandleFunc("/currency/{date}", c.currCtrl.GetRates).Methods(http.MethodGet)

	return c.router
}

func (c *controller) InitRoutes() {
	c.CurrencyControllerInit()
}

func (c *controller) CurrencyControllerInit() v1.CurrencyController {
	if c.currCtrl == nil {
		c.currCtrl = v1.NewCurrencyController(c.svc.CurrencyServiceInit())
	}

	return c.currCtrl
}
