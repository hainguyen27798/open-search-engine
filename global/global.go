package global

import (
	"github.com/go-playground/validator/v10"
	"github.com/hainguyen27798/open-search-engine/pkg/settings"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	Config   *settings.Config
	MongoDB  *mongo.Database
	Validate *validator.Validate
)
