package schemas

type RegisterUser struct {
	Email string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=32,strongpwd"`
}

type LoginUser struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32,strongpwd"`
}