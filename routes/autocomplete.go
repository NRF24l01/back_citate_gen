package routes

import (
	"quoter_back/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterAutocompleteRoutes(e *echo.Echo, h *handlers.Handler) {
	group := e.Group("/autocomplete")

	group.GET("/tags", h.GetTags)
	group.GET("/authors", h.GetAuthors)
}