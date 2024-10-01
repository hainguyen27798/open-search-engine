package global

import (
	"github.com/go-playground/validator/v10"
	"github.com/hainguyen27798/open-search-engine/pkg/settings"
	"github.com/qdrant/go-client/qdrant"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	Config       *settings.Config
	MongoDB      *mongo.Database
	Validate     *validator.Validate
	QDrantClient *qdrant.Client
)
