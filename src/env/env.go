package env

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Env struct {
	MongoClient    *mongo.Client
	UserCollection *mongo.Collection
	PostCollection *mongo.Collection
}

func InitializeEnvironment() *Env {
	mongoClientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URL"))
	mongoClient, err := mongo.Connect(context.TODO(), mongoClientOptions)

	if err != nil {
		log.Fatal("Failed to connect to mongodb, panic...")
		panic(err)
	}

	databaseName := "tmai"

	env := Env{
		MongoClient:    mongoClient,
		UserCollection: mongoClient.Database(databaseName).Collection("users"),
		PostCollection: mongoClient.Database(databaseName).Collection("posts"),
	}

	return &env
}

func ShutdownEnvironment(env *Env) {
	if env == nil {
		return
	}

	if err := env.MongoClient.Disconnect(context.TODO()); err != nil {
		log.Fatal("Failed to disconnect from mongodb, panic...")
		panic(err)
	}
}
