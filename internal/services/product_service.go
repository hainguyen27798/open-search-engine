package services

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen27798/open-search-engine/internal/dto"
	"github.com/hainguyen27798/open-search-engine/internal/pkg/response"
	"github.com/hainguyen27798/open-search-engine/internal/repos"
	"github.com/hainguyen27798/open-search-engine/internal/services/impl"
)

type (
	IProductAdminService interface {
		Create(c *gin.Context) *response.ServerCode
		Get(c *gin.Context) (*dto.ProductDto, *response.ServerCode)
		Delete(c *gin.Context) *response.ServerCode
		Update(c *gin.Context) *response.ServerCode
	}
	IProductSearchService interface{}
)

func NewProductAdminService(repo repos.IProductRepo) IProductAdminService {
	return impl.NewProductAdminServiceImpl(repo)
}
