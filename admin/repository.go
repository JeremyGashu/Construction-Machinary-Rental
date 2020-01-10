package admin

import "github.com/ermiasgashu/Construction-Machinary-Rental/entity"

//CompanyRepository ..
type CompanyRepository interface {
	Companies() ([]entity.Company, error)
	Company(id int) (entity.Company, error)
	UpdateCompany(Company entity.Company) error
	DeleteCompany(id int) error
	StoreCompany(Company entity.Company) error
}

//AdminRepository ..
type AdminRepository interface {
	Admins() ([]entity.Admin, error)
	Admin(uname string) (entity.Admin, error)
	UpdateAdmin(Admin entity.Admin) error
	DeleteAdmin(uname string) error
	StoreAdmin(Admin entity.Admin) error
}

//UserRepository ..
type UserRepository interface {
	Users() ([]entity.User, error)
	User(uname string) (entity.User, error)
	UpdateUser(User entity.User) error
	DeleteUser(uname string) error
	StoreUser(User entity.User) error
}

//CommentRepository ..
type CommentRepository interface {
	Comments() ([]entity.Comment, error)
	Comment(id int) (entity.Comment, error)
	UpdateComment(Comment entity.Comment) error
	DeleteComment(id int) error
	StoreComment(User entity.Comment) error
}
