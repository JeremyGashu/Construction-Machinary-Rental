package service

import (
	"github.com/ermiasgashu/Construction-Machinary-Rental/admin"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

// UserServiceImpl implements menu.UserService interface
type UserServiceImpl struct {
	UserRepo admin.UserRepository
}

// NewUserServiceImpl will create new UserService object
func NewUserServiceImpl(CatRepo admin.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{UserRepo: CatRepo}
}

// Users ..() returns list of users
func (cs *UserServiceImpl) Users() ([]entity.User, error) {

	users, err := cs.UserRepo.Users()

	if err != nil {
		return nil, err
	}

	return users, nil
}

// StoreUser persists new User information
func (cs *UserServiceImpl) StoreUser(User entity.User) error {

	err := cs.UserRepo.StoreUser(User)

	if err != nil {
		return err
	}

	return nil
}

// User returns a User object with a given id
func (cs *UserServiceImpl) User(id string) (entity.User, error) {

	c, err := cs.UserRepo.User(id)

	if err != nil {
		return c, err
	}

	return c, nil
}

// UpdateUser updates a cateogory with new data
func (cs *UserServiceImpl) UpdateUser(User entity.User) error {

	err := cs.UserRepo.UpdateUser(User)

	if err != nil {
		return err
	}

	return nil
}

// DeleteUser delete a User by its id
func (cs *UserServiceImpl) DeleteUser(id string) error {

	err := cs.UserRepo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

//AuthUser -
func (cs *UserServiceImpl) AuthUser(username string, password string) bool {
	c := cs.UserRepo.AuthUser(username, password)
	return c
}
