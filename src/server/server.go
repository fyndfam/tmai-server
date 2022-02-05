package server

import (
	"github.com/fyndfam/tmai-server/src/controller"
	"github.com/fyndfam/tmai-server/src/middleware"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewApp(mongoClient *mongo.Client) *fiber.App {
	app := fiber.New()

	controller.MountHealthCheckRoutes(mongoClient, app)
	controller.MountUserRoutes(mongoClient, app)
	controller.MountPostRoutes(mongoClient, app)

	app.Get("/", middleware.GetJwtMiddleware(), middleware.GetPostJwtMiddleware(mongoClient), restricted)

	return app
}

func restricted(context *fiber.Ctx) error {
	return context.JSON(map[string]string{"status": "ok"})
}
