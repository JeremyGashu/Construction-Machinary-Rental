package material

import "github.com/ermiasgashu/Construction-Machinary-Rental/entity"

type MaterialService interface {
	Materials() ([]entity.Material, error)
	Material(id string) (entity.Material, error)
	UpdateMaterial(material entity.Material) error
	DeleteMaterial(entity.Material) error
	AddMaterial(material entity.Material) error
}
