package services

import (
	"Intern-project/conf"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type AuthService struct {
}

type IAuthService interface {
	GenJWT(user_id *uuid.UUID) (string, error)
}

func NewAuthService() IAuthService {
	return &AuthService{}
}

func (s *AuthService) GenJWT(user_id *uuid.UUID) (string, error) {
	fmt.Println("vao genjwt")
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	claims["iss"] = "TokenWithUserID"
	claims["user_id"] = user_id.String()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(conf.LoadEnv().SecretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
