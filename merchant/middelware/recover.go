package middelware

import "github.com/labstack/echo/middleware"

//config the logger
var recoConfig = middleware.RecoverConfig{
	//Skipper:           DefaultSkipper,
	StackSize:         4 << 10, // 4 KB
	DisableStackAll:   false,
	DisablePrintStack: false,
}
