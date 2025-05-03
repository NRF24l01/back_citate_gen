package handlers

import (
	"log"
	"net/http"

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

func (h *Handler) ModerationReview(c echo.Context) error {
	moderation_review_data := c.Get("validatedBody").(*schemas.QuoteReview)
	user_id := c.Get("user_id").(string)

	// Get user and check permissions
	var user models.User
	if err := h.DB.Where("id = ?", user_id).First(&user).Error; err != nil {
		log.Printf("User with ID %s not found: %v", user_id, err)
		return c.JSON(http.StatusUnauthorized, schemas.ErrorMessage{Error: "Unauthorized"})
	}

	if user.Role != "moderator" {
		log.Printf("User with ID %s does not have sufficient permissions", user_id)
		return c.JSON(http.StatusForbidden, schemas.ErrorMessage{Error: "Insufficient permissions"})
	}

	// Get quote
	quote_id := moderation_review_data.QuoteID
	var quote models.Quote
	if err := h.DB.Where("id = ?", quote_id).First(&quote).Error; err != nil {
		log.Printf("Quote with ID %s not found: %v", quote_id, err)
		return c.JSON(422, schemas.ErrorMessage{Error: "No such quote"})
	}

	// Check if moderation already exists
	var existingModeration models.Moderation
	if err := h.DB.Where("quote_id = ?", quote_id).First(&existingModeration).Error; err == nil {
		log.Printf("Moderation for quote with ID %s already exists: %v", quote_id, err)
		return c.JSON(422, schemas.ErrorMessage{Error: "Moderation already exists"})
	}

	moderation := models.Moderation{
		Quote:            quote,
		Status:           moderation_review_data.Status,
		Moderator:        user,
		ModeratorComment: moderation_review_data.Comment,
	}

	if err := h.DB.Create(&moderation).Error; err != nil {
		log.Printf("Error creating moderation for quote with ID %s: %v", quote_id, err)
		return c.JSON(500, schemas.ErrorMessage{Error: "An error occurred while creating moderation"})
	}
	return c.JSON(200, schemas.SuccessCreateMessage{Status: "Moderation created successfully", ID: moderation.ID.String()})
}