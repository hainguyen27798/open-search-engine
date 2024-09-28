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
	global.Config = &settings.Config{
		Server: settings.ServerSettings{
			Port: serverPort,
			Mode: os.Getenv("SERVER_MODE"),
		},
		Typesense: settings.TypesenseSettings{
			Host:   os.Getenv("TYPESENSE_HOST"),
			ApiKey: os.Getenv("TYPESENSE_API_KEY"),
		},
	}
}
