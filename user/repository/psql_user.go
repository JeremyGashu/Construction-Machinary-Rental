package repository

import (
	"database/sql"
	"errors"

	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

//PsqlUserRepository -
type PsqlUserRepository struct {
	conn *sql.DB
}

//NewPsqlUserRepository -
func NewPsqlUserRepository(Conn *sql.DB) *PsqlUserRepository {
	return &PsqlUserRepository{conn: Conn}
}

//User -
func (pr *PsqlUserRepository) User(username string) (entity.User, error) {
	//returning a specifiv user or and error report
	user := entity.User{}
	query := "select firstname, lastname,email,phone,address from users where username = $1"
	err := pr.conn.QueryRow(query, username).Scan(&user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.DeliveryAddress) //all the user info
	if err != nil {
		return user, errors.New("No user found with this id")
	}
	return user, nil
}

//Users -
func (pr *PsqlUserRepository) Users() ([]entity.User, error) {
	users := make([]entity.User, 0)
	query := "select firstname, lastname,email,phone,address from users"
	data, err := pr.conn.Query(query)
	if err != nil {
		return users, errors.New("No user is found")
	}
	for data.Next() {
		var user entity.User
		data.Scan(&user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.DeliveryAddress) //all the datas that will be added in the category
		users = append(users, user)
	}
	if err := data.Err(); err != nil {
		return users, errors.New("Some error is occured")
	}
	return users, nil
}

//UpdateUser -
func (pr *PsqlUserRepository) UpdateUser(user entity.User) error {
	// u := entity.User{}
	query := "update users set firstname=$1,lastname=$2,email=$3,phone=$4,deliveryaddress=$5,imagename=$6,account=$7 where username=$9"
	_, err := pr.conn.Exec(query, user.FirstName, user.LastName, user.Email, user.Phone, user.DeliveryAddress, user.ImagePath, user.Account, user.Username)
	if err != nil {
		return err
	}
	return nil
}

//DeleteUser -
func (pr *PsqlUserRepository) DeleteUser(username string) error {
	query := "delete from users where username=$1"
	_, err := pr.conn.Exec(query, username)
	if err != nil {
		return (err)
	}
	return nil
}

//AddUser -
func (pr *PsqlUserRepository) AddUser(user entity.User) error {
	//We add on it after we figure all the co;umns out...
	query := "insert into users(username,firstname, lastname,email,phone,address,imagepath) values($1,$2,$3,$4,$5,$5,$7)"
	_, err := pr.conn.Exec(query, user.Username, user.FirstName, user.LastName, user.Email, user.Phone, user.DeliveryAddress, user.ImagePath)
	if err != nil {
		return err
	}
	return nil
}
