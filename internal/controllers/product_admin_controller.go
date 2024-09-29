package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/open-typesense-search/internal/pkg/response"
	"github.com/hainguyen27798/open-typesense-search/internal/services"
)

type ProductAdminController struct {
	service services.IProductAdminService
}

func NewProductAdminController(service services.IProductAdminService) *ProductAdminController {
	return &ProductAdminController{service}
}

func (pac *ProductAdminController) CreateProduct(c *gin.Context) {
	code := pac.service.Create(c)
	response.MessageResponse(c, code.Code())
}

func (pac *ProductAdminController) GetProduct(c *gin.Context) {
	product, code := pac.service.Get(c)
	if product == nil {
		response.NotFoundException(c, code.Code())
		return
	}
	response.OkResponse(c, response.CodeSuccess, *product)
}

func (pac *ProductAdminController) DeleteProduct(c *gin.Context) {
	code := pac.service.Delete(c)
	response.MessageResponse(c, code.Code())
}
