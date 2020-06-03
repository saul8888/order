package main

import (
	"fmt"

	"github.com/orderforme/catalog/catalog"
	"github.com/orderforme/catalog/config"
	"github.com/orderforme/catalog/middelware"
	"github.com/orderforme/catalog/router"
)

func main() {
	// create a new echo instance
	r := router.New()

	route := r.Group("/api")

	middelware.ConfigMiddelware(route)

	handler, err := catalog.NewHandler()

	if err != nil {
		panic(err)
	}

	defer handler.Done()

	handler.Register(route)

	r.Logger.Fatal(r.Start(fmt.Sprint(":", config.AppConfig.ServerPort)))
}
