package schemas

type QuoteReview struct {
	QuoteID string `json:"id" validate:"required"`
	Status  string `json:"result" validate:"required,oneof=approved rejected"`
	Comment string `json:"comment" validate:"max=256"`
}