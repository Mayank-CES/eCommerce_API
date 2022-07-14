package main

import (
	"eCommerce/Config"
	"eCommerce/Models"
	"eCommerce/Routes"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer func(DB *gorm.DB) {
		err := DB.Close()
		if err != nil {
			panic(err)
		}
	}(Config.DB)

	Config.DB.AutoMigrate(&Models.Customer{}, &Models.Retailer{}, &Models.Product{}, &Models.Transaction{})



	r := Routes.SetupRouter()
	//Services.NewCustomerService(Repository.NewRepo())



	err := r.Run()
	if err != nil {
		panic(err)
		return
	}
}


