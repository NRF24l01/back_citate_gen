package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"quoter_back/models"
	"quoter_back/schemas"
)


func (h *Handler) GetAuthors(c echo.Context) error {
	var response []string

	var authors []string
	if err := h.DB.Model(&models.Quote{}).Distinct().Pluck("author_name", &authors).Error; err != nil {
		log.Println("Error fetching authors:", err)
		return c.JSON(http.StatusInternalServerError, schemas.ErrorMessage{Error: "Failed to fetch authors"})
	}
	response = authors

	return c.JSON(200, response)
}

func (h *Handler) GetTags(c echo.Context) error {
	var response []string

	rows, err := h.DB.Raw(`SELECT DISTINCT jsonb_array_elements_text(tags) AS tag FROM quotes`).Rows()
	if err != nil {
		log.Println("Error fetching tags:", err)
		return c.JSON(http.StatusInternalServerError, schemas.ErrorMessage{Error: "Failed to fetch tags"})
	}
	defer rows.Close()

	tagSet := make(map[string]struct{})
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			log.Println("Error scanning tag:", err)
			continue
		}
		tagSet[tag] = struct{}{}
	}

	for tag := range tagSet {
		response = append(response, tag)
	}

	return c.JSON(200, response)
}
