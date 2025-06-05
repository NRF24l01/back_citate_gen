package routes

import (
	"quoter_back/handlers"
	"quoter_back/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterAutocompleteRoutes(e *echo.Echo, h *handlers.Handler) {
	group := e.Group("/autocomplete")

	group.GET("/tags", h.GetTags, middleware.JWTMiddleware())
	group.GET("/authors", h.GetAuthors, middleware.JWTMiddleware())
}