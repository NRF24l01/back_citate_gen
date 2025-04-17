package routes

import (
	"quoter_back/handlers"
	"quoter_back/middleware"
	"quoter_back/schemas"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	group := e.Group("/auth")

	group.POST("/register", handlers.UserRegister, middleware.ValidationMiddleware(func() interface{} {
		return &schemas.RegisterUser{}
	}))
}
