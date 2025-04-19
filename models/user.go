package models

type User struct {
	BaseModel
	Email        string `json:"email" gorm:"uniqueIndex;not null"`
	Password     string `json:"password" gorm:"not null"`
	Username     string `json:"username" gorm:"uniqueIndex;not null"`
	RefreshToken string `json:"refresh_token"`
}