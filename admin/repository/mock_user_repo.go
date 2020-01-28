package repository

import (
	"database/sql"
	"errors"

	"github.com/ermiasgashu/Construction-Machinary-Rental/admin"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

//MockUserRepository implements the menu.UserRepository interface
type MockUserRepository struct {
	conn *sql.DB
}

// NewMockUserRepository will create an object of
func NewMockUserRepository(Conn *sql.DB) admin.UserRepository {
	return &MockUserRepository{conn: Conn}
}

// Users returns all Users from the database
func (cri *MockUserRepository) Users() ([]entity.User, error) {

	res := []entity.User{entity.UserMock}
	return res, nil
}

// User returns a User with a given id
func (cri *MockUserRepository) User(id string) (entity.User, error) {
	res := entity.UserMock
	if id == res.Username {
		return res, nil
	}
	return res, errors.New("Not found")
}

// UpdateUser updates a given object with a new data
func (cri *MockUserRepository) UpdateUser(c entity.User) error {
	res := entity.UserMock
	res.Username = c.Username
	res.FirstName = c.FirstName
	res.LastName = c.LastName
	res.Password = c.Password
	res.Account = c.Account
	res.DeliveryAddress = c.DeliveryAddress
	res.Email = c.Email
	res.ImagePath = c.ImagePath
	return nil
}

// DeleteUser removes a User from a database by its id
func (cri *MockUserRepository) DeleteUser(id string) error {
	res := entity.UserMock
	if id != res.Username {
		return errors.New("Not found")
	}
	return nil
}

// StoreUser stores new User information to database
func (cri *MockUserRepository) StoreUser(c entity.User) error {

	res := entity.UserMock
	res.Username = c.Username
	res.FirstName = c.FirstName
	res.LastName = c.LastName
	res.Password = c.Password
	res.Account = c.Account
	res.DeliveryAddress = c.DeliveryAddress
	res.Email = c.Email
	res.ImagePath = c.ImagePath
	return nil
}




















//AuthUser -
func (cri *MockUserRepository) AuthUser(username string, password string) bool {
	query := "select username from users where username=$1 and password=$2"
	var name string
	row := cri.conn.QueryRow(query, username, password)
	err := row.Scan(&name)

	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
	}
	return true
}
func (cri *MockUserRepository) GetRentedMaterials(uname string) ([]entity.RentInformation, error) {
	query := "select * from materials_rented where borrower=$1"
	infos := make([]entity.RentInformation, 0)
	data, err := cri.conn.Query(query, uname)
	if err != nil {
		return infos, errors.New("No user is found")
	}
	for data.Next() {
		var info entity.RentInformation
		data.Scan(&info.MaterialID, &info.CompanyID, &info.RentDate, &info.DueDate, &info.TransactionMade, &info.Username) //all the datas that will be added in the category
		infos = append(infos, info)
	}
	if err := data.Err(); err != nil {
		return infos, errors.New("Some error is occured")
	}
	return infos, nil
}

//Pay - Payment for the user
func (cri *MockUserRepository) Pay(uname string, amount float64, cid int) bool {

	user, err := cri.User(uname)
	// fmt.Println(user)
	if err != nil {
		return false
	}
	user.Account = user.Account - amount
	err = cri.UpdateUser(user)
	if err != nil {
		return false
	}
	return true
}
