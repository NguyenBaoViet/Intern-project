package services

import (
	"Intern-project/pkg/model"
	"Intern-project/pkg/repo"
	"fmt"

	"Intern-project/pkg/utils"
)

type UserService struct {
	Repo repo.IRepo
}

type IUserService interface {
	SignUp(email, password string) (*model.User, error)
	CheckUserPassword(email string, password string) (*model.User, error)
	GetUserInfo(id string) (*model.User, error)
	ChangePassword(id string, newPW string, oldPW string) error
	DeleteAccount(id string, PW string) error
}

func NewUserService(rp repo.IRepo) IUserService {
	return &UserService{
		Repo: rp,
	}
}

func (us *UserService) SignUp(email, password string) (*model.User, error) {
	//go to db to check email exist or not

	user, err := us.Repo.CheckEmailExist(email)
	if err != nil && !utils.IsErrNotFound(err) {
		//database err
		return nil, err
	}

	if user != nil {
		//user existed
		//return err
		return nil, fmt.Errorf("email existed")
	}

	//hashpassword
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}
	//create user
	newUser := &model.User{
		Email:    email,
		Password: hashedPassword,
	}

	user, err = us.Repo.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) CheckUserPassword(email string, password string) (*model.User, error) {
	//get user by email
	user, err := us.Repo.CheckEmailExist(email)
	if err != nil {
		return nil, err
	}

	// check password
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, fmt.Errorf("wrong password")
	}
	return user, nil
}

func (us *UserService) GetUserInfo(id string) (*model.User, error) {
	user, err := us.Repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) ChangePassword(id string, newPW string, oldPW string) error {
	user, err := us.Repo.GetUserByID(id)
	if err != nil {
		return err
	}

	if utils.CheckPasswordHash(oldPW, user.Password) {
		PW, err := utils.HashPassword(newPW)
		if err != nil {
			return err
		}
		err = us.Repo.UpdateUser(id, PW)
		if err != nil {
			return err
		}
	}

	return nil
}

func (us *UserService) DeleteAccount(id string, PW string) error {
	user, err := us.Repo.GetUserByID(id)
	if err != nil {
		return err
	}
	if utils.CheckPasswordHash(PW, user.Password) {
		err = us.Repo.DeleteUser(id)
		if err != nil {
			return err
		}
	}
	return nil
}
