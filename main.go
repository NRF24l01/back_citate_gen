package main

import (
	"quoter_back/handlers"
	"quoter_back/middleware"
	"quoter_back/models"
	"quoter_back/routes"
	"quoter_back/schemas"

	"github.com/go-playground/validator"

	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	echoMw "github.com/labstack/echo/v4/middleware"
)

func main() {
	if os.Getenv("RUNTIME_PRODUCTION") != "true" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("failed to load .env: %v", err)
		}
	}

	db := models.RegisterPostgres()

	validater := validator.New()
	schemas.RegisterCustomValidations(validater)

	e := echo.New()

	e.Validator = &middleware.CustomValidator{Validator: validater}

	e.Use(echoMw.Logger())
    e.Use(echoMw.Recover())

	if os.Getenv("RUNTIME_PRODUCTION") != "true" {
		// Add CORS middleware
		e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
			AllowOrigins: []string{"http://localhost:5173"}, // Replace with your frontend's origin
			AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
			AllowCredentials: true,
		}))
	} else {
		// Add CORS middleware
		e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
			AllowOrigins: []string{"https://quoter.snnlab.ru/"}, // Replace with your frontend's origin
			AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
			AllowCredentials: true,
		}))
	}

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, schemas.Message{Status: "QUOTES!"})
	})

	handler := &handlers.Handler{DB: db}
	routes.RegisterRoutes(e, handler)
	
	e.Logger.Fatal(e.Start(":1323"))
}