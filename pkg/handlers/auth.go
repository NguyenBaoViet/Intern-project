package handlers

import (
	"Intern-project/pkg/model"
	"Intern-project/pkg/services"
	"net/http"

	"gitlab.com/goxp/cloud0/ginext"
)

type AuthHandler struct {
	UserSrv IUser
	AuthSrv services.IAuthService
}

type IUser interface {
	CheckUserPassword(email string, password string) (*model.User, error)
}

func NewAuthHandler(usr IUser, auth services.IAuthService) *AuthHandler {
	return &AuthHandler{
		UserSrv: usr,
		AuthSrv: auth,
	}
}

func (auth *AuthHandler) Login(c *ginext.Request) (*ginext.Response, error) {
	//get request
	req := model.LoginRequest{}
	c.MustBind(&req)

	//
	user, err := auth.UserSrv.CheckUserPassword(req.Email, req.Password)
	if err != nil {
		return nil, ginext.NewError(http.StatusUnauthorized, err.Error())
	}

	//gen JWT token
	token, err := auth.AuthSrv.GenJWT(&user.ID)
	if err != nil {
		return nil, ginext.NewError(http.StatusInternalServerError, err.Error())
	}
	return ginext.NewResponseData(http.StatusOK, token), nil
}
