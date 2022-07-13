package Models

type Retailer struct {
	RetailerId  	string  `json:"id" gorm:"primaryKey" binding:"required"`
	RetailerName 	string 	`json:"retailer_name" binding:"required"`
	RetailerAddress string 	`json:"retailer_address" binding:"required"`
}

func (r *Retailer) TableName() string {
	return "retailer"
}
