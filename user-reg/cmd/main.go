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
	defer app.Stop(ctx)
	err = app.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
