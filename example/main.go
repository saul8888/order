package main

import (
	"fmt"

	"github.com/orderforme/example/config"
	"github.com/orderforme/example/employee"
	"github.com/orderforme/example/middelware"
	"github.com/orderforme/example/router"
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
