package service

import (
	"context"
	"log"
	"time"

	"github.com/fyndfam/tmai-server/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(mongoClient *mongo.Client, email string) (*model.UserModel, error) {
	user := model.UserModel{
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

func GetUserByEmail(mongoClient *mongo.Client, email string) (*model.UserModel, error) {
	collection := mongoClient.Database("tmai").Collection("users")

	var user model.UserModel

	err := collection.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		log.Print(err)

		return nil, err
	}

	return &user, nil
}

func UpdateUsername(mongoClient *mongo.Client, username string, emailAddress string) error {
	collection := mongoClient.Database("tmai").Collection("users")

	filter := bson.D{{"email", emailAddress}}
	update := bson.D{{"$set", bson.D{{"username", username}}}}

	var user model.UserModel

	if err := collection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		log.Print("error when getting user", err)

		return err
	}

	if user.Username != nil {
		log.Println("username already exists, not updating")

		return nil
	}

	if _, err := collection.UpdateOne(context.TODO(), filter, update); err != nil {
		log.Println("error update username", err)

		return err
	}

	return nil
}
