package schemas

type Message struct {
	Status string `json:"messsage"`
}

type ErrorMessage struct {
	Error string `json:"error"`
}

type SuccessCreateMessage struct {
	Status string `json:"message"`
	ID     string `json:"id"`
}