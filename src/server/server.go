package server

import (
	"github.com/fyndfam/tmai-server/src/controller"
	"github.com/fyndfam/tmai-server/src/env"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewApp(env *env.Env) *fiber.App {
	app := fiber.New()

	app.Use(cors.New())

	controller.MountHealthCheckRoutes(env, app)
	controller.MountUserRoutes(env, app)
	controller.MountPostRoutes(env, app)

	return app
}
