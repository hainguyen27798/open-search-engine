package repos

import (
	"context"
	"github.com/hainguyen27798/open-typesense-search/internal/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
)

type IProductRepo interface {
	InsertProduct(product models.Product) error
}

type productRepo struct {
	collection *mongo.Collection
}

func NewProductRepo(db *mongo.Database) IProductRepo {
	return &productRepo{
		db.Collection("products"),
	}
}

func (repo *productRepo) InsertProduct(product models.Product) error {
	if err := product.Creating(); err != nil {
		return err
	}
	_, err := repo.collection.InsertOne(context.Background(), product)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
