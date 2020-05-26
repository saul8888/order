package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func New() *echo.Echo {
	r := echo.New()
	r.Logger.SetLevel(log.DEBUG)
	//r.Use(middleware.Logger())
	r.Validator = NewValiDB()
	return r
}
