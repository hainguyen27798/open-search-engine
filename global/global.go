package global

import (
	"github.com/hainguyen27798/open-typesense-search/pkg/settings"
	"github.com/typesense/typesense-go/v2/typesense"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	Config    *settings.Config
	Typesense *typesense.Client
	MongoDB   *mongo.Database
)
