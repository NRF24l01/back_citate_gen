package middleware

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

// CustomValidator wraps the go-playground validator
type CustomValidator struct {
    Validator *validator.Validate
}

// Validate implements the echo.Validator interface
func (cv *CustomValidator) Validate(i interface{}) error {
    return cv.Validator.Struct(i)
}

// ValidationMiddleware is a reusable middleware for validating request payloads
func ValidationMiddleware(schemaFactory func() interface{}) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            schema := schemaFactory()

            if err := c.Bind(schema); err != nil {
                return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
            }

            if err := c.Validate(schema); err != nil {
                return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
            }

            c.Set("validatedBody", schema)
            return next(c)
        }
    }
}