package admin

import "github.com/ermiasgashu/Construction-Machinary-Rental/entity"

type AdminRepository interface {
	Admins() ([]entity.Admin, error)
	Admin(id string) (entity.Admin, error)
	AddAdmin(admin entity.Admin) error
	UpdateAdmin(admin entity.Admin) error
	DeleteAdmin(id string) error
}
