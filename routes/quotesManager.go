package routes

import (
	"quoter_back/handlers"
	"quoter_back/middleware"
	"quoter_back/schemas"

	"github.com/labstack/echo/v4"
)

func RegisterQuoteManageRoutes(e *echo.Echo, h *handlers.Handler) {
	group := e.Group("/quotes")

	group.POST("", h.QuoteCreate, middleware.ValidationMiddleware(func() interface{} {
		return &schemas.QuoteCreate{}
	}), middleware.JWTMiddleware())
	
	group.GET("/user", h.QuotesByUser, middleware.JWTMiddleware())
}