package schemas

type JwtAccessToken struct {
	AccessToken string `json:"access_token"`
	Message string `json:"message"`
}

type JwtTokenPair struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Message string `json:"message"`
}