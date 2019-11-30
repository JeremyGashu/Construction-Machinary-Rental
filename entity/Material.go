package entity

type Material struct {
	Name        string
	Type        byte   //type of instrument wheather it is excavator or backhoe
	Owner       string //CompanyID from the Company information
	PricePerDay float64
	OnDiscount  bool
	Discount    float32 //in percent
	State       byte    //is the material new or used
	OnSale      bool    // is the material open for sale
}
