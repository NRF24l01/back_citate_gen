package schemas

type QuoteCreate struct {
	Text   string   `json:"text" validate:"required,min=3,max=256"`
	Author string   `json:"author" validate:"required,min=3,max=64"`
	Tags   []string `json:"tags" validate:"dive,min=1,max=30"`
}