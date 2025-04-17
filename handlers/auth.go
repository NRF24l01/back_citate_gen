package handlers

import (
	"github.com/labstack/echo/v4"

	"log"
	"quoter_back/schemas"
)

func UserRegister(c echo.Context) error {
	user := c.Get("validatedBody").(*schemas.RegisterUser)
	log.Printf("Registering user: %+v", user)
	return c.JSON(201, schemas.JwtToken{ Token: "fig tebe a ne token", Message: "User registered successfully" })
}