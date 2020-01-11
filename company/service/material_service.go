package service

import (
	"github.com/ermiasgashu/Construction-Machinary-Rental/company"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
	"github.com/ermiasgashu/Construction-Machinary-Rentalold/material"
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
