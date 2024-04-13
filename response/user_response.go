package response

type Login struct {
	Token               string `json:"token"`
	TokenExpirationTime string `json:"token_expiration_time"`
	RefreshToken        string `json:"refresh_token"`
	IsVerified          bool   `json:"is_verified"`
}
