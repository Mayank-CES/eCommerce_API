package Models

type Transaction struct {

	TransactionId	string  `json:"transaction_id" gorm:"primaryKey" gorm:"autoincrement"`
	ProductId    	string 	`json:"product_id" binding:"required"`
	Quantity		uint	`json:"transaction_quantity"`
	Status			string	`json:"order_status"`
	//ConsumerId		string	`json:"consumer_id" binding:"required"`
	//RetailerId		string	`json:"retailer_id" binding:"required"`
}

type Order struct{
	ProductId    	string 	`json:"product_id" binding:"required"`
	Quantity		uint	`json:"order_quantity"`
}

func (t *Transaction) TableName() string {
	return "transaction"
}
