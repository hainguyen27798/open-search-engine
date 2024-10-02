//go:build wireinject
// +build wireinject

package products

import (
	"github.com/google/wire"
	"github.com/hainguyen27798/open-search-engine/internal/domain/products/repos"
	"github.com/hainguyen27798/open-search-engine/internal/domain/products/services"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func InitProductHandler(*mongo.Database) (*ProductController, error) {
	wire.Build(
		repos.NewProductRepo,
		services.NewProductService,
		NewProductController,
	)
	return new(ProductController), nil
}
