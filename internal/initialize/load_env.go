package initialize

import (
	"github.com/hainguyen27798/open-search-engine/global"
	"github.com/hainguyen27798/open-search-engine/pkg/settings"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	serverPort, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))
	mongoPoolSize, _ := strconv.ParseUint(os.Getenv("MONGO_MAX_POOL_SIZE"), 10, 64)
	qdrantPort, _ := strconv.Atoi(os.Getenv("QDRANT_PORT"))
	global.Config = &settings.Config{
		Server: settings.ServerSettings{
			Port: serverPort,
			Mode: os.Getenv("SERVER_MODE"),
		},
		MongoDB: settings.MongoDBSettings{
			URL:         os.Getenv("MONGO_URL"),
			Username:    os.Getenv("MONGO_USERNAME"),
			Password:    os.Getenv("MONGO_PASSWORD"),
			Database:    os.Getenv("MONGO_BD_NAME"),
			MaxPoolSize: mongoPoolSize,
		},
		EmbeddingModel: settings.EmbeddingModelSettings{
			Auth:        os.Getenv("EMBEDDING_MODEL_TOKEN"),
			TextEMURL:   os.Getenv("TEXT_EMBEDDING_MODEL_URL"),
			TextEMName:  os.Getenv("TEXT_EMBEDDING_MODEL_NAME"),
			ImageEMURL:  os.Getenv("IMAGE_EMBEDDING_MODEL_URL"),
			ImageEMName: os.Getenv("IMAGE_EMBEDDING_MODEL_NAME"),
		},
		QDrant: settings.QDrantSettings{
			Host: os.Getenv("QDRANT_HOST"),
			Port: qdrantPort,
		},
	}
}
