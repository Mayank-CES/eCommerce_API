package Models

type Transaction struct {

	TransactionId	string  `json:"transaction_id" gorm:"primaryKey" gorm:"autoincrement"`
	ProdId    	string 	`json:"prod_id" binding:"required"`
	TransactionQuantity		uint	`json:"transaction_quantity"`
	//ConsumerId		string	`json:"consumer_id" binding:"required"`
	//RetailerId		string	`json:"retailer_id" binding:"required"`
}

func (t *Transaction) TableName() string {
	return "transaction"
}
