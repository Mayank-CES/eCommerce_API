package Repository

import (
	"eCommerce/Config"
	"eCommerce/Models"
)

type Repo interface {
	CreateCustomerAccount(customer *Models.Customer)	(err error)
	AddTransaction(order *Models.Order)	(err error)
	BuyProduct(product *Models.Product, order *Models.Order)	(err error)
	CheckOrderByID(transaction *Models.Transaction, id string) 	(err error)


	GetAllProducts(products *[]Models.Product) (err error)
	GetProductByID(product *Models.Product, id string) (err error)

	CreateRetailerAccount(retailer *Models.Retailer)	(err error)
	AddProduct(product *Models.Product)	(err error)
	PatchProduct(updatedProduct *Models.PatchProd,id string)	(err error)
	GetRHistoryByID(order *Models.Transaction,id string)	(err error)

}

type repository struct{}

func (r *repository) CreateRetailerAccount(retailer *Models.Retailer) (err error) {
	panic("implement me")
}

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
func (r *repository) AddTransaction(order *Models.Order)	( err error) {
	if err = Config.DB.Create(&order)/*Models.Transaction{ProductId: order.ProductId,Quantity: order.Quantity,Status: "order placed"})*/.Error; err != nil {
		return err
	}
	return nil
}

//BuyProduct ... Fetch only one product by Id
func (r *repository) BuyProduct(product *Models.Product,order *Models.Order) (err error) {
	//var prod *Models.Product
	//Config.DB.Model(&Models.Product{}).Where("product_id = ?", order.ProductId).Find(&prod)
	var quantity= order.Quantity
	if err = Config.DB.Model(&Models.Product{}).Where("product_id = ?", order.ProductId).Updates(map[string]interface{}{"Quantity":product.Quantity-quantity}).Error; err != nil {
		return err

	}
	return nil
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

// CreatRetailerAccount  ... Insert New Retailer data
func (r *repository)  CreatRetailerAccount(retailer *Models.Retailer)	(err error){
	if err = Config.DB.Create(retailer).Error; err != nil {
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
	if err = Config.DB.Model(&Models.Product{}).Where("product_id = ?", id).Updates(map[string]interface{}{"Price":updatedProduct.Price,"Quantity":updatedProduct.Quantity}).Error; err != nil {
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