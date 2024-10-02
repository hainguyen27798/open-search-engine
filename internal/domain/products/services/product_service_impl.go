package services

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/open-search-engine/internal/domain/products/dto"
	"github.com/hainguyen27798/open-search-engine/internal/domain/products/models"
	"github.com/hainguyen27798/open-search-engine/internal/domain/products/repos"
	"github.com/hainguyen27798/open-search-engine/pkg/response"
	"github.com/hainguyen27798/open-search-engine/pkg/utils"
)

type ProductServiceImpl struct {
	repo repos.IProductRepo
}

func NewProductServiceImpl(repo repos.IProductRepo) *ProductServiceImpl {
	return &ProductServiceImpl{repo}
}

func (pas *ProductServiceImpl) CreateProduct(c *gin.Context) *response.ServerCode {
	payload := utils.BodyToDto[dto.CreateProductDto](c)
	if payload == nil {
		return nil
	}

	product, errCode := utils.DtoToModel[models.Product](payload)
	if errCode != nil {
		return errCode
	}

	return pas.repo.InsertProduct(*product)
}

func (pas *ProductServiceImpl) GetProduct(c *gin.Context) (*dto.ProductDto, *response.ServerCode) {
	id := c.Param("id")
	product, code := pas.repo.GetProduct(id)
	if product == nil {
		return nil, code
	}
	return utils.ModelToDto[dto.ProductDto](*product), nil
}

func (pas *ProductServiceImpl) DeleteProduct(c *gin.Context) *response.ServerCode {
	id := c.Param("id")
	return pas.repo.DeleteProduct(id)
}

func (pas *ProductServiceImpl) UpdateProduct(c *gin.Context) *response.ServerCode {
	return response.ReturnCode(response.UpdatedSuccess)
}
