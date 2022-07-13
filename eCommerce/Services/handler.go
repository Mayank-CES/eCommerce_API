package Services

import (
	"eCommerce/Models"
	"eCommerce/Repository"
)

type Services interface {
	/*---Customer---*/

	CreateCustomerAccount(customer *Models.Customer)	(err error)
	BuyProduct(product *Models.Product, order *Models.Order)	(err error)
	CheckOrderByID(transaction *Models.Transaction, id string) ( err error)
	AddTransaction(order *Models.Order) (err error)

	/*---Product---*/
	GetAllProducts(products *[]Models.Product)	(err error)
	GetProductByID(product *Models.Product, id string)	(err error)

	/*---Retailer---*/
	CreateRetailerAccount(retailer *Models.Retailer)	(err error)
	AddProduct(product *Models.Product)	(err error)
	PatchProduct(updatedProduct *Models.PatchProd, id string)	(err error)
	GetRHistoryByID(transaction *Models.Transaction, id string) (err error)

}

type Server struct {
	Service Repository.Repo
}
func NewService(repo Repository.Repo) Services {
	return &Server{
		Service: repo,
	}
}



func (s *Server) CreateCustomerAccount(customer *Models.Customer)	(err error){
	return s.Service.CreateCustomerAccount(customer)
}


func (s *Server) BuyProduct(product *Models.Product, order *Models.Order)	(err error){
	return s.Service.BuyProduct(product,order)
}

func (s *Server) AddTransaction(order *Models.Order)	(err error){
	return s.Service.AddTransaction(order)
}


func (s *Server) CheckOrderByID(transaction *Models.Transaction, id string) 	(err error){
	return s.Service.CheckOrderByID(transaction, id)
}

func (s *Server) GetAllProducts(products *[]Models.Product)	(err error){
	return s.Service.GetAllProducts(products)
}

func (s *Server) GetProductByID(product *Models.Product, id string)	(err error){
	return s.Service.GetProductByID(product,id)
}

/*----Retailer----*/

func (s *Server) CreateRetailerAccount(retailer *Models.Retailer)	(err error){
	return s.Service.CreateRetailerAccount(retailer)
}

func (s *Server) AddProduct(product *Models.Product)	(err error){
	return s.Service.AddProduct(product)
}

func (s *Server) PatchProduct(updatedProduct *Models.PatchProd,id string)	(err error){
	return s.Service.PatchProduct(updatedProduct,id)
}

func (s *Server) GetRHistoryByID(transaction *Models.Transaction, id string) (err error){
	return s.Service.GetRHistoryByID(transaction,id)
}

