package service

import (
	"github.com/ermiasgashu/Construction-Machinary-Rental/company"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
	"github.com/ermiasgashu/Construction-Machinary-Rental/material"
	// "github.com/ermiasgashu/Construction-Machinary-Rental/material"
)

//MaterialService -
type MaterialService struct {
	repository company.MaterialRepository
}

//NewMaterialService -
func NewMaterialService(repo material.MaterialService) *MaterialService {
	return &MaterialService{repository: repo}
}

//Materials -
func (ms *MaterialService) Materials() ([]entity.Material, error) {
	materials, err := ms.repository.Materials()
	if err != nil {
		return materials, err
	}
	return materials, nil
}

//Material -
func (ms *MaterialService) Material(id int) (entity.Material, error) {
	material, err := ms.repository.Material(id)
	if err != nil {
		return material, err
	}
	return material, nil
}

//UpdateMaterial -
func (ms *MaterialService) UpdateMaterial(material entity.Material) error {
	err := ms.repository.UpdateMaterial(material)
	if err != nil {
		return err
	}
	return nil
}

//DeleteMaterial -
func (ms *MaterialService) DeleteMaterial(id int) error {
	err := ms.repository.DeleteMaterial(id)
	if err != nil {
		return err
	}
	return nil
}

//AddMaterial -
func (ms *MaterialService) AddMaterial(material entity.Material) error {
	err := ms.repository.AddMaterial(material)
	if err != nil {
		return err
	}
	return nil
}

//MaterialByCompanyOwner -
func (ms *MaterialService) MaterialByCompanyOwner(id int) ([]entity.Material, error) {
	materials, err := ms.repository.MaterialByCompanyOwner(id)
	if err != nil {
		return materials, err
	}
	return materials, nil
}

//GetOwner -
func (ms *MaterialService) GetOwner(id int) (entity.Company, error) {
	comp, err := ms.repository.GetOwner(id)
	if err != nil {
		return comp, err
	}
	return comp, nil
}

//RentMaterial -
func (ms *MaterialService) RentMaterial(rentInfo entity.RentInformation) error {
	err := ms.RentMaterial(rentInfo)
	if err != nil {
		return err
	}
	return nil
}

//MaterialSearch ..
func (ms *MaterialService) MaterialSearch(name string) ([]entity.Material, error) {
	materials, err := ms.repository.MaterialSearch(name)
	if err != nil {
		return materials, err
	}
	return materials, nil
}
