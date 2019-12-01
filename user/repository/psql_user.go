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
func (pr *PsqlUserRepository) User(id string) (entity.User, error) {
	//returning a specifiv user or and error report
	user := entity.User{}
	query := "select * from user where id = $1"
	err := pr.conn.QueryRow(query, id).Scan(&user.DeliveryAddress) //all the user info
	if err != nil {
		return user, errors.New("No user found with this id")
	}
	return user, nil
}

//Users -
func (pr *PsqlUserRepository) Users() ([]entity.User, error) {
	users := make([]entity.User, 0)
	query := "select * from user"
	data, err := pr.conn.Query(query)
	if err != nil {
		return users, errors.New("No user is found")
	}
	for data.Next() {
		var user entity.User
		data.Scan(&user.Email) //all the datas that will be added in the category
		users = append(users, user)
	}
	if err := data.Err(); err != nil {
		return users, errors.New("Some error is occured")
	}
	return users, nil
}

//UpdateUser -
func (pr *PsqlUserRepository) UpdateUser(entity.User) error {
	//update the user info given the user
	return nil
}

//DeleteUser -
func (pr *PsqlUserRepository) DeleteUser(id int) error {
	//delete user getting the id
	return nil
}

//AddUser -
func (pr *PsqlUserRepository) AddUser(entity.User) error {
	return nil
}
