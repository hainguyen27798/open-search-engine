package services

import "github.com/qdrant/go-client/qdrant"

type (
	ITextEmbeddingService interface {
		AddNewTextVector(collectionName string, id string, vectorObj map[string]*qdrant.Vector) bool
		InitTextCollection(collectionName string)
		CollectionExists(collectionName string) bool
	}
)

func NewTextEmbeddingService(client *qdrant.Client) ITextEmbeddingService {
	return NewTextEmbeddingServiceImpl(client)
}
