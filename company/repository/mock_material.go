package repository

import (
	"database/sql"
	"errors"

	"github.com/ermiasgashu/Construction-Machinary-Rental/company"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

//MockMaterialRepository implements the menu.MaterialRepository interface
type MockMaterialRepository struct {
	conn *sql.DB
}

// NewMockMaterialRepository will create an object of
func NewMockMaterialRepository(Conn *sql.DB) company.MaterialRepository {
	return &MockMaterialRepository{conn: Conn}
}

// Materials returns all Materials from the database
func (cri *MockMaterialRepository) Materials() ([]entity.Material, error) {

	res := []entity.Material{entity.MaterialMock}
	return res, nil
}

// Material returns a Material with a given id
func (cri *MockMaterialRepository) Material(id int) (entity.Material, error) {
	res := entity.MaterialMock
	if id == 1 {
		return res, nil
	}
	return res, errors.New("Not found")
}

// UpdateMaterial updates a given object with a new data
func (cri *MockMaterialRepository) UpdateMaterial(c entity.Material) error {
	res := entity.MaterialMock
	res.Discount = c.Discount
	res.ID = c.ID
	res.ImagePath = c.ImagePath
	res.Name = c.Name
	res.OnDiscount = c.OnDiscount
	res.OnSale = c.OnSale
	res.Owner = c.Owner
	res.PricePerDay = c.PricePerDay
	return nil
}

// DeleteMaterial removes a Material from a database by its id
func (cri *MockMaterialRepository) DeleteMaterial(id int) error {
	material := entity.MaterialMock
	if material.ID != 1 {
		return errors.New("Error found")
	}
	return nil
}

// AddMaterial stores new Material information to database
func (cri *MockMaterialRepository) AddMaterial(c entity.Material) error {
	res := entity.MaterialMock
	res.Discount = c.Discount
	res.ID = c.ID
	res.ImagePath = c.ImagePath
	res.Name = c.Name
	res.OnDiscount = c.OnDiscount
	res.OnSale = c.OnSale
	res.Owner = c.Owner
	res.PricePerDay = c.PricePerDay
	return nil
}

//MaterialByCompanyOwner -
func (mr *MockMaterialRepository) MaterialByCompanyOwner(id int) ([]entity.Material, error) {
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
func (mr *MockMaterialRepository) GetOwner(id int) (entity.Company, error) {
	query := "select id,name,email,address,phone,description,rating,imagepath from companies where id = $1"
	var Company entity.Company
	err := mr.conn.QueryRow(query, id).Scan(&Company.CompanyID, &Company.Name, &Company.Email, &Company.Address, &Company.PhoneNo, &Company.Description, &Company.Rating, &Company.ImagePath)
	if err != nil {
		return Company, err
	}
	return Company, nil
}

func (mr *MockMaterialRepository) RentMaterial(rentInfo entity.RentInformation) error {
	query := "insert into materials_rented(material_id,company_id,borrower,rent_date,due_date,transactionmade) values($1,$2,$3,$4,$5,$6)"
	// fmt.Println(rentInfo)

	_, err := mr.conn.Exec(query, rentInfo.MaterialID, rentInfo.CompanyID, rentInfo.Username, rentInfo.RentDate, rentInfo.DueDate, rentInfo.TransactionMade)
	if err != nil {
		return err
	}
	return nil
}

//MaterialSearch ..
func (mr *MockMaterialRepository) MaterialSearch(name string) ([]entity.Material, error) {
	materials := make([]entity.Material, 0)
	query := "select * from materials where name like '%' ||$1|| '%'"
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
