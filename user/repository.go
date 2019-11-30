package user

import "github.com/ermiasgashu/Construction-Machinary-Rental/entity"

type UserRepository interface {
	User(id string) (entity.User, error)
	UpdateUser(entity.User) error
	DeleteUser(entity.User) error
	AddUser(entity.User) error
	//There are things that will be added on the flow
}
