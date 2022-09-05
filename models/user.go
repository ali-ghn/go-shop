package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	UserId    primitive.ObjectID
	FirstName string
	LastName  string
	Email     string
	Password  string
	Roles     []string
}
