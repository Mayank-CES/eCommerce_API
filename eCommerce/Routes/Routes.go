
package Routes

import (
	"eCommerce/Controllers"
	"eCommerce/Repository"
	"eCommerce/Services"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	router := gin.Default()
	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"mayank":"passes",
	}))

	ctrl := Controllers.NewController(Services.NewService(Repository.NewRepo()))

	gp := authorized.Group("/")
	{
		customerRoutes := gp.Group("/customer")
		{
			customerRoutes.POST("/add", ctrl.CreateCustomerAccount)
			customerRoutes.GET("/order/:t_id", ctrl.CheckOrderByID)
			customerRoutes.PATCH("/:id", ctrl.BuyProduct)

		}
		retailerRoutes := gp.Group("/retailer")
		{
			retailerRoutes.POST("/add", ctrl.CreateRetailerAccount)
			retailerRoutes.PUT("/add", ctrl.AddProduct)
			retailerRoutes.PATCH("/:id", ctrl.PatchProduct)
			retailerRoutes.GET("/:id", ctrl.GetRHistoryByID)
		}
		productRoutes := gp.Group("/product")
		{
			productRoutes.GET("/", ctrl.GetAllProducts)
			productRoutes.GET("/:id", ctrl.GetProductByID)
		}
		//adminCheckRoutes := gp.Group("/admin")
		//{
		//	adminCheckRoutes.GET("/customer", ctrl.GetAllCustomer)
		//	adminCheckRoutes.GET("/retailer", ctrl.GetAllRetailer)
		//	adminCheckRoutes.DELETE("/customer", ctrl.DeleteCustomer)
		//	adminCheckRoutes.DELETE("/retailer", ctrl.DeleteRetailer)
		//	adminCheckRoutes.DELETE("/product", ctrl.DeleteProduct)
		//	adminCheckRoutes.DELETE("/transaction", ctrl.DeleteTransaction)
		//}
	}

	return router
}
