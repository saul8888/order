package main

import (
	"fmt"

	"github.com/orderforme/merchant/config"
	"github.com/orderforme/merchant/merchant"
	"github.com/orderforme/merchant/middelware"
	"github.com/orderforme/merchant/router"
)

func main() {
	// create a new echo instance
	r := router.New()

	route := r.Group("/api")

	middelware.ConfigMiddelware(route)

	handler, err := merchant.NewHandler()

	if err != nil {
		panic(err)
	}

	defer handler.Done()

	handler.Register(route)

	r.Logger.Fatal(r.Start(fmt.Sprint(":", config.AppConfig.ServerPort)))
}
