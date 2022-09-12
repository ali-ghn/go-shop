package repositories

import (
	"context"

	md "github.com/ali-ghn/go-shop/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IProductRepository interface {
	CreateProduct(p *md.Product) (*md.Product, error)
	ReadProduct(id string) (*md.Product, error)
	ReadProducts() (*[]md.Product, error)
	UpdateProduct(p *md.Product) (*md.Product, error)
}

type ProductRepository struct {
	client *mongo.Client
}

func NewProductRepository(client *mongo.Client) *ProductRepository {
	return &ProductRepository{
		client: client,
	}
}

func (pr ProductRepository) CreateProduct(p *md.Product) (*md.Product, error) {
	res, err := pr.client.Database(DatabaseName).Collection(ProductCollectionName).InsertOne(context.TODO(), p)
	if err != nil {
		return nil, err
	}
	p.ProductId = res.InsertedID.(primitive.ObjectID).Hex()
	return p, nil
}

func (pr ProductRepository) ReadProduct(id string) (*md.Product, error) {
	return nil, nil
}

func (pr ProductRepository) ReadProducts() (*[]md.Product, error) {
	return nil, nil
}

func (pr ProductRepository) UpdateProduct(p *md.Product) (*md.Product, error) {
	return nil, nil
}
