package entity

type User struct {
	FirstName       string
	LastName        string
	Username        string
	Email           string
	Phone           string
	DeliveryAddress string
	PostNum         string
	Password        string
	// Rating          byte // The rating will be made by the loaner company
	//This are infos that will be used in the system, another infos like password will be saved in db
}
