package main

import (
	"fmt"

	"github.com/orderforme/router"
	"github.com/saul8888/location/config"
	"github.com/saul8888/location/middelware"
)

func main() {
	// create a new echo instance
	r := router.New()

	route := r.Group("/api")

	middelware.ConfigMiddelware(route)

	handler, err := employee.NewHandler()

	if err != nil {
		panic(err)
	}

	defer handler.Done()

	handler.Register(route)

	r.Logger.Fatal(r.Start(fmt.Sprint(":", config.AppConfig.ServerPort)))
}
