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
