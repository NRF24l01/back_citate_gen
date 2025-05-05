package routes

import (
	"quoter_back/handlers"
	"quoter_back/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterQuotePersonalRoutes(e *echo.Echo, h *handlers.Handler) {
	group := e.Group("/quotes/personal")

	group.GET("", h.GetPersonalRandom, middleware.JWTMiddleware())
}