package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/open-search-engine/global"
	"github.com/hainguyen27798/open-search-engine/internal/wires"
)

type ProductRouter struct{}

func (p *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	productAdminHandler, _ := wires.InitProductAdminHandler(global.MongoDB)

	productAdminRouter := Router.Group("products")
	{
		productAdminRouter.POST("", productAdminHandler.CreateProduct)
		productAdminRouter.GET(":id", productAdminHandler.GetProduct)
		productAdminRouter.DELETE(":id", productAdminHandler.DeleteProduct)
	}
}
