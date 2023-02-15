package app

import (
	"context"
	"user-reg/config"
	"user-reg/repository/cli"
	"user-reg/repository/storage"

	"github.com/go-resty/resty/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type Envoriment interface {
	Stop(ctx context.Context)
	MongoCli() *mongo.Client
	RestyCli() *resty.Client
}

func NewEnv(ctx context.Context, cfg *config.Cfg) (Envoriment, error) {
	var err error
	e := env{}
	// mongo client
	e.storageCli, err = storage.NewMongo(ctx, cfg)
	if err != nil {
		return nil, err
	}
	// resty client
	e.restyCli = cli.NewResty(cfg)
	return &e, nil
}

type env struct {
	storageCli *mongo.Client
	restyCli   *resty.Client
}

func (e *env) MongoCli() *mongo.Client {
	return e.storageCli
}

func (e *env) RestyCli() *resty.Client {
	return e.restyCli
}

func (e *env) Stop(ctx context.Context) {
	if e.storageCli != nil {
		e.storageCli.Disconnect(ctx)
	}
}
