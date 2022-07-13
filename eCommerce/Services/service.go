package Services
//
//import (
//	"eCommerce/Models"
//	"eCommerce/Repository"
//)
//
//var CService CustomerServices
//
//type CustomerServices interface {
//	/*---Customer---*/
//
//	CreateCustomerAccount(customer *Models.Customer)	(err error)
//	//BuyProduct(product *Models.Product)	(err error)
//	CheckOrderByID(transaction *Models.Transaction, id string) (err error)
//
//	/*---Product---*/
//	GetAllProducts(products *[]Models.Product)	(err error)
//	GetProductByID(product *Models.Product, id string)	(err error)
//
//	/*---Retailer---*/
//	CreateRetailerAccount(retailer *Models.Retailer)	(err error)
//	AddProduct(product *Models.Product)	(err error)
//	PatchProduct(updatedProduct *Models.PatchProd, id string)	(err error)
//	GetRHistoryByID(transaction *Models.Transaction, id string) (err error)
//
//}
//
//type CustomerService struct {
//	cust Repository.Repo
//}
//func NewCustomerService(repo Repository.Repo) *CustomerService {
//	return &CustomerService{
//		cust: repo,
//	}
//}
//
//
//func (c *CustomerService) CreateCustomerAccount(customer *Models.Customer)	(err error){
//	return c.cust.CreateCustomerAccount(customer)
//}
//
///*
//func (c *CustomerService) BuyProduct(product *Models.Product,quantity)	(err error){
//	return c.cust.BuyProduct(product)
//}
//*/
//
//func (c *CustomerService) CheckOrderByID(transaction *Models.Transaction, id string) 	(err error){
//	return c.cust.CheckOrderByID(transaction, id)
//}
//
//func (c *CustomerService) GetAllProducts(products *[]Models.Product)	(err error){
//	return c.cust.GetAllProducts(products)
//}
//
//func (c *CustomerService) GetProductByID(product *Models.Product, id string)	(err error){
//	return c.cust.GetProductByID(product,id)
//}
//
///*----Retailer----*/
//
//func (c *CustomerService) CreateRetailerAccount(retailer *Models.Retailer)	(err error){
//	return c.cust.CreateRetailerAccount(retailer)
//}
//
//func (c *CustomerService) AddProduct(product *Models.Product)	(err error){
//	return c.cust.AddProduct(product)
//}
//
//func (c *CustomerService) PatchProduct(updatedProduct *Models.PatchProd,id string)	(err error){
//	return c.cust.PatchProduct(updatedProduct,id)
//}
//
//func (c *CustomerService) GetRHistoryByID(transaction *Models.Transaction, id string) (err error){
//	return c.cust.GetRHistoryByID(transaction,id)
//}
//
