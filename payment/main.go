package main

import (
	"fmt"

	"github.com/orderforme/payment/config"
	"github.com/orderforme/payment/middelware"
	"github.com/orderforme/payment/payment"
	"github.com/orderforme/payment/router"
)

func main() {
	// create a new echo instance
	r := router.New()

	route := r.Group("/api")

	middelware.ConfigMiddelware(route)

	handler, err := payment.NewHandler()

	if err != nil {
		panic(err)
	}

	defer handler.Done()

	handler.Register(route)

	r.Logger.Fatal(r.Start(fmt.Sprint(":", config.AppConfig.ServerPort)))
}
