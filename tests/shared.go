package tests

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/fyndfam/tmai-server/src/env"
	"github.com/fyndfam/tmai-server/src/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var environment *env.Env

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

func GetEnvironment() *env.Env {
	return environment
}

func ClearDB() {
	if environment == nil {
		log.Fatalln("environment is not setup")
	}

	environment.PostCollection.DeleteMany(context.TODO(), bson.M{})
	environment.UserCollection.DeleteMany(context.TODO(), bson.M{})
}

func SetupTests() {
	if environment == nil {
		environment = env.InitializeEnvironment()
	}
}

func TearDownTests() {
	if environment != nil {
		env.ShutdownEnvironment(environment)
		environment = nil
	}
}

func GivenUser() {
	if environment == nil {
		log.Fatalln("environment is not setup")
	}

	user := model.UserModel{
		ID:        primitive.NewObjectID(),
		Email:     email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if _, err := environment.UserCollection.InsertOne(context.TODO(), user); err != nil {
		log.Fatalln("Failed to insert test user data")
	}
}

func GivenUserWithUsername() {
	if environment == nil {
		log.Fatalln("environment is not setup")
	}

	username := "test"

	user := model.UserModel{
		ID:        primitive.NewObjectID(),
		Email:     email,
		Username:  &username,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if _, err := environment.UserCollection.InsertOne(context.TODO(), user); err != nil {
		log.Fatalln("Failed to insert test user data")
	}
}
