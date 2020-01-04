package service

import (
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
	"github.com/ermiasgashu/Construction-Machinary-Rental/user"
)

//UserServiceImpl -
type UserServiceImpl struct {
	repo user.Repository
}

//NewUserServiceImpl -
func NewUserServiceImpl(repo user.Repository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

//User -
func (us *UserServiceImpl) User(username string) (entity.User, error) {
	user, err := us.repo.User(username)
	if err != nil {
		return user, err
	}
	return user, nil
}

//UpdateUser -
func (us *UserServiceImpl) UpdateUser(user entity.User) error {
	err := us.repo.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser -
func (us *UserServiceImpl) DeleteUser(id string) error {
	err := us.repo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

//AddUser -
func (us *UserServiceImpl) AddUser(user entity.User) error {
	err := us.repo.AddUser(user)
	if err != nil {
		return err
	}

	return nil
}

//Users -
func (us *UserServiceImpl) Users() ([]entity.User, error) {
	users, err := us.repo.Users()
	if err != nil {
		return users, err
	}
	return users, nil
}
