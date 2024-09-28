//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/google/wire"
	"github.com/hainguyen27798/open-typesense-search/internal/controllers"
	"github.com/hainguyen27798/open-typesense-search/internal/repos"
	"github.com/hainguyen27798/open-typesense-search/internal/services"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func InitProductAdminHandler(*mongo.Database) (*controllers.ProductAdminController, error) {
	wire.Build(
		repos.NewProductRepo,
		services.NewProductAdminService,
		controllers.NewProductAdminController,
	)
	return new(controllers.ProductAdminController), nil
}
