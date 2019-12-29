package repository

import (
	"database/sql"
	"errors"

	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

//AdminRepo -
type AdminRepo struct {
	conn *sql.DB
}

//NewAdminRepo -
func NewAdminRepo(Conn *sql.DB) *AdminRepo {
	return &AdminRepo{conn: Conn}
}

//Admins -
func (ar *AdminRepo) Admins() ([]entity.Admin, error) {
	admins := make([]entity.Admin, 0)
	query := "select firstname, lastname,username,email from admins"
	data, err := ar.conn.Query(query)
	if err != nil {
		return admins, errors.New("No user is found")
	}
	for data.Next() {
		var admin entity.Admin
		data.Scan(&admin.FirstName, &admin.LastName, &admin.Username, &admin.Email) //all the datas that will be added in the category
		admins = append(admins, admin)
	}
	if err := data.Err(); err != nil {
		return admins, errors.New("Some error is occured")
	}
	return admins, nil
}

//Admin -
func (ar *AdminRepo) Admin(username string) (entity.Admin, error) {
	var admin entity.Admin
	query := "select firstname, lastname,username,email from admins where username=$1"
	err := ar.conn.QueryRow(query, username).Scan(&admin.FirstName, &admin.LastName, &admin.Username, &admin.Email)
	if err != nil {
		return admin, err
	}
	return admin, nil
}

//AddAdmin -
func (ar *AdminRepo) AddAdmin(admin entity.Admin) error {
	query := "insert into admins(firstname, lastname, username, email, password) values($1,$2,$3,$4,$5)"
	_, err := ar.conn.Exec(query, admin.FirstName, admin.LastName, admin.Username, admin.Email, admin.Password)
	if err != nil {
		return err
	}
	return nil
}

//UpdateAdmin -
func (ar *AdminRepo) UpdateAdmin(admin entity.Admin) error {
	query := "update admins set firstname=$1,lastname=$2,email=$3,password=$4"
	_, err := ar.conn.Exec(query, admin.FirstName, admin.LastName, admin.Email, admin.Password)
	if err != nil {
		return err
	}
	return nil
}

//DeleteAdmin -
func (ar *AdminRepo) DeleteAdmin(id string) error {
	query := "delete from admins where username=$1"
	_, err := ar.conn.Exec(query, id)
	if err != nil {
		return err
	}
	return err
}
