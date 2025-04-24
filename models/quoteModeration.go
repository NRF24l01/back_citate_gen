package models

type Moderation struct {
    BaseModel
    QuoteID         string `json:"quote_id" gorm:"not null"` // ID цитаты
    Quote           Quote  `json:"quote" gorm:"foreignKey:QuoteID"`
    Status          string `json:"status" gorm:"not null"` // Статус модерации (e.g., "pending", "approved", "rejected")
    ModeratorID     string `json:"moderator_id"`           // ID модератора
    Moderator       User   `json:"moderator" gorm:"foreignKey:ModeratorID"`
    ModeratorComment string `json:"moderator_comment"`     // Комментарий модератора
}