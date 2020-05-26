package middelware

import (
	"net/http"

	"github.com/labstack/echo/middleware"
)

var corsConfig = middleware.CORSConfig{
	AllowOrigins: []string{"*"},
	//AllowOrigins: []string{"https://localhost:8080"},
	AllowMethods: []string{http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
	//AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	//ExposedHeaders:   []string{"Link"},
	AllowCredentials: true,
	MaxAge:           300,
}
