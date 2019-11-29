package entity

type Material struct {
	Name         string
	Owner        string //CompanyID from the Company information
	PricePerDay  float64
	IsOnDiscount bool
	Discount     float32 //in percent
	State        byte    //is the material new or used
	OnSale       bool    // is the material open for sale
}
