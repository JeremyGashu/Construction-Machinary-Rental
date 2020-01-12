package repository

import (
	"database/sql"
	"errors"

	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

// CompanyRepositoryImpl implements the menu.CompanyRepository interface
type CompanyRepositoryImpl struct {
	conn *sql.DB
}

// NewCompanyRepositoryImpl will create an object of PsqlCompanyRepository
func NewCompanyRepositoryImpl(Conn *sql.DB) *CompanyRepositoryImpl {
	return &CompanyRepositoryImpl{conn: Conn}
}

// Companies returns all cateogories from the database
func (cri *CompanyRepositoryImpl) Companies() ([]entity.Company, error) {

	rows, err := cri.conn.Query("SELECT * FROM companies where activated=true")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	ctgs := []entity.Company{}

	for rows.Next() {
		Company := entity.Company{}
		err = rows.Scan(&Company.CompanyID, &Company.Name, &Company.Email, &Company.Address, &Company.PhoneNo, &Company.Description, &Company.Password, &Company.ImagePath, &Company.Rating, &Company.Account, &Company.Activated)
		if err != nil {
			return nil, err
		}
		ctgs = append(ctgs, Company)
	}

	return ctgs, nil
}

// Company returns a Company with a given id
func (cri *CompanyRepositoryImpl) Company(id int) (entity.Company, error) {

	row := cri.conn.QueryRow("SELECT * FROM companies WHERE id = $1", id)

	Company := entity.Company{}

	err := row.Scan(&Company.CompanyID, &Company.Name, &Company.Email, &Company.Address, &Company.PhoneNo, &Company.Description, &Company.Password, &Company.ImagePath, &Company.Rating, &Company.Account, &Company.Activated)
	if err != nil {
		return Company, err
	}

	return Company, nil
}

// UpdateCompany updates a given object with a new data
func (cri *CompanyRepositoryImpl) UpdateCompany(c entity.Company) error {

	_, err := cri.conn.Exec("UPDATE companies SET name=$1,description=$2, imagepath=$3,email=$4,phone=$5,address=$6,rating=$7,password=$8,activated=$9 WHERE id=$10", c.Name, c.Description, c.ImagePath, c.Email, c.PhoneNo, c.Address, c.Rating, c.Password, c.Activated, c.CompanyID)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}

// DeleteCompany removes a Company from a database by its id
func (cri *CompanyRepositoryImpl) DeleteCompany(id int) error {

	_, err := cri.conn.Exec("DELETE FROM companies WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

// StoreCompany stores new Company information to database
func (cri *CompanyRepositoryImpl) StoreCompany(c entity.Company) error {

	_, err := cri.conn.Exec("INSERT INTO companies (name,email,phone,address,description,imagepath,password) values($1, $2, $3,$4,$5,$6,$7)", c.Name, c.Email, c.PhoneNo, c.Address, c.Description, c.ImagePath, c.Password)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}

//AuthCompany -
func (cri *CompanyRepositoryImpl) AuthCompany(email string, password string) bool {
	query := "select name from companies where email=$1 and password=$2"
	var name string
	row := cri.conn.QueryRow(query, email, password)
	err := row.Scan(&name)

	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
	}
	return true
}

//UnactivatedCompanies -
func (cri *CompanyRepositoryImpl) UnactivatedCompanies() ([]entity.Company, error) {
	rows, err := cri.conn.Query("SELECT * FROM companies where activated=false")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	ctgs := []entity.Company{}

	for rows.Next() {
		Company := entity.Company{}
		err = rows.Scan(&Company.CompanyID, &Company.Name, &Company.Email, &Company.Address, &Company.PhoneNo, &Company.Description, &Company.Password, &Company.ImagePath, &Company.Rating, &Company.Account, &Company.Activated)
		if err != nil {
			return nil, err
		}
		ctgs = append(ctgs, Company)
	}

	return ctgs, nil
}

//ApproveCompany -
func (cri *CompanyRepositoryImpl) ApproveCompany(id int) error {
	query := "update companies set activated=true where id=$1"
	_, err := cri.conn.Exec(query, id)
	if err != nil {
		return err
	}
	return nil

}

// func (cri *CompanyRepositoryImpl) GetCompanyIDByEmail(email string) (int, error){
// 	query := "select"
// }
