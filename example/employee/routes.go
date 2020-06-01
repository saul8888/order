package employee

import "github.com/labstack/echo"

// Register a new user
func (handler *Handler) Register(route *echo.Group) {

	employee := route.Group("/employee")
	employee.GET("", handler.GetEmployeeByID)
	employee.GET("/total", handler.GetEmployees)
	employee.POST("", handler.CreateEmployee)
	employee.PUT("", handler.UpdateEmployee)
	employee.DELETE("", handler.DeleteEmployee)
	employee.GET("/search", handler.SearchEmployee)
	employee.POST("/prueba", handler.PruebaEmployee)
}
