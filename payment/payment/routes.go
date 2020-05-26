package payment

import "github.com/labstack/echo"

// Register a new user
func (handler *Handler) Register(route *echo.Group) {

	payment := route.Group("/payment")
	payment.GET("", handler.GetPaymentByID)
	payment.GET("/total", handler.GetPayments)
	payment.POST("", handler.CreatePayment)
	payment.PUT("", handler.UpdatePayment)
	payment.DELETE("", handler.DeletePayment)
	payment.GET("/search", handler.SearchPayment)
}
