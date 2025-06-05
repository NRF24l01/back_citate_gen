package handlers

import (
	"log"
	"net/http"

	"gorm.io/gorm"

	"github.com/labstack/echo/v4"

	"quoter_back/models"
	"quoter_back/schemas"
)

func (h *Handler) GetPersonalRandom(c echo.Context) error {
	userID := c.Get("user_id").(string)

	var quote struct {
		ID     string   `json:"id"`
		Text   string   `json:"text"`
		Author string   `json:"author"`
		Tags   []string `json:"tags"`
	}

	if err := h.DB.Table("quotes").
		Select("quotes.id, quotes.quote_text AS text, quotes.author_name AS author, COALESCE(quotes.tags, '[]') AS tags").
		Joins("JOIN moderations ON quotes.id = moderations.quote_id").
		Where("moderations.status = ?", "approved").
		Where("quotes.id NOT IN (?)", h.DB.Table("quote_for_users").Select("quote_id").Where("asker_id = ?", userID)).
		Order("RANDOM()").
		Limit(1).
		Scan(&quote).Error; err != nil {
		if err == gorm.ErrRecordNotFound || quote.ID == "" {
			return c.JSON(http.StatusNotFound, schemas.ErrorMessage{Error: "No quotes found"})
		}
		log.Printf("Error fetching random unseen quote: %v", err)
		return c.JSON(http.StatusInternalServerError, schemas.ErrorMessage{Error: "An error occurred while fetching a random quote"})
	}

	// Record that the user has seen this quote
	if quote.ID == "" {
		log.Printf("Error: Quote ID is empty, cannot record quote for user")
		return c.JSON(http.StatusNotFound, schemas.ErrorMessage{Error: "No quotes more for this user"})
	}

	quoteForUser := models.QuoteForUser{
		AskerID: userID,
		QuoteID: quote.ID,
	}
	if err := h.DB.Create(&quoteForUser).Error; err != nil {
		log.Printf("Error recording quote for user: %v", err)
		return c.JSON(http.StatusInternalServerError, schemas.ErrorMessage{Error: "An error occurred while recording the quote"})
	}

	return c.JSON(http.StatusOK, quote)
}

func (h *Handler) GetPersonalQuotesHistory(c echo.Context) error {
	userID := c.Get("user_id").(string)

	var quotes []struct {
		ID     string   `json:"id"`
		Text   string   `json:"text"`
		Author string   `json:"author"`
		Tags   []string `json:"tags"`
	}

	if err := h.DB.Table("quotes").
		Select("quotes.id, quotes.quote_text AS text, quotes.author_name AS author, quotes.tags").
		Joins("JOIN quote_for_users ON quotes.id = quote_for_users.quote_id").
		Where("quote_for_users.asker_id = ?", userID).
		Order("quote_for_users.created_at DESC").
		Scan(&quotes).Error; err != nil {
		log.Printf("Error fetching user's quote history: %v", err)
		return c.JSON(http.StatusInternalServerError, schemas.ErrorMessage{Error: "An error occurred while fetching the quote history"})
	}
	
	if quotes == nil {
		quotes = make([]struct {
			ID     string   `json:"id"`
			Text   string   `json:"text"`
			Author string   `json:"author"`
			Tags   []string `json:"tags"`
		}, 0)
	}

	return c.JSON(http.StatusOK, quotes)
}