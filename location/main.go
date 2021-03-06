package main

import (
	"fmt"

	"github.com/orderforme/location/config"
	"github.com/orderforme/location/location"
	"github.com/orderforme/location/middelware"
	"github.com/orderforme/location/router"
)

func main() {
	// create a new echo instance
	r := router.New()

	route := r.Group("/api")

	middelware.ConfigMiddelware(route)

	handler, err := location.NewHandler()

	if err != nil {
		panic(err)
	}

	defer handler.Done()

	handler.Register(route)

	r.Logger.Fatal(r.Start(fmt.Sprint(":", config.AppConfig.ServerPort)))
}
