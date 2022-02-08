package controller

import (
	"context"

	"github.com/fyndfam/tmai-server/src/env"
	"github.com/gofiber/fiber/v2"
)

func MountHealthCheckRoutes(env *env.Env, app *fiber.App) {
	app.Get("/health-check", createHealthCheckEndpoint(env))
}

func createHealthCheckEndpoint(env *env.Env) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := env.MongoClient.Ping(context.TODO(), nil)

		if err == nil {
			ctx.Status(200).SendString("OK")
		}

		return err
	}
}
