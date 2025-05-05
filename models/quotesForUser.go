package models

type QuoteForUser struct {
	BaseModel
	AskerID string `json:"creator_id" gorm:"not null"`
	Asker   User `json:"creator" gorm:"foreignKey:AskerID"`

	QuoteID string `json:"quote_id" gorm:"not null"`
	Quote   Quote `json:"quote" gorm:"foreignKey:QuoteID"`
}