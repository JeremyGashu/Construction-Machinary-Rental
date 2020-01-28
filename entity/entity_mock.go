package entity

//AdminMock is
var AdminMock = Admin{
	FirstName: "Moke fname",
	LastName:  "Moke lname",
	Username:  "Moke uName",
	Email:     "Moke email",
	Password:  "Moke password",
	ImagePath: "Moke image",
}

//CommentMock ..
var CommentMock = Comment{
	ID:       1,
	UserName: "Mokeuname",
	Message:  "Moke message",
	Email:    "Moke email",
	PlacedAt: "Moke placedate",
}

//CompanyMock -
var CompanyMock = Company{
	CompanyID:   1,
	Name:        "Moke name",
	Email:       "Moke email",
	PhoneNo:     "Moke phone",
	Address:     "Moke address",
	Description: "Moke desc",
	Rating:      4,
	ImagePath:   "Moke img",
	Password:    "Moke password",
	Account:     12121.2,
	Activated:   true,
}

//MaterialMock -
var MaterialMock = Material{
	ID:          1,
	Name:        "Moke name",
	Owner:       1,
	PricePerDay: 12.12,
	OnDiscount:  false,
	Discount:    1.1,
	OnSale:      false,
	ImagePath:   "Moke img",
}

//RentInformationMock - struct contains user info
var RentInformationMock = RentInformation{
	MaterialID:      1,
	CompanyID:       1,
	Username:        "Moke uname",
	RentDate:        "Moke rentdata",
	DueDate:         "Moke duedata",
	TransactionMade: 12.12,
}

//UserMock -
var UserMock = User{
	Username:        "Mokeuname",
	FirstName:       "Moke fname",
	LastName:        "Moke lname",
	Email:           "Moke email",
	Phone:           "Moke phone",
	DeliveryAddress: "Moke address", //Address
	Password:        "Moke password",
	ImagePath:       "Moke Img", //
	Account:         121232,
}
