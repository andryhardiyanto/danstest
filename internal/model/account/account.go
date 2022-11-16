package account

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	AccessExp    int64  `json:"access_exp"`
	RefreshToken string `json:"refresh_token"`
	RefreshExp   int64  `json:"refresh_exp"`
}
