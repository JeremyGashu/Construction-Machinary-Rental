package entity

//Admin -
type Admin struct {
	Username  string
	FirstName string
	LastName  string
	Email     string
	Password  string //this will be hashed to give the system more security
}
