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
	db *mongo.Database
}

func NewProductRepo(db *mongo.Database) IProductRepo {
	return &productRepo{
		db,
	}
}

func (repo *productRepo) InsertProduct(product models.Product) error {
	_, err := repo.db.Collection("products").InsertOne(context.Background(), product)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
