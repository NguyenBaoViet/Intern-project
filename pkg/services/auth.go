package services

import (
	"Intern-project/conf"
	"Intern-project/pkg/model"
	"Intern-project/pkg/repo"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"gitlab.com/goxp/cloud0/ginext"
)

type AuthService struct {
	Repo repo.IRepo
}

type IAuthService interface {
	GenJWT(user_id *uuid.UUID) (string, error)
	CheckJWT(c *ginext.Request) (string, error)
	GetUserInfo(id string) (*model.User, error)
}

func NewAuthService(rp repo.IRepo) IAuthService {
	return &AuthService{
		Repo: rp,
	}
}

func (s *AuthService) GenJWT(user_id *uuid.UUID) (string, error) {
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

func (s *AuthService) CheckJWT(c *ginext.Request) (string, error) {
	//check token
	token := c.GinCtx.Request.Header.Get("Authorization")

	// check bearer
	tmp := strings.Split(token, " ")
	if tmp[0] != "Bearer" {
		return "", fmt.Errorf("no permission")
	}

	// get info from token
	token = tmp[1]
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(
		token,
		claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(conf.LoadEnv().SecretKey), nil
		},
	)
	if err != nil {
		return "", err
	}

	rs := claims["user_id"].(string)
	return rs, nil
}

func (s *AuthService) GetUserInfo(id string) (*model.User, error) {
	user, err := s.Repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
