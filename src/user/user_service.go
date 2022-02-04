package user

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(mongoClient *mongo.Client, email string) (*UserModel, error) {
	user := UserModel{
		ID:        primitive.NewObjectID(),
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	collection := mongoClient.Database("tmai").Collection("users")

	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(mongoClient *mongo.Client, email string) (*UserModel, error) {
	collection := mongoClient.Database("tmai").Collection("users")

	var user UserModel

	err := collection.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		log.Print(err)

		return nil, err
	}

	return &user, nil
}
