package app

import (
	"context"
	"log"
	"net/http"
	"user-reg/config"
	"user-reg/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server interface {
	Start() error
	Stop(ctx context.Context)
}

func NewSrv(cfg *config.Cfg, env Envoriment, ctl controller.Controller) Server {
	s := server{
		cfg: cfg,
		r:   chi.NewRouter(),
		c:   ctl,
		env: env,
	}
	s.r.Use(middleware.Logger)
	s.registerCtl()
	return &s
}

type server struct {
	cfg *config.Cfg
	r   *chi.Mux
	c   controller.Controller
	env Envoriment
}

func (s *server) Start() error {
	log.Printf("starting at port: %s\n", s.cfg.Addr)
	return http.ListenAndServe(s.cfg.Addr, s.r)
}

func (s *server) Stop(ctx context.Context) {
	log.Println("stopping server")
	s.env.Stop(ctx)

}

func (s *server) registerCtl() {
	s.r.Post("/create-user", s.c.CreateUser)
	s.r.Get("/get-user/{email}", s.c.GetUser)
}
