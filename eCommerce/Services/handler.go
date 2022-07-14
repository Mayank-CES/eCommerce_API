package Services

import (
	"eCommerce/Models"
	"eCommerce/Repository"
)

type Services interface {
	/*---Customer---*/

	CreateCustomerAccount(customer *Models.Customer)	(err error)
	BuyProduct(product *Models.Product, transaction *Models.Transaction)	(err error)
	CheckOrderByID(transaction *Models.Transaction, id string) ( err error)
	BuyProduct2(order *Models.Transaction)	(err error)
	AddTransaction(order *Models.Transaction) (err error)

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


func (s *Server) BuyProduct(product *Models.Product, transaction *Models.Transaction)	(err error){
	return s.Service.BuyProduct(product,transaction)
}

func (s *Server) BuyProduct2(transaction *Models.Transaction)	(err error){
	return s.Service.BuyProduct2(transaction)
}

func (s *Server) AddTransaction(transaction *Models.Transaction)	(err error){
	return s.Service.AddTransaction(transaction)
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

