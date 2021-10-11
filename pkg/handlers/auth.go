package handlers

import (
	"Intern-project/pkg/model"

	"gitlab.com/goxp/cloud0/ginext"
)

type AuthHandler struct {
	UserSrv IUser
}

type IUser interface {
	CheckUserPassword(email string, password string) (*model.User, error)
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (auth *AuthHandler) Login(c *ginext.Request) (*ginext.Response, error) {
	//get request
	req := model.LoginRequest{}
	c.MustBind(&req)

	//gen JWT token

	return nil, nil
}
