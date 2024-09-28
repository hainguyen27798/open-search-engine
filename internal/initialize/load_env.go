package initialize

import (
	"github.com/hainguyen27798/open-typesense-search/global"
	"github.com/hainguyen27798/open-typesense-search/pkg/settings"
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
	global.Config = &settings.Config{
		Server: settings.ServerSettings{
			Port: serverPort,
			Mode: os.Getenv("SERVER_MODE"),
		},
		Typesense: settings.TypesenseSettings{
			Host:   os.Getenv("TYPESENSE_HOST"),
			ApiKey: os.Getenv("TYPESENSE_API_KEY"),
		},
		MongoDB: settings.MongoDBSettings{
			URL:         os.Getenv("MONGO_URL"),
			Username:    os.Getenv("MONGO_USERNAME"),
			Password:    os.Getenv("MONGO_PASSWORD"),
			MaxPoolSize: mongoPoolSize,
		},
	}
}
