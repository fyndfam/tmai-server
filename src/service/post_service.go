package service

import (
	"context"
	"log"
	"time"

	"github.com/fyndfam/tmai-server/src/env"
	"github.com/fyndfam/tmai-server/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPostByID(env *env.Env, postID string) (*model.PostModel, error) {
	objectId, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		log.Println("Invalid post ID")
		return nil, err
	}

	filter := bson.M{"_id": objectId}

	var post model.PostModel

	env.PostCollection.FindOne(context.TODO(), filter).Decode(&post)

	return &post, nil
}

func GetLatestPosts(env *env.Env, limit int64, offset int64) ([]*model.PostModel, error) {
	docs := make([]*model.PostModel, limit)

	options := options.Find()
	options.SetSort(bson.M{"createdAt": -1})
	options.SetSkip(offset)
	options.SetLimit(limit)

	cursor, err := env.PostCollection.Find(context.TODO(), bson.M{}, options)
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

func CreatePost(env *env.Env, user *model.UserModel, content string) (*model.PostModel, error) {
	var tags []string

	post := model.PostModel{
		ID:            primitive.NewObjectID(),
		Content:       content,
		Tags:          tags,
		View:          0,
		CreatedBy:     model.CreatedBy{Username: *user.Username, Avatar: *user.Avatar},
		ContentEdited: false,
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
	}

	_, err := env.PostCollection.InsertOne(context.TODO(), post)
	if err != nil {
		log.Println("error inserting post", err)
		return nil, err
	}

	return &post, nil
}

func ReplyPost(env *env.Env, user *model.UserModel, content string, replyPostId string) (*model.PostModel, error) {
	var tags []string

	replyPostObjectId, err := primitive.ObjectIDFromHex(replyPostId)
	if err != nil {
		log.Println("invalid reply post id", err)
		return nil, err
	}

	post := model.PostModel{
		ID:            primitive.NewObjectID(),
		Content:       content,
		Tags:          tags,
		View:          0,
		CreatedBy:     model.CreatedBy{Username: *user.Username, Avatar: *user.Avatar},
		ContentEdited: false,
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
		ReplyTo:       replyPostObjectId,
	}

	_, err = env.PostCollection.InsertOne(context.TODO(), post)
	if err != nil {
		log.Println("error inserting post", err)
		return nil, err
	}

	return &post, nil
}

func IncrementPostView(env *env.Env, postID string) error {
	objectId, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		log.Println("Invalid post ID")
		return err
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{"$inc": bson.M{"view": 1}}

	if _, err := env.PostCollection.UpdateOne(context.TODO(), filter, update); err != nil {
		log.Println("Error when incrementing post view count", err)
		return err
	}

	return nil
}
