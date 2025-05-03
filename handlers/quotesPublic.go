package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"quoter_back/schemas"
)


func (h *Handler) PublicGetList(c echo.Context) error {
	var quotes []struct {
		ID     string `json:"id"`
		Text   string `json:"text"`
		Author string `json:"author"`
	}

	if err := h.DB.Table("quotes").
		Select("quotes.id, quotes.quote_text AS text, quotes.author_name AS author").
		Joins("JOIN moderations ON quotes.id = moderations.quote_id").
		Where("moderations.status = ?", "approved").
		Scan(&quotes).Error; err != nil {
		log.Printf("Error fetching approved quotes: %v", err)
		return c.JSON(http.StatusInternalServerError, schemas.ErrorMessage{Error: "An error occurred while fetching quotes"})
	}

	return c.JSON(http.StatusOK, quotes)
}