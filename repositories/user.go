package repositories

import (
	"context"
	"fmt"

	"github.com/ali-ghn/go-shop/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	client *mongo.Client
}

type IUserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	ReadUser(id string) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{
		client: client,
	}
}

func (ur UserRepository) CreateUser(user *models.User) (*models.User, error) {
	res, err := ur.client.Database(DatabaseName).Collection(UserCollectionName).InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Printf("Could not insert document %v\n", user)
		return nil, err
	}
	user.UserId = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (ur UserRepository) ReadUser(id string) (*models.User, error) {
	user := models.User{}
	err := ur.client.Database(DatabaseName).Collection(UserCollectionName).FindOne(context.TODO(), bson.D{{"UserId", id}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (ur UserRepository) UpdateUser(user *models.User) (*models.User, error) {
	newUser := models.User{}
	err := ur.client.Database(DatabaseName).Collection(UserCollectionName).FindOneAndReplace(context.TODO(), bson.D{{"UserId", user.UserId}}, user).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return user, nil
}
