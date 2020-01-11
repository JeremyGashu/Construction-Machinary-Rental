package company

import "github.com/ermiasgashu/Construction-Machinary-Rental/entity"

type MaterialService interface {
	Materials() ([]entity.Material, error)
	Material(id int) (entity.Material, error)
	UpdateMaterial(material entity.Material) error
	DeleteMaterial(id int) error
	AddMaterial(material entity.Material) error
	MaterialByCompanyOwner(id int) ([]entity.Material, error)
}
type CompanyService interface {
	Companies() ([]entity.Company, error)
	Company(id int) (entity.Company, error)
	UpdateCompany(company entity.Company) error
	DeleteCompany(id int) error
	AddCompany(company entity.Company) error
}