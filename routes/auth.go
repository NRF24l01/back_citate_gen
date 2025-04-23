package routes

import (
	"quoter_back/handlers"
	"quoter_back/middleware"
	"quoter_back/schemas"

	"github.com/labstack/echo/v4"
)

func RegisterAuthRoutes(e *echo.Echo, h *handlers.Handler) {
	group := e.Group("/auth")

	group.POST("/register", h.UserRegister, middleware.ValidationMiddleware(func() interface{} {
		return &schemas.RegisterUser{}
	}))

	group.POST("/login", h.UserLogin, middleware.ValidationMiddleware(func() interface{} { 
		return &schemas.LoginUser{}
	}))

	group.POST("/refresh", h.TokenRefresh)
}