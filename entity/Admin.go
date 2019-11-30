package entity

type Admin struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string //this will be hashed to give the system more security
}
