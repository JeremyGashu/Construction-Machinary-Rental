package entity

type User struct {
	FirstName       string
	LastName        string
	Username        string
	Email           string
	Phone           string
	DeliveryAddress string //Address
	PostNum         string //Optional
	Password        string
	ImagePath       string //
	Account         int    //By default we will give them 200000 as starting point, so dont include it in the field
	// Rating          byte // The rating will be made by the loaner company
	//This are infos that will be used in the system, another infos like password will be saved in db
}
