package entity

//User -
type User struct {
  Username        string json:"username"
  FirstName       string json:"firstname"
  LastName        string json:"lastname"
  Email           string json:"email"
  Phone           string json:"phone"
  DeliveryAddress string json:"address" //Address
  Password        string json:"password"
  ImagePath       string json:"image"   //
  Account         int    json:"account" //By default we will give them 200000 as starting point, so dont include it in the field
  // Rating          byte // The rating will be made by the loaner company
  //This are infos that will be used in the system, another infos like password will be saved in db
}
