package service

import (
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
	"github.com/ermiasgashu/Construction-Machinary-Rental/user"
)

type UserServiceImpl struct {
	repo user.Repository
}

func NewUserServiceImpl(repo user.Repository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

// User(id string) (entity.User, error)
// 	UpdateUser(entity.User) error
// 	DeleteUser(entity.User) error
// 	AddUser(entity.User) error
func (us *UserServiceImpl) User(username string) (entity.User, error) {
	user, err := us.repo.User(username)
	if err != nil {
		return user, err
	}
	return user, nil
}
