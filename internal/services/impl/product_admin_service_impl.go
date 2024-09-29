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

func (pas *ProductAdminServiceImpl) Create(c *gin.Context) (*dto.ProductDto, *response.ServerCode) {
	payload := utils.BodyToDto[dto.CreateProductDto](c)
	if payload == nil {
		return nil, nil
	}

	product, errCode := utils.DtoToModel[models.Product](payload)
	if errCode != nil {
		return nil, errCode
	}

	if err := pas.repo.InsertProduct(*product); err != nil {
		return nil, response.ReturnCode(response.ErrBadRequest)
	}
	return nil, response.ReturnCode(response.CreatedSuccess)
}

func (pas *ProductAdminServiceImpl) Get(c *gin.Context) (*dto.ProductDto, *response.ServerCode) {
	id := c.Param("id")
	product, err := pas.repo.GetProduct(id)
	if err != nil {
		return nil, response.ReturnCode(response.ErrNotFound)
	}
	return utils.ModelToDto[dto.ProductDto](*product), nil
}

func (pas *ProductAdminServiceImpl) Delete(c *gin.Context) bool {
	//TODO implement me
	panic("implement me")
}

func (pas *ProductAdminServiceImpl) Update(c *gin.Context) dto.ProductDto {
	//TODO implement me
	panic("implement me")
}
