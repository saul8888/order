package user

import "github.com/labstack/echo"

// Register a new user
func (handler *Handler) Register(route *echo.Group) {

	user := route.Group("/user")
	user.GET("", handler.GetUserByID)
	user.GET("/total", handler.GetUsers)
	user.POST("", handler.CreateUser)
	user.PUT("", handler.UpdateUser)
	user.DELETE("", handler.DeleteUser)
	user.GET("/search", handler.SearchUser)
}
