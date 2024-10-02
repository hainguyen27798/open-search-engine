package initialize

import (
	"context"
	"github.com/hainguyen27798/open-search-engine/global"
	"github.com/hainguyen27798/open-search-engine/internal/domain/core_embedding/services"
	"github.com/qdrant/go-client/qdrant"
	"log"
)

func InitQDrant() {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: global.Config.QDrant.Host,
		Port: global.Config.QDrant.Port,
	})
	if err != nil {
		panic(err)
	}
	check, err := client.HealthCheck(context.Background())
	if err != nil {
		panic(err)
	}
	log.Println(check.String())
	global.QDrantClient = client

	textEmbeddingService := services.NewTextEmbeddingService(client)
	textEmbeddingService.InitTextCollection("products")
}
