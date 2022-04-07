package env

import (
	"context"
	"log"
	"os"

	"github.com/fyndfam/tmai-server/src/avatar"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Env struct {
	MongoClient    *mongo.Client
	UserCollection *mongo.Collection
	PostCollection *mongo.Collection
	AvatarService  avatar.GenerateUserAvatarer
}

func getDatabaseName() string {
	databaseName := os.Getenv("DATABASE_NAME")

	if len(databaseName) == 0 {
		return "tmai"
	}
	return databaseName
}

type TasEnvironmentVariable struct {
	TasURL    string
	TasApiKey string
}

func getTasEnvironmentVariables() *TasEnvironmentVariable {
	tasURL := os.Getenv("TAS_URL")
	apiKey := os.Getenv("TAS_API_KEY")

	if len(tasURL) == 0 {
		log.Panic("TAS_URL is not set")
	}

	if len(apiKey) == 0 {
		log.Panic("TAS_API_KEY is not set")
	}

	return &TasEnvironmentVariable{
		TasURL:    tasURL,
		TasApiKey: apiKey,
	}
}

func InitializeEnvironment() *Env {
	tasEnvVar := getTasEnvironmentVariables()
	avatarService := avatar.DefaultAvatarService{TasURL: tasEnvVar.TasURL, ApiKey: tasEnvVar.TasApiKey}

	mongoClientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URL"))
	mongoClient, err := mongo.Connect(context.TODO(), mongoClientOptions)

	if err != nil {
		log.Fatal("Failed to connect to mongodb, panic...")
		panic(err)
	}

	databaseName := getDatabaseName()

	env := Env{
		MongoClient:    mongoClient,
		UserCollection: mongoClient.Database(databaseName).Collection("users"),
		PostCollection: mongoClient.Database(databaseName).Collection("posts"),
		AvatarService:  &avatarService,
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
