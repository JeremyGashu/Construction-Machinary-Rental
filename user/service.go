package user

import "github.com/ermiasgashu/Construction-Machinary-Rental/entity"

//Service -
type Service interface {
	User(id string) (entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUser(id string) error
	AddUser(user entity.User) error
	Users() ([]entity.User, error)
	//There are things that will be added on the flow
}
