package repo

import (
	"Intern-project/pkg/model"

	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

type IRepo interface {
	CheckEmailExist(email string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	GetUserByID(id string) (*model.User, error)
	UpdateUser(id string, newPW string) error
	DeleteUser(id string) error
}

func NewReop(db *gorm.DB) IRepo {
	return &Repo{
		DB: db,
	}
}

func (rp *Repo) CheckEmailExist(email string) (*model.User, error) {
	rs := &model.User{}

	if err := rp.DB.Where("email = ?", email).First(rs).Error; err != nil {
		return nil, err
	}
	return rs, nil
}

func (rp *Repo) CreateUser(user *model.User) (*model.User, error) {
	if err := rp.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (rp *Repo) GetUserByID(id string) (*model.User, error) {
	rs := &model.User{}
	if err := rp.DB.Where("id = ?", id).First(rs).Error; err != nil {
		return nil, err
	}
	return rs, nil
}

func (rp *Repo) UpdateUser(id string, newPW string) error {
	err := rp.DB.Table("user").Where("id = ?", id).Update("password", newPW).Error
	if err != nil {
		return err
	}
	return nil
}

func (rp *Repo) DeleteUser(id string) error {
	err := rp.DB.Table("user").Where("id = ?", id).Delete(&model.User{}).Error
	if err != nil {
		return err
	}
	return nil
}
