package products

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/open-search-engine/internal/domain/products/services"
	"github.com/hainguyen27798/open-search-engine/pkg/response"
)

type ProductController struct {
	service services.IProductService
}

func NewProductController(service services.IProductService) *ProductController {
	return &ProductController{service}
}

func (pac *ProductController) CreateProduct(c *gin.Context) {
	code := pac.service.CreateProduct(c)
	if code == nil {
		return
	}
	response.MessageResponse(c, code.Code())
}

func (pac *ProductController) GetProduct(c *gin.Context) {
	product, code := pac.service.GetProduct(c)
	if product == nil {
		response.NotFoundException(c, code.Code())
		return
	}
	response.OkResponse(c, response.CodeSuccess, *product)
}

func (pac *ProductController) DeleteProduct(c *gin.Context) {
	code := pac.service.DeleteProduct(c)
	response.MessageResponse(c, code.Code())
}
