package main

import (
	"context"
	"log"
	"os"

	"github.com/fyndfam/tmai-server/src/server"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoClientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URL"))
	mongoClient, err := mongo.Connect(context.TODO(), mongoClientOptions)

	if err != nil {
		log.Fatal("Failed to connect to mongodb, panic...")
		panic(err)
	}

	defer func() {
		err := mongoClient.Disconnect(context.TODO())

		if err != nil {
			log.Fatal("Failed to disconnect from mongodb, panic...")
			panic(err)
		}
	}()

	app := server.NewApp(mongoClient)

	app.Listen(":8088")
}
