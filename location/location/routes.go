package location

import "github.com/labstack/echo"

// Register a new user
func (handler *Handler) Register(route *echo.Group) {

	location := route.Group("/location")
	location.GET("", handler.GetLocationByID)
	location.GET("/total", handler.GetLocations)
	location.POST("", handler.CreateLocation)
	location.PUT("", handler.UpdateLocation)
	location.DELETE("", handler.DeleteLocation)
	location.GET("/search", handler.SearchLocation)
}
