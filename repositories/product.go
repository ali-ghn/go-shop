package repositories

import (
	"myapp/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var collectionName string = "Products"
var DbName string = "go-shop"

type ProductRepository struct {
	session    *mgo.Session
	Collection *mgo.Collection
}

func NewProductRepository(session *mgo.Session) *ProductRepository {
	return &ProductRepository{
		session:    session,
		Collection: session.DB(DbName).C(collectionName),
	}
}

func (pr ProductRepository) CreateProduct(name string, price float32) error {
	p := models.Product{
		Name:  name,
		Id:    bson.NewObjectId(),
		Price: price,
	}
	err := pr.Collection.Insert(p)
	if err != nil {
		return err
	}
	return nil
}
