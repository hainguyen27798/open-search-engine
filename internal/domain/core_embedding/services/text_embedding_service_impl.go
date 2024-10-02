package services

import (
	"fmt"
	"github.com/qdrant/go-client/qdrant"
	"log"
)

type TextEmbeddingServiceImpl struct {
	client *qdrant.Client
}

func NewTextEmbeddingServiceImpl(client *qdrant.Client) *TextEmbeddingServiceImpl {
	return &TextEmbeddingServiceImpl{
		client,
	}
}

func (te *TextEmbeddingServiceImpl) CollectionExists(collectionName string) bool {
	exists, err := te.client.CollectionExists(ctx, collectionName)
	if err != nil {
		panic(err)
		return false
	}
	return exists
}

func (te *TextEmbeddingServiceImpl) InitTextCollection(collectionName string) {
	if !te.CollectionExists(collectionName) {
		if err := te.client.CreateCollection(ctx, &qdrant.CreateCollection{
			CollectionName: collectionName,
			VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
				Size:     768,
				Distance: qdrant.Distance_Cosine,
			}),
		}); err != nil {
			panic(err)
		}
	} else {
		log.Println(fmt.Sprintf("collection %s does not exist", collectionName))
	}
}

func (te *TextEmbeddingServiceImpl) AddNewTextVector(collectionName string, id string, vectorObj map[string]*qdrant.Vector) bool {
	_, err := te.client.Upsert(ctx, &qdrant.UpsertPoints{
		CollectionName: collectionName,
		Points: []*qdrant.PointStruct{
			{
				Id:      qdrant.NewIDUUID(id),
				Vectors: qdrant.NewVectorsMap(vectorObj),
			},
		},
	})
	if err != nil {
		panic(err)
		return false
	}
	return true
}
