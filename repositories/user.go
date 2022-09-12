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
	ReadUserByEmail(email string) (*models.User, error)
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{
		client: client,
	}
}

func (ur UserRepository) CreateUser(user *models.User) (*models.User, error) {
	user.Id = primitive.NewObjectID()
	res, err := ur.client.Database(DatabaseName).Collection(UserCollectionName).InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Printf("Could not insert document %v\n", user)
		fmt.Println(err)
		return nil, err
	}
	user.Id = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (ur UserRepository) ReadUser(id string) (*models.User, error) {
	user := models.User{}
	bid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", bid}}
	err := ur.client.Database(DatabaseName).Collection(UserCollectionName).FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (ur UserRepository) UpdateUser(user *models.User) (*models.User, error) {
	newUser := models.User{}
	filter := bson.D{{"_id", user.Id}}
	err := ur.client.Database(DatabaseName).Collection(UserCollectionName).FindOneAndReplace(context.TODO(), filter, user).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur UserRepository) ReadUserByEmail(email string) (*models.User, error) {
	filter := bson.D{{"email", email}}
	user := models.User{}
	err := ur.client.Database(DatabaseName).Collection(UserCollectionName).FindOne(context.TODO(), filter).Decode(&user)
	return &user, err
}
