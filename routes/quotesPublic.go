package routes

import (
	"quoter_back/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterQuotePublicRoutes(e *echo.Echo, h *handlers.Handler) {
	group := e.Group("/quotes/public")

	group.GET("", h.PublicGetList)
}