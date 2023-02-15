package main

import (
	"context"
	"log"
	"user-reg/app"
	"user-reg/config"
)

func main() {
	ctx := context.Background()

	cfg := config.Cfg{}
	cfg.Read()

	app, err := app.NewApp(ctx, &cfg)
	if err != nil {
		log.Fatalln(err)
	}
	err = app.Start(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
