package services

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/open-search-engine/internal/domain/products/dto"
	"github.com/hainguyen27798/open-search-engine/internal/domain/products/repos"
	"github.com/hainguyen27798/open-search-engine/pkg/response"
)

type (
	IProductService interface {
		CreateProduct(c *gin.Context) *response.ServerCode
		GetProduct(c *gin.Context) (*dto.ProductDto, *response.ServerCode)
		DeleteProduct(c *gin.Context) *response.ServerCode
		UpdateProduct(c *gin.Context) *response.ServerCode
	}
)

func NewProductService(repo repos.IProductRepo) IProductService {
	return NewProductServiceImpl(repo)
}
