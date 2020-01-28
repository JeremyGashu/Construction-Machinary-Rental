package repository

import (
	"database/sql"
	"errors"

	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

// UserRepositoryImpl implements the menu.UserRepository interface
type UserRepositoryImpl struct {
	conn *sql.DB
}

// NewUserRepositoryImpl will create an object of PsqlUserRepository
func NewUserRepositoryImpl(Conn *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{conn: Conn}
}

// Users returns all cateogories from the database
func (cri *UserRepositoryImpl) Users() ([]entity.User, error) {

	rows, err := cri.conn.Query("SELECT * FROM users;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	ctgs := []entity.User{}

	for rows.Next() {
		User := entity.User{}
		err = rows.Scan(&User.Username, &User.FirstName, &User.LastName, &User.Email, &User.Phone, &User.DeliveryAddress, &User.Password, &User.ImagePath, &User.Account)
		if err != nil {
			return nil, err
		}
		ctgs = append(ctgs, User)
	}

	return ctgs, nil
}

// User returns a User with a given id
func (cri *UserRepositoryImpl) User(id string) (entity.User, error) {

	row := cri.conn.QueryRow("SELECT * FROM users WHERE username=$1", id)

	User := entity.User{}

	err := row.Scan(&User.Username, &User.FirstName, &User.LastName, &User.Email, &User.Phone, &User.DeliveryAddress, &User.Password, &User.ImagePath, &User.Account)
	if err != nil {
		return User, err
	}

	return User, nil
}

// UpdateUser updates a given object with a new data
func (cri *UserRepositoryImpl) UpdateUser(c entity.User) error {

	_, err := cri.conn.Exec("UPDATE users SET firstname=$1,lastname=$2, email=$3,phone=$4,address=$5,password=$6,imagepath=$7,account=$8 WHERE username=$9", c.FirstName, c.LastName, c.Email, c.Phone, c.DeliveryAddress, c.Password, c.ImagePath, c.Account, c.Username)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}

// DeleteUser removes a User from a database by its id
func (cri *UserRepositoryImpl) DeleteUser(id string) error {

	_, err := cri.conn.Exec("DELETE FROM users WHERE username=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

// StoreUser stores new User information to database
func (cri *UserRepositoryImpl) StoreUser(c entity.User) error {

	_, err := cri.conn.Exec("INSERT INTO Users (username,firstname,lastname,email,phone,address,password,imagepath,account) values($1, $2, $3,$4,$5,$6,$7,$8,$9)", c.Username, c.FirstName, c.LastName, c.Email, c.Phone, c.DeliveryAddress, c.Password, c.ImagePath, c.Account)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}

//Pay - Payment for the user
func (cri *UserRepositoryImpl) Pay(uname string, amount float64) bool {
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

//AuthUser -
func (cri *UserRepositoryImpl) AuthUser(username string, password string) bool {
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

//GetRentedMaterials -
func (cri *UserRepositoryImpl) GetRentedMaterials(uname string) ([]entity.RentInformation, error) {
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
