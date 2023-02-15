package main

import (
	"flag"
	"log"
	"salt-gen/app"
)

var saltLength int = 12

func main() {
	var addr string
	flag.StringVar(&addr, "addr", "", "port address")
	flag.Parse()

	app := app.NewApp(saltLength)
	err := app.Start(addr)
	if err != nil {
		log.Fatalln(err)
	}
}
