package admin

import "github.com/ermiasgashu/Construction-Machinary-Rental/entity"

//CompanyService ..
type CompanyService interface {
	Companies() ([]entity.Company, error)
	Company(id int) (entity.Company, error)
	UpdateCompany(Company entity.Company) error
	DeleteCompany(id int) error
	StoreCompany(Company entity.Company) error
	AuthCompany(email string, password string) bool
	UnactivatedCompanies() ([]entity.Company, error)
	ApproveCompany(id int) error
	CompanyByEmail(email string) (entity.Company, error)
	GetRentedMaterials(id int) ([]entity.RentInformation, error)
	DeleteMaterialsRented(companyid int, materiaid int, username string) bool
	// GetCompanyIDByEmail(email string) (int, error)
}

//AdminService ..
type AdminService interface {
	Admins() ([]entity.Admin, error)
	Admin(uname string) (entity.Admin, error)
	UpdateAdmin(Admin entity.Admin) error
	DeleteAdmin(uname string) error
	StoreAdmin(Admin entity.Admin) error
	AuthAdmin(username string, password string) bool
}

//UserService ..
type UserService interface {
	Users() ([]entity.User, error)
	User(uname string) (entity.User, error)
	UpdateUser(User entity.User) error
	DeleteUser(uname string) error
	StoreUser(User entity.User) error
	Pay(uname string, amount float64, cid int) bool
	GetRentedMaterials(uname string) ([]entity.RentInformation, error)
	AuthUser(username string, password string) bool
}

//CommentService ..
type CommentService interface {
	Comments() ([]entity.Comment, error)
	Comment(id int) (entity.Comment, error)
	UpdateComment(Comment entity.Comment) error
	DeleteComment(id int) error
	StoreComment(User entity.Comment) error
}
