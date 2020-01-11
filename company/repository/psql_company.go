package repository

import (
	"database/sql"
	"errors"

	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
	
)

//CompanyRepository -
type CompanyRepository struct {
	conn *sql.DB
}

//NewCompanyRepository -
func NewCompanyRepository(Conn *sql.DB) *CompanyRepository {
	return &CompanyRepository{conn: Conn}
}

//Companies -
func (cr *CompanyRepository) Companies() ([]entity.Company, error) {
	companies := make([]entity.Company, 0)
	query := "select * from companies"
	data, err := cr.conn.Query(query)
	if err != nil {
		return companies, errors.New("No company is found")
	}
	for data.Next() {
		var company entity.Company
		data.Scan(&company.CompanyID, &company.Name, &company.Email, &company.PhoneNo, &company.Address, &company.Description, &company.Rating, &company.ImagePath, &company.Password, &company.Account, &company.Activated) //all the datas that will be added in the category
		companies = append(companies, company)
		
	}
	if err := data.Err(); err != nil {
		return companies, errors.New("Some error is occured")
	}
	return companies, nil
}

//Company -
func (cr *CompanyRepository) Company(id int) (entity.Company, error) {
	company := entity.Company{}
	query := "select * from companies where id=$1"
	err := cr.conn.QueryRow(query, id).Scan(&company.CompanyID, &company.Name, &company.Email, &company.PhoneNo, &company.Address, &company.Description, &company.Rating, &company.ImagePath, &company.Password, &company.Account, &company.Activated)
	if err != nil {
		return company, err
	}
	return company, nil
}

//UpdateCompany -
func (cr *CompanyRepository) UpdateCompany(company entity.Company) error {
	query := "update companies set name=$1,email=$2,phone=$3,address=$4,discription=$5,rating=$6,imagepath=$7,password=$5,account=$6,activated=$7 where id=$8"
	_, err := cr.conn.Exec(query, company.Name, company.Email, company.PhoneNo, company.Address, company.Description, company.Rating, company.ImagePath, company.Password, company.Account, company.Activated, company.CompanyID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteCompany -
func (cr *CompanyRepository) DeleteCompany(id int) error {
	query := "delete from companies where id=$1"
	_, err := cr.conn.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

//AddCompany -
func (cr *CompanyRepository) AddCompany(company entity.Company) error {
	query := "insert into companies(name, email,phone, address, description, rating, imagepath, password, account, activated) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)"

	_, err := cr.conn.Exec(query, company.Name, company.Email, company.PhoneNo, company.Address, company.Description, company.Rating, company.ImagePath, company.Password, company.Account, company.Activated )
	if err != nil {
		return err
	}
	return nil
}
