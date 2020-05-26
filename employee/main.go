package main

import (
	"fmt"

	"github.com/orderforme/config"
	"github.com/orderforme/employee"
	"github.com/orderforme/middelware"
	"github.com/orderforme/router"
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
