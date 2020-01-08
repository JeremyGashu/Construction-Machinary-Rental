package company

import "github.com/ermiasgashu/Construction-Machinary-Rental/entity"

//Service -
type Service interface {
	Company(id int) (entity.Company, error)
	UpdateCompany(company entity.Company) error
	DeleteCompany(id int) error
	AddCompany(company entity.Company) error
	Companies() ([]entity.Company, error)
	//There are things that will be added on the flow
}
