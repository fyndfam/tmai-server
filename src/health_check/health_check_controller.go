package health_check

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func MountRoutes(mongoClient *mongo.Client, app *fiber.App) {
	app.Get("/health-check", createHealthCheckEndpoint(mongoClient))
}

func createHealthCheckEndpoint(mongoClient *mongo.Client) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := mongoClient.Ping(context.TODO(), nil)

		if err == nil {
			ctx.Status(200).SendString("OK")
		}

		return err
	}
}
