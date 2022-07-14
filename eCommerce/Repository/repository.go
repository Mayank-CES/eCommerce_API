package Repository

import (
	"eCommerce/Config"
	"eCommerce/Models"
	"fmt"
)

type Repo interface {
	CreateCustomerAccount(customer *Models.Customer)	(err error)
	AddTransaction(transaction *Models.Transaction)	(err error)
	BuyProduct(product *Models.Product, transaction *Models.Transaction)	(err error)
	BuyProduct2(transaction *Models.Transaction)	(err error)
	CheckOrderByID(transaction *Models.Transaction, id string) 	(err error)


	GetAllProducts(products *[]Models.Product) (err error)
	GetProductByID(product *Models.Product, id string) (err error)


	CreateRetailerAccount(retailer *Models.Retailer)	(err error)
	AddProduct(product *Models.Product)	(err error)
	PatchProduct(updatedProduct *Models.PatchProd,id string)	(err error)
	GetRHistoryByID(order *Models.Transaction,id string)	(err error)

}

type repository struct{}

func NewRepo() Repo {
	return &repository{}
}

var err error

// CreateCustomerAccount  ... Insert New Customer data
func (r *repository)  CreateCustomerAccount(customer *Models.Customer)	(err error){
	if err = Config.DB.Create(&customer).Error; err != nil {
		return err
	}
	return nil
}

// AddTransaction ... Fetch only one product by Id
func (r *repository) AddTransaction(transaction *Models.Transaction)	( err error) {
	if err = Config.DB.Create(&transaction)/*Models.Transaction{ProductId: order.ProductId,Quantity: order.Quantity,Status: "order placed"})*/.Error; err != nil {
		return err
	}
	return nil
}

//BuyProduct ... Fetch only one product by Id
func (r *repository) BuyProduct(product *Models.Product,transaction *Models.Transaction) (err error) {
	var id string
	id = transaction.ProdId
	if err = Config.DB.Where("product_id = ?", id).First(&product).Error; err != nil {
		fmt.Printf("Error in finding the product %v", err)
		return err
	}else{
		//fmt.Println("Order Placed")
		newVal :=product.Quantity-transaction.TransactionQuantity
		//fmt.Printf("NewVal : %d", newVal)
		if newVal>0{
			//fmt.Println("NewVal>0")
			if err = Config.DB.Model(&Models.Product{}).Where("product_id = ?", id).Updates(map[string]interface{}{"Quantity":newVal}).Error; err != nil {
				return err
			}
		}else{
			fmt.Printf("Order Value is greater than the available stock. Please check the quantity.")
		}
	return nil
	}
}

//BuyProduct2 ... Fetch only one product by Id
func (r *repository) BuyProduct2(transaction *Models.Transaction) (err error) {
	var product Models.Product
	var id string
	id = transaction.ProdId
	if err = Config.DB.Where("product_id = ?", id).Find(product).Error; err != nil {
		fmt.Println("Error in finding the product")
		return err
	} else {
		fmt.Println("Order Placed")
		return nil
	}
}

//CheckOrderByID ... Fetch only one Order by Id
func (r *repository) CheckOrderByID(transaction *Models.Transaction, id string) (err error) {
	if err = Config.DB.Where("transaction_id = ?", id).First(transaction).Error; err != nil {
		return err
	}
	return nil
}

// GetAllProducts -----  Fetch all product data
func (r *repository) GetAllProducts(products *[]Models.Product) (err error){
	if err = Config.DB.Find(products).Error; err != nil {
		return err
	}
	return nil
}


//GetProductByID ... Fetch only one product by Id
func (r *repository) GetProductByID(product *Models.Product, id string) (err error) {
	if err = Config.DB.Where("product_id = ?", id).Find(product).Error; err != nil {
		return err
	}
	return nil
}

// CreateRetailerAccount  ... Insert New Retailer data
func (r *repository)  CreateRetailerAccount(retailer *Models.Retailer)	(err error){
	if err = Config.DB.Save(retailer).Error; err != nil {
		return err
	}
	return nil
}

// AddProduct ... Insert New Retailer data
func (r *repository)  AddProduct(product *Models.Product)	(err error){
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

// PatchProduct ... Insert New Retailer data
func (r *repository)  PatchProduct(updatedProduct *Models.PatchProd, id string)	(err error){
	if err = Config.DB.Model(&Models.Product{}).Where("product_id = ?", id).Updates(map[string]interface{}{"Price":updatedProduct.Price,"Quantity":updatedProduct.PatchQuantity}).Error; err != nil {
		return err
	}
	return nil
}

// GetRHistoryByID ... Insert New Retailer data
func (r *repository)  GetRHistoryByID(order *Models.Transaction,id string)	(err error){
	if err = Config.DB.Where("retailer_id = ?", id).Find(order).Error; err != nil {
		return err
	}
	return nil
}


/*

// UpdateInfo ... Update student
func UpdateInfo(student *Student, id string) (err error) {
	fmt.Println(student)
	Config.DB.Save(student)
	return nil
}

//DeleteStudent ... Delete student
func DeleteStudent(student *Student, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(student)
	return nil
}


//GetAllMarks ... Fetch all Marks
func GetAllMarks(marks *[]Marks) (err error) {
	if err = Config.DB.Find(marks).Error; err != nil {
		return err
	}
	return nil
}

//GetMarksByID ... Fetch only one user by Id
func GetMarksByID(marks *Marks, id string) (err error) {
	if err = Config.DB.Where("student_id = ?", id).Find(marks).Error; err != nil {
		return err
	}
	return nil
}

//UpdateMarks ... Update marks
func UpdateMarks(marks *Marks, id string) (err error) {
	fmt.Println(marks)
	Config.DB.Save(marks)
	return nil
}

//DeleteMarks ... Delete marks
func DeleteMarks(marks *Marks, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(marks)
	return nil
}

//DeleteAllMarks ... Delete marks
func DeleteAllMarks(marks *[]Marks, id string) (err error) {
	Config.DB.Where("student_id = ?", id).Delete(marks)
	return nil
}

 */