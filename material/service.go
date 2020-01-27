package material

import "github.com/ermiasgashu/Construction-Machinary-Rental/entity"

type MaterialService interface {
	Materials() ([]entity.Material, error)
	Material(id int) (entity.Material, error)
	UpdateMaterial(material entity.Material) error
	DeleteMaterial(id int) error
	AddMaterial(material entity.Material) error
	GetOwner(id int) (entity.Company, error)
	MaterialByCompanyOwner(id int) ([]entity.Material, error)
	RentMaterial(rentInfo entity.RentInformation) error
	MaterialSearch(name string) ([]entity.Material, error)

}
