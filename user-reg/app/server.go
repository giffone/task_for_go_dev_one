package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"user-reg/config"
	"user-reg/controller"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server interface {
	Start(ctx context.Context) error
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

func (s *server) Start(ctx context.Context) error {
	defer s.Stop(ctx)

	chQuit := make(chan os.Signal, 1)
	chErr := make(chan error, 1)
	signal.Notify(chQuit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("starting at port: %s\n", s.cfg.Addr)
		chErr <- http.ListenAndServe(s.cfg.Addr, s.r)
	}()

	select {
	case err := <-chErr:
		return err
	case quit := <-chQuit:
		log.Printf("server stopped by \"%s\" signal\n", quit)
	}
	return nil
}

func (s *server) Stop(ctx context.Context) {
	log.Println("gracefully stop server")
	s.env.Stop(ctx)
	log.Println("gracefully stop server... done")
}

func (s *server) registerCtl() {
	s.r.Post("/create-user", s.c.CreateUser)
	s.r.Get("/get-user/{email}", s.c.GetUser)
}
