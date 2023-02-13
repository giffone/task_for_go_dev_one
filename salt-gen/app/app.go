package app

import (
	"salt-gen/controller"
)

type App struct {
	srv Server
}

func NewApp(saltLength int) Server {
	ctl := controller.New(saltLength)
	app := App{
		srv: NewSrv(ctl),
	}
	return app.srv
}
