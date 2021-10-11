package services

type AuthService struct {
}

type IAuthService interface {
}

func NewAuthService() IAuthService {
	return &AuthService{}
}
