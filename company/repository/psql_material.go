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
		data.Scan(&material.ID, &material.Name, &material.Owner, &material.PricePerDay, &material.OnDiscount, &material.Discount, &material.OnSale, &material.ImagePath) //all the datas that will be added in the category
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
	err := mr.conn.QueryRow(query, id).Scan(&material.ID, &material.Name, &material.Owner, &material.PricePerDay, &material.OnDiscount, &material.Discount, &material.OnSale, &material.ImagePath)
	if err != nil {
		return material, err
	}
	return material, nil
}

//UpdateMaterial -
func (mr *MaterialRepository) UpdateMaterial(material entity.Material) error {
	query := "update materials set name=$1,owner=$2,priceperday=$3,ondiscount=$4,discount=$5,onsale=$6,imagepath=$7 where id=$8"
	_, err := mr.conn.Exec(query, material.Name, material.Owner, material.PricePerDay, material.OnDiscount, material.Discount, material.OnSale, material.ImagePath, material.ID)
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
	query := "insert into materials(name, priceperday, ondiscount, discount, onsale, imagepath,owner) values($1,$2,$3,$4,$5,$6,$7)"

	_, err := mr.conn.Exec(query, material.Name, material.PricePerDay, material.OnDiscount, material.Discount, material.OnSale, material.ImagePath, material.Owner)
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
		data.Scan(&material.ID, &material.Name, &material.Owner, &material.PricePerDay, &material.OnDiscount, &material.Discount, &material.OnSale, &material.ImagePath) //all the datas that will be added in the category
		materials = append(materials, material)
	}
	if err := data.Err(); err != nil {
		return materials, errors.New("Some error is occured")
	}
	return materials, nil
}

//GetOwner -
func (mr *MaterialRepository) GetOwner(id int) (entity.Company, error) {
	query := "select id,name,email,address,phone,description,rating,imagepath,account,password,activated from companies where id = $1"
	var Company entity.Company
	err := mr.conn.QueryRow(query, id).Scan(&Company.CompanyID, &Company.Name, &Company.Email, &Company.Address, &Company.PhoneNo, &Company.Description, &Company.Rating, &Company.ImagePath, &Company.Account, &Company.Password, &Company.Activated)
	if err != nil {
		return Company, err
	}
	return Company, nil
}

func (mr *MaterialRepository) RentMaterial(rentInfo entity.RentInformation) error {
	query := "insert into materials_rented values($1,$2,$3,$4,$5,$6)"
	// fmt.Println(rentInfo)

	_, err := mr.conn.Exec(query, rentInfo.MaterialID, rentInfo.CompanyID, rentInfo.RentDate, rentInfo.DueDate, rentInfo.TransactionMade, rentInfo.Username)
	if err != nil {
		return err
	}
	return nil
}

//MaterialSearch ..
func (mr *MaterialRepository) MaterialSearch(name string) ([]entity.Material, error) {
	materials := make([]entity.Material, 0)
	query := "select * from materials where name ilike '%' ||$1|| '%'"
	data, err := mr.conn.Query(query, name)
	if err != nil {
		return materials, errors.New("No user is found")
	}
	for data.Next() {
		var material entity.Material
		data.Scan(&material.ID, &material.Name, &material.Owner, &material.PricePerDay, &material.OnDiscount, &material.Discount, &material.OnSale, &material.ImagePath) //all the datas that will be added in the category
		materials = append(materials, material)
	}
	if err := data.Err(); err != nil {
		return materials, errors.New("Some error is occured")
	}
	return materials, nil
}

//MaterialsOnDiscount -
func (mr *MaterialRepository) MaterialsOnDiscount() ([]entity.Material, error) {
	materials := make([]entity.Material, 0)
	query := "select * from materials where ondiscount='true'"
	data, err := mr.conn.Query(query)
	if err != nil {
		return materials, errors.New("No user is found")
	}
	for data.Next() {
		var material entity.Material
		data.Scan(&material.ID, &material.Name, &material.Owner, &material.PricePerDay, &material.OnDiscount, &material.Discount, &material.OnSale, &material.ImagePath) //all the datas that will be added in the category
		materials = append(materials, material)
	}
	if err := data.Err(); err != nil {
		return materials, errors.New("Some error is occured")
	}
	return materials, nil
}
