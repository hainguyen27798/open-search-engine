package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/open-typesense-search/internal/dto"
	"github.com/hainguyen27798/open-typesense-search/internal/models"
	"github.com/hainguyen27798/open-typesense-search/internal/pkg/response"
	"github.com/hainguyen27798/open-typesense-search/internal/repos"
	"github.com/hainguyen27798/open-typesense-search/pkg/utils"
)

type ProductAdminServiceImpl struct {
	repo repos.IProductRepo
}

func NewProductAdminServiceImpl(repo repos.IProductRepo) *ProductAdminServiceImpl {
	return &ProductAdminServiceImpl{repo}
}

func (pas *ProductAdminServiceImpl) Create(c *gin.Context) *response.ServerCode {
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

func (pas *ProductAdminServiceImpl) Get(c *gin.Context) (*dto.ProductDto, *response.ServerCode) {
	id := c.Param("id")
	product, code := pas.repo.GetProduct(id)
	if product == nil {
		return nil, code
	}
	return utils.ModelToDto[dto.ProductDto](*product), nil
}

func (pas *ProductAdminServiceImpl) Delete(c *gin.Context) *response.ServerCode {
	id := c.Param("id")
	return pas.repo.DeleteProduct(id)
}

func (pas *ProductAdminServiceImpl) Update(c *gin.Context) *response.ServerCode {
	return response.ReturnCode(response.UpdatedSuccess)
}
