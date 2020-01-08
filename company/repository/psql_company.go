package repository

import (
	"database/sql"
	"errors"

	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

//CompanyRepoImpl -
type CompanyRepoImpl struct {
	conn *sql.DB
}

//NewCompanyRepo -
func NewCompanyRepo(c *sql.DB) *CompanyRepoImpl {
	return &CompanyRepoImpl{conn: c}
}

//Company -
func (cr *CompanyRepoImpl) Company(id int) (entity.Company, error) {
	//name, email, phone, address, password

	// query := "insert into companies(name, email, phone, address, password)"
	query := "select * from companies where id = $1"
	var company entity.Company
	err := cr.conn.QueryRow(query, id).Scan(&company.CompanyID, &company.Name, &company.Email, &company.Email, &company.PhoneNo, &company.Address, &company.Rating, &company.ImagePath, &company.Password, &company.Account)
	if err != nil {
		return company, err
	}
	return company, err
}

//Companies -
func (cr *CompanyRepoImpl) Companies() ([]entity.Company, error) {
	query := "select * from companies"
	companies := make([]entity.Company, 0)
	data, err := cr.conn.Query(query)
	if err != nil {
		return companies, err
	}
	for data.Next() {
		var co entity.Company
		data.Scan(&co.CompanyID, &co.Name, &co.Email, &co.Email, &co.PhoneNo, &co.Address, &co.Rating, &co.ImagePath, &co.Password, &co.Account)
		companies = append(companies, co)
	}
	err = data.Err()
	if err != nil {
		return companies, err
	}
	return companies, nil
}

//DeleteCompany -
func (cr *CompanyRepoImpl) DeleteCompany(id int) error {
	query := "delete from companies where id = $1"
	_, err := cr.conn.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

//UpdateCompany -
func (cr *CompanyRepoImpl) UpdateCompany(company entity.Company) error {
	query := "update companies set name=$1, email=$2, phone=$3,address=$4 where id = $1"
	_, err := cr.conn.Exec(query, company.Name, company.Email, company.PhoneNo, company.Address, company.CompanyID)
	if err != nil {
		return err
	}
	return nil
}

//Add company will be added when we finidh the w
func (cr *CompanyRepoImpl) AddCompany(company entity.Company) error {
	return errors.New("Error")
}
