package main

import (
	"quoter_back/schemas"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	validate := validator.New()
	schemas.RegisterCustomValidations(validate)

	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, schemas.Message{Status: "QUOTES!"})
	})
	
	e.Logger.Fatal(e.Start(":1323"))
}