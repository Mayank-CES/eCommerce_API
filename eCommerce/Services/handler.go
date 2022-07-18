package Services

import (
	"eCommerce/Models"
	"eCommerce/Repository"
)

type Services interface {
	/*---Customer---*/

	CreateCustomerAccount(customer *Models.Customer)	(c *Models.Customer, err error)
	BuyProduct(product *Models.Product, transaction *Models.Transaction)	(p *Models.Product, order *Models.Transaction, err error)
	CheckOrderByID(transaction *Models.Transaction, id string) (t *Models.Transaction, roll string, err error)
	BuyProduct2(order *Models.Transaction)	(o *Models.Transaction, err error)
	AddTransaction(order *Models.Transaction) (o *Models.Transaction,err error)

	/*---Product---*/
	GetAllProducts(products *[]Models.Product)	(p *[]Models.Product, err error)
	GetProductByID(product *Models.Product, id string)	(p *Models.Product, roll string, err error)

	/*---Retailer---*/
	CreateRetailerAccount(retailer *Models.Retailer)	(r *Models.Retailer, err error)
	AddProduct(product *Models.Product)	(p *Models.Product, err error)
	PatchProduct(updatedProduct *Models.PatchProd, id string)	(uP *Models.PatchProd, roll string,err error)
	GetRHistoryByID(transaction *Models.Transaction, id string) (o *Models.Transaction, roll string, err error)

}

type Server struct {
	Service Repository.Repo
}
func NewService(repo Repository.Repo) Services {
	return &Server{
		Service: repo,
	}
}



func (s *Server) CreateCustomerAccount(customer *Models.Customer)	(c *Models.Customer, err error){

	return customer, s.Service.CreateCustomerAccount(customer)
}


func (s *Server) BuyProduct(product *Models.Product, transaction *Models.Transaction)	(p *Models.Product, order *Models.Transaction, err error){
	return product, transaction, s.Service.BuyProduct(product,transaction)
}

func (s *Server) BuyProduct2(transaction *Models.Transaction)	(o  *Models.Transaction, err error){
	return transaction, s.Service.BuyProduct2(transaction)
}

func (s *Server) AddTransaction(transaction *Models.Transaction)	(o *Models.Transaction,err error){
	return transaction, s.Service.AddTransaction(transaction)
}


func (s *Server) CheckOrderByID(transaction *Models.Transaction, id string) (order *Models.Transaction, roll string, err error){
	return transaction, id, s.Service.CheckOrderByID(transaction, id)
}

func (s *Server) GetAllProducts(products *[]Models.Product)	(p *[]Models.Product, err error){
	return p, s.Service.GetAllProducts(products)
}

func (s *Server) GetProductByID(product *Models.Product, id string)	(p *Models.Product, roll string, err error){
	return product, id, s.Service.GetProductByID(product,id)
}

/*----Retailer----*/

func (s *Server) CreateRetailerAccount(retailer *Models.Retailer)	(r *Models.Retailer,err error){
	return retailer, s.Service.CreateRetailerAccount(retailer)
}

func (s *Server) AddProduct(product *Models.Product)	(p *Models.Product, err error){
	return p, s.Service.AddProduct(product)
}

func (s *Server) PatchProduct(updatedProduct *Models.PatchProd,id string)	(up *Models.PatchProd, roll string,err error){
	return updatedProduct, id, s.Service.PatchProduct(updatedProduct,id)
}

func (s *Server) GetRHistoryByID(transaction *Models.Transaction, id string) (o *Models.Transaction, roll string, err error){
	return transaction, id, s.Service.GetRHistoryByID(transaction,id)
}

