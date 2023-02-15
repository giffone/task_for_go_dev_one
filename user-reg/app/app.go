package app

import (
	"context"
	"user-reg/config"
	"user-reg/controller"
	"user-reg/helper"
	"user-reg/repository/cli"
	"user-reg/repository/storage"
	"user-reg/service"

	"github.com/go-playground/validator/v10"
)

func NewApp(ctx context.Context, cfg *config.Cfg) (Server, error) {
	// envoriments
	env, err := NewEnv(ctx, cfg)
	if err != nil {
		return nil, err
	}
	validate := validator.New()
	validate.RegisterValidation("email", func(fl validator.FieldLevel) bool {
		return helper.ValidateEmail(fl.Field().String())
	})

	// repo
	sto := storage.New(env.MongoCli())
	resty := cli.New(cfg, env.RestyCli())
	// service
	svc := service.New(sto, resty)
	// controller
	ctl := controller.New(svc, validate)

	return NewSrv(cfg, env, ctl), nil
}
