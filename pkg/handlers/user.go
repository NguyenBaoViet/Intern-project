package handlers

import (
	"Intern-project/pkg/model"
	"Intern-project/pkg/services"
	"net/http"

	"gitlab.com/goxp/cloud0/ginext"
)

type UserHandler struct {
	Service services.IUserService
}

type IUserHandler interface {
	SignUp(c *ginext.Request) (*ginext.Response, error)
}

func NewUserHandler(sv services.IUserService) IUserHandler {
	return &UserHandler{
		Service: sv,
	}
}

func (uh *UserHandler) SignUp(c *ginext.Request) (*ginext.Response, error) {
	//get request
	req := model.UserRequest{}
	c.MustBind(&req)

	//sign up upser
	rs, err := uh.Service.SignUp(req.Email, req.Password)
	if err != nil {
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}
	return ginext.NewResponseData(http.StatusOK, rs), nil
}
