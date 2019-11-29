package entity

type Company struct {
	CompanyID string
	Name      string
	Email     string
	PhoneNo   string
	Address   string
	Rating    int
	//This are infos that will be used in the system, another infos like password will be saved in db
}
