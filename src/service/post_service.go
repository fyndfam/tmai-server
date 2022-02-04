package service

import (
	"context"
	"log"
	"time"

	"github.com/fyndfam/tmai-server/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPostByID(mongoClient *mongo.Client, postID string) (*model.PostModel, error) {
	objectId, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		log.Println("Invalid post ID")
		return nil, err
	}

	filter := bson.D{{"_id", objectId}}

	var post model.PostModel

	collection := mongoClient.Database("tmai").Collection("posts")

	collection.FindOne(context.TODO(), filter).Decode(&post)

	return &post, nil
}

func GetLatestPosts(mongoClient *mongo.Client, limit int64, offset int64) ([]*model.PostModel, error) {
	collection := mongoClient.Database("tmai").Collection("posts")

	docs := make([]*model.PostModel, limit)

	options := options.Find()
	options.SetSort(bson.M{"createdAt": -1})
	options.SetSkip(offset)
	options.SetLimit(limit)

	cursor, err := collection.Find(context.TODO(), bson.M{}, options)
	if err != nil {
		log.Println("error finding latest post", err)
		return nil, err
	}
	if err = cursor.All(context.TODO(), &docs); err != nil {
		log.Println("error reading all docs", err)
		return nil, err
	}

	return docs, nil
}

func CreatePost(mongoClient *mongo.Client, username string, content string) (*model.PostModel, error) {
	collection := mongoClient.Database("tmai").Collection("posts")

	var tags []string

	post := model.PostModel{
		ID:            primitive.NewObjectID(),
		Content:       content,
		Tags:          tags,
		View:          0,
		CreatedBy:     username,
		ContentEdited: false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	_, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		log.Println("error inserting post", err)
		return nil, err
	}

	return &post, nil
}
