package server

import (
	"log"

	"github.com/fyndfam/tmai-server/src/health_check"
	"github.com/fyndfam/tmai-server/src/middleware"
	userModule "github.com/fyndfam/tmai-server/src/user"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewApp(mongoClient *mongo.Client) *fiber.App {
	// TODO: load all modules and pass in mongo client for each module

	app := fiber.New()

	health_check.MountRoutes(mongoClient, app)
	app.Get("/", middleware.GetJwtMiddleware(), middleware.GetPostJwtMiddleware(mongoClient), restricted)

	return app
}

func restricted(context *fiber.Ctx) error {
	user := context.Locals("user").(userModule.UserModel)
	log.Printf("user email is %s", user.Email)

	return context.JSON(map[string]string{"status": "ok"})
}
