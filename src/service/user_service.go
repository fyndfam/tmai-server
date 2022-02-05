package service

import (
	"context"
	"log"
	"time"

	"github.com/fyndfam/tmai-server/src/env"
	"github.com/fyndfam/tmai-server/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(env *env.Env, email string) (*model.UserModel, error) {
	user := model.UserModel{
		ID:        primitive.NewObjectID(),
		Email:     email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	_, err := env.UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(env *env.Env, email string) (*model.UserModel, error) {
	var user model.UserModel

	err := env.UserCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		log.Print(err)

		return nil, err
	}

	return &user, nil
}

func UpdateUsername(env *env.Env, username string, emailAddress string) error {
	filter := bson.M{"email": emailAddress}
	update := bson.M{"$set": bson.M{"username": username, "updatedAt": time.Now().UTC()}}

	var user model.UserModel

	if err := env.UserCollection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		log.Print("error when getting user", err)

		return err
	}

	if user.Username != nil {
		log.Println("username already exists, not updating")

		return nil
	}

	if _, err := env.UserCollection.UpdateOne(context.TODO(), filter, update); err != nil {
		log.Println("error update username", err)

		return err
	}

	return nil
}
