package entity

//Comment ..
type Comment struct {
	ID       uint   `json:"id"`
	UserName string `json:"fullname" gorm:"type:varchar(255)"`
	Message  string `json:"message"`
	Email    string `json:"email" gorm:"type:varchar(255);not null; unique"`
	PlacedAt string `json:"placedat"`
}
