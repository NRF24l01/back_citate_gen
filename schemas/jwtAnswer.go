package schemas

type JwtAccessToken struct {
	AccessToken string `json:"access_token"`
	Message string `json:"message"`
}
