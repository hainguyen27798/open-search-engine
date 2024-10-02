package products

import (
	"context"
	"fmt"
	"github.com/hainguyen27798/open-search-engine/global"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func InitProductChangeStream() {
	collection := global.MongoDB.Collection("products")
	changeStream, err := collection.Watch(context.TODO(), mongo.Pipeline{})
	if err != nil {
		panic(err)
	}
	defer func(changeStream *mongo.ChangeStream, ctx context.Context) {
		_ = changeStream.Close(ctx)
	}(changeStream, context.TODO())

	for changeStream.Next(context.TODO()) {
		fmt.Println(changeStream.Current)
	}
}
