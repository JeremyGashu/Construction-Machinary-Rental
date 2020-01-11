package entity

//Company -
type Company struct {
	CompanyID   int     `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	PhoneNo     string  `json:"phone"`
	Address     string  `json:"address"`
	Description string  `json:"description"`
	Rating      int     `json:"rating"`
	ImagePath   string  `json:"image"` //the logo picture path of the company
	Password    string  `json:"password"`
	Account     float32 `json:"account"`
	Activated   bool    `json:"activated"`
	//This are infos that will be used in the system, another infos like password will be saved in db
}
