package entity

//Company -
type Company struct {
	CompanyID   int
	Name        string
	Email       string
	PhoneNo     string
	Address     string
	Description string
	Rating      int
	ImagePath   string //the logo picture path of the company
	Password    string
	Account     float32
	//This are infos that will be used in the system, another infos like password will be saved in db
}
