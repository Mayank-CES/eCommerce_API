package Models

type Product struct {
	ProductId 		string	`json:"product_id" gorm:"primaryKey" binding:"required"`
	ProductName    	string 	`json:"product_name" binding:"required"`
	Price			uint	`json:"price" binding:"required"`
	Quantity		uint	`json:"quantity" binding:"required"`
}

type PatchProd struct {
	Price			uint	`json:"price" binding:"required"`
	PatchQuantity		uint	`json:"patch_quantity" binding:"required"`
}


func (p *Product) TableName() string {
	return "product"
}
