package service

import (
	"github.com/ermiasgashu/Construction-Machinary-Rental/company"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

//CompanyService -
type CompanyService struct {
	repo company.Repository
}

//NewCompanyService -
func NewCompanyService(r company.Repository) *CompanyService {
	return &CompanyService{repo: r}
}

//Company -
func (cs *CompanyService) Company(id int) (entity.Company, error) {
	company, err := cs.repo.Company(id)
	if err != nil {
		return company, err
	}
	return company, nil
}

//Companies -
func (cs *CompanyService) Companies() ([]entity.Company, error) {
	companies, err := cs.repo.Companies()
	if err != nil {
		return companies, err
	}
	return companies, nil
}

//DeleteCompany -
func (cs *CompanyService) DeleteCompany(id int) error {
	err := cs.repo.DeleteCompany(id)
	if err != nil {
		return err
	}
	return nil
}

//UpdateCompany -
func (cs *CompanyService) UpdateCompany(company entity.Company) error {
	err := cs.repo.UpdateCompany(company)
	if err != nil {
		return err
	}
	return nil
}
func (cs *CompanyService) AddCompany(company entity.Company) error {
	err := cs.repo.AddCompany(company)
	if err != nil {
		return err
	}
	return nil
}

//TODO add add company function
