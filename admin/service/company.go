package service

import (
	"github.com/ermiasgashu/Construction-Machinary-Rental/admin"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

// CompanyServiceImpl implements menu.CompanyService interface
type CompanyServiceImpl struct {
	CompanyRepo admin.CompanyRepository
}

// NewCompanyServiceImpl will create new CompanyService object
func NewCompanyServiceImpl(CatRepo admin.CompanyRepository) *CompanyServiceImpl {
	return &CompanyServiceImpl{CompanyRepo: CatRepo}
}

// Companies ..() returns list of companies
func (cs *CompanyServiceImpl) Companies() ([]entity.Company, error) {

	companies, err := cs.CompanyRepo.Companies()

	if err != nil {
		return nil, err
	}

	return companies, nil
}

// StoreCompany persists new Company information
func (cs *CompanyServiceImpl) StoreCompany(Company entity.Company) error {

	err := cs.CompanyRepo.StoreCompany(Company)

	if err != nil {
		return err
	}

	return nil
}

// Company returns a Company object with a given id
func (cs *CompanyServiceImpl) Company(id int) (entity.Company, error) {

	c, err := cs.CompanyRepo.Company(id)

	if err != nil {
		return c, err
	}

	return c, nil
}

// UpdateCompany updates a cateogory with new data
func (cs *CompanyServiceImpl) UpdateCompany(Company entity.Company) error {

	err := cs.CompanyRepo.UpdateCompany(Company)

	if err != nil {
		return err
	}

	return nil
}

// DeleteCompany delete a Company by its id
func (cs *CompanyServiceImpl) DeleteCompany(id int) error {

	err := cs.CompanyRepo.DeleteCompany(id)
	if err != nil {
		return err
	}
	return nil
}
