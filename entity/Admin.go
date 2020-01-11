package entity

//Admin is
type Admin struct {
	FirstName string `json:"Firstname"`
	LastName  string `json:"Lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"` //this will be hashed to give the system more security
	ImagePath string `json:"image"`    // material image path
}
