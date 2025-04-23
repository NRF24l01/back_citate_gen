package handlers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"

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