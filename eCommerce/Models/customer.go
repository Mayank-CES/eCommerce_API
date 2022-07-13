package Models

type Customer struct {

	CustomerId         	string  `json:"customer_id" gorm:"primaryKey" binding:"required"`
	CustomerName 		string 	`json:"customer_name" binding:"required"`
	CustomerAddress    	string 	`json:"customer_address" binding:"required"`
}

func (c *Customer) TableName() string {
	return "customer"
}
