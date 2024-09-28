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
	_, code := pac.service.Create(c)
	response.MessageResponse(c, code.Code())
}
