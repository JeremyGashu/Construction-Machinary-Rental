package repository

import (
	"database/sql"
	"errors"

	"github.com/ermiasgashu/Construction-Machinary-Rental/admin"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

//MockAdminRepo -
type MockAdminRepo struct {
	conn *sql.DB
}

//NewMockAdminRepo -
func NewMockAdminRepo(con *sql.DB) admin.AdminRepo {
	return &MockAdminRepo{conn: con}
}

//Admins -
func (mar *MockAdminRepo) Admins() ([]entity.Admin, error) {
	admins := []entity.Admin{entity.AdminMock}
	return admins, nil
}

//Admin -
func (mar *MockAdminRepo) Admin(uname string) (entity.Admin, error) {
	admin := entity.AdminMock
	if admin.Username == "Moke uName" {
		return admin, nil
	}
	return admin, errors.New("Cant Get Admin")
}

//UpdateAdmin -
func (mar *MockAdminRepo) UpdateAdmin(admin entity.Admin) error {
	// ad := entity.AdminMock
	return nil
}

//DeleteAdmin -
func (mar *MockAdminRepo) DeleteAdmin(uname string) error {
	admin := entity.AdminMock
	if admin.Username != "Moke uName" {
		return errors.New("Cant Delete Admin")
	}
	return nil
}

//StoreAdmin -
func (mar *MockAdminRepo) StoreAdmin(Admin entity.Admin) error {
	// adm := Admin
	return nil
}

//AuthAdmin -
func (mar *MockAdminRepo) AuthAdmin(username string, password string) bool {
	if username == "Moke uName" && password == "Moke password" {
		return true
	}
	return false
}
