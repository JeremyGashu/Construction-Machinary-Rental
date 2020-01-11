package service

import (
	"github.com/ermiasgashu/Construction-Machinary-Rental/company"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

//CompanyService -
type CompanyService struct {
	repository company.CompanyRepository
}

//NewCompanyService -
func NewCompanyService(repo company.CompanyService) *CompanyService {
	return &CompanyService{repository: repo}
}

//Companies -
func (cs *CompanyService) Companies() ([]entity.Company, error) {
	companies, err := cs.repository.Companies()
	if err != nil {
		return companies, err
	}
	return companies, nil
}

//Company -
func (cs *CompanyService) Company(id int) (entity.Company, error) {
	company, err := cs.repository.Company(id)
	if err != nil {
		return company, err
	}
	return company, nil
}

//UpdateCompany -
func (cs *CompanyService) UpdateCompany(company entity.Company) error {
	err := cs.repository.UpdateCompany(company)
	if err != nil {
		return err
	}
	return nil
}

//DeleteCompany -
func (cs *CompanyService) DeleteCompany(id int) error {
	err := cs.repository.DeleteCompany(id)
	if err != nil {
		return err
	}
	return nil
}

//AddCompany -
func (cs *CompanyService) AddCompany(company entity.Company) error {
	err := cs.repository.AddCompany(company)
	if err != nil {
		return err
	}
	return nil
}

