package catalog

import "github.com/labstack/echo"

// Register a new user
func (handler *Handler) Register(route *echo.Group) {

	catalog := route.Group("/catalog")
	catalog.GET("", handler.GetCatalogByID)
	catalog.GET("/total", handler.GetCatalogs)
	catalog.POST("", handler.CreateCatalog)
	catalog.PUT("", handler.UpdateCatalog)
	catalog.DELETE("", handler.DeleteCatalog)
	catalog.GET("/search", handler.SearchCatalog)
	catalog.PUT("/test", handler.AddCatalog)
}
