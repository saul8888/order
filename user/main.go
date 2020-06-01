package main

import (
	"fmt"

	"github.com/orderforme/user/config"
	"github.com/orderforme/user/middelware"
	"github.com/orderforme/user/router"
	"github.com/orderforme/user/user"
)

func main() {
	// create a new echo instance
	r := router.New()

	route := r.Group("/api")

	middelware.ConfigMiddelware(route)

	handler, err := user.NewHandler()

	if err != nil {
		panic(err)
	}

	defer handler.Done()

	handler.Register(route)

	r.Logger.Fatal(r.Start(fmt.Sprint(":", config.AppConfig.ServerPort)))
}
