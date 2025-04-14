package models

type User struct {
	BaseModel
	Email string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"password"`
	Username string `json:"username"`
}