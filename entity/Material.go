package entity

//Material -
type Material struct {
	ID          int
	Name        string
	Owner       int //CompanyID from the Company information
	PricePerDay float64
	OnDiscount  bool
	Discount    float32 //in percent
	OnSale      bool    // is the material open for sale
	ImagePath   string  // material image path
}
