package initialize

import (
	"context"
	"github.com/hainguyen27798/open-typesense-search/global"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"log"
)

func InitMongo() {
	client, err := mongo.Connect(
		options.
			Client().
			ApplyURI(global.Config.MongoDB.URL).
			SetMaxPoolSize(global.Config.MongoDB.MaxPoolSize).
			SetAuth(options.Credential{
				AuthSource: "admin",
				Username:   global.Config.MongoDB.Username,
				Password:   global.Config.MongoDB.Password,
			}),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	global.MongoDB = client
}
