package server

import (
	"context"
	"currency/internal/config"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Server struct {
	server          *http.Server
	ServerErrorChan chan error
}

func NewServer(router *mux.Router) *Server {
	return &Server{
		server: &http.Server{
			Addr:           ":" + config.Conf.Port,
			Handler:        router,
			MaxHeaderBytes: config.Conf.MaxHeaderBytes << 20,
			ReadTimeout:    time.Duration(config.Conf.ReadTimeout) * time.Second,
			WriteTimeout:   time.Duration(config.Conf.WriteTimeout) * time.Second,
			IdleTimeout:    time.Duration(config.Conf.IdleTimeout) * time.Second,
		},
		ServerErrorChan: make(chan error, 1),
	}
}

func (s *Server) Run() {
	s.ServerErrorChan <- s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) Notify() <-chan error {
	return s.ServerErrorChan
}
