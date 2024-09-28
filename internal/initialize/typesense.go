package initialize

import (
	"context"
	"github.com/hainguyen27798/open-typesense-search/global"
	"github.com/typesense/typesense-go/v2/typesense"
	"log"
	"time"
)

func InitTypesense() {
	global.Typesense = typesense.NewClient(
		typesense.WithServer(global.Config.Typesense.Host),
		typesense.WithAPIKey(global.Config.Typesense.ApiKey),
		typesense.WithNumRetries(5),
		typesense.WithRetryInterval(1*time.Second),
		typesense.WithHealthcheckInterval(2*time.Minute),
	)
	health, err := global.Typesense.Health(context.Background(), time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Typesense health: %v", health)
}
