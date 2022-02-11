package services

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (srv *AuthService) Login(username string, password string) bool {
	return username == "test" && password == "password"
}
