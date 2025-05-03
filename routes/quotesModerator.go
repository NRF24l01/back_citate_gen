package routes

import (
	"quoter_back/handlers"
	"quoter_back/middleware"
	"quoter_back/schemas"

	"github.com/labstack/echo/v4"
)

func RegisterQuoteModerationRoutes(e *echo.Echo, h *handlers.Handler) {
	group := e.Group("/moderation")

	group.GET("/quotes", h.ModerationGet, middleware.JWTMiddleware("moderator"))

	group.POST("/review", h.ModerationReview, middleware.JWTMiddleware("moderator"), middleware.ValidationMiddleware(func() interface{} {
		return &schemas.QuoteReview{}
	}))
}