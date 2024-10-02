package repos

import (
	"context"
	"github.com/hainguyen27798/open-search-engine/internal/domain/products/models"
	"github.com/hainguyen27798/open-search-engine/pkg/response"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
)

type IProductRepo interface {
	InsertProduct(product models.Product) *response.ServerCode
	GetProduct(id string) (*models.Product, *response.ServerCode)
	DeleteProduct(id string) *response.ServerCode
}

type productRepo struct {
	collection *mongo.Collection
}

func NewProductRepo(db *mongo.Database) IProductRepo {
	return &productRepo{
		db.Collection("products"),
	}
}

func (repo *productRepo) InsertProduct(product models.Product) *response.ServerCode {
	if err := product.Creating(); err != nil {
		return response.ReturnCode(response.ErrInternalError)
	}
	_, err := repo.collection.InsertOne(context.TODO(), product)
	if err != nil {
		log.Println(err)
		return response.ReturnCode(response.ErrBadRequest)
	}
	return response.ReturnCode(response.CreatedSuccess)
}

func (repo *productRepo) GetProduct(id string) (*models.Product, *response.ServerCode) {
	var product models.Product
	objId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, response.ReturnCode(response.ErrCodeParamInvalid)
	}
	filter := bson.D{{"_id", objId}}
	if err := repo.collection.FindOne(context.TODO(), filter).Decode(&product); err != nil {
		log.Println(err)
		return nil, response.ReturnCode(response.ErrNotFound)
	}
	return &product, nil
}

func (repo *productRepo) DeleteProduct(id string) *response.ServerCode {
	objId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return response.ReturnCode(response.ErrBadRequest)
	}
	filter := bson.D{{"_id", objId}}
	rs, err := repo.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		return response.ReturnCode(response.ErrBadRequest)
	}
	if rs.DeletedCount == 0 {
		return response.ReturnCode(response.ErrNotFound)
	}
	return response.ReturnCode(response.DeletedSuccess)
}
