package services

type LoginScheme struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type AuthService interface {
	Login(LoginScheme) (*LoginResponse, error)
}
