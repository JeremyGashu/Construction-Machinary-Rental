package service

import (
	"github.com/ermiasgashu/Construction-Machinary-Rental/admin"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

// AdminServiceImpl implements menu.AdminService interface
type AdminServiceImpl struct {
	AdminRepo admin.AdminRepository
}

// NewAdminServiceImpl will create new AdminService object
func NewAdminServiceImpl(CatRepo admin.AdminRepository) *AdminServiceImpl {
	return &AdminServiceImpl{AdminRepo: CatRepo}
}

// Admins ..() returns list of Admins
func (cs *AdminServiceImpl) Admins() ([]entity.Admin, error) {

	Admins, err := cs.AdminRepo.Admins()

	if err != nil {
		return nil, err
	}

	return Admins, nil
}

// StoreAdmin persists new Admin information
func (cs *AdminServiceImpl) StoreAdmin(Admin entity.Admin) error {

	err := cs.AdminRepo.StoreAdmin(Admin)

	if err != nil {
		return err
	}

	return nil
}

// Admin returns a Admin object with a given id
func (cs *AdminServiceImpl) Admin(id string) (entity.Admin, error) {

	c, err := cs.AdminRepo.Admin(id)

	if err != nil {
		return c, err
	}

	return c, nil
}

// UpdateAdmin updates a cateogory with new data
func (cs *AdminServiceImpl) UpdateAdmin(Admin entity.Admin) error {

	err := cs.AdminRepo.UpdateAdmin(Admin)

	if err != nil {
		return err
	}

	return nil
}

// DeleteAdmin delete a Admin by its id
func (cs *AdminServiceImpl) DeleteAdmin(id string) error {

	err := cs.AdminRepo.DeleteAdmin(id)
	if err != nil {
		return err
	}
	return nil
}
