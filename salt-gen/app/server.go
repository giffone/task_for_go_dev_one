package app

import (
	"log"
	"net/http"
	"salt-gen/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server interface {
	Start(addr string) error
}

func NewSrv(ctl controller.Controller) Server {
	s := server{
		r: chi.NewRouter(),
		c: ctl,
	}
	s.r.Use(middleware.Logger)
	s.registerCtl()
	return &s
}

type server struct {
	r *chi.Mux
	c controller.Controller
}

func (s *server) Start(addr string) error {
	log.Printf("starting at port: %s", addr)
	return http.ListenAndServe(addr, s.r)
}

func (s *server) registerCtl() {
	s.r.Post("/generate-salt", s.c.SaltGen)
}
