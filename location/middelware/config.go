package middelware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func ConfigMiddelware(route *echo.Group) {
	route.Use(middleware.LoggerWithConfig(logConfig))
	route.Use(middleware.RecoverWithConfig(recoConfig))
	route.Use(middleware.CORSWithConfig(corsConfig))
	//route.Use(middleware.JWTWithConfig(autoConfig))
}
