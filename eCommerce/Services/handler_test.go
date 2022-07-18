package Services

import (
	"eCommerce/Models"
	repomock "eCommerce/Repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCustomerAccount(t *testing.T){
	assertion := assert.New(t)

	ctrl := gomock.NewController(t)
	repo := repomock.NewMockRepo(ctrl)
	orderService := NewService(repo)

	data := Models.Customer{
	CustomerId:     	"1",
	CustomerName:		"Pulkit",
	CustomerAddress:	"Jaipur",
	}

	repo.EXPECT().CreateCustomerAccount(gomock.Any()).Return(nil).Times(1)

	res, err := orderService.CreateCustomerAccount(&data)
	assertion.Nil(err)
	assertion.Equal(res.CustomerId, data.CustomerId)
	assertion.Equal(res.CustomerName, data.CustomerName)
	assertion.Equal(res.CustomerAddress, data.CustomerAddress)

}

func TestCreateRetailerAccount(t *testing.T){
	assertion := assert.New(t)

	ctrl := gomock.NewController(t)
	repo := repomock.NewMockRepo(ctrl)
	orderService := NewService(repo)

	data := Models.Retailer{
		RetailerId:     	"2",
		RetailerName:		"Mayank",
		RetailerAddress:	"Delhi",
	}

	repo.EXPECT().CreateRetailerAccount(gomock.Any()).Return(nil).Times(1)

	res, err := orderService.CreateRetailerAccount(&data)
	assertion.Nil(err)
	assertion.Equal(res.RetailerId, data.RetailerId)
	assertion.Equal(res.RetailerName, data.RetailerName)
	assertion.Equal(res.RetailerAddress, data.RetailerAddress)

}

func TestCreateBuyProduct(t *testing.T){
	assertion := assert.New(t)

	ctrl := gomock.NewController(t)
	repo := repomock.NewMockRepo(ctrl)
	orderService := NewService(repo)

	data1 := Models.Product{
		ProductId:     			"2",
		Price:					10000,
		Quantity:				10,
	}
	data2 := Models.Transaction{
		TransactionId:     		"2",
		ProdId:					"2",
		TransactionQuantity:	10000,
	}

	repo.EXPECT().BuyProduct(gomock.Any(),gomock.Any()).Return(nil, nil).Times(1)

	res1, res2, err := orderService.BuyProduct(&data1,&data2)
	assertion.Nil(err)
	assertion.Equal(res1.ProductId, data1.ProductId)
	assertion.Equal(res1.Price, data1.Price)
	assertion.Equal(res1.Quantity, data1.Quantity)

	assertion.Equal(res2.TransactionId, data2.TransactionId)
	assertion.Equal(res2.ProdId, data2.ProdId)
	assertion.Equal(res2.TransactionQuantity, data2.TransactionQuantity)

}
