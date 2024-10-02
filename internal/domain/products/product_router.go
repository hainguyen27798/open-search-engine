package products

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/open-search-engine/global"
)

func InitProductRouter(Router *gin.RouterGroup) {
	productAdminHandler, _ := InitProductHandler(global.MongoDB)

	productRouter := Router.Group("products")
	{
		productRouter.POST("", productAdminHandler.CreateProduct)
		productRouter.GET(":id", productAdminHandler.GetProduct)
		productRouter.DELETE(":id", productAdminHandler.DeleteProduct)
	}
}
