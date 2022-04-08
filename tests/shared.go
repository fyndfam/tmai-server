package tests

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fyndfam/tmai-server/src/env"
	"github.com/fyndfam/tmai-server/src/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Bearer string = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNjQzODg5Njc4LCJleHAiOjQxMDI0NDQ4MDAsImF1ZCI6InRtYWlzZXJ2ZXIiLCJpc3MiOiJ0bWFpc2VydmVyIiwiZW1haWwiOiJ0ZXN0QHRtYWkuY28ifQ.P0H878UgorhlE3zT9l9HaAiX4fg0Esd35SZNfKjyJRs"

const email string = "test@tmai.co"
const email2 string = "test_user@tmai.co"

var UserId string = ""
var UserId2 string = ""

func FiberToHandler(app *fiber.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := app.Test(r)
		if err != nil {
			panic(err)
		}

		// copy headers
		for k, vv := range resp.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(resp.StatusCode)

		if _, err := io.Copy(w, resp.Body); err != nil {
			panic(err)
		}
	}
}

func ClearDB(environment *env.Env) {
	if environment == nil {
		log.Fatalln("environment is not setup")
	}

	environment.PostCollection.DeleteMany(context.TODO(), bson.M{})
	environment.UserCollection.DeleteMany(context.TODO(), bson.M{})
}

type TestAvatarService struct{}

func (svc *TestAvatarService) GenerateUserAvatar(userId string) (*string, error) {
	avatarURL := fmt.Sprintf("avatar/avatar_%v.png", userId)
	return &avatarURL, nil
}

func SetupTests() *env.Env {
	mongoClientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URL"))
	mongoClient, err := mongo.Connect(context.TODO(), mongoClientOptions)

	if err != nil {
		log.Fatal("Failed to connect to mongodb, panic...")
		panic(err)
	}

	databaseName := "tmai"

	env := env.Env{
		MongoClient:    mongoClient,
		UserCollection: mongoClient.Database(databaseName).Collection("users"),
		PostCollection: mongoClient.Database(databaseName).Collection("posts"),
		AvatarService:  &TestAvatarService{},
	}

	return &env
}

func TearDownTests(environment *env.Env) {
	if environment != nil {
		env.ShutdownEnvironment(environment)
	}
}

func GivenUser(environment *env.Env) {
	if environment == nil {
		log.Fatalln("environment is not setup")
	}

	avatar := "avatar/avatar_1.png"

	user := model.UserModel{
		ID:             primitive.NewObjectID(),
		ExternalUserId: "1234567890",
		Avatar:         &avatar,
		Email:          email,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}

	if _, err := environment.UserCollection.InsertOne(context.TODO(), user); err != nil {
		log.Fatalln("Failed to insert test user data")
	}
}

func GivenUserWithUsername(environment *env.Env) {
	if environment == nil {
		log.Fatalln("environment is not setup")
	}

	username := "test"
	avatar := "avatar/avatar_1.png"

	user := model.UserModel{
		ID:             primitive.NewObjectID(),
		ExternalUserId: "1234567890",
		Email:          email,
		Username:       &username,
		Avatar:         &avatar,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}

	if _, err := environment.UserCollection.InsertOne(context.TODO(), user); err != nil {
		log.Fatalln("Failed to insert test user data")
	}
}

func GivenPost(environment *env.Env, postContent string) string {
	if environment == nil {
		log.Fatalln("environment is not setup")
	}

	postID := primitive.NewObjectID()
	var tags []string

	post := model.PostModel{
		ID:            postID,
		Content:       postContent,
		Tags:          tags,
		CreatedBy:     model.CreatedBy{Username: "test", Avatar: "avatar/avatar_1.png"},
		View:          0,
		ContentEdited: false,
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
	}

	if _, err := environment.PostCollection.InsertOne(context.TODO(), post); err != nil {
		log.Fatalln("Failed to insert test post data")
	}

	return postID.Hex()
}

func GivenPosts(environment *env.Env) {
	if environment == nil {
		log.Fatalln("environment is not setup")
	}

	var tags []string

	postContent := [2]string{"hey, this is the first post", "oh, this is interesting, i want it!"}
	dates := [2]time.Time{time.Date(2022, time.January, 5, 0, 0, 0, 0, time.UTC), time.Date(2022, time.February, 3, 0, 0, 0, 0, time.UTC)}

	for i := 0; i < len(postContent); i++ {
		post := model.PostModel{
			ID:            primitive.NewObjectID(),
			Content:       postContent[i],
			Tags:          tags,
			CreatedBy:     model.CreatedBy{Username: "test", Avatar: "avatar/avatar_1.png"},
			View:          0,
			ContentEdited: false,
			CreatedAt:     dates[i],
			UpdatedAt:     dates[i],
		}

		if _, err := environment.PostCollection.InsertOne(context.TODO(), post); err != nil {
			log.Fatalln("Failed to insert test posts data")
		}
	}
}
