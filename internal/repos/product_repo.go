package repos

import (
	"context"
	"github.com/hainguyen27798/open-typesense-search/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
)

type IProductRepo interface {
	InsertProduct(product models.Product) error
	GetProduct(id string) (*models.Product, error)
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
	_, err := repo.collection.InsertOne(context.TODO(), product)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (repo *productRepo) GetProduct(id string) (*models.Product, error) {
	var product models.Product
	objId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{{"_id", objId}}
	if err := repo.collection.FindOne(context.TODO(), filter).Decode(&product); err != nil {
		log.Println(err)
		return nil, err
	}
	return &product, nil
}
