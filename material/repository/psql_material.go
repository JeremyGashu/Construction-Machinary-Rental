package repository

import (
	"database/sql"
	"errors"

	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

//MaterialRepository -
type MaterialRepository struct {
	conn *sql.DB
}

//NewMaterialRepository -
func NewMaterialRepository(Conn *sql.DB) *MaterialRepository {
	return &MaterialRepository{conn: Conn}
}

//Materials -
func (mr *MaterialRepository) Materials() ([]entity.Material, error) {
	materials := make([]entity.Material, 0)
	query := "select * from materials"
	data, err := mr.conn.Query(query)
	if err != nil {
		return materials, errors.New("No user is found")
	}
	for data.Next() {
		var material entity.Material
		data.Scan(&material.ID, &material.Name, &material.Type, &material.Owner, &material.PricePerDay, &material.OnDiscount, &material.Discount, &material.State, &material.OnSale, &material.ImagePath) //all the datas that will be added in the category
		materials = append(materials, material)
	}
	if err := data.Err(); err != nil {
		return materials, errors.New("Some error is occured")
	}
	return materials, nil
}

//Material -
func (mr *MaterialRepository) Material(id int) (entity.Material, error) {
	material := entity.Material{}
	query := "select * from materials where id=$1"
	err := mr.conn.QueryRow(query, id).Scan(&material.ID, &material.Name, &material.Type, &material.Owner, &material.PricePerDay, &material.OnDiscount, &material.Discount, &material.State, &material.OnSale, &material.ImagePath)
	if err != nil {
		return material, err
	}
	return material, nil
}

//UpdateMaterial -
func (mr *MaterialRepository) UpdateMaterial(material entity.Material) error {
	query := "update materials set name=$1,type=$2,owner=$3,priceperday=$4,ondiscount=$5,discount=$6,state=$7,onsale=$8,imagepath=$9 where id=$10"
	_, err := mr.conn.Exec(query, material.Name, material.Type, material.Owner, material.PricePerDay, material.OnDiscount, material.Discount, material.State, material.OnSale, material.ImagePath, material.ID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteMaterial -
func (mr *MaterialRepository) DeleteMaterial(id int) error {
	query := "delete from materials where id=$1"
	_, err := mr.conn.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

//AddMaterial -
func (mr *MaterialRepository) AddMaterial(material entity.Material) error {
	query := "insert into materials(name, type, owner,priceperday, ondiscount, discount, state, onsale, imagepath) values($1,$2,$3,$4,$5,$6,$7,$8,$9)"

	_, err := mr.conn.Exec(query, material.Name, material.Type, material.Owner, material.PricePerDay, material.OnDiscount, material.Discount, material.State, material.OnSale, material.ImagePath)
	if err != nil {
		return err
	}
	return nil
}

//MaterialByCompanyOwner -
func (mr *MaterialRepository) MaterialByCompanyOwner(id int) ([]entity.Material, error) {
	materials := make([]entity.Material, 0)
	query := "select * from materials where owner=$1"
	data, err := mr.conn.Query(query, id)
	if err != nil {
		return materials, errors.New("No user is found")
	}
	for data.Next() {
		var material entity.Material
		data.Scan(&material.ID, &material.Name, &material.Type, &material.Owner, &material.PricePerDay, &material.OnDiscount, &material.Discount, &material.State, &material.OnSale, &material.ImagePath) //all the datas that will be added in the category
		materials = append(materials, material)
	}
	if err := data.Err(); err != nil {
		return materials, errors.New("Some error is occured")
	}
	return materials, nil
}
