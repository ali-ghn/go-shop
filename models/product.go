package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	Name  string        `json:"name"`
	Id    bson.ObjectId `json:"id"`
	Price float32       `json:"price"`
}
