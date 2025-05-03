package handlers

import (
	"log"

	"github.com/labstack/echo/v4"

	"quoter_back/models"
	"quoter_back/schemas"
	"quoter_back/utils"
)


func (h *Handler) ModerationGet(c echo.Context) error {
	var quotes []models.Quote
	if err := h.DB.Preload("Moderation").
		Joins("LEFT JOIN moderations ON quotes.id = moderations.quote_id AND moderations.status = 'approved'").
		Where("moderations.id IS NULL").
		Find(&quotes).Error; err != nil {
		log.Printf("Error fetching unmoderated quotes: %v", err)
		return c.JSON(500, schemas.ErrorMessage{Error: "An error occurred while fetching quotes"})
	}

	type QuoteResponse struct {
		ID     string   `json:"id"`
		Author string   `json:"author"`
		Text   string   `json:"text"`
		Tags   []string `json:"tags"`
	}

	var response []QuoteResponse
	for _, quote := range quotes {
		response = append(response, QuoteResponse{
			ID:     quote.ID.String(),
			Author: quote.AuthorName,
			Text:   quote.QuoteText,
			Tags:   utils.FromJSONToStringArray(quote.Tags),
		})
	}

	return c.JSON(200, response)
}