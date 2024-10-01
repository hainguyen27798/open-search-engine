//go:build wireinject
// +build wireinject

package wires

import (
	"github.com/google/wire"
	"github.com/hainguyen27798/open-search-engine/internal/controllers"
	"github.com/hainguyen27798/open-search-engine/internal/repos"
	"github.com/hainguyen27798/open-search-engine/internal/services"
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
