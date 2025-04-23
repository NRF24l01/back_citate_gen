package models

import (
	"gorm.io/datatypes"
)

type Quote struct {
	BaseModel
	AuthorName string `json:"author_name" gorm:"not null"`
	QuoteText  string `json:"quote_text" gorm:"not null"`
	Tags       datatypes.JSON `json:"tags" gorm:"type:jsonb;not null"`

	CreatorID string `json:"creator_id" gorm:"not null"`
	Creator   User `json:"creator" gorm:"foreignKey:CreatorID"`
}