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
	SignUp(email, password string) (*model.User, error)
	GetUserInfo(id string) (*model.User, error)
	ChangePassword(id string, newPW string, oldPW string) error
	DeleteAccount(id string, PW string) error
}

func NewAuthHandler(usr IUser, auth services.IAuthService) *AuthHandler {
	return &AuthHandler{
		UserSrv: usr,
		AuthSrv: auth,
	}
}

// CREATE USER
func (auth *AuthHandler) SignUp(c *ginext.Request) (*ginext.Response, error) {
	//get request
	req := model.UserRequest{}
	c.MustBind(&req)

	//sign up upser
	_, err := auth.UserSrv.SignUp(req.Email, req.Password)
	if err != nil {
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}
	rs := "Sign up success"
	return ginext.NewResponseData(http.StatusOK, rs), nil
}

// GEN JWT - login
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

// GET USER
func (auth *AuthHandler) GetUserInfo(c *ginext.Request) (*ginext.Response, error) {
	user_id, err := auth.AuthSrv.CheckJWT(c)
	if err != nil {
		return nil, ginext.NewError(http.StatusInternalServerError, err.Error())
	}
	user, err := auth.UserSrv.GetUserInfo(user_id)
	if err != nil {
		return nil, ginext.NewError(http.StatusInternalServerError, err.Error())
	}
	return ginext.NewResponseData(http.StatusOK, user), nil
}

//  UPDATE USER - Change password
func (auth *AuthHandler) ChangePassword(c *ginext.Request) (*ginext.Response, error) {
	//check token
	userID, err := auth.AuthSrv.CheckJWT(c)
	if err != nil {
		return nil, ginext.NewError(http.StatusInternalServerError, err.Error())
	}

	//get request
	req := model.ChangePassword{}
	c.MustBind(&req)

	err = auth.UserSrv.ChangePassword(userID, req.NewPassword, req.Password)
	if err != nil {
		return nil, ginext.NewError(http.StatusInternalServerError, err.Error())
	}
	return ginext.NewResponseData(http.StatusOK, "Change Password Success"), nil
}

// DELETE USER
func (auth *AuthHandler) DeleteUser(c *ginext.Request) (*ginext.Response, error) {
	//check token
	userID, err := auth.AuthSrv.CheckJWT(c)
	if err != nil {
		return nil, ginext.NewError(http.StatusInternalServerError, err.Error())
	}

	//get request
	req := model.DeleteRequest{}
	c.MustBind(&req)

	err = auth.UserSrv.DeleteAccount(userID, req.Password)
	if err != nil {
		return nil, ginext.NewError(http.StatusInternalServerError, err.Error())
	}
	return ginext.NewResponseData(http.StatusOK, "Your account has been deleted"), nil

}
