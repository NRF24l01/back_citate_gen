package schemas

type Message struct {
	Status string `json:"messsage"`
}

type JwtToken struct {
	Token string `json:"token"`
	Message string `json:"message"`
}