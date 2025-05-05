package routes

import (
	"quoter_back/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, h *handlers.Handler) {
	RegisterAuthRoutes(e, h)
	RegisterQuoteManageRoutes(e, h)
	RegisterQuoteModerationRoutes(e, h)
	RegisterQuotePublicRoutes(e, h)
	RegisterQuotePersonalRoutes(e, h)
}