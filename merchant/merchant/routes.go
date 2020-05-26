package merchant

import "github.com/labstack/echo"

// Register a new user
func (handler *Handler) Register(route *echo.Group) {

	merchant := route.Group("/merchant")
	merchant.GET("", handler.GetMerchantByID)
	merchant.GET("/total", handler.GetMerchants)
	merchant.POST("", handler.CreateMerchant)
	merchant.PUT("", handler.UpdateMerchant)
	merchant.DELETE("", handler.DeleteMerchant)
	merchant.GET("/search", handler.SearchMerchant)
}
