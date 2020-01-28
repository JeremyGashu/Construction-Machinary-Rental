package repository

import (
	"database/sql"
	"errors"

	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

// AdminRepositoryImpl implements the menu.AdminRepository interface
type AdminRepositoryImpl struct {
	conn *sql.DB
}

// NewAdminRepositoryImpl will create an object of PsqlAdminRepository
func NewAdminRepositoryImpl(Conn *sql.DB) *AdminRepositoryImpl {
	return &AdminRepositoryImpl{conn: Conn}
}

// Admins returns all Admins from the database
func (cri *AdminRepositoryImpl) Admins() ([]entity.Admin, error) {

	rows, err := cri.conn.Query("SELECT * FROM admin")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	ctgs := []entity.Admin{}

	for rows.Next() {
		Admin := entity.Admin{}
		err = rows.Scan(&Admin.Username, &Admin.FirstName, &Admin.LastName, &Admin.Email, &Admin.Password, &Admin.ImagePath)
		if err != nil {
			return nil, err
		}
		ctgs = append(ctgs, Admin)
	}

	return ctgs, nil
}

// Admin returns a Admin with a given id
func (cri *AdminRepositoryImpl) Admin(id string) (entity.Admin, error) {

	row := cri.conn.QueryRow("select * from admin where username=$1", id)

	Admin := entity.Admin{}

	err := row.Scan(&Admin.Username, &Admin.FirstName, &Admin.LastName, &Admin.Email, &Admin.Password, &Admin.ImagePath)
	if err != nil {
		return Admin, err
	}

	return Admin, nil
}

// UpdateAdmin updates a given object with a new data
func (cri *AdminRepositoryImpl) UpdateAdmin(c entity.Admin) error {
	// fmt.Println(c)

	_, err := cri.conn.Exec("UPDATE admin SET firstname=$1,email=$2,password=$3,imagepath=$4,lastname=$5 WHERE username=$6", c.FirstName, c.Email, c.Password, c.ImagePath, c.LastName, c.Username)
	if err != nil {
		// fmt.Println(err)
		return errors.New("Update has failed")
	}

	return nil
}

// DeleteAdmin removes a Admin from a database by its id
func (cri *AdminRepositoryImpl) DeleteAdmin(id string) error {

	_, err := cri.conn.Exec("DELETE FROM admin WHERE username=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

// StoreAdmin stores new Admin information to database
func (cri *AdminRepositoryImpl) StoreAdmin(c entity.Admin) error {

	_, err := cri.conn.Exec("INSERT INTO admin (username,firstname,lastname,email,password,imagepath) values($1, $2, $3,$4,$5,$6)", c.Username, c.FirstName, c.LastName, c.Email, c.Password, c.ImagePath)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}

//AuthUser -
func (cri *AdminRepositoryImpl) AuthAdmin(username string, password string) bool {
	query := "select username from admin where username=$1 and password=$2"
	var name string
	row := cri.conn.QueryRow(query, username, password)
	err := row.Scan(&name)

	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}
	return true
}
