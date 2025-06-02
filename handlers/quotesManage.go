package handlers

import (
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"quoter_back/models"
	"quoter_back/schemas"
	"quoter_back/utils"
)


func (h *Handler) QuoteCreate(c echo.Context) error {
	quote_data := c.Get("validatedBody").(*schemas.QuoteCreate)
	user_id := c.Get("user_id").(string)

	quote := models.Quote{
		AuthorName: quote_data.Author,
		QuoteText:  quote_data.Text,
		Tags:       datatypes.JSON(utils.ToJSON(quote_data.Tags)),
		CreatorID:  user_id,
	}

	if err := h.DB.Create(&quote).Error; err != nil {
		return c.JSON(500, schemas.ErrorMessage{Error: "An error occurred while creating the quote"})
	}
	return c.JSON(201, schemas.SuccessCreateMessage{Status: "Quote created successfully", ID: quote.ID.String()})
}
func (h *Handler) QuotesByUser(c echo.Context) error {
	user_id := c.Get("user_id").(string)

	var quotes []models.Quote
	if err := h.DB.Preload("Moderation", func(db *gorm.DB) *gorm.DB {
		return db.Order("moderations.updated_at DESC").Limit(1)
	}).Where("creator_id = ?", user_id).Find(&quotes).Error; err != nil {
		log.Printf("Error fetching quotes for user %s: %v", user_id, err)
		return c.JSON(500, schemas.ErrorMessage{Error: "An error occurred while fetching quotes"})
	}

	type ModerationResponse struct {
		ID               string `json:"id,omitempty"`
		Status           string `json:"status"`
		ModeratorComment string `json:"comment,omitempty"`
	}

	type QuoteResponse struct {
		ID         string              `json:"id"`
		Author     string              `json:"author"`
		Text       string              `json:"text"`
		Tags       []string            `json:"tags"`
		Moderation *ModerationResponse `json:"moderation"`
	}

	var response []QuoteResponse
	for _, quote := range quotes {
		var moderationResponse *ModerationResponse
		if len(quote.Moderation) > 0 {
			moderation := quote.Moderation[0]
			moderationResponse = &ModerationResponse{
				ID:               moderation.ID.String(),
				Status:           moderation.Status,
				ModeratorComment: moderation.ModeratorComment,
			}
		} else {
			moderationResponse = &ModerationResponse{
				Status: "pending",
			}
		}

		response = append(response, QuoteResponse{
			ID:         quote.ID.String(),
			Author:     quote.AuthorName,
			Text:       quote.QuoteText,
			Tags:       utils.FromJSONToStringArray(quote.Tags),
			Moderation: moderationResponse,
		})
	}

	return c.JSON(200, response)
}