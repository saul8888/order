package customer

import "github.com/labstack/echo"

// Register a new user
func (handler *Handler) Register(route *echo.Group) {

	customer := route.Group("/customer")
	customer.GET("", handler.GetCustomerByID)
	customer.GET("/total", handler.GetCustomers)
	customer.POST("", handler.CreateCustomer)
	customer.PUT("", handler.UpdateCustomer)
	customer.DELETE("", handler.DeleteCustomer)
	customer.GET("/search", handler.SearchCustomer)
}
