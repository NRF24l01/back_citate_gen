package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"log"

	"quoter_back/models"
	"quoter_back/schemas"
	"quoter_back/utils"
)

func (h *Handler) UserRegister(c echo.Context) error {
	user_data := c.Get("validatedBody").(*schemas.RegisterUser)
	log.Printf("Registering user: %+v", user_data)

	var test models.User
	email_result := h.DB.Where("email = ?", user_data.Email).First(&test)
	if email_result.Error != nil {
		if email_result.Error == gorm.ErrRecordNotFound {
			log.Println("Email is available.")
		} else {
			log.Fatalf("Error checking email: %v", email_result.Error)
			return c.JSON(500, schemas.ErrorMessage{ Error: "An error occurred while checking the email" })
		}
	} else {
		log.Printf("Such email already exists: %v", user_data.Email)
		return c.JSON(409, schemas.ErrorMessage{ Error: "A user with this email already exists" })
	}

	nick_result := h.DB.Where("username = ?", user_data.Username).First(&test)
	if nick_result.Error != nil {
		if nick_result.Error == gorm.ErrRecordNotFound {
			log.Println("Username is available.")
		} else {
			log.Fatalf("Error checking username: %v", nick_result.Error)
			return c.JSON(500, schemas.ErrorMessage{ Error: "An error occurred while checking the username" })
		}
	} else {
		log.Printf("Such username already exists: %v", user_data.Username)
		return c.JSON(409, schemas.ErrorMessage{ Error: "A user with this username already exists" })
	}

	hashedPassword := utils.HashPassword(user_data.Password)

	user := models.User{
		Email:    user_data.Email,
		Username: user_data.Username,
		Password: hashedPassword,
	}

	if err := h.DB.Create(&user).Error; err != nil {
		log.Fatalf("Error creating user: %v", err)
		return c.JSON(500, schemas.ErrorMessage{ Error: "An error occurred while creating the user" })
	}
	log.Printf("User created successfully: %+v", user)

	refreshToken, err := utils.GenerateRefreshToken(user.ID.String())
    if err != nil {
        log.Fatalf("Error generating refresh token: %v", err)
		return c.JSON(500, schemas.ErrorMessage{ Error: "An error occurred while generating refresh token" })
    }

	accessToken, err := utils.GenerateAccessToken(user.ID.String())
    if err != nil {
        log.Fatalf("Error generating access token: %v", err)
		return c.JSON(500, schemas.ErrorMessage{ Error: "An error occurred while generating access token" })
    }

	// Set the refresh token in an HttpOnly cookie
	cookie := new(http.Cookie)
	cookie.Name = "refresh_token"
	cookie.Value = refreshToken
	cookie.HttpOnly = true
	cookie.Path = "/"
	c.SetCookie(cookie)

	return c.JSON(201, schemas.JwtAccessToken{ AccessToken: accessToken, Message: "User registered successfully" })
}

func (h *Handler) UserLogin(c echo.Context) error {
	user_data := c.Get("validatedBody").(*schemas.LoginUser)
	log.Printf("Logging in user: %+v", user_data)

	var user models.User
	if err := h.DB.Where("email = ?", user_data.Email).First(&user).Error; err != nil {
		log.Printf("Error finding user: %v", err)
		return c.JSON(401, schemas.ErrorMessage{ Error: "Invalid email or password" })
	}

	if !utils.CheckPassword(user_data.Password, user.Password) {
		log.Println("Invalid password")
		return c.JSON(401, schemas.ErrorMessage{ Error: "Invalid email or password" })
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID.String())
	if err != nil {
		log.Fatalf("Error generating refresh token: %v", err)
		return c.JSON(500, schemas.ErrorMessage{ Error: "An error occurred while generating refresh token" })
	}

	accessToken, err := utils.GenerateAccessToken(user.ID.String())
	if err != nil {
		log.Fatalf("Error generating access token: %v", err)
		return c.JSON(500, schemas.ErrorMessage{ Error: "An error occurred while generating access token" })
	}

	// Set the refresh token in an HttpOnly cookie
	cookie := new(http.Cookie)
	cookie.Name = "refresh_token"
	cookie.Value = refreshToken
	cookie.HttpOnly = true
	cookie.Path = "/"
	c.SetCookie(cookie)

	return c.JSON(200, schemas.JwtAccessToken{ AccessToken: accessToken, Message: "User logged in successfully" })
}