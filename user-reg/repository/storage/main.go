package storage

import (
	"context"
	"log"
	"time"
	"user-reg/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongo(ctx context.Context, cfg *config.Cfg) (*mongo.Client, error) {
	cli, err := mongo.NewClient(options.Client().ApplyURI(cfg.StorageAddr))
	if err != nil {
		return nil, err
	}
	ctx, _ = context.WithTimeout(ctx, 10*time.Second)
	err = cli.Connect(ctx)
	if err != nil {
		return nil, err
	}
	err = cli.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	databases, _ := cli.ListDatabaseNames(ctx, bson.M{})
	log.Printf("mongo: database list: %v", databases)

	return cli, nil
}
